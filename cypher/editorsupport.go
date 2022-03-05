package cypher

import (
	"reflect"

	"github.com/abc-inc/merovingian/ast"
	"github.com/abc-inc/merovingian/comp"
	"github.com/abc-inc/merovingian/db/neo4j"
	"github.com/abc-inc/merovingian/lang"
	"github.com/abc-inc/merovingian/parser"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type EditorSupport struct {
	schema  neo4j.Schema
	input   string
	posConv PosConv

	parseTree           antlr.ParseTree
	parseErrors         []SynErr
	referencesProviders map[string]comp.Provider
	completion          comp.AutoCompletion
	statements          []parser.CypherPartContext
}

func NewEditorSupport(input string) *EditorSupport {
	e := &EditorSupport{}
	e.completion = *comp.NewAutoCompletion(neo4j.Schema{})
	e.Update(input)
	return e
}

func (es *EditorSupport) Update(input string) {
	es.posConv = *NewPosConv(input)

	es.input = input
	parseTree, referencesListener, errorListener, referencesProviders := Parse(input)
	es.parseTree = parseTree

	es.parseErrors = errorListener.errors

	es.statements = referencesListener.Statements
	es.referencesProviders = referencesProviders

	es.completion.UpdateReferenceProviders(es.referencesProviders)
}

func (es *EditorSupport) SetSchema(schema neo4j.Schema) {
	es.schema = schema
	es.completion.UpdateSchema(es.schema)
}

func (es EditorSupport) GetElement(line, column int) antlr.Tree {
	abs := es.posConv.ToAbsolute(line, column)
	return getElement(es.parseTree, abs)
}

func getElement(pt antlr.Tree, abs int) antlr.Tree {
	pos := ast.GetPosition(pt)
	if pos != nil && (abs < pos.GetStart() || abs > pos.GetStop()) {
		return nil
	}

	c := pt.GetChildCount()
	if c == 0 && pos != nil {
		return pt
	}

	for _, c := range pt.GetChildren() {
		if e := getElement(c, abs); e != nil {
			return e
		}
	}

	if pos != nil {
		return pt
	}
	return nil
}

func (es EditorSupport) GetReferences(line, column int) []antlr.ParserRuleContext {
	e := ast.FindAnyParent(es.GetElement(line, column), lang.SymbolicContexts)
	if e == nil {
		return nil
	}

	var query antlr.Tree
	typ := reflect.TypeOf(e).Elem().Name()
	if typ == lang.VARIABLE_CONTEXT {
		query = ast.FindAnyParent(e, []string{lang.QUERY_CONTEXT})
	}

	text := e.(antlr.ParseTree).GetText()
	if query == nil {
		return es.referencesProviders[typ].GetReferences(text, nil)
	}
	return es.referencesProviders[typ].GetReferences(text, query.(*parser.CypherQueryContext))
}

func (es EditorSupport) GetCompletionInfo(line, column int) comp.ComplInfo {
	element := es.GetElementForCompletion(line, column)
	query := ast.FindAnyParent(element, []string{lang.QUERY_CONTEXT})
	complInfo := comp.GetTypes(element)
	return comp.ComplInfo{
		Element: element,
		Query:   query,
		Found:   complInfo.Found,
		Types:   complInfo.Types,
	}
}

func (es EditorSupport) GetElementForCompletion(line, column int) antlr.Tree {
	e := es.GetElement(line, column)
	//for _, c := range e.GetParent().GetChildren() {
	//	fmt.Println(c, ">", c.GetPayload(), "<", c.(*antlr.TerminalNodeImpl).GetSymbol(), "-", reflect.TypeOf(c.(*antlr.TerminalNodeImpl).GetSymbol()))
	//}
	if p := ast.FindAnyParent(e, lang.CompletionCandidates); p != nil {
		return p
	}
	return e
}

func (es EditorSupport) GetCompletion(line, column int, doFilter bool) comp.Result {
	info := es.GetCompletionInfo(line, column)
	if !info.Found && column > 0 {
		if prevInfo := es.GetCompletionInfo(line, column-1); prevInfo.Found {
			info = prevInfo
		}
	}

	element, query, found, types := info.Element, info.Query, info.Found, info.Types

	replFrom := comp.LineCol{Line: line, Col: column}
	replTo := comp.LineCol{Line: line, Col: column}
	var filter string

	shouldBeReplaced := comp.ShouldBeReplaced(element)
	if found && shouldBeReplaced {
		// There are number of situations where we need to be smarter than default behavior
		pos := ast.GetPosition(element)
		smartReplaceRange := comp.CalculateSmartReplaceRange(element, pos.GetStart(), pos.GetStop())
		if smartReplaceRange.FilterText != "" {
			replFrom.Line, replFrom.Col = es.posConv.ToRelative(smartReplaceRange.Start)
			replTo.Line, replTo.Col = es.posConv.ToRelative(smartReplaceRange.Stop + 1)

			if smartReplaceRange.FilterText != "" {
				filter = smartReplaceRange.FilterText
			}
		} else {
			replFrom.Line, replFrom.Col = es.posConv.ToRelative(pos.GetStart())
			replTo.Line, replTo.Col = es.posConv.ToRelative(pos.GetStop() + 1)
		}
	}

	if filter == "" {
		if doFilter && found && shouldBeReplaced {
			switch ctx := element.(type) {
			case *antlr.BaseParserRuleContext:
				filter = ctx.GetText()
			case *parser.FunctionInvocationBodyContext:
				filter = ctx.GetText()
			case *parser.NodeLabelContext:
				filter = ctx.GetText()[1:]
			case *parser.RelationshipTypeContext:
				filter = ctx.GetText()
			case *antlr.TerminalNodeImpl:
				filter = ctx.GetText()
			case *parser.VariableContext:
				filter = ctx.GetText()
			case *antlr.ErrorNodeImpl:
				return comp.Result{}
			default:
				panic(reflect.TypeOf(element))
			}
		}
	}

	items := es.completion.GetItems(types, query, filter)
	return comp.Result{
		Items: items,
		Range: comp.Range{From: replFrom, To: replTo},
	}
}

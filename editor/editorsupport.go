// Copyright 2022 The persephone authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package editor

import (
	"reflect"

	"github.com/abc-inc/persephone/ast"
	"github.com/abc-inc/persephone/comp"
	"github.com/abc-inc/persephone/graph"
	"github.com/abc-inc/persephone/lang"
	"github.com/abc-inc/persephone/parser"
	"github.com/abc-inc/persephone/ref"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type EditorSupport struct {
	schema  graph.Schema
	Input   string
	posConv PosConv

	ParseTree           antlr.ParseTree
	ParseErrors         []SynErr
	referencesProviders map[string]ref.Provider
	completion          comp.AutoCompletion
	statements          []parser.CypherPartContext
}

func NewEditorSupport(input string) *EditorSupport {
	e := &EditorSupport{}
	e.completion = *comp.NewAutoCompletion(graph.Schema{})
	e.Update(input)
	return e
}

func (es *EditorSupport) Update(input string) {
	es.posConv = *NewPosConv(input)

	es.Input = input
	parseTree, referencesListener, errorListener, referencesProviders := Parse(input)
	es.ParseTree = parseTree

	es.ParseErrors = errorListener.errors

	es.statements = referencesListener.Statements
	es.referencesProviders = referencesProviders

	es.completion.UpdateReferenceProviders(es.referencesProviders)
}

func (es *EditorSupport) SetSchema(schema graph.Schema) {
	es.schema = schema
	es.completion.UpdateSchema(es.schema)
}

func (es EditorSupport) GetElement(line, column int) antlr.Tree {
	abs := es.posConv.ToAbsolute(line, column)
	return getElement(es.ParseTree, abs)
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

func (es EditorSupport) GetCompletionInfo(line, column int) comp.Info {
	element := es.GetElementForCompletion(line, column)
	query := ast.FindAnyParent(element, []string{lang.QUERY_CONTEXT})
	info := comp.GetTypes(element)
	return comp.Info{
		Element: element,
		Query:   query,
		Found:   info.Found,
		Types:   info.Types,
	}
}

func (es EditorSupport) GetElementForCompletion(line, column int) antlr.Tree {
	e := es.GetElement(line, column)
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
		if smartReplaceRange != nil {
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
			if e, ok := element.(interface{ GetText() string }); ok {
				filter = e.GetText()
			} else {
				panic(reflect.TypeOf(element).String() + " does not have text")
			}
		}
	}

	items := es.completion.GetItems(types, query, filter)
	return comp.Result{
		Items: items,
		Range: comp.Range{From: replFrom, To: replTo},
	}
}

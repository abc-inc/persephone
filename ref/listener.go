package ref

import (
	"github.com/abc-inc/persephone/lang"
	"github.com/abc-inc/persephone/parser"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type RefListener struct {
	Queries            []parser.CypherQueryContext
	QueriesAndCommands []antlr.ParserRuleContext
	Statements         []parser.CypherPartContext
	Raw                []antlr.ParserRuleContext
	Indexes            map[string]*Index
	InConsoleCommand   bool
	parser.BaseCypherListener
}

func NewRefListener() *RefListener {
	is := map[string]*Index{}
	for _, ctx := range lang.SymbolicContexts {
		is[ctx] = NewIndex()
	}
	return &RefListener{
		Indexes: is,
	}
}

func (l *RefListener) EnterRaw(ctx antlr.ParserRuleContext) {
	l.Raw = append(l.Raw, ctx)
}

func (l *RefListener) ExitRaw(ctx antlr.ParserRuleContext) {
	if len(l.Raw) == 0 {
		l.Raw = append(l.Raw, ctx)
	}
}

func (l *RefListener) EnterCypherPart(ctx *parser.CypherPartContext) {
	l.Statements = append(l.Statements, *ctx)
}

func (l *RefListener) ExitCypherPart(ctx *parser.CypherPartContext) {
	if len(l.Statements) == 0 {
		l.Statements = append(l.Statements, *ctx)
	}
}

func (l *RefListener) EnterCypherConsoleCommand(ctx *parser.CypherConsoleCommandContext) {
	l.QueriesAndCommands = append(l.QueriesAndCommands, ctx)
	for _, i := range l.Indexes {
		i.AddQuery()
	}
	l.InConsoleCommand = true
}

func (l *RefListener) ExitCypherConsoleCommand(ctx *parser.CypherConsoleCommandContext) {
	l.InConsoleCommand = false
}

func (l *RefListener) EnterCypherQuery(ctx *parser.CypherQueryContext) {
	l.Queries = append(l.Queries, *ctx)
	l.QueriesAndCommands = append(l.QueriesAndCommands, ctx)
	for _, i := range l.Indexes {
		i.AddQuery()
	}
}

func (l *RefListener) ExitVariable(ctx *parser.VariableContext) {
	if l.InConsoleCommand {
		return
	}
	l.Indexes[lang.VARIABLE_CONTEXT].AddVariable(ctx)
}

func (l *RefListener) ExitLabelName(ctx *parser.LabelNameContext) {
	if l.InConsoleCommand {
		return
	}
	l.Indexes[lang.LABEL_NAME_CONTEXT].Add(ctx, true)
}

func (l *RefListener) ExitRelTypeName(ctx *parser.RelTypeNameContext) {
	if l.InConsoleCommand {
		return
	}
	l.Indexes[lang.RELATIONSHIP_TYPE_NAME_CONTEXT].Add(ctx, true)
}

func (l *RefListener) ExitPropertyKeyName(ctx *parser.PropertyKeyNameContext) {
	if l.InConsoleCommand {
		return
	}
	l.Indexes[lang.PROPERTY_KEY_NAME_CONTEXT].Add(ctx, true)
}

func (l *RefListener) ExitParameterName(ctx *parser.ParameterNameContext) {
	if l.InConsoleCommand {
		return
	}
	l.Indexes[lang.PARAMETER_NAME_CONTEXT].Add(ctx, true)
}

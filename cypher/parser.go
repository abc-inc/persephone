package cypher

import (
	"github.com/abc-inc/merovingian/comp"
	"github.com/abc-inc/merovingian/lang"
	"github.com/abc-inc/merovingian/parser"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

func Parse(input string) (antlr.ParseTree, *comp.RefListener, *ErrorListener, map[string]comp.Provider) {
	refListener := comp.NewRefListener()
	errListener := &ErrorListener{}
	chars := antlr.NewInputStream(input)
	lexer := parser.NewCypherLexer(chars)
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(errListener)
	tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	p := parser.NewCypherParser(tokens)
	p.BuildParseTrees = true
	p.RemoveErrorListeners()
	p.AddErrorListener(errListener)
	p.AddParseListener(refListener)
	parseTree := p.Cypher()
	queries, indexes := refListener.Queries, refListener.Indexes
	refProvs := make(map[string]comp.Provider, len(lang.SymbolicContexts))
	for _, ctx := range lang.SymbolicContexts {
		refProvs[ctx] = *comp.NewProvider(queries, indexes[ctx])
	}
	return parseTree, refListener, errListener, refProvs
}
package rule

import (
	"reflect"

	"github.com/abc-inc/merovingian/ast"
	"github.com/abc-inc/merovingian/comp"
	"github.com/abc-inc/merovingian/parser"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

func ruleLiteralEntry(e antlr.ParseTree) (is []Info) {
	literalEntry := ast.FindParent(e, reflect.TypeOf(parser.LiteralEntryContext{}))
	if literalEntry == nil {
		return nil
	}
	if literalEntry.GetChildCount() < 2 {
		return is
	}

	doubleDots := literalEntry.GetChild(1)
	var space antlr.Tree
	if literalEntry.GetChildCount() > 2 {
		space = literalEntry.GetChild(2)
	}
	if doubleDots == e || space == e {
		for _, t := range comp.All {
			is = append(is, Info{Type: t.Type})
		}
	}
	return is
}

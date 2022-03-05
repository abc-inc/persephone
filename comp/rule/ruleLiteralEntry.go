package rule

import (
	"reflect"

	"github.com/abc-inc/merovingian/ast"
	"github.com/abc-inc/merovingian/parser"
	"github.com/abc-inc/merovingian/types"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

func ruleLiteralEntry(e antlr.ParseTree) (is []Info) {
	literalEntry := ast.FindParent(e, reflect.TypeOf(parser.LiteralEntryContext{}))
	if literalEntry == nil {
		return nil
	}
	doubleDots := literalEntry.GetChild(1)
	space := literalEntry.GetChild(2)
	if doubleDots == e || space == e {
		for _, t := range types.All {
			is = append(is, Info{Type: t})
		}
	}
	return is
}

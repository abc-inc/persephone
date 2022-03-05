package rule

import (
	"github.com/abc-inc/merovingian/parser"
	"github.com/abc-inc/merovingian/types"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

func rulePropertyLookup(e antlr.ParseTree) []Info {
	if lookupCtx := e.GetParent(); lookupCtx != nil {
		if _, ok := lookupCtx.(*parser.PropertyLookupContext); ok && e.GetText() == "." {
			return []Info{{Type: types.PropertyKey}}
		}
	}
	return nil
}

package rule

import (
	"github.com/abc-inc/persephone/parser"
	"github.com/abc-inc/persephone/types"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

func rulePropertyLookup(e antlr.ParseTree) []Info {
	if parentCtx := e.GetParent(); parentCtx != nil {
		if lookupCtx := parentCtx.GetParent(); lookupCtx != nil {
			if _, ok := lookupCtx.(*parser.PropertyLookupContext); ok && e.GetText() == "." {
				return []Info{{Type: types.PropertyKey}}
			}
			// TODO: why is this necessary? the JavaScript implementation does not need it
			if _, ok := lookupCtx.(*parser.PropertyExpressionContext); ok {
				return []Info{{Type: types.PropertyKey}}
			}
			// TODO: why is this necessary? the JavaScript implementation does not need it
			if _, ok := lookupCtx.(*parser.PropertyOrLabelsExpressionContext); ok {
				return []Info{{Type: types.PropertyKey}}
			}
		}
	}
	return nil
}

package rule

import (
	"github.com/abc-inc/persephone/ast"
	"github.com/abc-inc/persephone/parser"
	"github.com/abc-inc/persephone/types"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// If we are in call rule, and element is second child of call return procedure types
func ruleCallClauseBeginning(e antlr.ParseTree) []Info {
	parent := ast.GetParent(e)
	if parent == nil {
		return nil
	}

	if _, ok := parent.(*parser.CallContext); ok {
		if parent.GetChildCount() > 1 && parent.GetChild(1) == e {
			return []Info{{Type: types.ProcedureName}}
		}
	}
	return nil
}

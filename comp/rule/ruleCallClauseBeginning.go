package rule

import (
	"github.com/abc-inc/merovingian/ast"
	"github.com/abc-inc/merovingian/parser"
	"github.com/abc-inc/merovingian/types"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// If we are in call rule, and element is second child of call return procedure types
func ruleCallClauseBeginning(e antlr.ParseTree) []Info {
	parent := ast.GetParent(e)
	if parent == nil {
		return nil
	}

	if _, ok := parent.(*parser.CallContext); ok {
		secondChild := parent.GetChild(1)
		if secondChild == e {
			return []Info{{Type: types.ProcedureName}}
		}
	}
	return nil
}

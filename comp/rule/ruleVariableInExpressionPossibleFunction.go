package rule

import (
	"github.com/abc-inc/merovingian/ast"
	"github.com/abc-inc/merovingian/lang"
	"github.com/abc-inc/merovingian/types"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

func ruleVariableInExpressionPossibleFunction(e antlr.ParseTree) []Info {
	variable := ast.FindAnyParent(e, []string{lang.VARIABLE_CONTEXT})
	expression := ast.FindAnyParent(variable, []string{lang.EXPRESSION_CONTEXT})
	if variable != nil && expression != nil {
		return []Info{{Type: types.Variable}, {Type: types.FunctionName}}
	}
	return nil
}

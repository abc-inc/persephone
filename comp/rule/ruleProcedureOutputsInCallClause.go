package rule

import (
	"github.com/abc-inc/persephone/ast"
	"github.com/abc-inc/persephone/lang"
	"github.com/abc-inc/persephone/parser"
	"github.com/abc-inc/persephone/types"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Return procedure output completion if we are inside procedure
func ruleProcedureOutputsInCallClause(e antlr.ParseTree) []Info {
	call := ast.FindAnyParent(e, []string{lang.CALL_CONTEXT})
	if call == nil {
		return nil
	}

	proc := ast.FindChild(call, lang.PROCEDURE_NAME_CONTEXT)
	resOutput := ast.FindAnyParent(e, []string{lang.PROCEDURE_RESULTS_CONTEXT})
	if proc == nil || resOutput == nil {
		return nil
	}
	return []Info{{
		Type: types.ProcedureOutput,
		Path: []string{proc.(*parser.ProcedureInvocationBodyContext).GetText()},
	}}
}

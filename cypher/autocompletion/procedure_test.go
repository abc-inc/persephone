package autocompletion

import (
	"testing"

	"github.com/abc-inc/merovingian/types"
)

func TestReturnProcedureNameType(t *testing.T) {
	checkCompletionTypes(t, "call ▼db.proc()", true, []types.Type{types.ProcedureName})
}

func TestReturnProcedureNameTypeIfOnlyCallPresent(t *testing.T) {
	checkCompletionTypes(t, "call▼", true, []types.Type{types.ProcedureName})
}

func TestReturnOutputAtTheBeginningOfYield(t *testing.T) {
	checkCompletionTypes(t, "call db.proc() yield▼ ", true, []types.Type{types.ProcedureOutput})
}

func TestReturnOutputBeforeTheFirstTypedSymbol(t *testing.T) {
	checkCompletionTypes(t, "call db.proc() yield ▼a", true, []types.Type{types.ProcedureOutput})
}

func TestReturnOutputAtTheBeginningOfSecondOutput(t *testing.T) {
	checkCompletionTypes(t, "call db.proc() yield a as b,▼ ", true, []types.Type{types.ProcedureOutput})
}

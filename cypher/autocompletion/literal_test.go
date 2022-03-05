package autocompletion

import (
	"testing"

	"github.com/abc-inc/merovingian/comp"
	"github.com/abc-inc/merovingian/types"
)

func TestLiteralStringTypesYieldsNothingType(t *testing.T) {
	checkCompletionTypes(t, `RETURN "▼"`, true, []types.Type{types.Noop})
	checkCompletionTypes(t, `RETURN '▼'`, true, []types.Type{types.Noop})
}

func TestLiteralStringWithFiltersYieldsNoAcInString(t *testing.T) {
	expected := comp.Result{
		Items: nil,
		Range: comp.Range{
			From: comp.LineCol{Line: 1, Col:  7},
			To: comp.LineCol{Line: 1, Col:  10},
		},
	}
	checkCompletion(t, `RETURN ":▼"`, expected, true)
	checkCompletion(t, `RETURN ':▼'`, expected, true)
}

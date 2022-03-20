package comp_test

import (
	"testing"

	. "github.com/abc-inc/persephone/comp"
	"github.com/abc-inc/persephone/types"
)

func TestPropertyKeyTypesYieldsPropertyKeyWithNoFirstCharInsideMapWithinExpression(t *testing.T) {
	checkCompletionTypes(t, "MATCH ()-[▼{}]-()", true, []types.Type{types.PropertyKey})
}

func TestPropertyKeyTypesYieldsPropertyKeyWithNoFirstCharInsideMapWithinExpressionWithFirstSymbol(t *testing.T) {
	checkCompletionTypes(t, "WITH $param MATCH ()-[{▼p}]-()", true, []types.Type{types.PropertyKey, types.Parameter})
}

func TestPropertyKeyTypesYieldsPropertyKeyWithNoFirstCharInsideMapWithinExpressionWithoutClosingCurlyBrace(t *testing.T) {
	checkCompletionTypes(t, "WITH $param MATCH ()-[▼{", true, []types.Type{types.PropertyKey, types.Parameter})
}

func TestPropertyKeyTypesYieldsPropertyKeyWithFirstCharInsideMapWithinExpressionWithoutClosingCurlyBrace(t *testing.T) {
	checkCompletionTypes(t, "WITH $param MATCH ()-[{▼p", true, []types.Type{types.PropertyKey, types.Parameter})
}

func TestPropertyKeyAutoCompletionYieldsPropertyKeyListInsideMapWithinExpression(t *testing.T) {
	expected := Result{
		Items: []Item{
			{Type: types.PropertyKey, View: "prop1", Content: "prop1"},
			{Type: types.PropertyKey, View: "prop2", Content: "prop2"},
		},
		Range: Range{
			From: LineCol{Line: 1, Col: 11},
			To:   LineCol{Line: 1, Col: 11},
		},
	}
	checkCompletion(t, "MATCH ()-[{▼}]-()", expected, true)
}

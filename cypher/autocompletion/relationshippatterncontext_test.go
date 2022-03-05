package autocompletion

import (
	"testing"

	"github.com/abc-inc/merovingian/comp"
	"github.com/abc-inc/merovingian/types"
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
	expected := comp.Result{
		Items: []comp.Item{
			{Type: types.PropertyKey, View: "prop1", Content: "prop1"},
			{Type: types.PropertyKey, View: "prop2", Content: "prop2"},
		},
		Range: comp.Range{
			From: comp.LineCol{Line: 1, Col: 11},
			To:   comp.LineCol{Line: 1, Col: 11},
		},
	}
	checkCompletion(t, "MATCH ()-[{▼}]-()", expected, true)
}

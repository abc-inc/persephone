package autocompletion

import (
	"testing"

	"github.com/abc-inc/merovingian/comp"
	"github.com/abc-inc/merovingian/types"
)

func TestMixedTypesYieldsAnyAtTheBeginningOfAQuery(t *testing.T) {
	checkCompletionTypes(t, "▼", false, comp.AllTypes)
}

func TestMixedTypesYieldsFunctionNameAndVariableInExpression(t *testing.T) {
	checkCompletionTypes(t, "return ▼fun", true, []types.Type{types.Variable, types.FunctionName})
}

func TestMixedWithoutFiltersYieldsFunctionNameAndVariableListInExpression(t *testing.T) {
	expected := comp.Result{
		Items: []comp.Item{
			{Type: types.Variable, View: "fun", Content: "fun"},
			{Type: types.FunctionName, View: "toFloat", Content: "toFloat", Postfix: "expression"},
			{Type: types.FunctionName, View: "head", Content: "head", Postfix: "expression"},
		},
		Range: comp.Range{
			From: comp.LineCol{Line: 1, Col: 19},
			To:   comp.LineCol{Line: 1, Col: 22},
		},
	}
	checkCompletion(t, "match (fun) return ▼fun", expected, false)
	checkCompletion(t, "match (fun) return fun▼", expected, false)
}

func TestMixedWithoutFiltersYieldsOnlyKeywordsAtTheStartOfAQuery(t *testing.T) {
	expected := comp.Result{
		Items: []comp.Item{
			{Type: types.Parameter, View: "param1", Content: "param1"},
			{Type: types.Parameter, View: "param2", Content: "param2"},
			{Type: types.PropertyKey, View: "prop1", Content: "prop1"},
			{Type: types.PropertyKey, View: "prop2", Content: "prop2"},
			{Type: types.FunctionName, View: "toFloat", Content: "toFloat", Postfix: "expression"},
			{Type: types.FunctionName, View: "head", Content: "head", Postfix: "expression"},
		},
		Range: comp.Range{
			From: comp.LineCol{Line: 1, Col: 0},
			To:   comp.LineCol{Line: 1, Col: 0},
		},
	}
	expected.Items = append(expected.Items, comp.KEYWORD_ITEMS...)

	checkCompletion(t, "▼", expected, false)
}

func TestMixedWithFiltersYieldsFunctionNameAndVariableListInExpression(t *testing.T) {
	expected := comp.Result{
		Items: []comp.Item{
			{Type: types.Variable, View: "atern", Content: "atern"},
			{Type: types.FunctionName, View: "toFloat", Content: "toFloat", Postfix: "expression"},
		},
		Range: comp.Range{
			From: comp.LineCol{Line: 1, Col: 21},
			To:   comp.LineCol{Line: 1, Col: 23},
		},
	}
	checkCompletion(t, "MATCH (atern) RETURN at▼", expected, true)
	checkCompletion(t, "MATCH (atern) RETURN a▼t", expected, true)
	checkCompletion(t, "MATCH (atern) RETURN ▼at", expected, true)
}

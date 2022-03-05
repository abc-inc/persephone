package autocompletion

import (
	"testing"

	"github.com/abc-inc/merovingian/comp"
	"github.com/abc-inc/merovingian/types"
)

func TestTypesYieldsFunctionName(t *testing.T) {
	checkCompletionTypes(t, "return ▼fun()", true, []types.Type{types.FunctionName})
}

func TestWithoutFiltersYieldsFunctionNameList(t *testing.T) {
	expectedItems := comp.Result{
		Items: []comp.Item{
			{Type: types.FunctionName, View: "toFloat", Content: "toFloat", Postfix: "expression"},
			{Type: types.FunctionName, View: "head", Content: "head", Postfix: "expression"},
		},
		Range: comp.Range{
			From: comp.LineCol{Line: 1, Col: 7},
			To:   comp.LineCol{Line: 1, Col: 10},
		},
	}
	checkCompletion(t, "return ▼fun()", expectedItems, false)
	checkCompletion(t, "return fun▼()", expectedItems, false)
}

func TestWithoutFiltersYieldsLongFunctionNameList(t *testing.T) {
	expectedItems := comp.Result{
		Items: []comp.Item{
			{Type: types.FunctionName, View: "toFloat", Content: "toFloat", Postfix: "expression"},
			{Type: types.FunctionName, View: "head", Content: "head", Postfix: "expression"},
		},
		Range: comp.Range{
			From: comp.LineCol{Line: 1, Col: 7},
			To:   comp.LineCol{Line: 1, Col: 21},
		},
	}
	checkCompletion(t, "return ▼name.space.fun()", expectedItems, false)
	checkCompletion(t, "return name.space.fun▼()", expectedItems, false)
}

func TestWithFiltersYieldsFunctionNameList(t *testing.T) {
	expectedItems := comp.Result{
		Items: []comp.Item{
			{Type: types.FunctionName, View: "toFloat", Content: "toFloat", Postfix: "expression"},
			{Type: types.FunctionName, View: "head", Content: "head", Postfix: "expression"},
		},
		Range: comp.Range{
			From: comp.LineCol{Line: 1, Col: 7},
			To:   comp.LineCol{Line: 1, Col: 9},
		},
	}
	checkCompletion(t, "return ▼name.space.fun()", expectedItems, false)
	checkCompletion(t, "return name.space.fun▼()", expectedItems, false)
}

package comp_test

import (
	"testing"

	. "github.com/abc-inc/persephone/comp"
	"github.com/abc-inc/persephone/types"
)

func TestTypesYieldsFunctionName(t *testing.T) {
	checkCompletionTypes(t, "return ▼fun()", true, []types.Type{types.FunctionName})
}

func TestWithoutFiltersYieldsFunctionNameList(t *testing.T) {
	expectedItems := Result{
		Items: []Item{
			{Type: types.FunctionName, View: "toFloat", Content: "toFloat", Postfix: "expression"},
			{Type: types.FunctionName, View: "head", Content: "head", Postfix: "expression"},
		},
		Range: Range{
			From: LineCol{Line: 1, Col: 7},
			To:   LineCol{Line: 1, Col: 10},
		},
	}
	checkCompletion(t, "return ▼fun()", expectedItems, false)
	checkCompletion(t, "return fun▼()", expectedItems, false)
}

func TestWithoutFiltersYieldsLongFunctionNameList(t *testing.T) {
	expectedItems := Result{
		Items: []Item{
			{Type: types.FunctionName, View: "toFloat", Content: "toFloat", Postfix: "expression"},
			{Type: types.FunctionName, View: "head", Content: "head", Postfix: "expression"},
		},
		Range: Range{
			From: LineCol{Line: 1, Col: 7},
			To:   LineCol{Line: 1, Col: 21},
		},
	}
	checkCompletion(t, "return ▼name.space.fun()", expectedItems, false)
	checkCompletion(t, "return name.space.fun▼()", expectedItems, false)
}

func TestWithFiltersYieldsFunctionNameList(t *testing.T) {
	expectedItems := Result{
		Items: []Item{
			{Type: types.FunctionName, View: "head", Content: "head", Postfix: "expression"},
		},
		Range: Range{
			From: LineCol{Line: 1, Col: 7},
			To:   LineCol{Line: 1, Col: 9},
		},
	}
	checkCompletion(t, "return he▼()", expectedItems, true)
	checkCompletion(t, "return h▼e()", expectedItems, true)
	checkCompletion(t, "return ▼he()", expectedItems, true)
}

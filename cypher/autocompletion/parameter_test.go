package autocompletion

import (
	"testing"

	"github.com/abc-inc/merovingian/comp"
	"github.com/abc-inc/merovingian/types"
)

func TestParameterTypesYieldsParameterName(t *testing.T) {
	checkCompletionTypes(t, "RETURN $▼a", true, []types.Type{types.Parameter})
}

func TestParameterTypesYieldsParameterNameWOFirstCharTyped(t *testing.T) {
	checkCompletionTypes(t, "RETURN ▼$", true, []types.Type{types.Parameter})
}

func TestParameterTypesYieldsParameterNameWithFirstCharTypedAndBothCurlyBraces(t *testing.T) {
	checkCompletionTypes(t, "RETURN {▼p}", true, []types.Type{types.Parameter})
}

func TestParameterWithoutFiltersYieldsParameterNameList(t *testing.T) {
	expected := comp.Result{
		Items: []comp.Item{
			{Type: types.Parameter, View: "param1", Content: "param1"},
			{Type: types.Parameter, View: "param2", Content: "param2"},
		},
		Range: comp.Range{
			From: comp.LineCol{Line: 1, Col: 14},
			To:   comp.LineCol{Line: 1, Col: 15},
		},
	}
	checkCompletion(t, "RETURN {b} + $▼a", expected, false)
	checkCompletion(t, "RETURN {b} + $a▼", expected, false)
}

func TestParameterWithoutFiltersYieldsParameterNameListAfterFirstSymbol(t *testing.T) {
	expected := comp.Result{
		Items: []comp.Item{
			{Type: types.Parameter, View: "param1", Content: "param1"},
			{Type: types.Parameter, View: "param2", Content: "param2"},
		},
		Range: comp.Range{
			From: comp.LineCol{Line: 1, Col: 13},
			To:   comp.LineCol{Line: 1, Col: 14},
		},
	}
	checkCompletion(t, "RETURN $b + {▼a}", expected, false)
	checkCompletion(t, "RETURN $b + {a▼}", expected, false)
}

func TestWithFiltersYieldsParameterNameList(t *testing.T) {
	expected := comp.Result{
		Items: []comp.Item{
			{Type: types.Parameter, View: "param1", Content: "param1"},
			{Type: types.Parameter, View: "param2", Content: "param2"},
		},
		Range: comp.Range{
			From: comp.LineCol{Line: 1, Col: 14},
			To:   comp.LineCol{Line: 1, Col: 15},
		},
	}
	checkCompletion(t, "RETURN {b} + $▼a", expected, true)
	checkCompletion(t, "RETURN {b} + $a▼", expected, true)
}

func TestWithFiltersYieldsLegacyParameterNameList(t *testing.T) {
	expected := comp.Result{
		Items: []comp.Item{
			{Type: types.Parameter, View: "param1", Content: "param1"},
			{Type: types.Parameter, View: "param2", Content: "param2"},
		},
		Range: comp.Range{
			From: comp.LineCol{Line: 1, Col: 8},
			To:   comp.LineCol{Line: 1, Col: 9},
		},
	}
	checkCompletion(t, "RETURN {p▼}", expected, true)
}

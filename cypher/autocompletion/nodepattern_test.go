package autocompletion

import (
	"testing"

	"github.com/abc-inc/merovingian/comp"
	"github.com/abc-inc/merovingian/types"
)

func TestTypesYieldsLabelIfCharacterPresent(t *testing.T) {
	checkCompletionTypes(t, "MATCH (n) MATCH (a:▼b", true, []types.Type{types.Label})
}

func TestTypesYieldsLabelIfOnlyColonPresent(t *testing.T) {
	checkCompletionTypes(t, "MATCH (n) MATCH (a▼:", true, []types.Type{types.Label})
}

func TestTypesYieldsLabelIfOnlyColonParenthesisPresent(t *testing.T) {
	checkCompletionTypes(t, "MATCH (n) MATCH (a▼:)", true, []types.Type{types.Label})
}

func TestTypesYieldsLabelAndVariableIfBeginningOfNodePattern(t *testing.T) {
	checkCompletionTypes(t, "MATCH (n) MATCH ▼(", true, []types.Type{types.Variable, types.Label})
}

func TestTypesYieldsLabelAndVariableIfBeginningOfNodePatternInChain(t *testing.T) {
	checkCompletionTypes(t, "MATCH (n) MATCH ()--()--▼(", true, []types.Type{types.Variable, types.Label})
}

func TestTypesYieldsLabelTypeIfOnlyColonIsPresent(t *testing.T) {
	checkCompletionTypes(t, "MATCH (n) MATCH (▼:", true, []types.Type{types.Label})
}

func TestTypesYieldsLabelTypeWhenMultipleLabels(t *testing.T) {
	checkCompletionTypes(t, "MATCH (:SomeLabel▼: ", true, []types.Type{types.Label})
}

func TestWithoutFiltersYieldsLabelList(t *testing.T) {
	expected := comp.Result{
		Items: []comp.Item{
			{Type: types.Label, View: ":y", Content: ":y"},
			{Type: types.Label, View: ":x", Content: ":x"},
		},
		Range: comp.Range{
			From: comp.LineCol{Line: 1, Col: 18},
			To:   comp.LineCol{Line: 1, Col: 20},
		},
	}
	checkCompletion(t, "MATCH (n) MATCH (a:▼b", expected, false)
	checkCompletion(t, "MATCH (n) MATCH (a:b▼", expected, false)
}

func TestWithFiltersYieldsLabelList(t *testing.T) {
	expected := comp.Result{
		Items: []comp.Item{
			{Type: types.Label, View: ":y", Content: ":y"},
		},
		Range: comp.Range{
			From: comp.LineCol{Line: 1, Col: 18},
			To:   comp.LineCol{Line: 1, Col: 20},
		},
	}
	checkCompletion(t, "MATCH (n) MATCH (a:▼y", expected, true)
	checkCompletion(t, "MATCH (n) MATCH (a:y▼", expected, true)
}

func TestWithoutFiltersYieldsLabelListIfOnlyColonIsPresent(t *testing.T) {
	expected := comp.Result{
		Items: []comp.Item{
			{Type: types.Label, View: ":y", Content: ":y"},
			{Type: types.Label, View: ":x", Content: ":x"},
		},
		Range: comp.Range{
			From: comp.LineCol{Line: 1, Col: 18},
			To:   comp.LineCol{Line: 1, Col: 19},
		},
	}
	checkCompletion(t, "MATCH (n) MATCH (a:▼", expected, false)
	checkCompletion(t, "MATCH (n) MATCH (a:▼", expected, true)
}

func TestWithoutFiltersYieldsLabelListIfOnlyColonParenthesisIsPresent(t *testing.T) {
	expected := comp.Result{
		Items: []comp.Item{
			{Type: types.Label, View: ":y", Content: ":y"},
			{Type: types.Label, View: ":x", Content: ":x"},
		},
		Range: comp.Range{
			From: comp.LineCol{Line: 1, Col: 8},
			To:   comp.LineCol{Line: 1, Col: 9},
		},
	}
	checkCompletion(t, "MATCH (a:▼) MATCH ()", expected, false)
	checkCompletion(t, "MATCH (a:▼) MATCH ()", expected, true)
}

func TestWithoutFiltersYieldsLabelAndVariablesAtTheBeginningOfNodePattern(t *testing.T) {
	expected := comp.Result{
		Items: []comp.Item{
			{Type: types.Variable, View: "a", Content: "a"},
			{Type: types.Label, View: ":y", Content: ":y"},
			{Type: types.Label, View: ":x", Content: ":x"},
		},
		Range: comp.Range{
			From: comp.LineCol{Line: 1, Col: 17},
			To:   comp.LineCol{Line: 1, Col: 17},
		},
	}
	checkCompletion(t, "MATCH (a) MATCH (▼", expected, false)
	checkCompletion(t, "MATCH (a) MATCH (▼", expected, true)
}

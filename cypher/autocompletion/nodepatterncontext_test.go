package autocompletion

import (
	"testing"

	"github.com/abc-inc/merovingian/comp"
	"github.com/abc-inc/merovingian/types"
)

func TestPropertyKeyTypesYieldsPropertyKeyWithNoFirstCharTypedInsideMapLiteral(t *testing.T) {
	checkCompletionTypes(t, "MATCH (▼{});", true, []types.Type{types.PropertyKey})
}

func TestPropertyKeyTypesYieldsAllIfOnClosingCurlyBrace(t *testing.T) {
	checkCompletionTypes(t, "MATCH ({▼});", false, comp.AllTypes)
}

func TestPropertyKeyTypesYieldsPropertyKeyAndParameterIfInPropertiesContext(t *testing.T) {
	checkCompletionTypes(t, "WITH $someParam MATCH ({som▼e })", true, []types.Type{types.PropertyKey, types.Parameter})
}

func TestPropertyKeyTypesYieldsPropertyKeyOrParameterWithNoClosingCurlyBrace(t *testing.T) {
	checkCompletionTypes(t, "WITH $param MATCH (n ▼{", true, []types.Type{types.PropertyKey, types.Parameter})
}

func TestPropertyKeyTypesYieldsPropertyKeyWithFirstCharTypedInsideMapLiteralWOClosingCurlyBrace(t *testing.T) {
	checkCompletionTypes(t, "WITH $param MATCH (n {▼p", true, []types.Type{types.PropertyKey, types.Parameter})
}

func TestPropertyKeyTypesYieldsPropertyKeyWithFirstCharTypedInsideMapLiteralWithClosingCurlyBrace(t *testing.T) {
	checkCompletionTypes(t, "WITH $param MATCH (n {p▼})", false, comp.AllTypes)
}

func TestPropertyKeyTypesYieldsAllBeforeColon(t *testing.T) {
	checkCompletionTypes(t, "MATCH (n {key▼:})", true, comp.AllTypes)
}

func TestPropertyKeyTypesYieldsAllBeforeColonUnclosedMapLiteral(t *testing.T) {
	checkCompletionTypes(t, "MATCH (n {key▼:", true, comp.AllTypes)
}

func TestPropertyKeyTypesYieldsAllAfterColon(t *testing.T) {
	checkCompletionTypes(t, "MATCH (n {key:▼ })", true, comp.AllTypes)
}

func TestPropertyKeyTypesYieldsPropertyKeyWithNoFirstCharTypedAfterKeyAndWithClosingCurlyBrace(t *testing.T) {
	checkCompletionTypes(t, "MATCH (n {key: 1, ▼});", false, comp.AllTypes)
}

func TestPropertyKeyTypesYieldsPropertyKeyAfterKeyAndFirstCharTypedWOClosingCurlyBrace(t *testing.T) {
	checkCompletionTypes(t, "MATCH (n {key: 1, ▼k", true, []types.Type{types.PropertyKey})
}

func TestPropertyKeyTypesYieldsPropertyKeyAfterKeyAndFirstCharTypedWithClosingCurlyBrace(t *testing.T) {
	checkCompletionTypes(t, "MATCH (n {key: 1, ▼k});", true, []types.Type{types.PropertyKey})
}

func TestPropertyKeyAutoCompletionYieldsPropertyKeyListFromWithinMap(t *testing.T) {
	expected := comp.Result{
		Items: []comp.Item{
			{Type: types.PropertyKey, View: "prop1", Content: "prop1"},
			{Type: types.PropertyKey, View: "prop2", Content: "prop2"},
		},
		Range: comp.Range{
			From: comp.LineCol{Line: 1, Col: 8},
			To:   comp.LineCol{Line: 1, Col: 8},
		},
	}
	checkCompletion(t, "MATCH ({▼});", expected, true)
}

func TestPropertyKeyAutoCompletionYieldsPropertyKeysAndParamInPropertiesContext(t *testing.T) {
	expected := comp.Result{
		Items: []comp.Item{
			{Type: types.PropertyKey, View: "prop1", Content: "prop1"},
			{Type: types.Parameter, View: "param1", Content: "param1"},
		},
		Range: comp.Range{
			From: comp.LineCol{Line: 1, Col: 8},
			To:   comp.LineCol{Line: 1, Col: 10},
		},
	}
	checkCompletion(t, "MATCH ({p1▼ });", expected, true)
}

func TestPropertyKeyAutoCompletionYieldsPropertyKeyListInMapLiteralWOClosingCurlyBrace(t *testing.T) {
	expected := comp.Result{
		Items: []comp.Item{
			{Type: types.PropertyKey, View: "prop1", Content: "prop1"},
			{Type: types.PropertyKey, View: "prop2", Content: "prop2"},
			{Type: types.Parameter, View: "param1", Content: "param1"},
			{Type: types.Parameter, View: "param2", Content: "param2"},
		},
		Range: comp.Range{
			From: comp.LineCol{Line: 1, Col: 10},
			To:   comp.LineCol{Line: 1, Col: 10},
		},
	}
	checkCompletion(t, "MATCH (n {▼", expected, true)
}

func TestPropertyKeyAutoCompletionYieldsAllAfterColonInUnclosedMapLiteral(t *testing.T) {
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
			From: comp.LineCol{Line: 1, Col: 12},
			To:   comp.LineCol{Line: 1, Col: 12},
		},
	}
	expected.Items = append(expected.Items, comp.KEYWORD_ITEMS...)

	checkCompletion(t, "MATCH (n {▼", expected, true)
}

func TestPropertyKeyAutoCompletionYieldsAllAfterColonInMapLiteral(t *testing.T) {
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
			From: comp.LineCol{Line: 1, Col: 12},
			To:   comp.LineCol{Line: 1, Col: 12},
		},
	}
	expected.Items = append(expected.Items, comp.KEYWORD_ITEMS...)

	checkCompletion(t, "MATCH ({key:▼ })", expected, true)
}

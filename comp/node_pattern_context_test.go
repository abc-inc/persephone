// Copyright 2022 The Persephone authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package comp_test

import (
	"testing"

	. "github.com/abc-inc/persephone/comp"
	"github.com/abc-inc/persephone/types"
)

func TestPropertyKeyTypesYieldsPropertyKeyWithNoFirstCharTypedInsideMapLiteral(t *testing.T) {
	checkCompletionTypes(t, "MATCH (▼{});", true, []types.Type{types.PropertyKey})
}

func TestPropertyKeyTypesYieldsAllIfOnClosingCurlyBrace(t *testing.T) {
	checkCompletionTypes(t, "MATCH ({▼});", false, types.AllComp)
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
	checkCompletionTypes(t, "WITH $param MATCH (n {p▼})", false, types.AllComp)
}

func TestPropertyKeyTypesYieldsAllBeforeColon(t *testing.T) {
	checkCompletionTypes(t, "MATCH (n {key▼:})", true, types.AllComp)
}

func TestPropertyKeyTypesYieldsAllBeforeColonUnclosedMapLiteral(t *testing.T) {
	checkCompletionTypes(t, "MATCH (n {key▼:", true, types.AllComp)
}

func TestPropertyKeyTypesYieldsAllAfterColon(t *testing.T) {
	checkCompletionTypes(t, "MATCH (n {key:▼ })", true, types.AllComp)
}

func TestPropertyKeyTypesYieldsPropertyKeyWithNoFirstCharTypedAfterKeyAndWithClosingCurlyBrace(t *testing.T) {
	checkCompletionTypes(t, "MATCH (n {key: 1, ▼});", false, types.AllComp)
}

func TestPropertyKeyTypesYieldsPropertyKeyAfterKeyAndFirstCharTypedWOClosingCurlyBrace(t *testing.T) {
	checkCompletionTypes(t, "MATCH (n {key: 1, ▼k", true, []types.Type{types.PropertyKey})
}

func TestPropertyKeyTypesYieldsPropertyKeyAfterKeyAndFirstCharTypedWithClosingCurlyBrace(t *testing.T) {
	checkCompletionTypes(t, "MATCH (n {key: 1, ▼k});", true, []types.Type{types.PropertyKey})
}

func TestPropertyKeyAutoCompletionYieldsPropertyKeyListFromWithinMap(t *testing.T) {
	expected := Result{
		Items: []Item{
			{Type: types.PropertyKey, View: "prop1", Content: "prop1"},
			{Type: types.PropertyKey, View: "prop2", Content: "prop2"},
		},
		Range: Range{
			From: LineCol{Line: 1, Col: 8},
			To:   LineCol{Line: 1, Col: 8},
		},
	}
	checkCompletion(t, "MATCH ({▼});", expected, true)
}

func TestPropertyKeyAutoCompletionYieldsPropertyKeysAndParamInPropertiesContext(t *testing.T) {
	expected := Result{
		Items: []Item{
			{Type: types.PropertyKey, View: "prop1", Content: "prop1"},
			{Type: types.Parameter, View: "param1", Content: "param1"},
		},
		Range: Range{
			From: LineCol{Line: 1, Col: 8},
			To:   LineCol{Line: 1, Col: 10},
		},
	}
	checkCompletion(t, "MATCH ({p1▼ });", expected, true)
}

func TestPropertyKeyAutoCompletionYieldsPropertyKeyListInMapLiteralWOClosingCurlyBrace(t *testing.T) {
	expected := Result{
		Items: []Item{
			{Type: types.PropertyKey, View: "prop1", Content: "prop1"},
			{Type: types.PropertyKey, View: "prop2", Content: "prop2"},
			{Type: types.Parameter, View: "param1", Content: "param1"},
			{Type: types.Parameter, View: "param2", Content: "param2"},
		},
		Range: Range{
			From: LineCol{Line: 1, Col: 10},
			To:   LineCol{Line: 1, Col: 10},
		},
	}
	checkCompletion(t, "MATCH (n {▼", expected, true)
}

func TestPropertyKeyAutoCompletionYieldsAllAfterColonInUnclosedMapLiteral(t *testing.T) {
	expected := Result{
		Items: []Item{
			{Type: types.Parameter, View: "param1", Content: "param1"},
			{Type: types.Parameter, View: "param2", Content: "param2"},
			{Type: types.PropertyKey, View: "prop1", Content: "prop1"},
			{Type: types.PropertyKey, View: "prop2", Content: "prop2"},
			{Type: types.FunctionName, View: "toFloat", Content: "toFloat", Postfix: "expression"},
			{Type: types.FunctionName, View: "head", Content: "head", Postfix: "expression"},
		},
		Range: Range{
			From: LineCol{Line: 1, Col: 12},
			To:   LineCol{Line: 1, Col: 12},
		},
	}
	expected.Items = append(expected.Items, KeywordItems()...)

	checkCompletion(t, "MATCH ({key:▼", expected, true)
}

func TestPropertyKeyAutoCompletionYieldsAllAfterColonInMapLiteral(t *testing.T) {
	expected := Result{
		Items: []Item{
			{Type: types.Parameter, View: "param1", Content: "param1"},
			{Type: types.Parameter, View: "param2", Content: "param2"},
			{Type: types.PropertyKey, View: "prop1", Content: "prop1"},
			{Type: types.PropertyKey, View: "prop2", Content: "prop2"},
			{Type: types.FunctionName, View: "toFloat", Content: "toFloat", Postfix: "expression"},
			{Type: types.FunctionName, View: "head", Content: "head", Postfix: "expression"},
		},
		Range: Range{
			From: LineCol{Line: 1, Col: 12},
			To:   LineCol{Line: 1, Col: 12},
		},
	}
	expected.Items = append(expected.Items, KeywordItems()...)

	checkCompletion(t, "MATCH ({key:▼ })", expected, true)
}

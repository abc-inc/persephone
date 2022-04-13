// Copyright 2022 The persephone authors
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

func TestMixedTypesYieldsAnyAtTheBeginningOfAQuery(t *testing.T) {
	checkCompletionTypes(t, "▼", false, types.AllComp)
}

func TestMixedTypesYieldsFunctionNameAndVariableInExpression(t *testing.T) {
	checkCompletionTypes(t, "return ▼fun", true, []types.Type{types.Variable, types.FunctionName})
}

func TestMixedWithoutFiltersYieldsFunctionNameAndVariableListInExpression(t *testing.T) {
	expected := Result{
		Items: []Item{
			{Type: types.Variable, View: "fun", Content: "fun"},
			{Type: types.FunctionName, View: "toFloat", Content: "toFloat", Postfix: "expression"},
			{Type: types.FunctionName, View: "head", Content: "head", Postfix: "expression"},
		},
		Range: Range{
			From: LineCol{Line: 1, Col: 19},
			To:   LineCol{Line: 1, Col: 22},
		},
	}
	checkCompletion(t, "match (fun) return ▼fun", expected, false)
	checkCompletion(t, "match (fun) return fun▼", expected, false)
}

func TestMixedWithoutFiltersYieldsOnlyKeywordsAtTheStartOfAQuery(t *testing.T) {
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
			From: LineCol{Line: 1, Col: 0},
			To:   LineCol{Line: 1, Col: 0},
		},
	}
	expected.Items = append(expected.Items, KeywordItems...)

	checkCompletion(t, "▼", expected, false)
}

func TestMixedWithFiltersYieldsFunctionNameAndVariableListInExpression(t *testing.T) {
	expected := Result{
		Items: []Item{
			{Type: types.Variable, View: "atern", Content: "atern"},
			{Type: types.FunctionName, View: "toFloat", Content: "toFloat", Postfix: "expression"},
		},
		Range: Range{
			From: LineCol{Line: 1, Col: 21},
			To:   LineCol{Line: 1, Col: 23},
		},
	}
	checkCompletion(t, "MATCH (atern) RETURN at▼", expected, true)
	checkCompletion(t, "MATCH (atern) RETURN a▼t", expected, true)
	checkCompletion(t, "MATCH (atern) RETURN ▼at", expected, true)
}

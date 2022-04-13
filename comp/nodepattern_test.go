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
	expected := Result{
		Items: []Item{
			{Type: types.Label, View: ":y", Content: ":y"},
			{Type: types.Label, View: ":x", Content: ":x"},
		},
		Range: Range{
			From: LineCol{Line: 1, Col: 18},
			To:   LineCol{Line: 1, Col: 20},
		},
	}
	checkCompletion(t, "MATCH (n) MATCH (a:▼b", expected, false)
	checkCompletion(t, "MATCH (n) MATCH (a:b▼", expected, false)
}

func TestWithFiltersYieldsLabelList(t *testing.T) {
	expected := Result{
		Items: []Item{
			{Type: types.Label, View: ":y", Content: ":y"},
		},
		Range: Range{
			From: LineCol{Line: 1, Col: 18},
			To:   LineCol{Line: 1, Col: 20},
		},
	}
	checkCompletion(t, "MATCH (n) MATCH (a:▼y", expected, true)
	checkCompletion(t, "MATCH (n) MATCH (a:y▼", expected, true)
}

func TestWithoutFiltersYieldsLabelListIfOnlyColonIsPresent(t *testing.T) {
	expected := Result{
		Items: []Item{
			{Type: types.Label, View: ":y", Content: ":y"},
			{Type: types.Label, View: ":x", Content: ":x"},
		},
		Range: Range{
			From: LineCol{Line: 1, Col: 18},
			To:   LineCol{Line: 1, Col: 19},
		},
	}
	checkCompletion(t, "MATCH (n) MATCH (a:▼", expected, false)
	checkCompletion(t, "MATCH (n) MATCH (a:▼", expected, true)
}

func TestWithoutFiltersYieldsLabelListIfOnlyColonParenthesisIsPresent(t *testing.T) {
	expected := Result{
		Items: []Item{
			{Type: types.Label, View: ":y", Content: ":y"},
			{Type: types.Label, View: ":x", Content: ":x"},
		},
		Range: Range{
			From: LineCol{Line: 1, Col: 8},
			To:   LineCol{Line: 1, Col: 9},
		},
	}
	checkCompletion(t, "MATCH (a:▼) MATCH ()", expected, false)
	checkCompletion(t, "MATCH (a:▼) MATCH ()", expected, true)
}

func TestWithoutFiltersYieldsLabelAndVariablesAtTheBeginningOfNodePattern(t *testing.T) {
	expected := Result{
		Items: []Item{
			{Type: types.Variable, View: "a", Content: "a"},
			{Type: types.Label, View: ":y", Content: ":y"},
			{Type: types.Label, View: ":x", Content: ":x"},
		},
		Range: Range{
			From: LineCol{Line: 1, Col: 17},
			To:   LineCol{Line: 1, Col: 17},
		},
	}
	checkCompletion(t, "MATCH (a) MATCH (▼", expected, false)
	checkCompletion(t, "MATCH (a) MATCH (▼", expected, true)
}

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

func TestVariableWithoutFiltersYieldsVariableList(t *testing.T) {
	expected := Result{
		Items: []Item{
			{Type: types.Variable, View: "n", Content: "n"},
			{Type: types.Variable, View: "a", Content: "a"},
		},
		Range: Range{
			From: LineCol{Line: 1, Col: 17},
			To:   LineCol{Line: 1, Col: 18},
		},
	}

	checkCompletion(t, "MATCH (n) MATCH (▼a", expected, false)
	checkCompletion(t, "MATCH (n) MATCH (a▼", expected, false)
}

func TestVariableWithoutFiltersYieldsVariableListInMultipleQueriesFirstQuery(t *testing.T) {
	expected := Result{
		Items: []Item{
			{Type: types.Variable, View: "x", Content: "x"},
		},
		Range: Range{
			From: LineCol{Line: 1, Col: 7},
			To:   LineCol{Line: 1, Col: 8},
		},
	}

	checkCompletion(t, "MATCH (▼x); MATCH (n) MATCH (a", expected, false)
	checkCompletion(t, "MATCH (x▼); MATCH (n) MATCH (a", expected, false)
}

func TestVariableWithoutFiltersYieldsVariableListInMultipleQueriesSecondQuery(t *testing.T) {
	expected := Result{
		Items: []Item{
			{Type: types.Variable, View: "n", Content: "n"},
			{Type: types.Variable, View: "a", Content: "a"},
		},
		Range: Range{
			From: LineCol{Line: 1, Col: 28},
			To:   LineCol{Line: 1, Col: 29},
		},
	}

	checkCompletion(t, "MATCH (x); MATCH (n) MATCH (▼a", expected, false)
	checkCompletion(t, "MATCH (x); MATCH (n) MATCH (a▼", expected, false)
}

func TestVariableWithFiltersYieldsVariableList(t *testing.T) {
	expected := Result{
		Items: []Item{
			{Type: types.Variable, View: "markus", Content: "markus"},
		},
		Range: Range{
			From: LineCol{Line: 1, Col: 36},
			To:   LineCol{Line: 1, Col: 40},
		},
	}

	checkCompletion(t, "MATCH (penny) MATCH (markus) RETURN mark▼", expected, true)
	checkCompletion(t, "MATCH (penny) MATCH (markus) RETURN mar▼k", expected, true)
}

func TestVariableWithFiltersYieldsVariableListWithoutVariableUnderCursor(t *testing.T) {
	expected := Result{
		Items: []Item{
			{Type: types.Variable, View: "var", Content: "var"},
		},
		Range: Range{
			From: LineCol{Line: 1, Col: 19},
			To:   LineCol{Line: 1, Col: 21},
		},
	}

	checkCompletion(t, "MATCH (var) RETURN va▼", expected, true)
}

func TestVariableWithFiltersYieldsVariableListUnderCursorMatches(t *testing.T) {
	expected := Result{
		Items: []Item{
			{Type: types.Variable, View: "var", Content: "var"},
		},
		Range: Range{
			From: LineCol{Line: 1, Col: 19},
			To:   LineCol{Line: 1, Col: 22},
		},
	}

	checkCompletion(t, "MATCH (var) RETURN var▼", expected, true)
}

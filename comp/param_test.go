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
	expected := Result{
		Items: []Item{
			{Type: types.Parameter, View: "param1", Content: "param1"},
			{Type: types.Parameter, View: "param2", Content: "param2"},
		},
		Range: Range{
			From: LineCol{Line: 1, Col: 14},
			To:   LineCol{Line: 1, Col: 15},
		},
	}
	checkCompletion(t, "RETURN {b} + $▼a", expected, false)
	checkCompletion(t, "RETURN {b} + $a▼", expected, false)
}

func TestParameterWithoutFiltersYieldsParameterNameListAfterFirstSymbol(t *testing.T) {
	expected := Result{
		Items: []Item{
			{Type: types.Parameter, View: "param1", Content: "param1"},
			{Type: types.Parameter, View: "param2", Content: "param2"},
		},
		Range: Range{
			From: LineCol{Line: 1, Col: 13},
			To:   LineCol{Line: 1, Col: 14},
		},
	}
	checkCompletion(t, "RETURN $b + {▼a}", expected, false)
	checkCompletion(t, "RETURN $b + {a▼}", expected, false)
}

func TestWithFiltersYieldsParameterNameList(t *testing.T) {
	expected := Result{
		Items: []Item{
			{Type: types.Parameter, View: "param1", Content: "param1"},
			{Type: types.Parameter, View: "param2", Content: "param2"},
		},
		Range: Range{
			From: LineCol{Line: 1, Col: 14},
			To:   LineCol{Line: 1, Col: 15},
		},
	}
	checkCompletion(t, "RETURN {b} + $▼a", expected, true)
	checkCompletion(t, "RETURN {b} + $a▼", expected, true)
}

func TestWithFiltersYieldsLegacyParameterNameList(t *testing.T) {
	expected := Result{
		Items: []Item{
			{Type: types.Parameter, View: "param1", Content: "param1"},
			{Type: types.Parameter, View: "param2", Content: "param2"},
		},
		Range: Range{
			From: LineCol{Line: 1, Col: 8},
			To:   LineCol{Line: 1, Col: 9},
		},
	}
	checkCompletion(t, "RETURN {p▼}", expected, true)
}

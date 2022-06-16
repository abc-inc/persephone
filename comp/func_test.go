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

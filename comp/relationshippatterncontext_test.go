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

func TestPropertyKeyTypesYieldsPropertyKeyWithNoFirstCharInsideMapWithinExpression(t *testing.T) {
	checkCompletionTypes(t, "MATCH ()-[▼{}]-()", true, []types.Type{types.PropertyKey})
}

func TestPropertyKeyTypesYieldsPropertyKeyWithNoFirstCharInsideMapWithinExpressionWithFirstSymbol(t *testing.T) {
	checkCompletionTypes(t, "WITH $param MATCH ()-[{▼p}]-()", true, []types.Type{types.PropertyKey, types.Parameter})
}

func TestPropertyKeyTypesYieldsPropertyKeyWithNoFirstCharInsideMapWithinExpressionWithoutClosingCurlyBrace(t *testing.T) {
	checkCompletionTypes(t, "WITH $param MATCH ()-[▼{", true, []types.Type{types.PropertyKey, types.Parameter})
}

func TestPropertyKeyTypesYieldsPropertyKeyWithFirstCharInsideMapWithinExpressionWithoutClosingCurlyBrace(t *testing.T) {
	checkCompletionTypes(t, "WITH $param MATCH ()-[{▼p", true, []types.Type{types.PropertyKey, types.Parameter})
}

func TestPropertyKeyAutoCompletionYieldsPropertyKeyListInsideMapWithinExpression(t *testing.T) {
	expected := Result{
		Items: []Item{
			{Type: types.PropertyKey, View: "prop1", Content: "prop1"},
			{Type: types.PropertyKey, View: "prop2", Content: "prop2"},
		},
		Range: Range{
			From: LineCol{Line: 1, Col: 11},
			To:   LineCol{Line: 1, Col: 11},
		},
	}
	checkCompletion(t, "MATCH ()-[{▼}]-()", expected, true)
}

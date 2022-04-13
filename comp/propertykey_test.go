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

func TestTypesYieldsPropertyKey(t *testing.T) {
	checkCompletionTypes(t, "MATCH (a) RETURN a.▼b", true, []types.Type{types.PropertyKey})
}

func TestTypesYieldsPropertyKeyWithNoFirstCharTypedInSetClause(t *testing.T) {
	checkCompletionTypes(t, "MATCH (a) SET a▼.", true, []types.Type{types.PropertyKey})
}

func TestTypesYieldsPropertyKeyWithNoFirstCharTypedAfterWhereKeyword(t *testing.T) {
	checkCompletionTypes(t, "MATCH (n) where n▼.", true, []types.Type{types.PropertyKey})
}

func TestTypesYieldsPropertyKeyWithNoFirstCharTypedAfterAnExpression(t *testing.T) {
	checkCompletionTypes(t, `MATCH (a) WHERE a.name > "name" AND a▼. `, true, []types.Type{types.PropertyKey})
}

func TestWithoutFiltersYieldsPropertyKeyList(t *testing.T) {
	expected := Result{
		Items: []Item{
			{Type: types.PropertyKey, View: "prop1", Content: "prop1"},
			{Type: types.PropertyKey, View: "prop2", Content: "prop2"},
		},
		Range: Range{
			From: LineCol{Line: 1, Col: 19},
			To:   LineCol{Line: 1, Col: 20},
		},
	}
	checkCompletion(t, "MATCH (a) RETURN a.▼b", expected, false)
	checkCompletion(t, "MATCH (a) RETURN a.b▼", expected, false)
}

func TestWithoutFiltersYieldsPropertyKeyListWithoutFirstCharTypedInASetClause(t *testing.T) {
	expected := Result{
		Items: []Item{
			{Type: types.PropertyKey, View: "prop1", Content: "prop1"},
			{Type: types.PropertyKey, View: "prop2", Content: "prop2"},
		},
		Range: Range{
			From: LineCol{Line: 1, Col: 16},
			To:   LineCol{Line: 1, Col: 16},
		},
	}
	checkCompletion(t, "MATCH (a) SET a.▼", expected, false)
}

func TestWithoutFiltersYieldsPropertyKeyListWithoutFirstCharTypedAfterWhereKeyword(t *testing.T) {
	expected := Result{
		Items: []Item{
			{Type: types.PropertyKey, View: "prop1", Content: "prop1"},
			{Type: types.PropertyKey, View: "prop2", Content: "prop2"},
		},
		Range: Range{
			From: LineCol{Line: 1, Col: 18},
			To:   LineCol{Line: 1, Col: 18},
		},
	}
	checkCompletion(t, "MATCH (a) WHERE n.▼", expected, false)
}

func TestWithFiltersYieldsPropertyKeyList(t *testing.T) {
	expected := Result{
		Items: []Item{
			{Type: types.PropertyKey, View: "prop1", Content: "prop1"},
		},
		Range: Range{
			From: LineCol{Line: 1, Col: 19},
			To:   LineCol{Line: 1, Col: 21},
		},
	}
	checkCompletion(t, "MATCH (a) RETURN a.p1▼", expected, true)
	checkCompletion(t, "MATCH (a) RETURN a.p▼1", expected, true)
	checkCompletion(t, "MATCH (a) RETURN a.▼p1", expected, true)
}

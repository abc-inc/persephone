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

func TestTypesYieldsRelationshipTypeIfCharacterPresent(t *testing.T) {
	checkCompletionTypes(t, "MATCH (a)-[a:▼b]-", true, []types.Type{types.RelationshipType})
}

func TestTypesYieldsRelationshipTypeIfStartOfRelationshipDetails(t *testing.T) {
	checkCompletionTypes(t, "MATCH (a)-▼[", true, []types.Type{types.Variable, types.RelationshipType})
}

func TestTypesYieldsRelationshipTypeIfSecondAndOnlyColonPresent(t *testing.T) {
	checkCompletionTypes(t, "MATCH (a)-[▼:", true, []types.Type{types.RelationshipType})
}

func TestTypesYieldsRelationshipTypeIfSecondAndColonPresent(t *testing.T) {
	checkCompletionTypes(t, "MATCH (a)-[:q|▼:", true, []types.Type{types.RelationshipType})
}

func TestWithoutFiltersYieldsRelationshipTypeList(t *testing.T) {
	expected := Result{
		Items: []Item{
			{Type: types.RelationshipType, View: ":rel1", Content: ":rel1"},
			{Type: types.RelationshipType, View: ":rel 2", Content: ":`rel 2`"},
		},
		Range: Range{
			From: LineCol{Line: 1, Col: 12},
			To:   LineCol{Line: 1, Col: 14},
		},
	}
	checkCompletion(t, "MATCH (a)-[a:▼b]-", expected, false)
	checkCompletion(t, "MATCH (a)-[a:b▼]-", expected, false)
}

func TestWithoutFiltersYieldsRelationshipTypeListIfOnlyColonPresent(t *testing.T) {
	expected := Result{
		Items: []Item{
			{Type: types.RelationshipType, View: ":rel1", Content: ":rel1"},
			{Type: types.RelationshipType, View: ":rel 2", Content: ":`rel 2`"},
		},
		Range: Range{
			From: LineCol{Line: 1, Col: 12},
			To:   LineCol{Line: 1, Col: 13},
		},
	}
	checkCompletion(t, "MATCH (a)-[a▼:]-()", expected, false)
}

func TestWithFiltersYieldsRelationshipTypeList(t *testing.T) {
	expected := Result{
		Items: []Item{
			{Type: types.RelationshipType, View: ":rel1", Content: ":rel1"},
		},
		Range: Range{
			From: LineCol{Line: 1, Col: 12},
			To:   LineCol{Line: 1, Col: 15},
		},
	}
	checkCompletion(t, "MATCH (a)-[a:l1▼]-", expected, true)
	checkCompletion(t, "MATCH (a)-[a:l▼1]-", expected, true)
	checkCompletion(t, "MATCH (a)-[a:▼l1]-", expected, true)
}

func TestWithFiltersYieldsRelationshipTypeListIfOnlyColonPresent(t *testing.T) {
	expected := Result{
		Items: []Item{
			{Type: types.RelationshipType, View: ":rel1", Content: ":rel1"},
			{Type: types.RelationshipType, View: ":rel 2", Content: ":`rel 2`"},
		},
		Range: Range{
			From: LineCol{Line: 1, Col: 12},
			To:   LineCol{Line: 1, Col: 13},
		},
	}
	checkCompletion(t, "MATCH (a)-[a▼:]-()", expected, true)
}

func TestWithFiltersYieldsRelationshipTypeAndVariableListAtTheBeginningOfPattern(t *testing.T) {
	expected := Result{
		Items: []Item{
			{Type: types.Variable, View: "a", Content: "a"},
			{Type: types.RelationshipType, View: ":rel1", Content: ":rel1"},
			{Type: types.RelationshipType, View: ":rel 2", Content: ":`rel 2`"},
		},
		Range: Range{
			From: LineCol{Line: 1, Col: 10},
			To:   LineCol{Line: 1, Col: 10},
		},
	}
	checkCompletion(t, "MATCH (a)-▼[", expected, true)
}

func TestWithFiltersYieldsRelationshipTypeIfFirstLetterIsTypedIn(t *testing.T) {
	expected := Result{
		Items: []Item{
			{Type: types.RelationshipType, View: ":rel1", Content: ":rel1"},
			{Type: types.RelationshipType, View: ":rel 2", Content: ":`rel 2`"},
		},
		Range: Range{
			From: LineCol{Line: 1, Col: 11},
			To:   LineCol{Line: 1, Col: 13},
		},
	}
	checkCompletion(t, "MATCH (a)-[:r▼", expected, true)
}

func TestWithFiltersYieldsRelationshipTypeIfAfterColonWithSpace(t *testing.T) {
	expected := Result{
		Items: []Item{
			{Type: types.RelationshipType, View: ":rel1", Content: ":rel1"},
			{Type: types.RelationshipType, View: ":rel 2", Content: ":`rel 2`"},
		},
		Range: Range{
			From: LineCol{Line: 1, Col: 11},
			To:   LineCol{Line: 1, Col: 12},
		},
	}
	checkCompletion(t, "MATCH (a)-[:▼ return n;", expected, true)
}

package comp_test

import (
	"testing"

	. "github.com/abc-inc/merovingian/comp"
	"github.com/abc-inc/merovingian/types"
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

package comp_test

import (
	"testing"

	. "github.com/abc-inc/merovingian/comp"
	"github.com/abc-inc/merovingian/types"
)

func TestKeyword(t *testing.T) {
	expected := Result{
		Range: Range{
			From: LineCol{Line: 1, Col: 10},
			To:   LineCol{Line: 1, Col: 12},
		},
		Items: []Item{
			{Type: types.Keyword, View: "WHEN", Content: "WHEN"},
			{Type: types.Keyword, View: "WHERE", Content: "WHERE"},
			{Type: types.Keyword, View: "WITH", Content: "WITH"},
		},
	}
	checkCompletion(t, "MATCH (n) wH▼", expected, true)
	checkCompletion(t, "MATCH (n) w▼H", expected, true)
	checkCompletion(t, "MATCH (n) ▼wH", expected, true)
}

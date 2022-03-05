package autocompletion

import (
	"testing"

	"github.com/abc-inc/merovingian/comp"
	"github.com/abc-inc/merovingian/types"
)

func TestKeyword(t *testing.T) {
	expected := comp.Result{
		Range: comp.Range{
			From: comp.LineCol{Line: 1, Col: 0},
			To:   comp.LineCol{Line: 1, Col: 5},
		},
		Items: []comp.Item{
			{Type: types.Keyword, View: "MATCH", Content: "MATCH"},
		},
	}
	checkCompletion(t, "MAT▼", expected, true)

	expected = comp.Result{
		Range: comp.Range{
			From: comp.LineCol{Line: 1, Col: 10},
			To:   comp.LineCol{Line: 1, Col: 12},
		},
		Items: []comp.Item{
			{Type: types.Keyword, View: "WHEN", Content: "WHEN"},
			{Type: types.Keyword, View: "WHERE", Content: "WHERE"},
			{Type: types.Keyword, View: "WITH", Content: "WITH"},
		},
	}
	//checkCompletion(t, "MATCH (n) wH▼", expected, true)
	//checkCompletion(t, "MATCH (n) w▼H", expected, true)
	//checkCompletion(t, "MATCH (n) ▼wH", expected, true)
}

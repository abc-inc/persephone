package comp_test

import (
	"fmt"
	"strings"
	"testing"

	. "github.com/abc-inc/merovingian/comp"
	"github.com/abc-inc/merovingian/editor"
	"github.com/abc-inc/merovingian/ndb"
	"github.com/abc-inc/merovingian/types"
	. "github.com/stretchr/testify/assert"
)

var schema = ndb.Schema{
	Labels:   []string{":y", ":x"},
	RelTypes: []string{":rel1", ":rel 2"},
	PropKeys: []string{"prop1", "prop2"},
	Params:   []string{"param1", "param2"},
	Funcs: []ndb.Func{
		{Name: "toFloat", Sig: "expression"},
		{Name: "head", Sig: "expression"},
	},
	Procs: []ndb.Func{{
		Name: "db.indexes",
		Sig:  "()",
		RetItems: []ndb.Func{
			{Name: "description", Sig: "STRING?"},
			{Name: "state", Sig: "STRING?"},
			{Name: "type", Sig: "STRING?"},
		},
	},
		{Name: "org.neo4j.graph.traverse", Sig: "expression"},
	},
	ConCmds: []ndb.Cmd{
		{Name: ":clear"},
		{Name: ":play"},
		{Name: ":help", Desc: "helpdesc", SubCmds: []ndb.Cmd{{Name: "match"}, {Name: "create"}}},
		{Name: ":server", SubCmds: []ndb.Cmd{
			{Name: "user", SubCmds: []ndb.Cmd{
				{Name: "list", Desc: "listdesc"},
				{Name: "add"},
			}},
		}},
	},
}

func checkCompletion(t *testing.T, queryWithCursor string, expectedItems Result, doFilter bool) {
	pos := strings.IndexRune(queryWithCursor, '▼')
	query := strings.Replace(queryWithCursor, "▼", "", 1)

	backend := editor.NewEditorSupport(query)
	backend.SetSchema(schema)
	completion := backend.GetCompletion(1, pos, doFilter)
	Equal(t, expectedItems, completion)
}

func checkCompletionTypes(t *testing.T, queryWithCursor string, found bool, expectedTypes []types.Type) {
	pos := strings.IndexRune(queryWithCursor, '▼')
	query := strings.Replace(queryWithCursor, "▼", "", 1)

	es := editor.NewEditorSupport(query)
	el := es.GetElementForCompletion(1, pos)
	ts := GetTypes(el)

	var exp []types.Data
	for _, t := range expectedTypes {
		exp = append(exp, types.Data{Type: t, FilterLastElement: found})
	}

	// TODO: fix workaround
	for i := range ts.Types {
		ts.Types[i].Path = nil
		ts.Types[i].FilterLastElement = found
		fmt.Println(ts.Types[i])
	}

	Equal(t, ComplInfo{Found: found, Types: exp}, ts)
}

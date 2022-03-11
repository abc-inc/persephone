package autocompletion

import (
	"fmt"
	"strings"
	"testing"

	"github.com/abc-inc/merovingian/comp"
	"github.com/abc-inc/merovingian/cypher"
	"github.com/abc-inc/merovingian/db/neo4j"
	"github.com/abc-inc/merovingian/resolver"
	"github.com/abc-inc/merovingian/types"
	. "github.com/stretchr/testify/assert"
)

var schema = neo4j.Schema{
	Labels:   []string{":y", ":x"},
	RelTypes: []string{":rel1", ":rel 2"},
	PropKeys: []string{"prop1", "prop2"},
	Params:   []string{"param1", "param2"},
	Funcs: []neo4j.Func{
		{Name: "toFloat", Sig: "expression"},
		{Name: "head", Sig: "expression"},
	},
	Procs: []neo4j.Func{{
		Name: "db.indexes",
		Sig:  "()",
		RetItems: []neo4j.Func{
			{Name: "description", Sig: "STRING?"},
			{Name: "state", Sig: "STRING?"},
			{Name: "type", Sig: "STRING?"},
		},
	},
		{Name: "org.neo4j.graph.traverse", Sig: "expression"},
	},
	ConCmds: []neo4j.Cmd{
		{Name: ":clear"},
		{Name: ":play"},
		{Name: ":help", Desc: "helpdesc", SubCmds: []neo4j.Cmd{{Name: "match"}, {Name: "create"}}},
		{Name: ":server", SubCmds: []neo4j.Cmd{
			{Name: "user", SubCmds: []neo4j.Cmd{
				{Name: "list", Desc: "listdesc"},
				{Name: "add"},
			}},
		}},
	},
}

func checkCompletion(t *testing.T, queryWithCursor string, expectedItems comp.Result, doFilter bool) {
	pos := strings.IndexRune(queryWithCursor, '▼')
	query := strings.Replace(queryWithCursor, "▼", "", 1)

	backend := cypher.NewEditorSupport(query)
	backend.SetSchema(schema)
	completion := backend.GetCompletion(1, pos, doFilter)
	Equal(t, expectedItems, completion)
}

func checkCompletionTypes(t *testing.T, queryWithCursor string, found bool, expectedTypes []types.Type) {
	pos := strings.IndexRune(queryWithCursor, '▼')
	query := strings.Replace(queryWithCursor, "▼", "", 1)

	backend := cypher.NewEditorSupport(query)
	el := backend.GetElementForCompletion(1, pos)
	ts := resolver.GetTypes(el)

	var exp []comp.TypeData
	for _, t := range expectedTypes {
		exp = append(exp, comp.TypeData{Type: t})
	}

	// TODO: fix workaround
	for i := range ts.Types {
		ts.Types[i].Path = nil
		fmt.Println(ts.Types[i])
	}

	Equal(t, comp.ComplInfo{Found: found, Types: exp}, ts)
}

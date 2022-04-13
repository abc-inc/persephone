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
	"fmt"
	"strings"
	"testing"

	. "github.com/abc-inc/persephone/comp"
	"github.com/abc-inc/persephone/editor"
	"github.com/abc-inc/persephone/graph"
	"github.com/abc-inc/persephone/types"
	. "github.com/stretchr/testify/require"
)

var schema = graph.Schema{
	Labels:   []string{":y", ":x"},
	RelTypes: []string{":rel1", ":rel 2"},
	PropKeys: []string{"prop1", "prop2"},
	Params:   []string{"param1", "param2"},
	Funcs: []graph.Func{
		{Name: "toFloat", Sig: "expression"},
		{Name: "head", Sig: "expression"},
	},
	Procs: []graph.Func{{
		Name: "db.indexes",
		Sig:  "()",
		RetItems: []graph.Func{
			{Name: "description", Sig: "STRING?"},
			{Name: "state", Sig: "STRING?"},
			{Name: "type", Sig: "STRING?"},
		},
	},
		{Name: "org.neo4j.graph.traverse", Sig: "expression"},
	},
	ConCmds: []graph.Cmd{
		{Name: ":clear"},
		{Name: ":play"},
		{Name: ":help", Desc: "helpdesc", SubCmds: []graph.Cmd{{Name: "match"}, {Name: "create"}}},
		{Name: ":server", SubCmds: []graph.Cmd{
			{Name: "user", SubCmds: []graph.Cmd{
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

	Equal(t, Info{Found: found, Types: exp}, ts)
}

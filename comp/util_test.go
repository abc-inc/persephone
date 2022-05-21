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
	"strings"
	"testing"

	"github.com/abc-inc/go-data-neo4j/meta"
	. "github.com/abc-inc/persephone/comp"
	"github.com/abc-inc/persephone/editor"
	"github.com/abc-inc/persephone/types"
	. "github.com/stretchr/testify/require"
)

var schema = Metadata{
	Schema: meta.Schema{
		Labels:   []string{":y", ":x"},
		RelTypes: []string{":rel1", ":rel 2"},
		PropKeys: []string{"prop1", "prop2"},
		Funcs: []meta.Func{
			{Name: "toFloat", Sig: "expression"},
			{Name: "head", Sig: "expression"},
		},
		Procs: []meta.Func{{
			Name: "db.indexes",
			Sig:  "()",
			RetItems: []meta.Func{
				{Name: "description", Sig: "STRING?"},
				{Name: "state", Sig: "STRING?"},
				{Name: "type", Sig: "STRING?"},
			},
		},
			{Name: "org.neo4j.graph.traverse", Sig: "expression"},
		},
	},
	Params: []string{"param1", "param2"},
	ConCmds: []Cmd{
		{Name: ":clear"},
		{Name: ":play"},
		{Name: ":help", Desc: "helpdesc", SubCmds: []Cmd{{Name: "match"}, {Name: "create"}}},
		{Name: ":server", SubCmds: []Cmd{
			{Name: "user", SubCmds: []Cmd{
				{Name: "list", Desc: "listdesc"},
				{Name: "add"},
			}},
		}},
	},
}

func checkCompletion(t *testing.T, queryWithCursor string, expRes Result, doFilter bool) {
	pos := strings.IndexRune(queryWithCursor, '▼')
	query := strings.Replace(queryWithCursor, "▼", "", 1)

	e := editor.NewEditor(query)
	e.SetSchema(schema)
	completion := e.GetCompletion(1, pos, doFilter)
	Equal(t, expRes, completion)
}

func checkCompletionTypes(t *testing.T, queryWithCursor string, found bool, expTypes []types.Type) {
	exp := make([]types.Data, len(expTypes))
	for i, t := range expTypes {
		exp[i] = types.Data{Type: t, FilterLastElement: false}
	}

	checkCompletionTypesInfo(t, queryWithCursor, Info{Found: found, Types: exp})
}

func checkCompletionTypesInfo(t *testing.T, queryWithCursor string, expInfo Info) {
	pos := strings.IndexRune(queryWithCursor, '▼')
	query := strings.Replace(queryWithCursor, "▼", "", 1)

	e := editor.NewEditor(query)
	el := e.GetElementForCompletion(1, pos)
	ts := GetTypes(el)

	Equal(t, expInfo, ts)
}

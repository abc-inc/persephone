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

package editor_test

import (
	"testing"

	"github.com/abc-inc/persephone/editor"
	. "github.com/stretchr/testify/require"
)

func Test(t *testing.T) {
	tests := []struct {
		name   string
		cypher string
	}{
		{name: "param command with arg and other command", cypher: ":play http://something.com; :play;"},
		{name: "param command with url", cypher: ":play http://something.com/something.html;"},
		{name: "param command with function", cypher: ":param num => rand();"},
		{name: "param command with int", cypher: ":param myInt => 1;"},
		{name: "param command with double", cypher: ":param myDouble => 1.1;"},
		{name: "param command with string", cypher: `:param myString => "hello";`},
		{name: "param command with map", cypher: ":param obj => {x: 1, y: 2};"},
		{name: "param command with array", cypher: ":param arr => [1, 2, 3];"},
		{name: "command with json param and something else", cypher: `:play "http://link.com" {"hello": "world", "key": true, "pop": 125.45};`},
		{name: "command with json param", cypher: `:play {"hello": "world", "key": true, "pop": 125.45};`},
		{name: "simple command", cypher: ":play;"},
		{name: "command with param", cypher: ":play 1 'string' true;"},
		{name: "command and query", cypher: `:play "url"; match (n);`},
		{name: "command with variable", cypher: ":play variable;"},
		{name: "command with multiple variables", cypher: ":play variable anotherVariable;"},
		{name: "command with map literal", cypher: `:play {hello: "world", key: true, pop: 125.45};`},
		{name: "command with map literal and something else", cypher: `:play "http://link.com" {hello: "world", key: true, pop: 125.45};`},
		{name: "command with dashes", cypher: ":play-this-now;"},
		{name: "command with key value literal", cypher: ":config n: 'xxx';"},
		{name: "command variable-with-dashes", cypher: ":config variable-with-dashes;"},
		{name: "GET", cypher: ":GET /db/data/labels;"},
		{name: "DELETE", cypher: ":DELETE /db/data/transaction/2;"},
		{name: "POST", cypher: ":POST /db/data/node { name:\"Tiberius\" }"},
		{name: "PUT", cypher: ":PUT /db/data/node/198/properties/foo \"Delia\""},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			e := editor.NewEditor(test.cypher)
			Nil(t, e.ParseErrors)
		})
	}
}

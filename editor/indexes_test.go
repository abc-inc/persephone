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

func TestIndexes(t *testing.T) {
	tests := []struct {
		name   string
		cypher string
	}{
		{name: "create on single property", cypher: "CREATE INDEX ON :Person(name)"},
		{name: "create on compound property", cypher: "CREATE INDEX ON :Person(name, surname)"},
		{name: "drop on single property", cypher: "DROP INDEX ON :Person(name)"},
		{name: "drop on compound property", cypher: "DROP INDEX ON :Person(name, surname)"},
		{name: "hint on single property", cypher: "MATCH (f:Foo) USING INDEX f:Foo(bar)"},
		{name: "hint on compound property", cypher: "MATCH (f:Foo) USING INDEX f:Foo(bar,baz)"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			s := editor.NewEditorSupport(test.cypher)
			Nil(t, s.ParseErrors)
		})
	}
}

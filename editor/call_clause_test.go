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

func TestCallClause(t *testing.T) {
	tests := []struct {
		name   string
		cypher string
	}{
		{name: "with parenthesis", cypher: "CALL procedure()"},
		{name: "without parenthesis", cypher: "CALL procedure"},
		{name: "with where after yield", cypher: "CALL procedure() YIELD name WHERE true RETURN name"},
		{name: "string contains cypher", cypher: `CALL foo.bar("RETURN 1")`},
		{name: "string contains cypher with new lines", cypher: `CALL foo.bar("MATCH (n) \nRETURN n")`},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			e := editor.NewEditor(test.cypher)
			Nil(t, e.ParseErrors)
		})
	}
}

func TestCallClausMismatchedParenthesis(t *testing.T) {
	e := editor.NewEditor("CALL ()")
	Equal(t, 1, len(e.ParseErrors))
	Equal(t, 5, e.ParseErrors[0].Col)
}

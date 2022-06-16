// Copyright 2022 The Persephone authors
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

func TestSpecialFunctions(t *testing.T) {
	tests := []struct {
		name   string
		cypher string
	}{
		{name: "reduce", cypher: "RETURN reduce(totalAge = 0, n IN [p] | totalAge + n.age) AS reduction"},
		{name: "shortestPath", cypher: "RETURN shortestPath( ( f)-[]-( t) );"},
		{name: "allShortestPaths", cypher: "RETURN allShortestPaths((f)-[]-(t));"},
		{name: "exists", cypher: "RETURN n.prop AS prop1, exists((n)-[:SOMETHING]->()) AS something;"},
		{name: "3rd party", cypher: "RETURN a.b();"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			e := editor.NewEditor(test.cypher)
			Nil(t, e.ParseErrors)
		})
	}
}

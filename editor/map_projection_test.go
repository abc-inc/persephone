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

func TestMapProjection(t *testing.T) {
	tests := []struct {
		name   string
		cypher string
	}{
		{name: "all property selector", cypher: "RETURN person { .* };"},
		{name: "property selector", cypher: "RETURN person { .name };"},
		{name: "literal entry and map projection inside", cypher: "RETURN person { someProp: collect(moreProps { .variable1, .variable2 })};"},
		{name: "literal entry", cypher: "RETURN person { someProp: collect(expression)};"},
		{name: "variable", cypher: "RETURN person { person };"},
		{name: "variable without spaces", cypher: "RETURN person{person};"},
		{name: "multiple", cypher: "RETURN person{ person, .person, something: expression(), .*};"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			e := editor.NewEditor(test.cypher)
			Nil(t, e.ParseErrors)
		})
	}
}

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

func TestConstraints(t *testing.T) {
	tests := []struct {
		name   string
		cypher string
	}{
		{name: "create on single property", cypher: "CREATE CONSTRAINT ON (n:Person) ASSERT (n.email) IS NODE KEY"},
		{name: "drop on single property", cypher: "DROP CONSTRAINT ON (n:Person) ASSERT (n.email) IS NODE KEY"},
		{name: "create on composite property", cypher: "CREATE CONSTRAINT ON (n:User) ASSERT (n.firstname,n.lastname) IS NODE KEY"},
		{name: "drop on composite property", cypher: "DROP CONSTRAINT ON (n:Person) ASSERT (n.firstname,n.lastname) IS NODE KEY"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			s := editor.NewEditorSupport(test.cypher)
			Nil(t, s.ParseErrors)
		})
	}
}

func TestDropWithoutPropertyFails(t *testing.T) {
	s := editor.NewEditorSupport("DROP CONSTRAINT ON (n:Person) ASSERT (n)")
	Equal(t, 1, len(s.ParseErrors))
	Equal(t, len(s.Input), s.ParseErrors[0].Col)
}

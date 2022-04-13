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
	"strconv"
	"testing"

	"github.com/abc-inc/persephone/editor"
	. "github.com/stretchr/testify/require"
)

func TestCypherLegacyFiles(t *testing.T) {
	for i, test := range cypherLegacy {
		t.Run("cypher-legacy-query-"+strconv.Itoa(i), func(t *testing.T) {
			s := editor.NewEditorSupport(test)
			Nil(t, s.ParseErrors)
		})
	}
}

func TestOpenCypherFiles(t *testing.T) {
	for i, test := range cypherDefault {
		t.Run("cypher-query-"+strconv.Itoa(i), func(t *testing.T) {
			s := editor.NewEditorSupport(test)
			Nil(t, s.ParseErrors)
		})
	}
}

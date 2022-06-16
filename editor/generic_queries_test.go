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

func TestGenericQueries(t *testing.T) {
	tests := []string{
		"MERGE () ON CREATE SET connection.departure = 1445, connection.arrival = 1710",
		"MERGE () ON MATCH SET connection.departure = 1445, connection.arrival = 1710",
		"SET a=a,b=b",
		"SET a=a ,b=b",
		"SET a=a, b=b",
		"SET a=a , b=b",
		"RETURN a ORDER BY a,b;",
		"RETURN a ORDER BY a ,b;",
		"RETURN a ORDER BY a, b;",
		"RETURN a ORDER BY a , b;",
	}
	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			e := editor.NewEditor(test)
			Nil(t, e.ParseErrors)
		})
	}
}

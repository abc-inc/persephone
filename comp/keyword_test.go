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
	"testing"

	. "github.com/abc-inc/persephone/comp"
	"github.com/abc-inc/persephone/types"
)

func TestKeyword(t *testing.T) {
	expected := Result{
		Range: Range{
			From: LineCol{Line: 1, Col: 10},
			To:   LineCol{Line: 1, Col: 12},
		},
		Items: []Item{
			{Type: types.Keyword, View: "WHEN", Content: "WHEN"},
			{Type: types.Keyword, View: "WHERE", Content: "WHERE"},
			{Type: types.Keyword, View: "WITH", Content: "WITH"},
		},
	}
	checkCompletion(t, "MATCH (n) wH▼", expected, true)
	checkCompletion(t, "MATCH (n) w▼H", expected, true)
	checkCompletion(t, "MATCH (n) ▼wH", expected, true)
}

func TestFirstKeyword(t *testing.T) {
	expected := Result{
		Range: Range{
			From: LineCol{Line: 1, Col: 0},
			To:   LineCol{Line: 1, Col: 6},
		},
		Items: []Item{
			{Type: types.Keyword, View: "CREATE", Content: "CREATE"},
		},
	}
	checkCompletion(t, "CREATE▼", expected, true)
}

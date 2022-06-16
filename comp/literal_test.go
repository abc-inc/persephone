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

package comp_test

import (
	"testing"

	. "github.com/abc-inc/persephone/comp"
	"github.com/abc-inc/persephone/types"
)

func TestLiteralStringTypesYieldsNothingType(t *testing.T) {
	checkCompletionTypes(t, `RETURN "▼"`, true, []types.Type{types.Noop})
	checkCompletionTypes(t, `RETURN '▼'`, true, []types.Type{types.Noop})
}

func TestLiteralStringWithFiltersYieldsNoAcInString(t *testing.T) {
	expected := Result{
		Items: nil,
		Range: Range{
			From: LineCol{Line: 1, Col: 7},
			To:   LineCol{Line: 1, Col: 10},
		},
	}
	checkCompletion(t, `RETURN ":▼"`, expected, true)
	checkCompletion(t, `RETURN ':▼'`, expected, true)
}

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
	"github.com/antlr/antlr4/runtime/Go/antlr"
	. "github.com/stretchr/testify/require"
)

func TestGetElementIdentifyRuleAtCursor(t *testing.T) {
	s := editor.NewEditorSupport("MATCH (n)-[r]->(n) RETURN n")
	tree := s.GetElement(1, 12).GetParent()
	ctx := tree.(*antlr.BaseParserRuleContext)
	Equal(t, "[r]", ctx.GetText())
	Equal(t, 10, ctx.GetStart().GetColumn())
	Equal(t, 12, ctx.GetStop().GetColumn())
}

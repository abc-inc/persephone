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

package ref_test

import (
	"reflect"
	"testing"

	"github.com/abc-inc/persephone/editor"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	. "github.com/stretchr/testify/require"
)

func TestRefLabelsReturnsReferenceForSingleLabel(t *testing.T) {
	e := editor.NewEditorSupport("MATCH (n:Label)")
	refs := e.GetReferences(1, 10)

	ref := refs[0]
	Equal(t, "LabelNameContext", reflect.TypeOf(ref).Elem().Name())
	Equal(t, 9, ref.GetStart().GetStart())
	Equal(t, 1, ref.GetStart().GetLine())
	Equal(t, 13, ref.GetStop().GetStop())
	Equal(t, 1, ref.GetStop().GetLine())
	Equal(t, "Label", ref.GetText())
}

func TestRefLabelsReturnsReferencesForMultipleLabels(t *testing.T) {
	e := editor.NewEditorSupport("MATCH (n:Label) MATCH (m:Label)")
	refs := e.GetReferences(1, 10)

	ref := refs[0].(antlr.ParserRuleContext)
	Equal(t, "LabelNameContext", reflect.TypeOf(ref).Elem().Name())
	Equal(t, 9, ref.GetStart().GetStart())
	Equal(t, 1, ref.GetStart().GetLine())
	Equal(t, 13, ref.GetStop().GetStop())
	Equal(t, 1, ref.GetStop().GetLine())
	Equal(t, "Label", ref.GetText())

	ref = refs[1].(antlr.ParserRuleContext)
	Equal(t, "LabelNameContext", reflect.TypeOf(ref).Elem().Name())
	Equal(t, 25, ref.GetStart().GetStart())
	Equal(t, 1, ref.GetStart().GetLine())
	Equal(t, 29, ref.GetStop().GetStop())
	Equal(t, 1, ref.GetStop().GetLine())
	Equal(t, "Label", ref.GetText())
}

func TestRefLabelsReturnsReferencesForMultipleQueries(t *testing.T) {
	e := editor.NewEditorSupport("MATCH (n:Label); MATCH (n:Label);")
	refs := e.GetReferences(1, 10)

	ref := refs[0].(antlr.ParserRuleContext)
	Equal(t, "LabelNameContext", reflect.TypeOf(ref).Elem().Name())
	Equal(t, 9, ref.GetStart().GetStart())
	Equal(t, 1, ref.GetStart().GetLine())
	Equal(t, 13, ref.GetStop().GetStop())
	Equal(t, 1, ref.GetStop().GetLine())
	Equal(t, "Label", ref.GetText())

	ref = refs[1].(antlr.ParserRuleContext)
	Equal(t, "LabelNameContext", reflect.TypeOf(ref).Elem().Name())
	Equal(t, 26, ref.GetStart().GetStart())
	Equal(t, 1, ref.GetStart().GetLine())
	Equal(t, 30, ref.GetStop().GetStop())
	Equal(t, 1, ref.GetStop().GetLine())
	Equal(t, "Label", ref.GetText())
}

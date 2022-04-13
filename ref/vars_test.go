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
	"github.com/abc-inc/persephone/lang"
	. "github.com/stretchr/testify/require"
)

func TestVariablesReturnsReferenceForSingleVariable(t *testing.T) {
	es := editor.NewEditorSupport("RETURN n")
	refs := es.GetReferences(1, 7)

	ref := refs[0]
	Equal(t, lang.VARIABLE_CONTEXT, reflect.TypeOf(ref).Elem().Name())
	Equal(t, 7, ref.GetStart().GetStart())
	Equal(t, 1, ref.GetStart().GetLine())
	Equal(t, 7, ref.GetStop().GetStop())
	Equal(t, 1, ref.GetStop().GetLine())
	Equal(t, "n", ref.GetText())
}

func TestVariablesReturnsReferenceForMultipleVariables(t *testing.T) {
	es := editor.NewEditorSupport("MATCH (n)-[r]->(n) RETURN n")
	refs := es.GetReferences(1, 7)

	ref := refs[0]
	Equal(t, lang.VARIABLE_CONTEXT, reflect.TypeOf(ref).Elem().Name())
	Equal(t, 7, ref.GetStart().GetStart())
	Equal(t, 1, ref.GetStart().GetLine())
	Equal(t, 7, ref.GetStop().GetStop())
	Equal(t, 1, ref.GetStop().GetLine())
	Equal(t, "n", ref.GetText())
	ref = refs[1]
	Equal(t, lang.VARIABLE_CONTEXT, reflect.TypeOf(ref).Elem().Name())
	Equal(t, 16, ref.GetStart().GetStart())
	Equal(t, 1, ref.GetStart().GetLine())
	Equal(t, 16, ref.GetStop().GetStop())
	Equal(t, 1, ref.GetStop().GetLine())
	Equal(t, "n", ref.GetText())

	ref = refs[2]
	Equal(t, lang.VARIABLE_CONTEXT, reflect.TypeOf(ref).Elem().Name())
	Equal(t, 26, ref.GetStart().GetStart())
	Equal(t, 1, ref.GetStart().GetLine())
	Equal(t, 26, ref.GetStop().GetStop())
	Equal(t, 1, ref.GetStop().GetLine())
	Equal(t, "n", ref.GetText())
}

func TestVariablesReturnsReferenceForMultipleQueries(t *testing.T) {
	es := editor.NewEditorSupport("MATCH (n) RETURN n; MATCH (n) RETURN n")
	refs := es.GetReferences(1, 7)

	ref := refs[0]
	Equal(t, lang.VARIABLE_CONTEXT, reflect.TypeOf(ref).Elem().Name())
	Equal(t, 7, ref.GetStart().GetStart())
	Equal(t, 1, ref.GetStart().GetLine())
	Equal(t, 7, ref.GetStop().GetStop())
	Equal(t, 1, ref.GetStop().GetLine())
	Equal(t, "n", ref.GetText())
	ref = refs[1]
	Equal(t, lang.VARIABLE_CONTEXT, reflect.TypeOf(ref).Elem().Name())
	Equal(t, 17, ref.GetStart().GetStart())
	Equal(t, 1, ref.GetStart().GetLine())
	Equal(t, 17, ref.GetStop().GetStop())
	Equal(t, 1, ref.GetStop().GetLine())
	Equal(t, "n", ref.GetText())
}

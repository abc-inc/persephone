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

package ref_test

import (
	"reflect"
	"testing"

	"github.com/abc-inc/persephone/editor"
	"github.com/abc-inc/persephone/lang"
	. "github.com/stretchr/testify/require"
)

func TestPropertyKeysReturnsReferenceForSingleKey(t *testing.T) {
	e := editor.NewEditor("RETURN n.key")
	refs := e.GetReferences(1, 10)

	ref := refs[0]
	Equal(t, lang.PropertyKeyNameContext, reflect.TypeOf(ref).Elem().Name())
	Equal(t, 9, ref.GetStart().GetStart())
	Equal(t, 1, ref.GetStart().GetLine())
	Equal(t, 11, ref.GetStop().GetStop())
	Equal(t, 1, ref.GetStop().GetLine())
	Equal(t, "key", ref.GetText())
}

func TestPropertyKeysReturnsReferencesForMultipleKeys(t *testing.T) {
	e := editor.NewEditor("MATCH (n {key: 42}) SET n.key = 4 RETURN n.key;")
	refs := e.GetReferences(1, 10)

	ref := refs[0]
	Equal(t, lang.PropertyKeyNameContext, reflect.TypeOf(ref).Elem().Name())
	Equal(t, 10, ref.GetStart().GetStart())
	Equal(t, 1, ref.GetStart().GetLine())
	Equal(t, 12, ref.GetStop().GetStop())
	Equal(t, 1, ref.GetStop().GetLine())
	Equal(t, "key", ref.GetText())

	ref = refs[1]
	Equal(t, lang.PropertyKeyNameContext, reflect.TypeOf(ref).Elem().Name())
	Equal(t, 26, ref.GetStart().GetStart())
	Equal(t, 1, ref.GetStart().GetLine())
	Equal(t, 28, ref.GetStop().GetStop())
	Equal(t, 1, ref.GetStop().GetLine())
	Equal(t, "key", ref.GetText())

	ref = refs[2]
	Equal(t, lang.PropertyKeyNameContext, reflect.TypeOf(ref).Elem().Name())
	Equal(t, 43, ref.GetStart().GetStart())
	Equal(t, 1, ref.GetStart().GetLine())
	Equal(t, 45, ref.GetStop().GetStop())
	Equal(t, 1, ref.GetStop().GetLine())
	Equal(t, "key", ref.GetText())
}

func TestPropertyKeysReturnsReferencesForMultipleQueries(t *testing.T) {
	e := editor.NewEditor("MATCH (n {key: 42})\n" +
		"          SET n.key = 42\n" +
		"          RETURN n.key;\n" +
		"          MATCH (n {key: 42})\n" +
		"          SET n.key = 42\n" +
		"          RETURN n.key")

	refs := e.GetReferences(1, 10)

	ref := refs[0]
	Equal(t, lang.PropertyKeyNameContext, reflect.TypeOf(ref).Elem().Name())
	Equal(t, 10, ref.GetStart().GetStart())
	Equal(t, 1, ref.GetStart().GetLine())
	Equal(t, 12, ref.GetStop().GetStop())
	Equal(t, 1, ref.GetStop().GetLine())
	Equal(t, "key", ref.GetText())

	ref = refs[1]
	Equal(t, lang.PropertyKeyNameContext, reflect.TypeOf(ref).Elem().Name())
	Equal(t, 36, ref.GetStart().GetStart())
	Equal(t, 2, ref.GetStart().GetLine())
	Equal(t, 38, ref.GetStop().GetStop())
	Equal(t, 2, ref.GetStop().GetLine())
	Equal(t, "key", ref.GetText())

	ref = refs[2]
	Equal(t, lang.PropertyKeyNameContext, reflect.TypeOf(ref).Elem().Name())
	Equal(t, 64, ref.GetStart().GetStart())
	Equal(t, 3, ref.GetStart().GetLine())
	Equal(t, 66, ref.GetStop().GetStop())
	Equal(t, 3, ref.GetStop().GetLine())
	Equal(t, "key", ref.GetText())

	ref = refs[3]
	Equal(t, lang.PropertyKeyNameContext, reflect.TypeOf(ref).Elem().Name())
	Equal(t, 89, ref.GetStart().GetStart())
	Equal(t, 4, ref.GetStart().GetLine())
	Equal(t, 91, ref.GetStop().GetStop())
	Equal(t, 4, ref.GetStop().GetLine())
	Equal(t, "key", ref.GetText())

	ref = refs[4]
	Equal(t, lang.PropertyKeyNameContext, reflect.TypeOf(ref).Elem().Name())
	Equal(t, 115, ref.GetStart().GetStart())
	Equal(t, 5, ref.GetStart().GetLine())
	Equal(t, 117, ref.GetStop().GetStop())
	Equal(t, 5, ref.GetStop().GetLine())
	Equal(t, "key", ref.GetText())

	ref = refs[5]
	Equal(t, lang.PropertyKeyNameContext, reflect.TypeOf(ref).Elem().Name())
	Equal(t, 143, ref.GetStart().GetStart())
	Equal(t, 6, ref.GetStart().GetLine())
	Equal(t, 145, ref.GetStop().GetStop())
	Equal(t, 6, ref.GetStop().GetLine())
	Equal(t, "key", ref.GetText())
}

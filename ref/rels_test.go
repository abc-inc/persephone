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

func TestRelationshipsReturnsReferenceForRelationshipTypes(t *testing.T) {
	e := editor.NewEditor("MATCH ()-[:TYPE]-()")
	refs := e.GetReferences(1, 13)

	ref := refs[0]
	Equal(t, lang.RelationshipTypeNameContext, reflect.TypeOf(ref).Elem().Name())
	Equal(t, 11, ref.GetStart().GetStart())
	Equal(t, 1, ref.GetStart().GetLine())
	Equal(t, 14, ref.GetStop().GetStop())
	Equal(t, 1, ref.GetStop().GetLine())
	Equal(t, "TYPE", ref.GetText())
}

func TestRelationshipsReturnsReferencesForMultipleRelationshipTypes(t *testing.T) {
	e := editor.NewEditor("MATCH ()-[:TYPE]-() MATCH ()-[:TYPE]-()")
	refs := e.GetReferences(1, 13)

	ref := refs[0]
	Equal(t, lang.RelationshipTypeNameContext, reflect.TypeOf(ref).Elem().Name())
	Equal(t, 11, ref.GetStart().GetStart())
	Equal(t, 1, ref.GetStart().GetLine())
	Equal(t, 14, ref.GetStop().GetStop())
	Equal(t, 1, ref.GetStop().GetLine())
	Equal(t, "TYPE", ref.GetText())

	ref = refs[1]
	Equal(t, lang.RelationshipTypeNameContext, reflect.TypeOf(ref).Elem().Name())
	Equal(t, 31, ref.GetStart().GetStart())
	Equal(t, 1, ref.GetStart().GetLine())
	Equal(t, 34, ref.GetStop().GetStop())
	Equal(t, 1, ref.GetStop().GetLine())
	Equal(t, "TYPE", ref.GetText())
}

func TestRelationshipsReturnsReferencesForMultipleQueries(t *testing.T) {
	e := editor.NewEditor("MATCH ()-[:TYPE]-(); MATCH ()-[:TYPE]-()")
	refs := e.GetReferences(1, 13)

	ref := refs[0]
	Equal(t, lang.RelationshipTypeNameContext, reflect.TypeOf(ref).Elem().Name())
	Equal(t, 11, ref.GetStart().GetStart())
	Equal(t, 1, ref.GetStart().GetLine())
	Equal(t, 14, ref.GetStop().GetStop())
	Equal(t, 1, ref.GetStop().GetLine())
	Equal(t, "TYPE", ref.GetText())

	ref = refs[1]
	Equal(t, lang.RelationshipTypeNameContext, reflect.TypeOf(ref).Elem().Name())
	Equal(t, 32, ref.GetStart().GetStart())
	Equal(t, 1, ref.GetStart().GetLine())
	Equal(t, 35, ref.GetStop().GetStop())
	Equal(t, 1, ref.GetStop().GetLine())
	Equal(t, "TYPE", ref.GetText())
}

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

func TestParametersReturnsReferenceForSingleParameter(t *testing.T) {
	e := editor.NewEditor("RETURN $param;")
	refs := e.GetReferences(1, 10)

	ref := refs[0]
	Equal(t, lang.ParameterNameContext, reflect.TypeOf(ref).Elem().Name())
	Equal(t, 8, ref.GetStart().GetStart())
	Equal(t, 1, ref.GetStart().GetLine())
	Equal(t, 12, ref.GetStop().GetStop())
	Equal(t, 1, ref.GetStop().GetLine())
	Equal(t, "param", ref.GetText())
}

func TestParametersReturnsReferenceForMultipleParameters(t *testing.T) {
	e := editor.NewEditor("MATCH (n) SET n.key = $param SET n.key = {param}")
	refs := e.GetReferences(1, 45)

	ref := refs[0]
	Equal(t, lang.ParameterNameContext, reflect.TypeOf(ref).Elem().Name())
	Equal(t, 23, ref.GetStart().GetStart())
	Equal(t, 1, ref.GetStart().GetLine())
	Equal(t, 27, ref.GetStop().GetStop())
	Equal(t, 1, ref.GetStop().GetLine())
	Equal(t, "param", ref.GetText())

	ref = refs[1]
	Equal(t, lang.ParameterNameContext, reflect.TypeOf(ref).Elem().Name())
	Equal(t, 42, ref.GetStart().GetStart())
	Equal(t, 1, ref.GetStart().GetLine())
	Equal(t, 46, ref.GetStop().GetStop())
	Equal(t, 1, ref.GetStop().GetLine())
	Equal(t, "param", ref.GetText())
}

func TestParametersReturnsReferencesForMultipleQueries(t *testing.T) {
	e := editor.NewEditor("MATCH (n) SET n.key = $param SET n.key = {param};\n" +
		"          MATCH (n) SET n.key = $param SET n.key = {param};")
	refs := e.GetReferences(1, 45)

	ref := refs[0]
	Equal(t, lang.ParameterNameContext, reflect.TypeOf(ref).Elem().Name())
	Equal(t, 23, ref.GetStart().GetStart())
	Equal(t, 1, ref.GetStart().GetLine())
	Equal(t, 27, ref.GetStop().GetStop())
	Equal(t, 1, ref.GetStop().GetLine())
	Equal(t, "param", ref.GetText())

	ref = refs[1]
	Equal(t, lang.ParameterNameContext, reflect.TypeOf(ref).Elem().Name())
	Equal(t, 42, ref.GetStart().GetStart())
	Equal(t, 1, ref.GetStart().GetLine())
	Equal(t, 46, ref.GetStop().GetStop())
	Equal(t, 1, ref.GetStop().GetLine())
	Equal(t, "param", ref.GetText())

	ref = refs[2]
	Equal(t, lang.ParameterNameContext, reflect.TypeOf(ref).Elem().Name())
	Equal(t, 83, ref.GetStart().GetStart())
	Equal(t, 2, ref.GetStart().GetLine())
	Equal(t, 87, ref.GetStop().GetStop())
	Equal(t, 2, ref.GetStop().GetLine())
	Equal(t, "param", ref.GetText())

	ref = refs[3]
	Equal(t, lang.ParameterNameContext, reflect.TypeOf(ref).Elem().Name())
	Equal(t, 102, ref.GetStart().GetStart())
	Equal(t, 2, ref.GetStart().GetLine())
	Equal(t, 106, ref.GetStop().GetStop())
	Equal(t, 2, ref.GetStop().GetLine())
	Equal(t, "param", ref.GetText())
}

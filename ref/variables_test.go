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

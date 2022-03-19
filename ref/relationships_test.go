package ref_test

import (
	"reflect"
	"testing"

	"github.com/abc-inc/merovingian/editor"
	"github.com/abc-inc/merovingian/lang"
	. "github.com/stretchr/testify/assert"
)

func TestRelationshipsReturnsReferenceForRelationshipTypes(t *testing.T) {
	es := editor.NewEditorSupport("MATCH ()-[:TYPE]-()")
	refs := es.GetReferences(1, 13)

	ref := refs[0]
	Equal(t, lang.RELATIONSHIP_TYPE_NAME_CONTEXT, reflect.TypeOf(ref).Elem().Name())
	Equal(t, 11, ref.GetStart().GetStart())
	Equal(t, 1, ref.GetStart().GetLine())
	Equal(t, 14, ref.GetStop().GetStop())
	Equal(t, 1, ref.GetStop().GetLine())
	Equal(t, "TYPE", ref.GetText())
}

func TestRelationshipsReturnsReferencesForMultipleRelationshipTypes(t *testing.T) {
	es := editor.NewEditorSupport("MATCH ()-[:TYPE]-() MATCH ()-[:TYPE]-()")
	refs := es.GetReferences(1, 13)

	ref := refs[0]
	Equal(t, lang.RELATIONSHIP_TYPE_NAME_CONTEXT, reflect.TypeOf(ref).Elem().Name())
	Equal(t, 11, ref.GetStart().GetStart())
	Equal(t, 1, ref.GetStart().GetLine())
	Equal(t, 14, ref.GetStop().GetStop())
	Equal(t, 1, ref.GetStop().GetLine())
	Equal(t, "TYPE", ref.GetText())

	ref = refs[1]
	Equal(t, lang.RELATIONSHIP_TYPE_NAME_CONTEXT, reflect.TypeOf(ref).Elem().Name())
	Equal(t, 31, ref.GetStart().GetStart())
	Equal(t, 1, ref.GetStart().GetLine())
	Equal(t, 34, ref.GetStop().GetStop())
	Equal(t, 1, ref.GetStop().GetLine())
	Equal(t, "TYPE", ref.GetText())
}

func TestRelationshipsReturnsReferencesForMultipleQueries(t *testing.T) {
	es := editor.NewEditorSupport("MATCH ()-[:TYPE]-(); MATCH ()-[:TYPE]-()")
	refs := es.GetReferences(1, 13)

	ref := refs[0]
	Equal(t, lang.RELATIONSHIP_TYPE_NAME_CONTEXT, reflect.TypeOf(ref).Elem().Name())
	Equal(t, 11, ref.GetStart().GetStart())
	Equal(t, 1, ref.GetStart().GetLine())
	Equal(t, 14, ref.GetStop().GetStop())
	Equal(t, 1, ref.GetStop().GetLine())
	Equal(t, "TYPE", ref.GetText())

	ref = refs[1]
	Equal(t, lang.RELATIONSHIP_TYPE_NAME_CONTEXT, reflect.TypeOf(ref).Elem().Name())
	Equal(t, 32, ref.GetStart().GetStart())
	Equal(t, 1, ref.GetStart().GetLine())
	Equal(t, 35, ref.GetStop().GetStop())
	Equal(t, 1, ref.GetStop().GetLine())
	Equal(t, "TYPE", ref.GetText())
}

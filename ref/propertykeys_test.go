package ref

import (
	"reflect"
	"testing"

	"github.com/abc-inc/merovingian/cypher"
	"github.com/abc-inc/merovingian/lang"
	. "github.com/stretchr/testify/assert"
)

func TestPropertyKeysReturnsReferenceForSingleKey(t *testing.T) {
	es := cypher.NewEditorSupport("RETURN n.key")
	refs := es.GetReferences(1, 10)

	ref := refs[0]
	Equal(t, lang.PROPERTY_KEY_NAME_CONTEXT, reflect.TypeOf(ref).Elem().Name())
	Equal(t, 9, ref.GetStart().GetStart())
	Equal(t, 1, ref.GetStart().GetLine())
	Equal(t, 11, ref.GetStop().GetStop())
	Equal(t, 1, ref.GetStop().GetLine())
	Equal(t, "key", ref.GetText())
}

func TestPropertyKeysReturnsReferencesForMultipleKeys(t *testing.T) {
	es := cypher.NewEditorSupport("MATCH (n {key: 42}) SET n.key = 4 RETURN n.key;")
	refs := es.GetReferences(1, 10)

	ref := refs[0]
	Equal(t, lang.PROPERTY_KEY_NAME_CONTEXT, reflect.TypeOf(ref).Elem().Name())
	Equal(t, 10, ref.GetStart().GetStart())
	Equal(t, 1, ref.GetStart().GetLine())
	Equal(t, 12, ref.GetStop().GetStop())
	Equal(t, 1, ref.GetStop().GetLine())
	Equal(t, "key", ref.GetText())

	ref = refs[1]
	Equal(t, lang.PROPERTY_KEY_NAME_CONTEXT, reflect.TypeOf(ref).Elem().Name())
	Equal(t, 26, ref.GetStart().GetStart())
	Equal(t, 1, ref.GetStart().GetLine())
	Equal(t, 28, ref.GetStop().GetStop())
	Equal(t, 1, ref.GetStop().GetLine())
	Equal(t, "key", ref.GetText())

	ref = refs[2]
	Equal(t, lang.PROPERTY_KEY_NAME_CONTEXT, reflect.TypeOf(ref).Elem().Name())
	Equal(t, 43, ref.GetStart().GetStart())
	Equal(t, 1, ref.GetStart().GetLine())
	Equal(t, 45, ref.GetStop().GetStop())
	Equal(t, 1, ref.GetStop().GetLine())
	Equal(t, "key", ref.GetText())
}

func TestPropertyKeysReturnsReferencesForMultipleQueries(t *testing.T) {
	es := cypher.NewEditorSupport("MATCH (n {key: 42})\n" +
		"SET n.key = 42\n" +
		"RETURN n.key;\n" +
		"MATCH (n {key: 42})\n" +
		"SET n.key = 42\n" +
		"RETURN n.key")

	refs := es.GetReferences(1, 10)

	ref := refs[0]
	Equal(t, lang.PROPERTY_KEY_NAME_CONTEXT, reflect.TypeOf(ref).Elem().Name())
	Equal(t, 10, ref.GetStart().GetStart())
	Equal(t, 1, ref.GetStart().GetLine())
	Equal(t, 12, ref.GetStop().GetStop())
	Equal(t, 1, ref.GetStop().GetLine())
	Equal(t, "key", ref.GetText())

	ref = refs[1]
	Equal(t, lang.PROPERTY_KEY_NAME_CONTEXT, reflect.TypeOf(ref).Elem().Name())
	Equal(t, 36, ref.GetStart().GetStart())
	Equal(t, 2, ref.GetStart().GetLine())
	Equal(t, 38, ref.GetStop().GetStop())
	Equal(t, 2, ref.GetStop().GetLine())
	Equal(t, "key", ref.GetText())

	ref = refs[2]
	Equal(t, lang.PROPERTY_KEY_NAME_CONTEXT, reflect.TypeOf(ref).Elem().Name())
	Equal(t, 64, ref.GetStart().GetStart())
	Equal(t, 3, ref.GetStart().GetLine())
	Equal(t, 66, ref.GetStop().GetStop())
	Equal(t, 3, ref.GetStop().GetLine())
	Equal(t, "key", ref.GetText())

	ref = refs[3]
	Equal(t, lang.PROPERTY_KEY_NAME_CONTEXT, reflect.TypeOf(ref).Elem().Name())
	Equal(t, 89, ref.GetStart().GetStart())
	Equal(t, 4, ref.GetStart().GetLine())
	Equal(t, 91, ref.GetStop().GetStop())
	Equal(t, 4, ref.GetStop().GetLine())
	Equal(t, "key", ref.GetText())

	ref = refs[4]
	Equal(t, lang.PROPERTY_KEY_NAME_CONTEXT, reflect.TypeOf(ref).Elem().Name())
	Equal(t, 115, ref.GetStart().GetStart())
	Equal(t, 5, ref.GetStart().GetLine())
	Equal(t, 117, ref.GetStop().GetStop())
	Equal(t, 5, ref.GetStop().GetLine())
	Equal(t, "key", ref.GetText())

	ref = refs[5]
	Equal(t, lang.PROPERTY_KEY_NAME_CONTEXT, reflect.TypeOf(ref).Elem().Name())
	Equal(t, 143, ref.GetStart().GetStart())
	Equal(t, 6, ref.GetStart().GetLine())
	Equal(t, 145, ref.GetStop().GetStop())
	Equal(t, 6, ref.GetStop().GetLine())
	Equal(t, "key", ref.GetText())
}

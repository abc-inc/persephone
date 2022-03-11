package ref

import (
	"reflect"
	"testing"

	"github.com/abc-inc/merovingian/cypher"
	"github.com/abc-inc/merovingian/lang"
	. "github.com/stretchr/testify/assert"
)

func TestParametersReturnsReferenceForSingleParameter(t *testing.T) {
	es := cypher.NewEditorSupport("RETURN $param;")
	refs := es.GetReferences(1, 10)

	ref := refs[0]
	Equal(t, lang.PARAMETER_NAME_CONTEXT, reflect.TypeOf(ref).Elem().Name())
	Equal(t, 8, ref.GetStart().GetStart())
	Equal(t, 1, ref.GetStart().GetLine())
	Equal(t, 12, ref.GetStop().GetStop())
	Equal(t, 1, ref.GetStop().GetLine())
	Equal(t, "param", ref.GetText())
}

func TestParametersReturnsReferenceForMultipleParameters(t *testing.T) {
	es := cypher.NewEditorSupport("MATCH (n) SET n.key = $param SET n.key = {param}")
	refs := es.GetReferences(1, 45)

	ref := refs[0]
	Equal(t, lang.PARAMETER_NAME_CONTEXT, reflect.TypeOf(ref).Elem().Name())
	Equal(t, 23, ref.GetStart().GetStart())
	Equal(t, 1, ref.GetStart().GetLine())
	Equal(t, 27, ref.GetStop().GetStop())
	Equal(t, 1, ref.GetStop().GetLine())
	Equal(t, "param", ref.GetText())

	ref = refs[1]
	Equal(t, lang.PARAMETER_NAME_CONTEXT, reflect.TypeOf(ref).Elem().Name())
	Equal(t, 42, ref.GetStart().GetStart())
	Equal(t, 1, ref.GetStart().GetLine())
	Equal(t, 46, ref.GetStop().GetStop())
	Equal(t, 1, ref.GetStop().GetLine())
	Equal(t, "param", ref.GetText())
}

func TestParametersReturnsReferencesForMultipleQueries(t *testing.T) {
	es := cypher.NewEditorSupport("MATCH (n) SET n.key = $param SET n.key = {param};\n" +
		"          MATCH (n) SET n.key = $param SET n.key = {param};")
	refs := es.GetReferences(1, 45)

	ref := refs[0]
	Equal(t, lang.PARAMETER_NAME_CONTEXT, reflect.TypeOf(ref).Elem().Name())
	Equal(t, 23, ref.GetStart().GetStart())
	Equal(t, 1, ref.GetStart().GetLine())
	Equal(t, 27, ref.GetStop().GetStop())
	Equal(t, 1, ref.GetStop().GetLine())
	Equal(t, "param", ref.GetText())

	ref = refs[1]
	Equal(t, lang.PARAMETER_NAME_CONTEXT, reflect.TypeOf(ref).Elem().Name())
	Equal(t, 42, ref.GetStart().GetStart())
	Equal(t, 1, ref.GetStart().GetLine())
	Equal(t, 46, ref.GetStop().GetStop())
	Equal(t, 1, ref.GetStop().GetLine())
	Equal(t, "param", ref.GetText())

	ref = refs[2]
	Equal(t, lang.PARAMETER_NAME_CONTEXT, reflect.TypeOf(ref).Elem().Name())
	Equal(t, 83, ref.GetStart().GetStart())
	Equal(t, 2, ref.GetStart().GetLine())
	Equal(t, 87, ref.GetStop().GetStop())
	Equal(t, 2, ref.GetStop().GetLine())
	Equal(t, "param", ref.GetText())

	ref = refs[3]
	Equal(t, lang.PARAMETER_NAME_CONTEXT, reflect.TypeOf(ref).Elem().Name())
	Equal(t, 102, ref.GetStart().GetStart())
	Equal(t, 2, ref.GetStart().GetLine())
	Equal(t, 106, ref.GetStop().GetStop())
	Equal(t, 2, ref.GetStop().GetLine())
	Equal(t, "param", ref.GetText())
}

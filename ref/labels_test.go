package ref

import (
	"reflect"
	"testing"

	"github.com/abc-inc/merovingian/cypher"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	. "github.com/stretchr/testify/assert"
)

func TestRefLabelsReturnsReferenceForSingleLabel(t *testing.T) {
	e := cypher.NewEditorSupport("MATCH (n:Label)")
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
	e := cypher.NewEditorSupport("MATCH (n:Label) MATCH (m:Label)")
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
	e := cypher.NewEditorSupport("MATCH (n:Label); MATCH (n:Label);")
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

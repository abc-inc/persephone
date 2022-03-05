package cypher

import (
	"testing"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	. "github.com/stretchr/testify/require"
)

func TestGetElementIdentifyRuleAtCursor(t *testing.T) {
	s:=NewEditorSupport("MATCH (n)-[r]->(n) RETURN n")
	tree := s.GetElement(1, 12).GetParent()
	ctx := tree.(*antlr.BaseParserRuleContext)
	Equal(t, "[r]", ctx.GetText())
	Equal(t, 10, ctx.GetStart().GetColumn())
	Equal(t, 12, ctx.GetStop().GetColumn())
}
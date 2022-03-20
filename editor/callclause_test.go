package editor_test

import (
	"testing"

	"github.com/abc-inc/persephone/editor"
	. "github.com/stretchr/testify/require"
)

func TestCallClause(t *testing.T) {
	tests := []struct {
		name   string
		cypher string
	}{
		{name: "with parenthesis", cypher: "CALL procedure()"},
		{name: "without parenthesis", cypher: "CALL procedure"},
		{name: "with where after yield", cypher: "CALL procedure() YIELD name WHERE true RETURN name"},
		{name: "string contains cypher", cypher: `CALL foo.bar("RETURN 1")`},
		{name: "string contains cypher with new lines", cypher: `CALL foo.bar("MATCH (n) \nRETURN n")`},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			s := editor.NewEditorSupport(test.cypher)
			Nil(t, s.ParseErrors)
		})
	}
}

func TestCallClausMismatchedParenthesis(t *testing.T) {
	s := editor.NewEditorSupport("CALL ()")
	Equal(t, 1, len(s.ParseErrors))
	Equal(t, 5, s.ParseErrors[0].Col)
}

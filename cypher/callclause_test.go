package cypher

import (
	"testing"

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
			s := NewEditorSupport(test.cypher)
			Nil(t, s.parseErrors)
		})
	}
}

func TestCallClausMismatchedParenthesis(t *testing.T) {
	s := NewEditorSupport("CALL ()")
	Equal(t, 1, len(s.parseErrors))
	Equal(t, 5, s.parseErrors[0].Col)
}

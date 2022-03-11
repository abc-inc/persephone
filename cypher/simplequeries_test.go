package cypher

import (
	"testing"

	. "github.com/stretchr/testify/require"
)

func TestCorrectASTForSimpleQuery(t *testing.T) {
	s := NewEditorSupport("RETURN 42;")
	Nil(t, s.parseErrors)
}

func TestErrorsForIncorrectQuery(t *testing.T) {
	msg1 := "mismatched input 'POTATO' expecting {':', CYPHER, EXPLAIN, PROFILE, USING, CREATE, DROP, LOAD, WITH, OPTIONAL, MATCH, UNWIND, MERGE, SET, DETACH, DELETE, REMOVE, FOREACH, RETURN, START, CALL, SP}"
	s := NewEditorSupport("POTATO")
	Equal(t, 1, len(s.parseErrors))
	Equal(t, SynErr{1, 0, msg1}, s.parseErrors[0])
	Equal(t, "", s.parseTree)
}

func TestErrorsIfErrorInLexer(t *testing.T) {
	msg1 := "mismatched input '`' expecting {<EOF>, ';'}"
	s := NewEditorSupport("WITH a` WITH 1;")
	Equal(t, 1, len(s.parseErrors))
	Equal(t, SynErr{1, 6, msg1}, s.parseErrors[0])
}

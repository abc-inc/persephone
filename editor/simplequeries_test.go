package editor_test

import (
	"testing"

	"github.com/abc-inc/merovingian/editor"
	. "github.com/stretchr/testify/require"
)

func TestCorrectASTForSimpleQuery(t *testing.T) {
	s := editor.NewEditorSupport("RETURN 42;")
	Nil(t, s.ParseErrors)
}

func TestErrorsForIncorrectQuery(t *testing.T) {
	msg1 := "mismatched input 'POTATO' expecting {':', CYPHER, EXPLAIN, PROFILE, USING, CREATE, DROP, LOAD, WITH, OPTIONAL, MATCH, UNWIND, MERGE, SET, DETACH, DELETE, REMOVE, FOREACH, RETURN, START, CALL, SP}"
	s := editor.NewEditorSupport("POTATO")
	Equal(t, 1, len(s.ParseErrors))
	Equal(t, editor.SynErr{1, 0, msg1}, s.ParseErrors[0])
	Equal(t, "", s.ParseTree)
}

func TestErrorsIfErrorInLexer(t *testing.T) {
	msg1 := "mismatched input '`' expecting {<EOF>, ';'}"
	s := editor.NewEditorSupport("WITH a` WITH 1;")
	Equal(t, 1, len(s.ParseErrors))
	Equal(t, editor.SynErr{1, 6, msg1}, s.ParseErrors[0])
}

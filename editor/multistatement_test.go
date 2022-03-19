package editor_test

import (
	"testing"

	"github.com/abc-inc/merovingian/editor"
	. "github.com/stretchr/testify/require"
)

func TestCatchErrorInSecondStatement(t *testing.T) {
	msg1 := "mismatched input 'POTATO' expecting {':', CYPHER, EXPLAIN, PROFILE, USING, CREATE, DROP, LOAD, WITH, OPTIONAL, MATCH, UNWIND, MERGE, SET, DETACH, DELETE, REMOVE, FOREACH, RETURN, START, CALL}"
	cypher := `RETURN 1;
POTATO;
RETURN rand();`

	s := editor.NewEditorSupport(cypher)
	Equal(t, 1, len(s.ParseErrors))
	Equal(t, editor.SynErr{2, 0, msg1}, s.ParseErrors[0])
}

func TestParseCommonParam(t *testing.T) {
	msg1 := "mismatched input 'hello' expecting {':', CYPHER, EXPLAIN, PROFILE, USING, CREATE, DROP, LOAD, WITH, OPTIONAL, MATCH, UNWIND, MERGE, SET, DETACH, DELETE, REMOVE, FOREACH, RETURN, START, CALL}"
	msg2 := "mismatched input 'hello2' expecting {':', CYPHER, EXPLAIN, PROFILE, USING, CREATE, DROP, LOAD, WITH, OPTIONAL, MATCH, UNWIND, MERGE, SET, DETACH, DELETE, REMOVE, FOREACH, RETURN, START, CALL}"
	cypher := `:play;
hello;
:param x => 1;
hello2;
:play reco;`

	s := editor.NewEditorSupport(cypher)
	Equal(t, 2, len(s.ParseErrors))
	Equal(t, editor.SynErr{2, 0, msg1}, s.ParseErrors[0])
	Equal(t, editor.SynErr{4, 0, msg2}, s.ParseErrors[1])
}

func TestParseCommon(t *testing.T) {
	msg1 := "mismatched input 'hello' expecting {':', CYPHER, EXPLAIN, PROFILE, USING, CREATE, DROP, LOAD, WITH, OPTIONAL, MATCH, UNWIND, MERGE, SET, DETACH, DELETE, REMOVE, FOREACH, RETURN, START, CALL, SP}"
	msg2 := "mismatched input 'hello2' expecting {':', CYPHER, EXPLAIN, PROFILE, USING, CREATE, DROP, LOAD, WITH, OPTIONAL, MATCH, UNWIND, MERGE, SET, DETACH, DELETE, REMOVE, FOREACH, RETURN, START, CALL}"
	cypher := `hello;
:param x => 1;
hello2;
:play reco;`

	s := editor.NewEditorSupport(cypher)
	Equal(t, 2, len(s.ParseErrors))
	Equal(t, editor.SynErr{1, 0, msg1}, s.ParseErrors[0])
	Equal(t, editor.SynErr{3, 0, msg2}, s.ParseErrors[1])
}

func TestParseCommonParamCommand(t *testing.T) {
	cypher := `:play http://guides.neo4j.com/reco;
:param x => 1;
RETURN $x;
:play reco;`

	s := editor.NewEditorSupport(cypher)
	Nil(t, s.ParseErrors)
}

func TestParseMultipleParamCommandsWithQuery(t *testing.T) {
	cypher := `:param age => 25;
:param interests => ['football', 'fishing'];
MATCH (n)
WHERE n.age > $age
AND n.interest IN $interests
RETURN n;`

	s := editor.NewEditorSupport(cypher)
	Nil(t, s.ParseErrors)
}

func TestRecoverToSecondStatementAfterInvalidCommand(t *testing.T) {
	cypher := ":PUT ao*51 fagas 8(!; :play;"
	s := editor.NewEditorSupport(cypher)
	Equal(t, 1, len(s.ParseErrors))
	Equal(t, editor.SynErr{1, 7, "mismatched input '*' expecting {<EOF>, ';'}"}, s.ParseErrors[0])
}

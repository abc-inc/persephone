// Copyright 2022 The persephone authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package editor_test

import (
	"testing"

	"github.com/abc-inc/persephone/editor"
	. "github.com/stretchr/testify/require"
)

func TestCatchErrorInSecondStatement(t *testing.T) {
	msg1 := "mismatched input 'POTATO' expecting {':', CYPHER, EXPLAIN, PROFILE, USING, CREATE, DROP, LOAD, WITH, OPTIONAL, MATCH, UNWIND, MERGE, SET, DETACH, DELETE, REMOVE, FOREACH, RETURN, START, CALL, CATALOG, SHOW, STOP, ALTER, GRANT, DENY, REVOKE}"
	cypher := `RETURN 1;
POTATO;
RETURN rand();`

	e := editor.NewEditor(cypher)
	Equal(t, 1, len(e.ParseErrors))
	Equal(t, editor.SynErr{2, 0, msg1}, e.ParseErrors[0])
}

func TestParseCommonParam(t *testing.T) {
	msg1 := "mismatched input 'hello' expecting {':', CYPHER, EXPLAIN, PROFILE, USING, CREATE, DROP, LOAD, WITH, OPTIONAL, MATCH, UNWIND, MERGE, SET, DETACH, DELETE, REMOVE, FOREACH, RETURN, START, CALL, CATALOG, SHOW, STOP, ALTER, GRANT, DENY, REVOKE}"
	msg2 := "mismatched input 'hello2' expecting {':', CYPHER, EXPLAIN, PROFILE, USING, CREATE, DROP, LOAD, WITH, OPTIONAL, MATCH, UNWIND, MERGE, SET, DETACH, DELETE, REMOVE, FOREACH, RETURN, START, CALL, CATALOG, SHOW, STOP, ALTER, GRANT, DENY, REVOKE}"
	cypher := `:play;
hello;
:param x => 1;
hello2;
:play reco;`

	e := editor.NewEditor(cypher)
	Equal(t, 2, len(e.ParseErrors))
	Equal(t, editor.SynErr{2, 0, msg1}, e.ParseErrors[0])
	Equal(t, editor.SynErr{4, 0, msg2}, e.ParseErrors[1])
}

func TestParseCommon(t *testing.T) {
	msg1 := "mismatched input 'hello' expecting {':', CYPHER, EXPLAIN, PROFILE, USING, CREATE, DROP, LOAD, WITH, OPTIONAL, MATCH, UNWIND, MERGE, SET, DETACH, DELETE, REMOVE, FOREACH, RETURN, START, CALL, CATALOG, SHOW, STOP, ALTER, GRANT, DENY, REVOKE, SP}"
	msg2 := "mismatched input 'hello2' expecting {':', CYPHER, EXPLAIN, PROFILE, USING, CREATE, DROP, LOAD, WITH, OPTIONAL, MATCH, UNWIND, MERGE, SET, DETACH, DELETE, REMOVE, FOREACH, RETURN, START, CALL, CATALOG, SHOW, STOP, ALTER, GRANT, DENY, REVOKE}"
	cypher := `hello;
:param x => 1;
hello2;
:play reco;`

	e := editor.NewEditor(cypher)
	Equal(t, 2, len(e.ParseErrors))
	Equal(t, editor.SynErr{1, 0, msg1}, e.ParseErrors[0])
	Equal(t, editor.SynErr{3, 0, msg2}, e.ParseErrors[1])
}

func TestParseCommonParamCommand(t *testing.T) {
	cypher := `:play http://guides.neo4j.com/reco;
:param x => 1;
RETURN $x;
:play reco;`

	e := editor.NewEditor(cypher)
	Nil(t, e.ParseErrors)
}

func TestParseMultipleParamCommandsWithQuery(t *testing.T) {
	cypher := `:param age => 25;
:param interests => ['football', 'fishing'];
MATCH (n)
WHERE n.age > $age
AND n.interest IN $interests
RETURN n;`

	e := editor.NewEditor(cypher)
	Nil(t, e.ParseErrors)
}

func TestRecoverToSecondStatementAfterInvalidCommand(t *testing.T) {
	cypher := ":PUT ao*51 fagas 8(!; :play;"
	e := editor.NewEditor(cypher)
	Equal(t, 1, len(e.ParseErrors))
	Equal(t, editor.SynErr{1, 7, "mismatched input '*' expecting {<EOF>, ';'}"}, e.ParseErrors[0])
}

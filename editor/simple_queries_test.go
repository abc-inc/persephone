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

func TestCorrectASTForSimpleQuery(t *testing.T) {
	e := editor.NewEditor("RETURN 42;")
	Nil(t, e.ParseErrors)
}

func TestErrorsForIncorrectQuery(t *testing.T) {
	msg1 := "mismatched input 'POTATO' expecting {':', CYPHER, EXPLAIN, PROFILE, USING, CREATE, DROP, LOAD, WITH, OPTIONAL, MATCH, UNWIND, MERGE, SET, DETACH, DELETE, REMOVE, FOREACH, RETURN, START, CALL, CATALOG, SHOW, STOP, ALTER, GRANT, DENY, REVOKE, SP}"
	e := editor.NewEditor("POTATO")
	Equal(t, 1, len(e.ParseErrors))
	Equal(t, editor.SynErr{1, 0, msg1}, e.ParseErrors[0])
	Equal(t, "POTATO<EOF>", e.ParseTree.GetText())
}

func TestErrorsIfErrorInLexer(t *testing.T) {
	msg1 := "mismatched input '`' expecting {<EOF>, ';'}"
	e := editor.NewEditor("WITH a` WITH 1;")
	Equal(t, 1, len(e.ParseErrors))
	Equal(t, editor.SynErr{1, 6, msg1}, e.ParseErrors[0])
}

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

package editor

import "github.com/antlr/antlr4/runtime/Go/antlr"

type SynErr struct {
	Line int
	Col  int
	Msg  string
}
type ErrorListener struct {
	errors []SynErr
}

func (el *ErrorListener) SyntaxError(
	r antlr.Recognizer, offendingSymbol any, line, column int, msg string, e antlr.RecognitionException) {

	if msg == "mismatched input '<EOF>' expecting {';', SP}" {
		// suppress error about missing semicolon at the end of a query
		return
	}
	if msg == "missing ';' at '<EOF>'" {
		return
	}
	if msg == "mismatched input '<EOF>' expecting {':', CYPHER, EXPLAIN, PROFILE, USING, CREATE, DROP, LOAD, WITH, OPTIONAL, MATCH, UNWIND, MERGE, SET, DETACH, DELETE, REMOVE, FOREACH, RETURN, START, CALL}" {
		return
	}
	el.errors = append(el.errors, SynErr{line, column, msg})
}

func (el ErrorListener) ReportAmbiguity(antlr.Parser, *antlr.DFA, int, int, bool, *antlr.BitSet, antlr.ATNConfigSet) {
}

func (el ErrorListener) ReportAttemptingFullContext(antlr.Parser, *antlr.DFA, int, int, *antlr.BitSet, antlr.ATNConfigSet) {
}

func (el ErrorListener) ReportContextSensitivity(antlr.Parser, *antlr.DFA, int, int, int, antlr.ATNConfigSet) {
}

var _ antlr.ErrorListener = (*ErrorListener)(nil)

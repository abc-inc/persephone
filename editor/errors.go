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

func (el *ErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
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

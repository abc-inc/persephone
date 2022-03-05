package cypher

import (
	"testing"

	. "github.com/stretchr/testify/require"
)

func TestQuerySemicolonIgnoreLastQuery(t *testing.T) {
	s := NewEditorSupport("RETURN 1")
	Nil(t, s.parseErrors)
}

func TestQuerySemicolonIgnoreLastQuery2(t *testing.T) {
	s := NewEditorSupport("RETURN 1; RETURN 1")
	Nil(t, s.parseErrors)
}

func TestQuerySemicolonIgnoreLastQueryNewLine(t *testing.T) {
	s := NewEditorSupport("RETURN 1\n")
	Nil(t, s.parseErrors)
}

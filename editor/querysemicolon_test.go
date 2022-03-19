package editor_test

import (
	"testing"

	"github.com/abc-inc/merovingian/editor"
	. "github.com/stretchr/testify/require"
)

func TestQuerySemicolonIgnoreLastQuery(t *testing.T) {
	s := editor.NewEditorSupport("RETURN 1")
	Nil(t, s.ParseErrors)
}

func TestQuerySemicolonIgnoreLastQuery2(t *testing.T) {
	s := editor.NewEditorSupport("RETURN 1; RETURN 1")
	Nil(t, s.ParseErrors)
}

func TestQuerySemicolonIgnoreLastQueryNewLine(t *testing.T) {
	s := editor.NewEditorSupport("RETURN 1\n")
	Nil(t, s.ParseErrors)
}

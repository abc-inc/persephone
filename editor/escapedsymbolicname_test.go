package editor_test

import (
	"testing"

	"github.com/abc-inc/merovingian/editor"
	. "github.com/stretchr/testify/require"
)

func TestEscapedSymbolicNames(t *testing.T) {
	tests := []struct {
		name   string
		cypher string
	}{
		{name: "variable", cypher: "RETURN ` () some name \"`;"},
		{name: "label", cypher: "MATCH (:`Label()`);"},
		{name: "relationship type", cypher: "MATCH ()-[:` type`]-();"},
		{name: "function", cypher: "RETURN `func`();"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			s := editor.NewEditorSupport(test.cypher)
			Nil(t, s.ParseErrors)
		})
	}
}

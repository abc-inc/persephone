package editor_test

import (
	"strconv"
	"testing"

	"github.com/abc-inc/persephone/editor"
	. "github.com/stretchr/testify/require"
)

func TestCypherLegacyFiles(t *testing.T) {
	for i, test := range cypherLegacy {
		t.Run("cypher-legacy-query-"+strconv.Itoa(i), func(t *testing.T) {
			s := editor.NewEditorSupport(test)
			Nil(t, s.ParseErrors)
		})
	}
}

func TestOpenCypherFiles(t *testing.T) {
	for i, test := range cypherDefault {
		t.Run("cypher-query-"+strconv.Itoa(i), func(t *testing.T) {
			s := editor.NewEditorSupport(test)
			Nil(t, s.ParseErrors)
		})
	}
}

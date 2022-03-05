package cypher

import (
	"strconv"
	"testing"

	"github.com/abc-inc/merovingian/cypher/files"
	. "github.com/stretchr/testify/require"
)

func TestCypherLegacyFiles(t *testing.T) {
	for i, test := range files.CypherLegacy {
		t.Run("cypher-legacy-query-"+strconv.Itoa(i), func(t *testing.T) {
			s := NewEditorSupport(test)
			Nil(t, s.parseErrors)
		})
	}
}

func TestOpenCypherFiles(t *testing.T) {
	for i, test := range files.CypherDefault {
		t.Run("cypher-query-"+strconv.Itoa(i), func(t *testing.T) {
			s := NewEditorSupport(test)
			Nil(t, s.parseErrors)
		})
	}
}

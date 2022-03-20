package editor_test

import (
	"testing"

	"github.com/abc-inc/persephone/editor"
	. "github.com/stretchr/testify/require"
)

func TestGenericQueries(t *testing.T) {
	tests := []string{
		"MERGE () ON CREATE SET connection.departure = 1445, connection.arrival = 1710",
		"MERGE () ON MATCH SET connection.departure = 1445, connection.arrival = 1710",
		"SET a=a,b=b",
		"SET a=a ,b=b",
		"SET a=a, b=b",
		"SET a=a , b=b",
		"RETURN a ORDER BY a,b;",
		"RETURN a ORDER BY a ,b;",
		"RETURN a ORDER BY a, b;",
		"RETURN a ORDER BY a , b;",
	}
	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			s := editor.NewEditorSupport(test)
			Nil(t, s.ParseErrors)
		})
	}
}

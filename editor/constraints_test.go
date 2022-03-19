package editor_test

import (
	"testing"

	"github.com/abc-inc/merovingian/editor"
	. "github.com/stretchr/testify/require"
)

func TestConstraints(t *testing.T) {
	tests := []struct {
		name   string
		cypher string
	}{
		{name: "create on single property", cypher: "CREATE CONSTRAINT ON (n:Person) ASSERT (n.email) IS NODE KEY"},
		{name: "drop on single property", cypher: "DROP CONSTRAINT ON (n:Person) ASSERT (n.email) IS NODE KEY"},
		{name: "create on composite property", cypher: "CREATE CONSTRAINT ON (n:User) ASSERT (n.firstname,n.lastname) IS NODE KEY"},
		{name: "drop on composite property", cypher: "DROP CONSTRAINT ON (n:Person) ASSERT (n.firstname,n.lastname) IS NODE KEY"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			s := editor.NewEditorSupport(test.cypher)
			Nil(t, s.ParseErrors)
		})
	}
}

func TestDropWithoutPropertyFails(t *testing.T) {
	s := editor.NewEditorSupport("DROP CONSTRAINT ON (n:Person) ASSERT (n)")
	Equal(t, 1, len(s.ParseErrors))
	Equal(t, len(s.Input), s.ParseErrors[0].Col)
}

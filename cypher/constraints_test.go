package cypher

import (
	"testing"

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
			s := NewEditorSupport(test.cypher)
			Nil(t, s.parseErrors)
		})
	}
}

func TestDropWithoutPropertyFails(t *testing.T) {
	s := NewEditorSupport("DROP CONSTRAINT ON (n:Person) ASSERT (n)")
	Equal(t, 1, len(s.parseErrors))
	Equal(t, len(s.input), s.parseErrors[0].Col)
}

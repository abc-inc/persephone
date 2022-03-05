package cypher

import (
	"testing"

	. "github.com/stretchr/testify/require"
)

func TestIndexes(t *testing.T) {
	tests := []struct {
		name   string
		cypher string
	}{
		{name: "create on single property", cypher: "CREATE INDEX ON :Person(name)"},
		{name: "create on compound property", cypher: "CREATE INDEX ON :Person(name, surname)"},
		{name: "drop on single property", cypher: "DROP INDEX ON :Person(name)"},
		{name: "drop on compound property", cypher: "DROP INDEX ON :Person(name, surname)"},
		{name: "hint on single property", cypher: "MATCH (f:Foo) USING INDEX f:Foo(bar)"},
		{name: "hint on compound property", cypher: "MATCH (f:Foo) USING INDEX f:Foo(bar,baz)"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			s := NewEditorSupport(test.cypher)
			Nil(t, s.parseErrors)
		})
	}
}

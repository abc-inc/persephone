package editor_test

import (
	"testing"

	"github.com/abc-inc/merovingian/editor"
	. "github.com/stretchr/testify/require"
)

func TestSpecialFunctions(t *testing.T) {
	tests := []struct {
		name   string
		cypher string
	}{
		{name: "extract", cypher: "RETURN extract(n IN nodes(p) | n.age) AS extracted;"},
		{name: "reduce", cypher: "RETURN reduce(totalAge = 0, n IN nodes(p)| totalAge + n.age) AS reduction"},
		{name: "shortestPath", cypher: "RETURN shortestPath( ( f)-[]-( t) );"},
		{name: "allShortestPaths", cypher: "RETURN allShortestPaths((f)-[]-(t));"},
		{name: "exists", cypher: "RETURN n.prop AS prop1, exists((n)-[:SOMETHING]->()) AS something;"},
		{name: "3rd party", cypher: "RETURN a.b();"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			b := editor.NewEditorSupport(test.cypher)
			Nil(t, b.ParseErrors)
		})
	}
}

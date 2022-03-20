package editor_test

import (
	"testing"

	"github.com/abc-inc/persephone/editor"
	. "github.com/stretchr/testify/require"
)

func TestMapProjection(t *testing.T) {
	tests := []struct {
		name   string
		cypher string
	}{
		{name: "all property selector", cypher: "RETURN person { .* };"},
		{name: "property selector", cypher: "RETURN person { .name };"},
		{name: "literal entry and map projection inside", cypher: "RETURN person { someProp: collect(moreProps { .variable1, .variable2 })};"},
		{name: "literal entry", cypher: "RETURN person { someProp: collect(expression)};"},
		{name: "variable", cypher: "RETURN person { person };"},
		{name: "variable without spaces", cypher: "RETURN person{person};"},
		{name: "multiple", cypher: "RETURN person{ person, .person, something: expression(), .*};"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			s := editor.NewEditorSupport(test.cypher)
			Nil(t, s.ParseErrors)
		})
	}
}

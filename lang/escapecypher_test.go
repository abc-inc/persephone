package lang

import (
	"testing"

	. "github.com/stretchr/testify/require"
)

func TestOnlyBackticksStringsWithSpaces(t *testing.T) {
	Equal(t, "nospaces", EscapeCypher("nospaces"))
	Equal(t, "`with spaces`", EscapeCypher("with spaces"))
}

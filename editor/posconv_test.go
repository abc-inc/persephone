package editor_test

import (
	"testing"

	"github.com/abc-inc/persephone/editor"
	. "github.com/stretchr/testify/require"
)

func TestPosConv_ToAbsolute(t *testing.T) {
	pc := editor.NewPosConv("a\nbc")
	Equal(t, 0, pc.ToAbsolute(1, 0))
	Equal(t, 2, pc.ToAbsolute(2, 0))
	Equal(t, 3, pc.ToAbsolute(2, 1))
	Equal(t, 1, pc.ToAbsolute(1, 1))
}

func TestPosConv_ToRelative(t *testing.T) {
	type pair struct {
		line, col int
	}

	pc := editor.NewPosConv("a\nbc")
	line, col := pc.ToRelative(0)
	Equal(t, pair{1, 0}, pair{line, col})
	line, col = pc.ToRelative(1)
	Equal(t, pair{1, 1}, pair{line, col})
	line, col = pc.ToRelative(2)
	Equal(t, pair{2, 0}, pair{line, col})
	line, col = pc.ToRelative(3)
	Equal(t, pair{2, 1}, pair{line, col})
}

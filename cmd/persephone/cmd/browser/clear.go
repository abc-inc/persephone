package cmd

import (
	"github.com/abc-inc/persephone/cmd/persephone/cmd/types"
	"github.com/inancgumus/screen"
	"github.com/spf13/cobra"
)

var ClearCmd = &cobra.Command{
	Use:         ":clear",
	Short:       "Clear the screen",
	Annotations: types.Annotate(types.Offline),
	Run:         func(cmd *cobra.Command, args []string) { Clear() },
}

func Clear() {
	screen.Clear()
	screen.MoveTopLeft()
}

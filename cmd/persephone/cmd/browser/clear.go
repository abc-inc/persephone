package cmd

import (
	"github.com/abc-inc/persephone/cmd/persephone/cmd/persephone"
	"github.com/inancgumus/screen"
	"github.com/spf13/cobra"
)

var ClearCmd = &cobra.Command{
	Use:         ":clear",
	Short:       "Clear the screen",
	Annotations: cmd.Annotate(cmd.Offline),
	Run:         func(cmd *cobra.Command, args []string) { Clear() },
}

func Clear() {
	screen.Clear()
	screen.MoveTopLeft()
}

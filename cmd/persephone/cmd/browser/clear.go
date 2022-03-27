package cmd

import (
	"github.com/inancgumus/screen"
	"github.com/spf13/cobra"
)

var ClearCmd = &cobra.Command{
	Use:   ":clear",
	Short: "Clear the screen",
	Run:   clearCmd,
}

func clearCmd(cmd *cobra.Command, args []string) {
	screen.Clear()
	screen.MoveTopLeft()
}

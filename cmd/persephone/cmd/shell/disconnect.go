package cmd

import (
	"github.com/abc-inc/persephone/format"
	"github.com/abc-inc/persephone/graph"
	"github.com/spf13/cobra"
)

var DisconnectCmd = &cobra.Command{
	Use:   ":disconnect",
	Short: "Disconnects from database",
	Run:   disconnectCmd,
}

func disconnectCmd(cmd *cobra.Command, args []string) {
	if err := graph.GetConn().Close(); err != nil {
		format.Writeln(err)
		return
	}
	format.Writeln("Disconnected")
}

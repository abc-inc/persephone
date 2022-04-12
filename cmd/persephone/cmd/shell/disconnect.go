package cmd

import (
	"github.com/abc-inc/persephone/console"
	"github.com/abc-inc/persephone/graph"
	"github.com/spf13/cobra"
)

var DisconnectCmd = &cobra.Command{
	Use:   ":disconnect",
	Short: "Disconnect from database",
	Run:   func(cmd *cobra.Command, args []string) { Disconnect() },
}

func Disconnect() {
	if err := graph.GetConn().Close(); err != nil {
		console.WriteErr(err)
	}
}

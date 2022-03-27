package cmd

import (
	"github.com/abc-inc/persephone/format"
	"github.com/abc-inc/persephone/graph"
	"github.com/spf13/cobra"
)

var BeginCmd = &cobra.Command{
	Use:   ":begin",
	Short: "Open a transaction",
	Long:  "Start a transaction which will remain open until :commit or :rollback is called",
	Run:   beginCmd,
}

func beginCmd(cmd *cobra.Command, args []string) {
	_, _, err := graph.GetConn().GetTransaction()
	if err != nil {
		format.Writeln(err)
	}
}

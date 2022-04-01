package cmd

import (
	"github.com/abc-inc/persephone/format"
	"github.com/abc-inc/persephone/graph"
	"github.com/spf13/cobra"
)

var UseCmd = &cobra.Command{
	Use:   ":use database",
	Short: "Set the active database",
	Long:  "Set the active database that transactions are executed on",
	Run:   useCmd,
}

func useCmd(cmd *cobra.Command, args []string) {
	if err := graph.GetConn().UseDB(args[0]); err != nil {
		format.Writeln(err)
	}
}

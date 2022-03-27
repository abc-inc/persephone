package cmd

import (
	"github.com/abc-inc/persephone/format"
	"github.com/abc-inc/persephone/graph"
	"github.com/spf13/cobra"
)

var RollbackCmd = &cobra.Command{
	Use:   ":rollback",
	Short: "Rollback the currently open transaction",
	Long:  "Roll back and closes the currently open transaction",
	Run:   rollbackCmd,
}

func rollbackCmd(cmd *cobra.Command, args []string) {
	if err := graph.GetConn().Rollback(); err != nil {
		format.Writeln(err)
	}
}

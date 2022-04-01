package cmd

import (
	"errors"

	"github.com/abc-inc/persephone/format"
	"github.com/abc-inc/persephone/graph"
	"github.com/spf13/cobra"
)

var errNoTxRollback = errors.New("there is no open transaction to rollback")

var RollbackCmd = &cobra.Command{
	Use:   ":rollback",
	Short: "Rollback the currently open transaction",
	Long:  "Roll back and closes the currently open transaction",
	Run:   rollbackCmd,
}

func rollbackCmd(cmd *cobra.Command, args []string) {
	if ok, err := graph.GetConn().Rollback(); err != nil {
		format.Writeln(err)
	} else if !ok {
		format.Writeln(errNoTxRollback)
	}
}

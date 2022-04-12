package cmd

import (
	"errors"

	"github.com/abc-inc/persephone/console"
	"github.com/abc-inc/persephone/graph"
	"github.com/spf13/cobra"
)

var errNoTxRollback = errors.New("there is no open transaction to rollback")

var RollbackCmd = &cobra.Command{
	Use:   ":rollback",
	Short: "Rollback the currently open transaction",
	Long:  "Rollback and close the currently open transaction",
	Run:   func(cmd *cobra.Command, args []string) { Rollback() },
}

func Rollback() {
	if ok, err := graph.GetConn().Rollback(); err != nil {
		console.WriteErr(err)
	} else if !ok {
		console.WriteErr(errNoTxRollback)
	}
}

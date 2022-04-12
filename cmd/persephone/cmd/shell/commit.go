package cmd

import (
	"errors"

	"github.com/abc-inc/persephone/console"
	"github.com/abc-inc/persephone/graph"
	"github.com/spf13/cobra"
)

var errNoTxCommit = errors.New("there is no open transaction to commit")

var CommitCmd = &cobra.Command{
	Use:   ":commit",
	Short: "Commit the currently open transaction",
	Long:  "Commit and close the currently open transaction",
	Run:   func(cmd *cobra.Command, args []string) { Commit() },
}

func Commit() {
	if ok, err := graph.GetConn().Commit(); err != nil {
		console.WriteErr(err)
	} else if !ok {
		console.WriteErr(errNoTxCommit)
	}
}

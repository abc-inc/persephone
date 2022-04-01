package cmd

import (
	"errors"

	"github.com/abc-inc/persephone/format"
	"github.com/abc-inc/persephone/graph"
	"github.com/spf13/cobra"
)

var errNoTxCommit = errors.New("there is no open transaction to commit")

var CommitCmd = &cobra.Command{
	Use:   ":commit",
	Short: "Commit the currently open transaction",
	Long:  "Commit and close the currently open transaction",
	Run:   commitCmd,
}

func commitCmd(cmd *cobra.Command, args []string) {
	if ok, err := graph.GetConn().Commit(); err != nil {
		format.Writeln(err)
	} else if !ok {
		format.Writeln(errNoTxCommit)
	}
}

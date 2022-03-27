package cmd

import (
	"github.com/abc-inc/persephone/format"
	"github.com/abc-inc/persephone/graph"
	"github.com/spf13/cobra"
)

var CommitCmd = &cobra.Command{
	Use:   ":commit",
	Short: "Commit the currently open transaction",
	Long:  "Commit and close the currently open transaction",
	Run:   commitCmd,
}

func commitCmd(cmd *cobra.Command, args []string) {
	if err := graph.GetConn().Commit(); err != nil {
		format.Writeln(err)
	}
}

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
	Run:   func(cmd *cobra.Command, args []string) { Use(args[0]) },
}

func Use(dbName string) {
	if err := graph.GetConn().UseDB(dbName); err != nil {
		format.Writeln(err)
	}
}

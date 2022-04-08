package cmd

import (
	"errors"

	"github.com/abc-inc/persephone/console"
	"github.com/abc-inc/persephone/graph"
	"github.com/spf13/cobra"
)

var errTxActive = errors.New("there is already an open transaction")

var BeginCmd = &cobra.Command{
	Use:   ":begin",
	Short: "Open a transaction",
	Long:  "Start a transaction which will remain open until :commit or :rollback is called",
	Run:   func(cmd *cobra.Command, args []string) { Begin() },
}

func Begin() {
	_, created, err := graph.GetConn().GetTransaction()
	if err != nil {
		console.Writeln(err)
	} else if !created {
		console.Writeln(errTxActive)
	}
}

package cmd

import (
	"errors"

	"github.com/abc-inc/persephone/format"
	"github.com/abc-inc/persephone/graph"
	"github.com/spf13/cobra"
)

var errTxActive = errors.New("there is already an open transaction")

var BeginCmd = &cobra.Command{
	Use:   ":begin",
	Short: "Open a transaction",
	Long:  "Start a transaction which will remain open until :commit or :rollback is called",
	Run:   beginCmd,
}

func beginCmd(cmd *cobra.Command, args []string) {
	_, created, err := graph.GetConn().GetTransaction()
	if err != nil {
		format.Writeln(err)
	} else if !created {
		format.Writeln(errTxActive)
	}
}

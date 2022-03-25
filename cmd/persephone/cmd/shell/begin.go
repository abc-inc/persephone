package cmd

import "github.com/spf13/cobra"

var BeginCmd = &cobra.Command{
	Use:   ":begin",
	Short: "Open a transaction",
	Run:   beginCmd,
}

func beginCmd(cmd *cobra.Command, args []string) {
}

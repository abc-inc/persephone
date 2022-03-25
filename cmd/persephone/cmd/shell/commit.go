package cmd

import "github.com/spf13/cobra"

var CommitCmd = &cobra.Command{
	Use:   ":commit",
	Short: "Commit the currently open transaction",
	Run:   commitCmd,
}

func commitCmd(cmd *cobra.Command, args []string) {
}

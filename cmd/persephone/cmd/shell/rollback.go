package cmd

import "github.com/spf13/cobra"

var RollbackCmd = &cobra.Command{
	Use: ":rollback",
	Short: "Rollback the currently open transaction",
	Run: rollbackCmd,
}

func rollbackCmd(cmd *cobra.Command, args []string) {
}

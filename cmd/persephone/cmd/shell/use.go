package cmd

import "github.com/spf13/cobra"

var UseCmd = &cobra.Command{
	Use:   ":use database",
	Short: "Set the active database",
	Long:  "Set the active database that transactions are executed on",
	Run:   useCmd,
}

func useCmd(cmd *cobra.Command, args []string) {
}

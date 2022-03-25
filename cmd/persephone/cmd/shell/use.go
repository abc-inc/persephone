package cmd

import "github.com/spf13/cobra"

var UseCmd = &cobra.Command{
	Use:   ":use",
	Short: "Set the active database",
	Run:   useCmd,
}

func useCmd(cmd *cobra.Command, args []string) {
}

package cmd

import "github.com/spf13/cobra"

var SourceCmd = &cobra.Command{
	Use: ":source",
	Short: "Interactively executes cypher statements from a file",
	Run: sourceCmd,
}

func sourceCmd(cmd *cobra.Command, args []string) {
}

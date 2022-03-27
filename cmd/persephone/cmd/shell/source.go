package cmd

import "github.com/spf13/cobra"

var SourceCmd = &cobra.Command{
	Use:   ":source [filename]",
	Short: "Interactively executes cypher statements from a file",
	Long:  "Executes Cypher statements from a file",
	Run:   sourceCmd,
}

func sourceCmd(cmd *cobra.Command, args []string) {
}

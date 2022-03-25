package cmd

import "github.com/spf13/cobra"

var QueriesCmd = &cobra.Command{
	Use: ":queries",
	Short: "List your servers and clusters running queries",
	Run: queriesCmd,
}

func queriesCmd(cmd *cobra.Command, args []string) {
}

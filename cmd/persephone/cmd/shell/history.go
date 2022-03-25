package cmd

import "github.com/spf13/cobra"

var HistoryCmd = &cobra.Command{
	Use: ":history",
	Short: "Print a list of the last commands executed",
	Run: historyCmd,
}

func historyCmd(cmd *cobra.Command, args []string) {
}

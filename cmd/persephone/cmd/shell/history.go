package cmd

import "github.com/spf13/cobra"

var HistoryCmd = &cobra.Command{
	Use:   ":history",
	Short: "Print a list of the last commands executed",
	Long:  "Prints a list of the last statements executed",
	Run:   historyCmd,
}

func init() {
	HistoryCmd.AddCommand(&cobra.Command{
		Use:   "clear",
		Short: "Removes all entries from the history",
		Run:   historyClearCmd,
	})
}

func historyCmd(cmd *cobra.Command, args []string) {
}

func historyClearCmd(cmd *cobra.Command, args []string) {
}

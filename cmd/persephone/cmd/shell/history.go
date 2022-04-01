package cmd

import (
	"fmt"

	"github.com/abc-inc/persephone/hist"
	"github.com/spf13/cobra"
)

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
	for i, e := range hist.Get().Entries() {
		fmt.Printf("%4d  %s\n", i+1, e)
	}
}

func historyClearCmd(cmd *cobra.Command, args []string) {
	hist.Get().Clear()
}

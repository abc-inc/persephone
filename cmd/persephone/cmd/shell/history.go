package cmd

import (
	"fmt"

	"github.com/abc-inc/persephone/console"
	"github.com/abc-inc/persephone/event"
	"github.com/spf13/cobra"
)

type entry struct {
	Pos  int    `json:"Pos"`
	Stmt string `json:"Statement"`
}

var HistoryCmd = &cobra.Command{
	Use:   ":history",
	Short: "Print a list of the last commands executed",
	Long:  "Print a list of the last statements executed",
	Run:   func(cmd *cobra.Command, args []string) { History() },
}

func init() {
	HistoryCmd.AddCommand(&cobra.Command{
		Use:   "clear",
		Short: "Removes all entries from the history",
		Run:   func(cmd *cobra.Command, args []string) { HistoryClear() },
	})

	event.Subscribe(event.FormatEvent{}, func(e event.FormatEvent) {
		sep := e.Sep
		console.SetFormatter(entry{}, func(i interface{}) (string, error) {
			e := i.(entry)
			return fmt.Sprintf("%d%s%s", e.Pos+1, sep, e.Stmt), nil
		})
	})
}

func History() {
	var es []entry
	for i, e := range console.Get().Entries() {
		es = append(es, entry{i, e})
	}
	console.Writeln(es)
}

func HistoryClear() {
	console.Get().Clear()
}

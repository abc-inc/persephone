package repl

import (
	"strings"

	"github.com/spf13/cobra"
)

type Item struct {
	View    string `json:"view"`
	Content string `json:"content"`
}

func (i Item) String() string {
	return i.View
}

type CompFunc func(str string) []Item

func NoComp(string) []Item {
	return nil
}

func SubCmdComp(cmd *cobra.Command) func(str string) []Item {
	return func(s string) (its []Item) {
		for _, c := range cmd.Commands() {
			if strings.HasPrefix(c.Name(), s) {
				its = append(its, Item{View: c.Name()})
			}
		}
		return
	}
}

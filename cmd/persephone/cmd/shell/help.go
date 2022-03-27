package cmd

import (
	"strings"

	"github.com/spf13/cobra"
)

var HelpCmd = &cobra.Command{
	Use:   ":help [command]",
	Short: "Show this help message",
	Long:  "Show the list of available commands or help for a specific command",
	Run:   helpCmd,
}

func helpCmd(cmd *cobra.Command, args []string) {
	if len(args) > 0 {
		for _, c := range cmd.Root().Commands() {
			if strings.TrimPrefix(c.Name(), ":") == args[0] {
				c.Help()
				return
			}
		}
	}

	cmd.Usage()
}

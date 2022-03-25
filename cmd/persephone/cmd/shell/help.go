package cmd

import "github.com/spf13/cobra"

var HelpCmd = &cobra.Command{
	Use: ":help",
	Short: "Show this help message",
	Run: helpCmd,
}

func helpCmd(cmd *cobra.Command, args []string) {
	if len(args) != 1 {
		for _, c := range cmd.Root().Commands() {
			if c.Name() == args[0] {
				cmd = c
				break
			}
		}
	}

	cmd.Usage()
}

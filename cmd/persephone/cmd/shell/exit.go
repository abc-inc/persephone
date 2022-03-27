package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var ExitCmd = &cobra.Command{
	Use:   ":exit",
	Short: "Exit persephone",
	Run:   exitCmd,
}

func exitCmd(cmd *cobra.Command, args []string) {
	for _, c := range cmd.Root().Commands() {
		if c.Name() == ":disconnect" {
			c.Run(cmd, args)
			os.Exit(0)
		}
	}
}

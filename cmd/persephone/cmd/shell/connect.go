package cmd

import "github.com/spf13/cobra"

var ConnectCmd = &cobra.Command{
	Use:   ":connect",
	Short: "Connects to a database",
	Run:   connectCmd,
}

func connectCmd(cmd *cobra.Command, args []string) {
}

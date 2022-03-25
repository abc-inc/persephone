package cmd

import "github.com/spf13/cobra"

var DisconnectCmd = &cobra.Command{
	Use:   ":disconnect",
	Short: "Disconnects from database",
	Run:   disconnectCmd,
}

func disconnectCmd(cmd *cobra.Command, args []string) {
}

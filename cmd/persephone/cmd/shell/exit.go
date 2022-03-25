package cmd

import "github.com/spf13/cobra"

var ExitCmd = &cobra.Command{
	Use: ":exit",
	Short: "Exit persephone",
	Run: exitCmd,
}

func exitCmd(cmd *cobra.Command, args []string) {
}

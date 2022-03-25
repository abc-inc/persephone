package cmd

import "github.com/spf13/cobra"

var SysinfoCmd = &cobra.Command{
	Use: ":sysinfo",
	Short: "Print system information",
	Run: sysinfoCmd,
}

func sysinfoCmd(cmd *cobra.Command, args []string) {
}

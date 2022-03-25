package cmd

import "github.com/spf13/cobra"

var ParamsCmd = &cobra.Command{
	Use: ":params",
	Short: "Print all currently set query parameters and their values",
	Run: paramsCmd,
}

func paramsCmd(cmd *cobra.Command, args []string) {
}

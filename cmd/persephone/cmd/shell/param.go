package cmd

import "github.com/spf13/cobra"

var ParamCmd = &cobra.Command{
	Use: ":param",
	Short: "Set the value of a query parameter",
	Run: paramCmd,
}

func paramCmd(cmd *cobra.Command, args []string) {
}

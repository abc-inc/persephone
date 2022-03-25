package cmd

import "github.com/spf13/cobra"

var SchemaCmd = &cobra.Command{
	Use: ":schema",
	Short: "Shows information about database schema indexes and constraints",
	Run: schemaCmd,
}

func schemaCmd(cmd *cobra.Command, args []string) {
}

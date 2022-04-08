package cmd

import (
	"github.com/abc-inc/persephone/console"
	"github.com/spf13/cobra"
)

var FormatCmd = &cobra.Command{
	Use:         ":format FORMAT",
	Short:       "Change the output format (supported formats: auto, csv, json, jsonc, table, text, tsv, yaml, yamlc)",
	ValidArgs:   []string{"auto", "csv", "json", "jsonc", "table", "text", "tsv", "yaml", "yamlc"},
	Args:        cobra.ExactValidArgs(1),
	Annotations: Annotate(Offline),
	Run:         func(cmd *cobra.Command, args []string) { Format(args[0]) },
}

func Format(f string) {
	console.ChangeFmt(f)
}

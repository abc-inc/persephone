package cmd

import (
	"github.com/abc-inc/persephone/cmd/persephone/cmd/types"
	"github.com/abc-inc/persephone/format"
	"github.com/spf13/cobra"
)

var FormatCmd = &cobra.Command{
	Use:         ":format FORMAT",
	Short:       "Change the output format (supported formats: csv, json, jsonc, table, text, tsv, yaml, yamlc)",
	ValidArgs:   []string{"csv", "json", "jsonc", "table", "text", "tsv", "yaml", "yamlc"},
	Args:        cobra.ExactValidArgs(1),
	Annotations: types.Annotate(types.Offline),
	Run:         func(cmd *cobra.Command, args []string) { Format(args[0]) },
}

func Format(f string) {
	format.Change(f)
}

package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/abc-inc/persephone/graph"
	"github.com/abc-inc/persephone/internal"
	"github.com/spf13/cobra"
)

var ParamsCmd = &cobra.Command{
	Use:   ":params [parameter]",
	Short: "Print all currently set query parameters and their values",
	Long:  "Print a table of all currently set query parameters or the value for the given parameter",
	Run:   paramsCmd,
}

func paramsCmd(cmd *cobra.Command, args []string) {
	j := internal.Must(json.MarshalIndent(graph.GetConn().Params, "", "  "))
	fmt.Println(string(j))
}

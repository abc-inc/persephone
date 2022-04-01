package cmd

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/abc-inc/persephone/format"
	"github.com/abc-inc/persephone/graph"
	"github.com/spf13/cobra"
)

var ParamCmd = &cobra.Command{
	Use:   ":param name value",
	Short: "Set the value of a query parameter",
	Long:  "Set the specified query parameter to the value given",
	Args:  cobra.ExactArgs(2),
	Run:   paramCmd,
}

func paramCmd(cmd *cobra.Command, args []string) {
	var v map[string]interface{}
	err := json.Unmarshal([]byte(fmt.Sprintf(`{"%s": %s}`, args[0], args[1])), &v)
	if err != nil {
		format.Writeln(errors.New("failed to evaluate expression " + args[1]))
	}
	graph.GetConn().Params[args[0]] = v[args[0]]
}

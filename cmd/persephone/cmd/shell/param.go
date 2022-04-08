package cmd

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/abc-inc/persephone/console"
	"github.com/abc-inc/persephone/graph"
	"github.com/spf13/cobra"
)

var ParamCmd = &cobra.Command{
	Use:   ":param name value",
	Short: "Set the value of a query parameter",
	Long:  "Set the specified query parameter to the value given",
	Args:  cobra.ExactArgs(2),
	Run:   func(cmd *cobra.Command, args []string) { Param(args[0], args[1]) },
}

func Param(key, val string) {
	var m map[string]interface{}
	err := json.Unmarshal([]byte(fmt.Sprintf(`{"%s": %s}`, key, val)), &m)
	if err != nil {
		console.Writeln(errors.New("failed to evaluate expression " + val))
	}
	graph.GetConn().Params[key] = m[key]
}

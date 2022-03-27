package cmd

import (
	"encoding/json"
	"regexp"
	"strings"

	"github.com/abc-inc/persephone/graph"
	"github.com/abc-inc/persephone/internal"
	"github.com/spf13/cobra"
)

var ParamCmd = &cobra.Command{
	Use:   ":param name => value",
	Short: "Set the value of a query parameter",
	Long:  "Set the specified query parameter to the value given",
	Run:   paramCmd,
}

func paramCmd(cmd *cobra.Command, args []string) {
	r := regexp.MustCompile("\\s*=>\\s*")
	kv := r.Split(strings.Join(args, " "), 2)
	var v map[string]interface{}
	internal.MustNoErr(json.Unmarshal([]byte(kv[1]), &v))
	graph.GetConn().Params[kv[0]] = v
}

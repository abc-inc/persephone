// Copyright 2022 The persephone authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/abc-inc/go-data-neo4j/graph"
	"github.com/abc-inc/persephone/cmd/persephone/cmd/cmdutil"
	"github.com/abc-inc/persephone/console"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

func NewCmdParam(f *cmdutil.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use:   ":param <command>",
		Short: "Manage query parameters",
	}

	cmd.AddCommand(
		NewCmdParamList(f),
		NewCmdParamSet(f),
	)

	return cmd
}

func NewCmdParamList(f *cmdutil.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "Print all currently set query parameters and their values",
		Long:  "Print a table of all currently set query parameters or the value for the given parameter",
		Args:  cobra.ExactArgs(0),
		Run:   func(cmd *cobra.Command, args []string) { ListParams() },
	}

	return cmd
}

func NewCmdParamSet(f *cmdutil.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set name value",
		Short: "Set the value of a query parameter",
		Long:  "Set the specified query parameter to the value given",
		Args:  cobra.ExactArgs(2),
		Run:   func(cmd *cobra.Command, args []string) { _ = SetParam(args[0], args[1]) },
	}

	return cmd
}

func ListParams() {
	console.Write(graph.GetConn().Params)
}

func SetParam(key, val string) error {
	key = strings.Trim(key, `"`)
	var m map[string]any
	err := json.Unmarshal([]byte(fmt.Sprintf(`{"%s": %s}`, key, val)), &m)
	if err != nil {
		log.Err(err).Msgf("Failed to parse parameter: '%s'\n"+
			"The value must be a valid JSON string, number, object, etc.", val)
		return err
	}
	graph.GetConn().Params[key] = m[key]
	return err
}

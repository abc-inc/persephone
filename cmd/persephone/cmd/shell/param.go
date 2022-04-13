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

	"github.com/abc-inc/persephone/graph"
	"github.com/rs/zerolog/log"
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
	key = strings.Trim(key, `"`)
	var m map[string]interface{}
	err := json.Unmarshal([]byte(fmt.Sprintf(`{"%s": %s}`, key, val)), &m)
	if err != nil {
		log.Err(err).Msg("Failed to parse parameter")
		log.Info().Msg("The value must be a valid JSON string, number, object, etc.")
		return
	}
	graph.GetConn().Params[key] = m[key]
}

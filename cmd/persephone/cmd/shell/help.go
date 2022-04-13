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
	"strings"

	"github.com/spf13/cobra"
)

var HelpCmd = &cobra.Command{
	Use:         ":help [command]",
	Short:       "Show this help message",
	Long:        "Show the list of available commands or help for a specific command",
	Annotations: map[string]string{"offline": "true"},
	Run:         helpCmd,
}

func helpCmd(cmd *cobra.Command, args []string) {
	if len(args) > 0 {
		for _, c := range cmd.Root().Commands() {
			if strings.TrimPrefix(c.Name(), ":") == strings.TrimPrefix(args[0], ":") {
				_ = c.Help()
				return
			}
		}
	}

	_ = cmd.Usage()
}

// Copyright 2022 The Persephone authors
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
	"github.com/abc-inc/persephone/cmd/persephone/cmd/cmdutil"
	cmd "github.com/abc-inc/persephone/cmd/persephone/cmd/persephone"
	"github.com/spf13/cobra"
)

func NewCmdHelp(f *cmdutil.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use:         ":help [command]",
		Short:       "Show this help message",
		Long:        "Show the list of available commands or help for a specific command",
		Annotations: cmd.Annotate(cmdutil.SkipAuth),
		Run:         helpCmd,
	}

	return cmd
}

func helpCmd(cmd *cobra.Command, args []string) {
	if subCmd, _, err := cmd.Root().Find(args); err == nil && subCmd != nil {
		_ = subCmd.Help()
	} else {
		_ = cmd.Root().Help()
	}
}

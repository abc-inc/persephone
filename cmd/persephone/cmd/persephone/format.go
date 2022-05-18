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

	"github.com/abc-inc/persephone/cmd/persephone/cmd/cmdutil"
	"github.com/abc-inc/persephone/console"
	"github.com/abc-inc/persephone/console/repl"
	"github.com/spf13/cobra"
)

var formats = []string{"auto", "csv", "json", "jsonc", "raw", "rawc", "table", "text", "tsv", "yaml", "yamlc"}

func NewCmdFormat(f *cmdutil.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use:         ":format <format>",
		Short:       "Change the output format (supported formats: " + strings.Join(formats, ", ") + ")",
		ValidArgs:   formats,
		Args:        cobra.ExactValidArgs(1),
		Annotations: Annotate(cmdutil.SkipAuth),
		Run:         func(cmd *cobra.Command, args []string) { Format(args[0]) },
	}

	return cmd
}

func Format(f string) {
	console.ChangeFmt(f)
}

func FormatComp(s string) (its []repl.Item) {
	for _, f := range formats {
		if strings.HasPrefix(f, s) {
			its = append(its, repl.Item{View: f})
		}
	}
	return
}

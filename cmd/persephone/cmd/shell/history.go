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
	"fmt"

	"github.com/abc-inc/persephone/cmd/persephone/cmd/cmdutil"
	"github.com/abc-inc/persephone/console"
	"github.com/abc-inc/persephone/console/repl"
	"github.com/spf13/cobra"
)

type entry struct {
	Pos  int    `json:"Pos" view:"Pos" yaml:"Pos"`
	Stmt string `json:"Statement" view:"Statement" yaml:"Statement"`
}

func NewCmdHistory(f *cmdutil.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use:   ":history",
		Short: "Print a list of the last statements executed",
		Run:   func(cmd *cobra.Command, args []string) { History() },
	}

	cmd.AddCommand(&cobra.Command{
		Use:   "clear",
		Short: "Remove all entries from the history",
		Run:   func(cmd *cobra.Command, args []string) { HistoryClear() },
	})

	console.OnFormatChange(func(i console.FormatInfo) {
		sep := i.Sep
		console.SetFormatter(entry{}, func(i any) (string, error) {
			e := i.(entry)
			return fmt.Sprintf("%d%s%s", e.Pos+1, sep, e.Stmt), nil
		})
	})

	return cmd
}

func History() {
	es := make([]entry, len(repl.GetHistory().Entries()))
	for i, e := range repl.GetHistory().Entries() {
		es[i] = entry{i, e}
	}
	console.Write(es)
}

func HistoryClear() {
	repl.GetHistory().Clear()
}

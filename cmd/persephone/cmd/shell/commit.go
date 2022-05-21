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
	"errors"

	"github.com/abc-inc/go-data-neo4j/graph"
	"github.com/abc-inc/persephone/cmd/persephone/cmd/cmdutil"
	"github.com/abc-inc/persephone/console"
	"github.com/spf13/cobra"
)

var errNoTxCommit = errors.New("there is no open transaction to commit")

func NewCmdCommit(f *cmdutil.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use:   ":commit",
		Short: "Commit the currently open transaction",
		Long:  "Commit and close the currently open transaction",
		Run:   func(cmd *cobra.Command, args []string) { Commit() },
	}

	return cmd
}

func Commit() {
	if ok, err := graph.GetConn().Commit(); err != nil {
		console.WriteErr(err)
	} else if !ok {
		console.WriteErr(errNoTxCommit)
	}
}

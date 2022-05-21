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
	"github.com/abc-inc/go-data-neo4j/graph"
	"github.com/abc-inc/persephone/cmd/persephone/cmd/cmdutil"
	"github.com/abc-inc/persephone/console"
	"github.com/spf13/cobra"
)

func NewCmdDisconnect(f *cmdutil.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use:   ":disconnect",
		Short: "Disconnect from database",
		Run:   func(cmd *cobra.Command, args []string) { Disconnect() },
	}

	return cmd
}

func Disconnect() {
	if err := graph.GetConn().Close(); err != nil {
		console.WriteErr(err)
	}
}

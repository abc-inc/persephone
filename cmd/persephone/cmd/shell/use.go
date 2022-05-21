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
	"sync"

	"github.com/abc-inc/go-data-neo4j/graph"
	cmd "github.com/abc-inc/persephone/cmd/persephone/cmd/browser"
	"github.com/abc-inc/persephone/cmd/persephone/cmd/cmdutil"
	"github.com/abc-inc/persephone/console"
	"github.com/abc-inc/persephone/console/repl"
	"github.com/spf13/cobra"
)

var dbsCache sync.Once

func NewCmdUse(f *cmdutil.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use:   ":use database",
		Short: "Set the active database",
		Long:  "Set the active database that transactions are executed on",
		Run:   func(cmd *cobra.Command, args []string) { Use(args[0]) },
	}

	return cmd
}

func Use(dbName string) {
	if err := graph.GetConn().UseDB(dbName); err != nil {
		console.WriteErr(err)
	}
}

func UseCompFunc() func(string) []repl.Item {
	var dbs []cmd.DBInfo
	return func(s string) (its []repl.Item) {
		dbsCache.Do(func() {
			dbs = cmd.ListDBs()
		})

		for _, db := range dbs {
			if strings.HasPrefix(db.Name, s) || strings.Contains(db.Address, s) {
				its = append(its, repl.Item{View: db.Name, Content: db.Address})
			}
		}
		return
	}
}

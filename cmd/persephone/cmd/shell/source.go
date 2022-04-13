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
	"bufio"
	"os"
	"strings"

	"github.com/abc-inc/persephone/console"
	"github.com/abc-inc/persephone/graph"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"github.com/spf13/cobra"
)

var SourceCmd = &cobra.Command{
	Use:   ":source [filename]",
	Short: "Execute Cypher statements from a file",
	Args:  cobra.ExactArgs(1),
	Run:   func(cmd *cobra.Command, args []string) { Source(args[0]) },
}

func Source(path string) {
	f, err := os.Open(path)
	if err != nil {
		console.WriteErr(err)
		return
	}

	defer func(f *os.File) {
		_ = f.Close()
	}(f)

	tx, created, err := graph.GetConn().GetTransaction()
	if err != nil {
		console.WriteErr(err)
		return
	} else if created {
		defer func(tx neo4j.Transaction) {
			_, _ = graph.GetConn().Rollback()
		}(tx)
	}

	sc := bufio.NewScanner(f)
	for sc.Scan() {
		if strings.HasPrefix(sc.Text(), ":") {
			continue
		}

		err := console.Query(graph.Request{Query: sc.Text(), Params: graph.GetConn().Params})
		if err != nil {
			console.WriteErr(err)
			return
		}
	}

	_, _ = graph.GetConn().Commit()
}

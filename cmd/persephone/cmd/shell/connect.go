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
	"os"

	cmd "github.com/abc-inc/persephone/cmd/persephone/cmd/persephone"
	"github.com/abc-inc/persephone/console"
	"github.com/abc-inc/persephone/graph"
	"github.com/abc-inc/persephone/internal"
	"github.com/mattn/go-isatty"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var ConnectCmd = &cobra.Command{
	Use:         ":connect",
	Short:       "Connect to a database",
	Annotations: cmd.Annotate(cmd.Offline),
	Run:         connectCmd,
}

func init() {
	ConnectCmd.Flags().StringP("username", "u", "", "username to connect as (default: neo4j). (env: NEO4J_USERNAME)")
	ConnectCmd.Flags().StringP("password", "p", "", "password to connect with (default: ). (env: NEO4J_PASSWORD)")
	ConnectCmd.Flags().StringP("database", "d", "neo4j", "database to connect to (default: neo4j). (env: NEO4J_DATABASE)")
}

func connectCmd(cmd *cobra.Command, args []string) {
	u := internal.Must(cmd.Flags().GetString("username"))
	p := internal.Must(cmd.Flags().GetString("password"))
	Connect(u, p)
}

func Connect(user, pass string) {
	if graph.GetConn() != nil && graph.GetConn().Driver != nil {
		console.Write("Already connected")
		return
	}

	addr := viper.GetString("address")
	db := viper.GetString("database")

	if isatty.IsTerminal(os.Stdin.Fd()) {
		if user == "" {
			user, pass = console.Input("username:", "neo4j"), ""
		}
		if pass == "" {
			pass = console.Pwd("password:")
		}
	}

	log.Info().Str("db", db).Str("addr", addr).Str("user", user).
		Msg("Connecting to Neo4j database")
	_ = graph.NewConn(addr, user, neo4j.BasicAuth(user, pass, ""), db)
}

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
	"os"

	"github.com/abc-inc/go-data-neo4j/graph"
	"github.com/abc-inc/persephone/cmd/persephone/cmd/cmdutil"
	cmd "github.com/abc-inc/persephone/cmd/persephone/cmd/persephone"
	"github.com/abc-inc/persephone/console"
	"github.com/abc-inc/persephone/internal"
	"github.com/mattn/go-isatty"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

func NewCmdConnect(f *cmdutil.Factory) *cobra.Command {
	run := func(cmd *cobra.Command, args []string) {
		u := internal.Must(cmd.Flags().GetString("username"))
		p := internal.Must(cmd.Flags().GetString("password"))
		d := internal.Must(cmd.Flags().GetString("database"))
		if err := Connect(*f.SessionConfig().Address, u, p, d); err != nil {
			console.WriteErr(err)
		}
	}

	cmd := &cobra.Command{
		Use:         ":connect",
		Short:       "Connect to a database",
		Annotations: cmd.Annotate(cmdutil.SkipAuth),
		Run:         run,
	}

	cmd.Flags().StringP("username", "u", "", "Username to connect as (default: neo4j) (env: NEO4J_USERNAME)")
	cmd.Flags().StringP("password", "p", "", "Password to connect with (default: ) (env: NEO4J_PASSWORD)")
	cmd.Flags().StringP("database", "d", "neo4j", "Database to connect to (default: neo4j) (env: NEO4J_DATABASE)")

	return cmd
}

func Connect(addr string, user, pass, db string) error {
	if graph.IsConnected() {
		console.Write("Already connected")
		return nil
	}

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

	auth, u := graph.Auth("basic:" + user + ":" + pass)
	_, err := graph.NewConn(addr, u, auth, db, func(config *neo4j.Config) {
		config.UserAgent = "Persephone (" + neo4j.UserAgent + ")"
	})
	return err
}

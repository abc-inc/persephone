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

package cmdutil

import (
	"os"

	"github.com/abc-inc/go-data-neo4j/graph"
	"github.com/abc-inc/persephone/config"
	"github.com/abc-inc/persephone/console"
	"github.com/fatih/color"
	"github.com/mattn/go-isatty"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

const SkipAuth = "skipAuthCheck"

// IsAuthCheckEnabled returns true if the given command or any of its parents
// requires authentication.
func IsAuthCheckEnabled(cmd *cobra.Command) bool {
	switch cmd.Name() {
	case "help", cobra.ShellCompRequestCmd, cobra.ShellCompNoDescRequestCmd:
		return false
	}

	for _, f := range []string{"driver-version", "help", "version"} {
		if f, err := cmd.Root().Flags().GetBool(f); err == nil && f {
			return false
		}
	}

	for c := cmd; c.Parent() != nil; c = c.Parent() {
		if c.Annotations != nil && c.Annotations[SkipAuth] == "true" {
			return false
		}
	}

	return true
}

func Connect(sessCfg *config.SessionConfig) {
	addr := *sessCfg.Address
	u := *sessCfg.Username
	p := *sessCfg.Password
	db := *sessCfg.Database

	if u == "" && isatty.IsTerminal(os.Stdin.Fd()) {
		u = console.Input("username:", "neo4j")
	}
	if p == "" && isatty.IsTerminal(os.Stdin.Fd()) {
		p = console.Pwd("password:")
	}

	log.Info().Str("db", db).Str("addr", addr).Str("user", u).
		Msg("Connecting to Neo4j database")

	auth, u := graph.Auth(u + ":" + p)
	_, err := graph.NewConn(addr, u, auth, db, func(config *neo4j.Config) {
		config.UserAgent = "persephone (" + neo4j.UserAgent + ")"
	})
	if err != nil {
		console.WriteErr(err)
		os.Exit(1)
	}

	if isatty.IsTerminal(os.Stdin.Fd()) {
		consCmdCol := color.New(color.FgCyan).Sprint
		log.Info().Msgf("Type %s for a list of available commands or %s to exit the shell.",
			consCmdCol(":help"), consCmdCol(":exit"))
		log.Info().Msg("Note that Cypher queries must end with a semicolon.")
	}
}

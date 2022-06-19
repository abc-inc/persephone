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
	"github.com/abc-inc/roland/graph"
	"github.com/fatih/color"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

func NewCmdStatus(f *cmdutil.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use:         ":status",
		Short:       "Show metadata for the currently open connection",
		Annotations: cmd.Annotate(cmdutil.SkipAuth),
		Run:         func(cmd *cobra.Command, args []string) { Status() },
	}

	return cmd
}

func Status() {
	if graph.IsConnected() {
		t := graph.GetConn().Driver.Target()
		log.Info().Str("as", graph.GetConn().Username()).Str("to", t.String()).Msg("You are connected")
	} else {
		log.Info().Msgf("You are currently not connected to Neo4j.\n"+
			"Execute %s and enter your credentials to connect.", color.CyanString(":connect"))
	}
}

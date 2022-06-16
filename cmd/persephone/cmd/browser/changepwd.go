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
	"errors"
	"os"

	"github.com/abc-inc/go-data-neo4j/graph"
	"github.com/abc-inc/persephone/cmd/persephone/cmd/cmdutil"
	"github.com/abc-inc/persephone/console"
	"github.com/fatih/color"
	"github.com/mattn/go-isatty"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

const cypSetPass = "ALTER CURRENT USER SET PASSWORD FROM $old TO $new"

var errPwdEmpty = errors.New("current and new password must be provided")
var errPwdMismatch = errors.New("the two entered passwords must be the same")

func NewCmdChangePass(f *cmdutil.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use:   ":change-password",
		Short: "Change the user password",
		Run:   func(cmd *cobra.Command, args []string) { ChangePass("", "", "") },
	}

	return cmd
}

func ChangePass(p, newP1, newP2 string) {
	if isatty.IsTerminal(os.Stdin.Fd()) {
		if p == "" {
			p = console.Pwd("old password:")
		}
		if newP1 == "" {
			newP1 = console.Pwd("new password:")
			newP2 = console.Pwd("new password (repeat):")
		}
	}

	if p == "" || newP1 == "" {
		console.WriteErr(errPwdEmpty)
		return
	}
	if newP1 != newP2 {
		console.WriteErr(errPwdMismatch)
		return
	}

	_, err := graph.GetConn().Session().WriteTransaction(func(tx neo4j.Transaction) (any, error) {
		return tx.Run(cypSetPass, map[string]any{"old": p, "new": newP1})
	})
	if err != nil {
		console.WriteErr(err)
	} else {
		log.Info().Msg(color.GreenString("Password changed successfully"))
	}
}

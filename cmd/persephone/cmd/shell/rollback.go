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

	"github.com/abc-inc/persephone/console"
	"github.com/abc-inc/persephone/graph"
	"github.com/spf13/cobra"
)

var errNoTxRollback = errors.New("there is no open transaction to rollback")

var RollbackCmd = &cobra.Command{
	Use:   ":rollback",
	Short: "Rollback the currently open transaction",
	Long:  "Rollback and close the currently open transaction",
	Run:   func(cmd *cobra.Command, args []string) { Rollback() },
}

func Rollback() {
	if ok, err := graph.GetConn().Rollback(); err != nil {
		console.WriteErr(err)
	} else if !ok {
		console.WriteErr(errNoTxRollback)
	}
}

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

var errTxActive = errors.New("there is already an open transaction")

var BeginCmd = &cobra.Command{
	Use:   ":begin",
	Short: "Open a transaction",
	Long:  "Start a transaction which will remain open until :commit or :rollback is called",
	Run:   func(cmd *cobra.Command, args []string) { Begin() },
}

func Begin() {
	_, created, err := graph.GetConn().GetTransaction()
	if err != nil {
		console.WriteErr(err)
	} else if !created {
		console.WriteErr(errTxActive)
	}
}

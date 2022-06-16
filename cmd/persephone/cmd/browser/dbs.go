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
	"github.com/abc-inc/persephone/console"
	"github.com/spf13/cobra"
)

func NewCmdDBs(f *cmdutil.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use:   ":dbs",
		Short: "Show databases available for the current user",
		Run:   func(cmd *cobra.Command, args []string) { DBs() },
	}

	return cmd
}

func DBs() {
	dbs := ListDBs()
	dbNames := make([]string, len(dbs))
	for i, db := range dbs {
		dbNames[i] = db.Name
	}
	console.Write(dbNames)
}

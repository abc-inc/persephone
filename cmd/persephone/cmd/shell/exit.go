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
	"path/filepath"

	"github.com/abc-inc/persephone/cmd/persephone/cmd/cmdutil"
	cmd "github.com/abc-inc/persephone/cmd/persephone/cmd/persephone"
	"github.com/abc-inc/persephone/config"
	"github.com/abc-inc/persephone/console"
	"github.com/abc-inc/persephone/console/repl"
	"github.com/abc-inc/persephone/internal"
	"github.com/spf13/cobra"
)

func NewCmdExit(f *cmdutil.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use:         ":exit",
		Short:       "Exit persephone",
		Annotations: cmd.Annotate(cmdutil.SkipAuth),
		Run:         func(cmd *cobra.Command, args []string) { Exit(f.Config()) },
	}

	return cmd
}

func Exit(cfg config.Config) {
	// Make sure that exit succeeds even if disconnect would fail.
	defer func() {
		if ex := recover(); ex != nil {
			os.Exit(1)
		}
		os.Exit(0)
	}()

	f := filepath.Join(internal.Must(os.UserCacheDir()), "persephone", "history")
	if err := repl.GetHistory().Save(f); err != nil {
		console.WriteErr(err)
	}

	console.WriteErr(cfg.Save())
	Disconnect()
}

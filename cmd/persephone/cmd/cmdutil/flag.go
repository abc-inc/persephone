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
	"github.com/abc-inc/persephone/internal"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

// ResetFlags visits all local flags and sets them to the default value.
// Additionally, it resets the "help" flag of the root command.
func ResetFlags(cmd *cobra.Command, _ []string) {
	f := cmd.Root().PersistentFlags().Lookup("help")
	internal.MustNoErr(f.Value.Set(f.DefValue))
	if cmd == cmd.Root() {
		return
	}

	cmd.LocalFlags().VisitAll(func(f *pflag.Flag) {
		internal.MustNoErr(f.Value.Set(f.DefValue))
	})
}
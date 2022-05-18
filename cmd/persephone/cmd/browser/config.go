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
	"github.com/abc-inc/persephone/cmd/persephone/cmd/cmdutil"
	cmd "github.com/abc-inc/persephone/cmd/persephone/cmd/persephone"
	"github.com/abc-inc/persephone/config"
	"github.com/abc-inc/persephone/console"
	"github.com/abc-inc/persephone/internal"
	"github.com/spf13/cobra"
)

func NewCmdConfig(f *cmdutil.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use:         ":config COMMAND",
		Short:       "Get and set config options",
		Annotations: cmd.Annotate(cmdutil.SkipAuth),
	}

	cmd.AddCommand(
		NewCmdConfigGet(f),
		NewCmdConfigList(f),
		NewCmdConfigSet(f),
	)

	return cmd
}

func NewCmdConfigGet(f *cmdutil.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use:         "get <key>",
		Short:       "Print the value of a given configuration key",
		Args:        cobra.ExactArgs(1),
		Annotations: cmd.Annotate(cmdutil.SkipAuth),
		Run: func(cmd *cobra.Command, args []string) {
			console.Write(GetConfig(f.Config(), args[0]))
		},
	}

	return cmd
}

func NewCmdConfigList(f *cmdutil.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use:         "list",
		Short:       "Print a list of given configuration keys and values",
		Annotations: cmd.Annotate(cmdutil.SkipAuth),
		Run: func(cmd *cobra.Command, args []string) {
			console.Write(ListConfig(f.Config()))
		},
	}

	return cmd
}

func NewCmdConfigSet(f *cmdutil.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use:         "set <key> <value>",
		Short:       "Update configuration with a value for the given key",
		Args:        cobra.ExactArgs(2),
		Annotations: cmd.Annotate(cmdutil.SkipAuth),
		Run: func(cmd *cobra.Command, args []string) {
			SetConfig(f.Config(), args[0], args[1])
		},
	}

	return cmd
}

func ListConfig(cfg config.Config) map[string]any {
	return cfg.List()
}

func GetConfig(cfg config.Config, key string) any {
	return cfg.Get(key, nil)
}

func SetConfig(cfg config.Config, key, val string) {
	switch key {
	case "format":
		console.ChangeFmt(val)
	default:
		cfg.Set(key, internal.Parse(val))
	}
}

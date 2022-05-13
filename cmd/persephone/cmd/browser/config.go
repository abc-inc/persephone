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

	"github.com/abc-inc/persephone/cmd/persephone/cmd/cmdutil"
	cmd "github.com/abc-inc/persephone/cmd/persephone/cmd/persephone"
	"github.com/abc-inc/persephone/config"
	"github.com/abc-inc/persephone/console"
	"github.com/abc-inc/persephone/internal"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var errInvalidArgs = errors.New("invalid arguments")

func NewCmdConfig(f *cmdutil.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use:         ":config [name [value]]",
		Short:       "Get and set config options",
		Annotations: cmd.Annotate(cmdutil.SkipAuth),
		Run: func(cmd *cobra.Command, args []string) {
			configCmd(f.Config(), cmd, args)
		},
	}

	return cmd
}

func configCmd(cfg config.Config, cmd *cobra.Command, args []string) {
	switch len(args) {
	case 0:
		console.Write(ListConfig())
	case 1:
		console.Write(GetConfig(args[0]))
	case 2:
		SetConfig(args[0], args[1])
	default:
		console.WriteErr(errInvalidArgs)
	}
}

func ListConfig() map[string]any {
	return viper.AllSettings()
}

func GetConfig(key string) any {
	return viper.Get(key)
}

func SetConfig(key, val string) {
	viper.Set(key, internal.Parse(val))
}

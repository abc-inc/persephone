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
	"strings"

	"github.com/abc-inc/persephone/cmd/persephone/cmd/cmdutil"
	"github.com/abc-inc/persephone/config"
	"github.com/abc-inc/persephone/console/repl"
	"github.com/alecthomas/chroma/styles"
	"github.com/spf13/cobra"
	"golang.org/x/exp/slices"
)

var errSchemaNotFound = errors.New("color scheme not found")

func NewCmdColors(f *cmdutil.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use:         ":colors <scheme>",
		Short:       "Set the color scheme for the prompt input",
		Args:        cobra.ExactArgs(1),
		Annotations: Annotate(cmdutil.SkipAuth),
		RunE: func(cmd *cobra.Command, args []string) error {
			return Colors(f.Config(), args[0])
		},
	}

	return cmd
}

func Colors(cfg config.Config, s string) error {
	if !slices.Contains(styles.Names(), s) {
		return errSchemaNotFound
	}
	cfg.Set("colors", s)
	return nil
}

func ColorsComp(s string) (its []repl.Item) {
	for _, n := range styles.Names() {
		if strings.HasPrefix(n, s) {
			its = append(its, repl.Item{View: n})
		}
	}
	return
}

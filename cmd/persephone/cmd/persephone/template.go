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
	"path/filepath"
	"strings"

	"github.com/abc-inc/browser"
	"github.com/abc-inc/persephone/console"
	"github.com/abc-inc/persephone/internal"
	"github.com/spf13/cobra"
)

var errTmplOpen = errors.New("cannot open template")
var errTmplNotFound = errors.New("template does not exist")

var TemplateCmd = &cobra.Command{
	Use:         ":template",
	Short:       "Define a template",
	Args:        cobra.MinimumNArgs(1),
	Annotations: Annotate(Offline),
}

var TemplateEditCmd = &cobra.Command{
	Use:         "edit",
	Short:       "Open a template in the default editor",
	Args:        cobra.ExactArgs(1),
	Annotations: Annotate(Offline),
	Run: func(cmd *cobra.Command, args []string) {
		console.WriteErr(TemplateEdit(args[0]))
	},
}

var TemplateGetCmd = &cobra.Command{
	Use:         "get",
	Short:       "Print the templates",
	Args:        cobra.ExactArgs(1),
	Annotations: Annotate(Offline),
	Run: func(cmd *cobra.Command, args []string) {
		console.WriteErr(TemplateGet(args[0]))
	},
}

var TemplateListCmd = &cobra.Command{
	Use:         "list",
	Short:       "List all templates",
	Args:        cobra.ExactArgs(0),
	Annotations: Annotate(Offline),
	Run: func(cmd *cobra.Command, args []string) {
		TemplateList()
	},
}

var TemplateSetCmd = &cobra.Command{
	Use:         "set",
	Short:       "Define a template for the current session",
	Args:        cobra.MinimumNArgs(2),
	Annotations: Annotate(Offline),
	Run: func(cmd *cobra.Command, args []string) {
		console.WriteErr(TemplateSet(args[0], strings.Join(args[1:], " ")))
	},
}

var TemplateWriteCmd = &cobra.Command{
	Use:         "write",
	Short:       "Define a template and save it",
	Args:        cobra.MinimumNArgs(2),
	Annotations: Annotate(Offline),
	Run: func(cmd *cobra.Command, args []string) {
		console.WriteErr(TemplateWrite(args[0], strings.Join(args[1:], " ")))
	},
}

func init() {
	TemplateCmd.Flags().StringP("template", "t", "", "template")
	TemplateCmd.AddCommand(
		TemplateEditCmd,
		TemplateGetCmd,
		TemplateListCmd,
		TemplateSetCmd,
		TemplateWriteCmd,
	)
}

func TemplateEdit(path string) error {
	if !browser.Open(filepath.Join(console.TmplDir, filepath.Base(path))) {
		return errTmplOpen
	}
	return nil
}

func TemplateGet(path string) error {
	if t := console.GetTmplMgr().Get(path); t != nil {
		console.Write(t.Root.String())
		return nil
	}
	return errTmplNotFound
}

func TemplateList() {
	ts := []console.NamedTemplate{}
	for p, t := range console.GetTmplMgr().TmplsByPath {
		b := filepath.Base(p)
		n := strings.TrimSuffix(b, console.TmplExt)
		t := console.NamedTemplate{n, t.Root.String(), b != p}
		ts = append(ts, t)
	}
	console.Write(ts)
}

func TemplateSet(name, text string) error {
	return internal.Second(console.GetTmplMgr().Set(filepath.Base(name), text))
}

func TemplateWrite(name, text string) error {
	p := filepath.Join(console.TmplDir, filepath.Base(name))
	return internal.Second(console.GetTmplMgr().Set(p, text))
}

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
	"path/filepath"
	"strings"

	"github.com/abc-inc/browser"
	"github.com/abc-inc/persephone/cmd/persephone/cmd/cmdutil"
	"github.com/abc-inc/persephone/console"
	"github.com/abc-inc/persephone/internal"
	"github.com/spf13/cobra"
)

var errTmplOpen = errors.New("cannot open template")
var errTmplNotFound = errors.New("template does not exist")

func NewCmdTemplate(f *cmdutil.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use:         ":template",
		Short:       "Define a template",
		Args:        cobra.MinimumNArgs(1),
		Annotations: Annotate(cmdutil.SkipAuth),
	}

	cmd.Flags().StringP("template", "t", "", "template")
	cmd.AddCommand(
		NewCmdTemplateEdit(f),
		NewCmdTemplateGet(f),
		NewCmdTemplateList(f),
		NewCmdTemplateSet(f),
		NewCmdTemplateWrite(f),
	)

	return cmd
}

func NewCmdTemplateEdit(f *cmdutil.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use:         "edit",
		Short:       "Open a template in the default editor",
		Args:        cobra.ExactArgs(1),
		Annotations: Annotate(cmdutil.SkipAuth),
		Run: func(cmd *cobra.Command, args []string) {
			console.WriteErr(TemplateEdit(args[0]))
		},
		Hidden: true, // hidden for now, because the cache is not synced
	}

	return cmd
}

func NewCmdTemplateGet(f *cmdutil.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use:         "get",
		Short:       "Print the templates",
		Args:        cobra.ExactArgs(1),
		Annotations: Annotate(cmdutil.SkipAuth),
		Run: func(cmd *cobra.Command, args []string) {
			console.WriteErr(TemplateGet(args[0]))
		},
	}

	return cmd
}

func NewCmdTemplateList(f *cmdutil.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use:         "list",
		Short:       "List all templates",
		Args:        cobra.ExactArgs(0),
		Annotations: Annotate(cmdutil.SkipAuth),
		Run: func(cmd *cobra.Command, args []string) {
			TemplateList()
		},
	}

	return cmd
}

func NewCmdTemplateSet(f *cmdutil.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use:         "set",
		Short:       "Define a template for the current session",
		Args:        cobra.MinimumNArgs(2),
		Annotations: Annotate(cmdutil.SkipAuth),
		Run: func(cmd *cobra.Command, args []string) {
			console.WriteErr(TemplateSet(args[0], strings.Join(args[1:], " ")))
		},
	}

	return cmd
}

func NewCmdTemplateWrite(f *cmdutil.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use:         "write",
		Short:       "Define a template and save it",
		Args:        cobra.MinimumNArgs(2),
		Annotations: Annotate(cmdutil.SkipAuth),
		Run: func(cmd *cobra.Command, args []string) {
			console.WriteErr(TemplateWrite(args[0], strings.Join(args[1:], " ")))
		},
	}

	return cmd
}

func TemplateEdit(path string) error {
	if !browser.Open(filepath.Join(console.TmplDir, console.TmplFileName(path))) {
		return errTmplOpen
	}
	return nil
}

func TemplateGet(path string) error {
	if t := console.GetTmplMgr().Get(path); t != nil {
		console.Write(strings.TrimSuffix(t.Root.String(), "\n"))
		return nil
	}
	return errTmplNotFound
}

func TemplateList() {
	ts := []console.NamedTemplate{}
	for p, t := range console.GetTmplMgr().TmplsByPath {
		n := strings.TrimSuffix(p, console.TmplExt)
		txt := strings.TrimSuffix(t.Root.String(), "\n")
		ts = append(ts, console.NamedTemplate{Name: n, Tmpl: txt, Persistent: n != p})
	}
	console.Write(ts)
}

func TemplateSet(name, text string) error {
	name = strings.TrimSuffix(filepath.Base(name), console.TmplExt)
	return internal.Second(console.GetTmplMgr().Set(name, text))
}

func TemplateWrite(name, text string) error {
	return internal.Second(console.GetTmplMgr().Set(console.TmplFileName(name), text))
}

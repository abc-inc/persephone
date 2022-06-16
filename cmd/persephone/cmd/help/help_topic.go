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

package help

import (
	"fmt"
	"path"
	"strings"

	"github.com/abc-inc/persephone/docs"
	"github.com/abc-inc/persephone/internal"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

type Topic struct {
	File string
	Desc string
}

var Topics = []Topic{
	{File: "help/environment.md", Desc: "Environment variables that can be used with Persephone"},
	{File: "help/formatting.md", Desc: "Formatting options for CSV/JSON/YAML data exported from Persephone"},
	{File: "contributing.md", Desc: "Information to get started"},
	{File: "faq.md", Desc: "Frequently Asked Questions"},
	{File: "source.md", Desc: "Instructions for building Persephone from source"},
	{File: "why.md", Desc: fmt.Sprintf("Why %s?", color.New(color.Italic).Sprint("Persephone"))},
}

func NewCmdHelpTopic(t Topic) *cobra.Command {
	n := strings.TrimSuffix(path.Base(t.File), ".md")
	cmd := &cobra.Command{
		Use:    n,
		Short:  t.Desc,
		Hidden: true,
	}

	cmd.SetHelpFunc(func(cmd *cobra.Command, args []string) {
		md := string(internal.Must(docs.FS.ReadFile(t.File)))
		cmd.Print(internal.Must(docs.Render(md)))
	})

	return cmd
}

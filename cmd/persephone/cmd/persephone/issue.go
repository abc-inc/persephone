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
	"fmt"
	"os"

	"github.com/abc-inc/browser"
	"github.com/abc-inc/persephone/cmd/persephone/cmd/cmdutil"
	"github.com/abc-inc/persephone/internal"
	"github.com/spf13/cobra"
)

const ghIssuesURL = "https://github.com/abc-inc/persephone/issues"

func NewCmdIssue(f *cmdutil.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use:         ":issue",
		Short:       "Report an issue on GitHub",
		Annotations: Annotate(cmdutil.SkipAuth),
		Run:         func(cmd *cobra.Command, args []string) { Issue() },
	}

	return cmd
}

func Issue() {
	if !browser.Open(ghIssuesURL) {
		internal.Must(fmt.Fprintln(os.Stderr, "Please report issues to "+ghIssuesURL))
	}
}

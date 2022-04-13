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
	"fmt"
	"os"

	"github.com/abc-inc/browser"
	"github.com/abc-inc/persephone/internal"
	"github.com/spf13/cobra"
)

const ghIssuesUrl = "https://github.com/abc-inc/persephone/issues"

var IssueCmd = &cobra.Command{
	Use:         ":issue",
	Short:       "Report an issue on GitHub",
	Annotations: Annotate(Offline),
	Run:         func(cmd *cobra.Command, args []string) { Issue() },
}

func Issue() {
	if !browser.Open(ghIssuesUrl) {
		internal.Must(fmt.Fprintln(os.Stderr, "Please report issues to "+ghIssuesUrl))
	}
}

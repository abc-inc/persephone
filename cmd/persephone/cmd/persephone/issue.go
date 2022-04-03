package cmd

import (
	"fmt"
	"os"

	"github.com/abc-inc/browser"
	"github.com/abc-inc/persephone/cmd/persephone/cmd/types"
	"github.com/abc-inc/persephone/internal"
	"github.com/spf13/cobra"
)

const ghIssuesUrl = "https://github.com/abc-inc/persephone/issues"

var IssueCmd = &cobra.Command{
	Use:         ":issue",
	Short:       "Report an issue on GitHub",
	Annotations: types.Annotate(types.Offline),
	Run:         func(cmd *cobra.Command, args []string) { Issue() },
}

func Issue() {
	if !browser.Open(ghIssuesUrl) {
		internal.Must(fmt.Fprintln(os.Stderr, "Please report issues to "+ghIssuesUrl))
	}
}

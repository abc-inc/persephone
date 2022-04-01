package cmd

import (
	"fmt"
	"os"

	"github.com/abc-inc/browser"
	"github.com/abc-inc/persephone/cmd/persephone/cmd/types"
	"github.com/abc-inc/persephone/internal"
	"github.com/spf13/cobra"
)

var IssueCmd = &cobra.Command{
	Use:         ":issue",
	Short:       "Report an issue on GitHub",
	Annotations: types.Annotate(types.Offline),
	Run:         issueCmd,
}

func issueCmd(cmd *cobra.Command, args []string) {
	var url = "https://github.com/abc-inc/persephone/issues"
	if !browser.Open(url) {
		internal.Must(fmt.Fprintln(os.Stderr, "Please report issues to "+url))
	}
}

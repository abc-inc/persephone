package cmd

import (
	"fmt"
	"os"

	"github.com/abc-inc/browser"
	"github.com/spf13/cobra"
)

var IssueCmd = &cobra.Command{
	Use: ":issue",
	Short: "Report an issue on GitHub",
	Run: issueCmd,
}

func issueCmd(cmd *cobra.Command, args []string) {
	var url = "https://github.com/abc-inc/persephone/issues/"
	if !browser.Open(url) {
		fmt.Fprintln(os.Stderr, "Please report issues to: "+url)
	}
}

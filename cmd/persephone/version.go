package main

import (
	"fmt"

	"github.com/abc-inc/persephone/cmd/persephone/cmd/persephone"
	"github.com/spf13/cobra"
)

var version string

var VersionCmd = &cobra.Command{
	Use:         ":version",
	Aliases:     []string{"version"},
	Short:       "Print version information and exit.",
	Annotations: cmd.Annotate(cmd.Offline),
	Run:         versionCmd,
}

func versionCmd(cmd *cobra.Command, args []string) {
	fmt.Println("persephone ", version)
}

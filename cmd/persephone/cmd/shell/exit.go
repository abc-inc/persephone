package cmd

import (
	"os"

	"github.com/abc-inc/persephone/cmd/persephone/cmd/types"
	"github.com/abc-inc/persephone/hist"
	"github.com/spf13/cobra"
)

var ExitCmd = &cobra.Command{
	Use:         ":exit",
	Short:       "Exit persephone",
	Annotations: types.Annotate(types.Offline),
	Run:         exitCmd,
}

func exitCmd(cmd *cobra.Command, args []string) {
	// Make sure that exit succeeds even if disconnect would fail.
	defer func() {
		if ex := recover(); ex != nil {
			os.Exit(0)
		}
	}()

	hist.Get().Save()

	for _, c := range cmd.Root().Commands() {
		if c.Name() == ":disconnect" {
			c.Run(cmd, args)
			os.Exit(0)
		}
	}
}

package cmd

import (
	"os"
	"path/filepath"

	"github.com/abc-inc/persephone/cmd/persephone/cmd/persephone"
	"github.com/abc-inc/persephone/console"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var ExitCmd = &cobra.Command{
	Use:         ":exit",
	Short:       "Exit persephone",
	Annotations: cmd.Annotate(cmd.Offline),
	Run:         func(cmd *cobra.Command, args []string) { Exit() },
}

func Exit() {
	// Make sure that exit succeeds even if disconnect would fail.
	defer func() {
		if ex := recover(); ex != nil {
			os.Exit(0)
		}
	}()

	console.Get().Save()
	if f := viper.GetViper().ConfigFileUsed(); f != "" {
		os.MkdirAll(filepath.Dir(f), 0700)
		if err := viper.WriteConfig(); err != nil {
			console.Writeln(err)
		}
	}

	Disconnect()
	os.Exit(0)
}

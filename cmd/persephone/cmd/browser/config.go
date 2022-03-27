package cmd

import (
	"fmt"

	"github.com/abc-inc/persephone/format"
	"github.com/abc-inc/persephone/internal"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var ConfigCmd = &cobra.Command{
	Use:   ":config [name [value]]",
	Short: "Print the config or change it.",
	Run:   configCmd,
}

func configCmd(cmd *cobra.Command, args []string) {
	switch len(args) {
	case 0:
		cfg := viper.AllSettings()
		delete(cfg, "config")
		delete(cfg, "password")
		format.Writeln(cfg)
		break
	case 1:
		format.Writeln(viper.Get(args[0]))
		break
	case 2:
		viper.Set(args[0], internal.Parse(args[1]))
		break
	default:
		fmt.Println("error: invalid arguments")
	}
}

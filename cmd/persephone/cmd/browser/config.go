package cmd

import (
	"errors"

	"github.com/abc-inc/persephone/cmd/persephone/cmd/persephone"
	"github.com/abc-inc/persephone/console"
	"github.com/abc-inc/persephone/internal"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var errInvalidArgs = errors.New("invalid arguments")

var ConfigCmd = &cobra.Command{
	Use:         ":config [name [value]]",
	Short:       "Get and set config options",
	Annotations: cmd.Annotate(cmd.Offline),
	Run:         configCmd,
}

func configCmd(cmd *cobra.Command, args []string) {
	switch len(args) {
	case 0:
		console.Writeln(ListConfig())
	case 1:
		console.Writeln(GetConfig(args[0]))
	case 2:
		SetConfig(args[0], args[1])
	default:
		console.Writeln(errInvalidArgs)
	}
}

func ListConfig() map[string]interface{} {
	cfg := viper.AllSettings()
	delete(cfg, "config")
	delete(cfg, "password")
	return cfg
}

func GetConfig(key string) interface{} {
	return viper.Get(key)
}

func SetConfig(key, val string) {
	viper.Set(key, internal.Parse(val))
}

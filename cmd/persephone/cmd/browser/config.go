package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/abc-inc/persephone/internal"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var ConfigCmd = &cobra.Command{
	Use: ":config",
	Short: "Print the config or change it.",
	Run: configCmd,
}

func configCmd(cmd *cobra.Command, args []string) {
	cfg := viper.AllSettings()
	delete(cfg, "password")
	j := internal.Must(json.MarshalIndent(cfg, "", "  "))
	fmt.Println(string(j))
}

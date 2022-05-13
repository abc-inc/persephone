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

package main

import (
	"errors"
	stdlog "log"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/abc-inc/persephone/cmd/persephone/cmd/cmdutil"
	cmd "github.com/abc-inc/persephone/cmd/persephone/cmd/root"
	"github.com/abc-inc/persephone/config"
	"github.com/abc-inc/persephone/console"
	"github.com/abc-inc/persephone/console/repl"
	"github.com/abc-inc/persephone/internal"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

// init initializes the all loggers to meaningful default values.
func init() {
	stdlog.SetFlags(0)
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	zerolog.TimeFieldFormat = "15:04:05"
	log.Logger = log.Output(zerolog.ConsoleWriter{
		Out:        os.Stderr,
		NoColor:    false,
		TimeFormat: zerolog.TimeFieldFormat,
	})

	console.ChangeFmt("")
}

func main() {
	sessCfg := config.NewSessionConfig()
	cfg := config.FromFile(*sessCfg.CfgFile)
	f := cmdutil.NewFactory(cfg, sessCfg)
	rootCmd := cmd.NewCmdRoot(f)

	var o sync.Once
	cobra.OnInitialize(func() {
		o.Do(func() { onCfgLoaded(cfg, sessCfg, rootCmd) })
	})

	defer func(cfg config.Config) {
		if err := cfg.Save(); err != nil {
			log.Warn().Err(err).Msg("Cannot save config")
		}
	}(cfg)
	cobra.CheckErr(rootCmd.Execute())
}

// onCfgLoaded is invoked after the CLI is initialized and all configs are loaded.
func onCfgLoaded(cfg config.Config, sessCfg *config.SessionConfig, rootCmd *cobra.Command) {
	// Reconfigure Logging
	ll := strings.ToLower(internal.Must(rootCmd.Flags().GetString("log-level")))
	l := internal.Must(zerolog.ParseLevel(ll))
	zerolog.SetGlobalLevel(l)

	// Make sure that config is loaded
	if err := cfg.Load(); !errors.Is(err, os.ErrNotExist) {
		internal.MustNoErr(err)
	}

	// Setup output format
	if f := internal.Must(rootCmd.PersistentFlags().GetString("format")); f != "" {
		console.ChangeFmt(f)
	} else {
		console.ChangeFmt(cfg.Get("format", "auto").(string))
	}
	console.OnFormatChange(func(i console.FormatInfo) {
		cfg.Set("format", i.Format)
		*sessCfg.Format = i.Format
	})

	// Set credentials, if available
	if p, ok := os.LookupEnv("NEO4J_PASSWORD"); ok {
		*sessCfg.Password = p
	}

	f := filepath.Join(internal.Must(os.UserCacheDir()), "persephone", "history")
	_ = repl.GetHistory().Load(f)
}

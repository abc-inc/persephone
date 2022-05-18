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

package config

import (
	"os"
	"path/filepath"

	"github.com/abc-inc/persephone/internal"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

// Config holds persistent configuration.
type Config interface {
	Get(string, any) any
	Set(string, any)
	List() map[string]any
	Load() error
	Save() error
}

// viperConfig uses Viper to lookup and save configuration.
type viperConfig struct {
}

// SessionConfig holds config, which should be forgotten upon exit.
type SessionConfig struct {
	CfgFile  *string
	Address  *string
	Database *string
	Username *string
	Password *string
	Format   *string
}

// NewSessionConfig initializes a new SessionConfig with default values.
func NewSessionConfig() *SessionConfig {
	cfgFile := filepath.Join(internal.Must(os.UserConfigDir()), "persephone", "config.yaml")
	address := "neo4j://localhost:7687"
	database := "neo4j"
	username := "neo4j"
	password := ""
	format := ""

	return &SessionConfig{
		CfgFile:  &cfgFile,
		Address:  &address,
		Database: &database,
		Username: &username,
		Password: &password,
		Format:   &format,
	}
}

// FromFile creates a new Config with environment lookup capabilities.
// Note that Load() must be called before using it.
func FromFile(path string) Config {
	cfg := &viperConfig{}
	viper.AddConfigPath(path)
	if path != "" {
		viper.SetConfigFile(path)
	}
	viper.SetEnvPrefix("NEO4J")
	viper.AutomaticEnv()
	return cfg
}

// Get returns the value for the given key.
// If the key is not set, def is returned instead.
func (c *viperConfig) Get(key string, def any) any {
	if !viper.IsSet(key) {
		return def
	}
	return viper.Get(key)
}

// Set changes the value for the given key.
func (c *viperConfig) Set(key string, val any) {
	viper.Set(key, val)
}

// List returns a copy of all settings.
func (c *viperConfig) List() map[string]any {
	return viper.AllSettings()
}

// Load loads the config from the pre-configured file.
func (c *viperConfig) Load() error {
	log.Debug().Str("file", viper.ConfigFileUsed()).Msg("Loading config")
	return viper.ReadInConfig()
}

// Save writes the config to the pre-configured file.
func (c *viperConfig) Save() error {
	f := viper.GetViper().ConfigFileUsed()
	if f == "" || len(viper.AllKeys()) == 0 {
		return nil
	}
	log.Debug().Str("file", viper.ConfigFileUsed()).Msg("Saving config")
	if err := os.MkdirAll(filepath.Dir(f), 0700); err != nil {
		return err
	}
	return viper.WriteConfig()
}

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

package cmdutil

import "github.com/abc-inc/persephone/config"

type Factory struct {
	Config        func() config.Config
	SessionConfig func() *config.SessionConfig
}

func NewFactory(cfg config.Config, sessCfg *config.SessionConfig) *Factory {
	return &Factory{
		func() config.Config { return cfg },
		func() *config.SessionConfig { return sessCfg },
	}
}

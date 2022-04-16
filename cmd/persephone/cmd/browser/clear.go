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

package cmd

import (
	cmd "github.com/abc-inc/persephone/cmd/persephone/cmd/persephone"
	"github.com/inancgumus/screen"
	"github.com/spf13/cobra"
)

var ClearCmd = &cobra.Command{
	Use:         ":clear",
	Short:       "Clear the screen",
	Annotations: cmd.Annotate(cmd.Offline),
	Run:         func(cmd *cobra.Command, args []string) { Clear() },
}

func Clear() {
	screen.Clear()
	screen.MoveTopLeft()
}

// Copyright 2022 The Persephone authors
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

package repl

import (
	"strings"

	"github.com/spf13/cobra"
)

type Item struct {
	View    string
	Content string
}

func (i Item) String() string {
	return i.View
}

type CompFunc func(str string) []Item

func NoComp(string) []Item {
	return nil
}

func SubCmdComp(cmd *cobra.Command) func(str string) []Item {
	return func(s string) (its []Item) {
		for _, c := range cmd.Commands() {
			if strings.HasPrefix(c.Name(), s) && !c.Hidden {
				its = append(its, Item{View: c.Name()})
			}
		}
		return
	}
}

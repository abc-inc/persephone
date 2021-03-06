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

package comp

import (
	"github.com/abc-inc/persephone/comp/rule"
	"github.com/abc-inc/persephone/types"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// evaluateRules applies the Rules to the given ParseTree and returns the first
// non-empty result of Rule.
func evaluateRules(e antlr.ParseTree) []rule.Info {
	for _, r := range rule.OrderedRules {
		if items := r(e); len(items) > 0 {
			return items
		}
	}
	return nil
}

// GetTypes returns the completion types of the first Rule matching the given
// ParseTree.
func GetTypes(e antlr.Tree) Info {
	// If element is nil, then no types
	if e == nil {
		return Info{
			Found: false,
			Types: types.AllCompData,
		}
	}

	// Retrieve types from rules
	if infos := evaluateRules(e.(antlr.ParseTree)); len(infos) > 0 {
		ts := make([]types.Data, len(infos))
		for i, it := range infos {
			ts[i] = types.Data{Type: it.Type, Path: it.Path, FilterLastElement: it.Found}
		}
		return Info{
			Found: true,
			Types: ts,
		}
	}

	// If no types found, then no types
	return Info{
		Found: false,
		Types: types.AllCompData,
	}
}

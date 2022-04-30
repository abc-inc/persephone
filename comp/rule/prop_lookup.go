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

package rule

import (
	"github.com/abc-inc/persephone/types"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/gschauer/cypher2go/v4/parser"
)

// rulePropertyLookup checks whether we are in any property context, and then
// returns property key completion.
func rulePropertyLookup(e antlr.ParseTree) []Info {
	if parentCtx := e.GetParent(); parentCtx != nil {
		if lookupCtx := parentCtx.GetParent(); lookupCtx != nil {
			if _, ok := lookupCtx.(*parser.PropertyLookupContext); ok && e.GetText() == "." {
				return []Info{{Type: types.PropertyKey}}
			}
			// TODO: why is this necessary? the JavaScript implementation does not need it
			if _, ok := lookupCtx.(*parser.PropertyExpressionContext); ok {
				return []Info{{Type: types.PropertyKey}}
			}
			// TODO: why is this necessary? the JavaScript implementation does not need it
			if _, ok := lookupCtx.(*parser.PropertyOrLabelsExpressionContext); ok {
				return []Info{{Type: types.PropertyKey}}
			}
		}
	}
	return nil
}

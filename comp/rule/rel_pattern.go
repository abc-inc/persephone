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

package rule

import (
	"reflect"

	"github.com/abc-inc/persephone/ast"
	"github.com/abc-inc/persephone/types"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/gschauer/cypher2go/v4/parser"
)

// ruleRelationshipPattern checks if we are in relationship pattern, and then
// returns variables and types.
func ruleRelationshipPattern(e antlr.ParseTree) []Info {
	parent := ast.FindParent(e, reflect.TypeOf(parser.RelationshipPatternContext{}))
	if parent == nil {
		return nil
	}

	// We are at the beginning, so allow variables too
	if e.GetText() == "[" {
		return []Info{{Type: types.Variable}, {Type: types.RelationshipType}}
	}
	// We are at the end, fail and allow algorithm to get back by 1 char
	if e.GetText() == "]" {
		return nil
	}
	return nil
}

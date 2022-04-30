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
	"reflect"
	"strings"

	"github.com/abc-inc/persephone/ast"
	"github.com/abc-inc/persephone/types"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/gschauer/cypher2go/v4/parser"
)

// rulePropInMapLiteral checks whether we are in a map, and then returns
// parameter and/or property key completion.
func rulePropInMapLiteral(e antlr.ParseTree) []Info {
	mapLitCtx := ast.FindParent(e, reflect.TypeOf(parser.MapLiteralContext{}))
	if mapLitCtx != nil {
		if e.GetText() == "}" {
			return nil
		}
		return []Info{{Type: types.PropertyKey}}
	}

	propCtx := ast.FindParent(e, reflect.TypeOf(parser.PropertiesContext{}))
	if propCtx != nil {
		if e.GetText() == "}" || strings.TrimSpace(e.GetText()) == "" {
			return nil
		}
		return []Info{{Type: types.PropertyKey}, {Type: types.Parameter}}
	}

	return nil
}

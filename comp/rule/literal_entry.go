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

func ruleLiteralEntry(e antlr.ParseTree) (is []Info) {
	literalEntry := ast.FindParent(e, reflect.TypeOf(parser.LiteralEntryContext{}))
	if literalEntry == nil {
		return nil
	}
	if literalEntry.GetChildCount() < 2 {
		return is
	}

	doubleDots := literalEntry.GetChild(1)
	var space antlr.Tree
	if literalEntry.GetChildCount() > 2 {
		space = literalEntry.GetChild(2)
	}
	if doubleDots == e || space == e {
		for _, t := range types.AllComp {
			is = append(is, Info{Type: t})
		}
	}
	return is
}

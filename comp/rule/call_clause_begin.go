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
	"github.com/abc-inc/persephone/ast"
	"github.com/abc-inc/persephone/types"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/gschauer/cypher2go/v4/parser"
)

// If we are in call rule, and element is second child of call return procedure types
func ruleCallClauseBeginning(e antlr.ParseTree) []Info {
	parent := ast.GetParent(e)
	if parent == nil {
		return nil
	}

	if _, ok := parent.(*parser.CallContext); ok {
		if parent.GetChildCount() > 1 && parent.GetChild(1) == e {
			return []Info{{Type: types.ProcedureName}}
		}
	}
	return nil
}
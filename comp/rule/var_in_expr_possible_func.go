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
	"github.com/abc-inc/persephone/ast"
	"github.com/abc-inc/persephone/lang"
	"github.com/abc-inc/persephone/types"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// ruleVariableInExpressionPossibleFunction checks if a variable is inside an
// expression context, and then it might be a variable or a function.
func ruleVariableInExpressionPossibleFunction(e antlr.ParseTree) []Info {
	variable := ast.FindAnyParent(e, []string{lang.VariableContext})
	expression := ast.FindAnyParent(variable, []string{lang.ExpressionContext})
	if variable != nil && expression != nil {
		return []Info{{Type: types.Variable}, {Type: types.FunctionName}}
	}
	return nil
}

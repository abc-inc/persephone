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
	"github.com/abc-inc/persephone/lang"
	"github.com/abc-inc/persephone/types"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/gschauer/cypher2go/v4/parser"
)

// Return procedure output completion if we are inside procedure
func ruleProcedureOutputsInCallClause(e antlr.ParseTree) []Info {
	call := ast.FindAnyParent(e, []string{lang.CALL_CONTEXT})
	if call == nil {
		return nil
	}

	proc := ast.FindChild(call, lang.PROCEDURE_NAME_CONTEXT)
	resOutput := ast.FindAnyParent(e, []string{lang.PROCEDURE_RESULTS_CONTEXT})
	if proc == nil || resOutput == nil {
		return nil
	}
	return []Info{{
		Type: types.ProcedureOutput,
		Path: []string{proc.(*parser.ProcedureInvocationBodyContext).GetText()},
	}}
}

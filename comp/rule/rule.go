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
)

type Rule func(e antlr.ParseTree) []Info

type Info struct {
	Type  types.Type
	Path  []string
	Found bool
}

// OrderedRules are sorted starting with specific ones, and finishing with more generic ones.
var OrderedRules = []Rule{
	ruleNoop,
	ruleVariableInExpressionPossibleFunction,
	ruleLiteralEntry,
	rulePropInMapLiteral,
	ruleParamStartsWithDollar,
	ruleSpecificParent,
	ruleNodePattern,
	ruleRelationshipPattern,
	ruleProcedureOutputsInCallClause,
	ruleCallClauseBeginning,
	ruleConsoleCommandSubcommands,
	rulePropertyLookup,
	rulePossibleKeyword,
}

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

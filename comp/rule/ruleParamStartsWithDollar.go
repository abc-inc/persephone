package rule

import (
	"github.com/abc-inc/merovingian/types"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

func ruleParamStartsWithDollar(e antlr.ParseTree) []Info {
	if e.GetText() == "$" {
		return []Info{{Type: types.Parameter}}
	}
	return nil
}

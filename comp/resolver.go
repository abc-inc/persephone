package comp

import (
	"github.com/abc-inc/merovingian/comp/rule"
	"github.com/abc-inc/merovingian/types"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)


func evaluateRules(element antlr.ParseTree) []rule.Info {
	for _, r := range rule.OrderedRules {
		if items := r(element); len(items) > 0 {
			return items
		}
	}
	return nil
}

func GetTypes(element antlr.Tree) ComplInfo {
	// If element is nil, then no types
	if element == nil {
		return ComplInfo{
			Found: false,
			Types: types.All,
		}
	}

	// Retrieve types from rules
	if infos := evaluateRules(element.(antlr.ParseTree)); len(infos) > 0 {
		ts := make([]types.Type, len(infos))
		for i, it := range infos {
			ts[i] = it.Type
		}
		return ComplInfo{
			Found: true,
			Types: ts,
		}
	}

	// If no types found, then no types
	return ComplInfo{
		Found: false,
		Types: types.All,
	}
}

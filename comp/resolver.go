package comp

import (
	"github.com/abc-inc/persephone/comp/rule"
	"github.com/abc-inc/persephone/types"
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

func GetTypes(element antlr.Tree) Info {
	// If element is nil, then no types
	if element == nil {
		return Info{
			Found: false,
			Types: types.AllCompData,
		}
	}

	// Retrieve types from rules
	if infos := evaluateRules(element.(antlr.ParseTree)); len(infos) > 0 {
		ts := make([]types.Data, len(infos))
		for i, it := range infos {
			ts[i] = types.Data{Type: it.Type, Path: it.Path, FilterLastElement: it.Found}
		}
		return Info{
			Found: true,
			Types: ts,
		}
	}

	// If no types found, then no types
	return Info{
		Found: false,
		Types: types.AllCompData,
	}
}

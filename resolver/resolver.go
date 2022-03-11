package resolver

import (
	"github.com/abc-inc/merovingian/comp"
	"github.com/abc-inc/merovingian/comp/rule"
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

func GetTypes(element antlr.Tree) comp.ComplInfo {
	// If element is nil, then no types
	if element == nil {
		return comp.ComplInfo{
			Found: false,
			Types: comp.All,
		}
	}

	// Retrieve types from rules
	if infos := evaluateRules(element.(antlr.ParseTree)); len(infos) > 0 {
		ts := make([]comp.TypeData, len(infos))
		for i, it := range infos {
			ts[i] = comp.TypeData{Type: it.Type, Path: it.Path, FilterLastElement: it.Found}
		}
		return comp.ComplInfo{
			Found: true,
			Types: ts,
		}
	}

	// If no types found, then no types
	return comp.ComplInfo{
		Found: false,
		Types: comp.All,
	}
}

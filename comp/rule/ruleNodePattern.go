package rule

import (
	"github.com/abc-inc/persephone/ast"
	"github.com/abc-inc/persephone/parser"
	"github.com/abc-inc/persephone/types"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// If we are in node pattern then return variables and types
func ruleNodePattern(e antlr.ParseTree) []Info {
	parent := ast.GetParent(e)
	text := e.GetText()

	// Special case. We are at the beginning of first node pattern.
	if parent != nil {
		if _, ok := parent.(*parser.PatternElementContext); ok && text == "(" {
			return []Info{{Type: types.Variable}, {Type: types.Label}}
		}
	}

	if _, ok := parent.(*parser.NodePatternContext); ok {
		// We are at the beginning of node pattern.
		if text == "(" {
			return []Info{{Type: types.Variable}, {Type: types.Label}}
		}
		if text == ":" {
			return []Info{{Type: types.Label}}
		}
	}
	return nil
}

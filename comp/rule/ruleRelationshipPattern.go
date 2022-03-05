package rule

import (
	"reflect"

	"github.com/abc-inc/merovingian/ast"
	"github.com/abc-inc/merovingian/parser"
	"github.com/abc-inc/merovingian/types"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// If we are in relationship pattern then return variables and types
func ruleRelationshipPattern(e antlr.ParseTree) []Info {
	parent := ast.FindParent(e, reflect.TypeOf(parser.RelationshipPatternContext{}))
	if parent == nil {
		return nil
	}

	// We are at the beginning, so allow variables too
	if e.GetText() == "[" {
		return []Info{{Type: types.Variable}, {Type: types.RelationshipType}}
	}
	// We are at the end, fail and allow algorithm to get back by 1 char
	if e.GetText() == "]" {
		return nil
	}
	return nil
}

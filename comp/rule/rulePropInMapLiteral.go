package rule

import (
	"reflect"
	"strings"

	"github.com/abc-inc/persephone/ast"
	"github.com/abc-inc/persephone/parser"
	"github.com/abc-inc/persephone/types"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

func rulePropInMapLiteral(e antlr.ParseTree) []Info {
	mapLitContext := ast.FindParent(e, reflect.TypeOf(parser.MapLiteralContext{}))
	propContext := ast.FindParent(e, reflect.TypeOf(parser.PropertiesContext{}))

	if mapLitContext != nil {
		if e.GetText() == "}" {
			return nil
		}
		return []Info{{Type: types.PropertyKey}}
	}

	if propContext != nil {
		if e.GetText() == "}" || strings.TrimSpace(e.GetText()) == "" {
			return nil
		}
		return []Info{{Type: types.PropertyKey}, {Type: types.Parameter}}
	}

	return nil
}

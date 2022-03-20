package rule

import (
	"reflect"

	"github.com/abc-inc/persephone/lang"
	"github.com/abc-inc/persephone/types"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Specify place where no autocompletion should be triggered
func ruleNoop(e antlr.ParseTree) []Info {
	if reflect.TypeOf(e).Elem().Name() == lang.STRING_LITERAL_CONTEXT {
		return []Info{{Type: types.Noop}}
	}
	return nil
}

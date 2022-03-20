package comp

import (
	"github.com/abc-inc/persephone/lang"
	"github.com/abc-inc/persephone/parser"
	"github.com/abc-inc/persephone/ref"
	"github.com/abc-inc/persephone/types"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type QueryBased struct {
	refProvs map[string]ref.Provider
}

func NewQueryBased(refProvs map[string]ref.Provider) *QueryBased {
	return &QueryBased{refProvs: refProvs}
}

func (q QueryBased) Complete(ts []types.Data, query antlr.Tree) (is []Item) {
	if query == nil {
		return is
	}

	for _, t := range ts {
		if t.Type != types.Variable {
			continue
		}
		ns := q.refProvs[lang.VARIABLE_CONTEXT].GetNames(query.(*parser.CypherQueryContext))
		for _, n := range ns {
			is = append(is, Item{
				Type:    types.Variable,
				View:    n,
				Content: n,
			})
		}
	}
	return is
}

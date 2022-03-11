package comp

import (
	"github.com/abc-inc/merovingian/lang"
	"github.com/abc-inc/merovingian/parser"
	"github.com/abc-inc/merovingian/types"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type QueryBased struct {
	refProvs map[string]Provider
}

func NewQueryBased(refProvs map[string]Provider) *QueryBased {
	return &QueryBased{refProvs: refProvs}
}

func (q QueryBased) Complete(ts []TypeData, query antlr.Tree) (is []Item) {
	if query == nil {
		return is
	}

	for _, t:=range ts {
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

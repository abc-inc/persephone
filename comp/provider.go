package comp

import (
	"sort"

	"github.com/abc-inc/merovingian/parser"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type Provider struct {
	Queries []parser.CypherQueryContext

	Names                    []string
	NamesByQuery             [][]string
	ReferencesByName         map[string][]antlr.ParserRuleContext
	ReferencesByQueryAndName []map[string][]antlr.ParserRuleContext
}

func NewProvider(queries []parser.CypherQueryContext, index *Index) *Provider {
	namesByQuery := make([][]string, len(index.NamesByQuery))
	for i, names := range index.NamesByQuery {
		ns := keys(names)
		namesByQuery[i] = ns
	}

	return &Provider{
		Queries:                  queries,
		Names:                    keys(index.Names),
		NamesByQuery:             namesByQuery,
		ReferencesByName:         index.ReferencesByName,
		ReferencesByQueryAndName: index.ReferencesByQueryAndName,
	}
}

func (p Provider) GetReferences(name string, query *parser.CypherQueryContext) []antlr.ParserRuleContext {
	if query == nil {
		return p.ReferencesByName[name]
	}
	for i, q := range p.Queries {
		if q == *query {
			return p.ReferencesByQueryAndName[i][name]
		}
	}
	return nil
}

func (p Provider) GetNames(query *parser.CypherQueryContext) []string {
	if query == nil {
		return p.Names
	}
	for i, q := range p.Queries {
		if q == *query {
			return p.NamesByQuery[i]
		}
	}
	return nil
}

func keys(m map[string]interface{}) (ks []string) {
	for y := range m {
		ks = append(ks, y)
	}
	sort.Strings(ks)
	return ks
}

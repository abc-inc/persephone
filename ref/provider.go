// Copyright 2022 The persephone authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package ref

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/gschauer/cypher2go/v4/parser"
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
		ns := names
		namesByQuery[i] = ns
	}

	return &Provider{
		Queries:                  queries,
		Names:                    index.Names,
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

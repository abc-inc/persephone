// Copyright 2022 The Persephone authors
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

package comp

import (
	"github.com/abc-inc/persephone/lang"
	"github.com/abc-inc/persephone/ref"
	"github.com/abc-inc/persephone/types"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/gschauer/cypher2go/v4/parser"
)

// QueryBased completion provides Items based on the input text.
type QueryBased struct {
	refProvByCtx map[string]ref.Provider
}

var _ Comp = (*QueryBased)(nil)

// NewQueryBased creates a new QueryBased completer using the given Providers.
func NewQueryBased(refProvByCtx map[string]ref.Provider) *QueryBased {
	return &QueryBased{refProvByCtx: refProvByCtx}
}

// Complete returns all variables from the entire CypherQueryContext.
func (q QueryBased) Complete(ts []types.Data, query antlr.Tree) (its []Item) {
	if query == nil {
		return its
	}

	for _, t := range ts {
		its = append(its, q.CalculateItems(t, query)...)
	}
	return its
}

// CalculateItems returns all variables from the entire CypherQueryContext.
func (q QueryBased) CalculateItems(t types.Data, query antlr.Tree) (its []Item) {
	if t.Type != types.Variable {
		return
	}
	ns := q.refProvByCtx[lang.VariableContext].GetNames(query.(*parser.CypherQueryContext))
	for _, n := range ns {
		its = append(its, Item{
			Type:    types.Variable,
			View:    n,
			Content: n,
		})
	}
	return
}

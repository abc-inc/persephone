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

package comp

import (
	"github.com/abc-inc/persephone/lang"
	"github.com/abc-inc/persephone/ref"
	"github.com/abc-inc/persephone/types"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/gschauer/cypher2go/v4/parser"
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
		ns := q.refProvs[lang.VariableContext].GetNames(query.(*parser.CypherQueryContext))
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

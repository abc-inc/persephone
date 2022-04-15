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

var Present = struct{}{}

type Index struct {
	Names                    []string
	NamesByQuery             [][]string
	ReferencesByName         map[string][]antlr.ParserRuleContext
	ReferencesByQueryAndName []map[string][]antlr.ParserRuleContext
}

func NewIndex() *Index {
	i := &Index{}
	i.Names = make([]string, 0)
	i.NamesByQuery = make([][]string, 0)
	i.ReferencesByName = make(map[string][]antlr.ParserRuleContext)
	i.ReferencesByQueryAndName = make([]map[string][]antlr.ParserRuleContext, 0)
	return i
}

func (i *Index) AddQuery() {
	i.NamesByQuery = append(i.NamesByQuery, make([]string, 0))
	i.ReferencesByQueryAndName = append(i.ReferencesByQueryAndName, make(map[string][]antlr.ParserRuleContext))
}

func (i *Index) Add(ctx antlr.ParserRuleContext, addName bool) {
	idx := len(i.NamesByQuery) - 1
	if addName {
		if !contains(i.Names, ctx.GetText()) {
			i.Names = append(i.Names, ctx.GetText())
			i.NamesByQuery[idx] = append(i.NamesByQuery[idx], ctx.GetText())
		}
	}
	i.ReferencesByName[ctx.GetText()] = append(i.ReferencesByName[ctx.GetText()], ctx)
	index := i.ReferencesByQueryAndName[idx]
	index[ctx.GetText()] = append(index[ctx.GetText()], ctx)
}

// AddVariable registers a new variable context.
// Variables have specific rules, because they participate in autocompletion.
// We should not add to the names list variables that are in expression.
func (i *Index) AddVariable(ctx *parser.VariableContext) {
	addName := true
	p := ctx.GetParent()
	if _, ok := p.(*parser.AtomContext); p != nil && ok {
		addName = false
	}
	i.Add(ctx, addName)
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

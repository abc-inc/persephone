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
	"regexp"
	"strings"

	"github.com/abc-inc/go-fuzzaldrin-plus"
	"github.com/abc-inc/persephone/ast"
	"github.com/abc-inc/persephone/lang"
	"github.com/abc-inc/persephone/ref"
	"github.com/abc-inc/persephone/types"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/gschauer/cypher2go/v4/parser"
)

// KeywordItems lists all keywords as completion items.
func KeywordItems() []Item {
	its := make([]Item, len(lang.Keywords))
	for i, kw := range lang.Keywords {
		its[i] = Item{Type: types.Keyword, View: kw, Content: kw}
	}
	return its
}

type Comp interface {
	CalculateItems(ts types.Data, query antlr.Tree) []Item
	Complete(ts []types.Data, query antlr.Tree) (its []Item)
}

type AutoCompletion struct {
	QueryBased  *QueryBased
	SchemaBased *SchemaBased
}

func NewAutoCompletion(schema Metadata) *AutoCompletion {
	a := &AutoCompletion{}
	a.UpdateSchema(schema)
	return a
}

func (a AutoCompletion) GetItems(types []types.Data, query antlr.Tree, filter string) (its []Item) {
	text := strings.ToLower(filter)
	filteredText := filterText(text)

	if a.QueryBased != nil {
		its = append(its, a.QueryBased.Complete(types, query)...)
	}
	if a.SchemaBased != nil {
		its = append(its, a.SchemaBased.Complete(types, query)...)
	}

	if len(filteredText) > 0 {
		return fuzzaldrin.Filter(its, filteredText, func(i Item) string { return i.View })
	}
	if len(text) > 0 {
		return fuzzaldrin.Filter(its, text, func(i Item) string { return i.View })
	}
	return its
}

func (a *AutoCompletion) UpdateSchema(schema Metadata) {
	a.SchemaBased = NewSchemaBased(schema)
}

func (a *AutoCompletion) UpdateReferenceProviders(refProvByCtx map[string]ref.Provider) {
	a.QueryBased = NewQueryBased(refProvByCtx)
}

// ShouldBeReplaced defines whether element should be replaced or not.
func ShouldBeReplaced(e antlr.Tree) bool {
	if e == nil {
		return false
	}

	text := e.(antlr.ParseTree).GetText()
	parent := ast.GetParent(e)

	// If element is whitespace
	if ok, err := regexp.MatchString("^\\s+$", text); err == nil && ok {
		return false
	}
	// If element is opening bracket (e.g. start of relationship pattern)
	if text == "[" {
		return false
	}
	// If element is opening brace (e.g. start of node pattern)
	if text == "(" {
		return false
	}
	if text == "." {
		return false
	}
	if text == "{" {
		return false
	}
	if text == "$" {
		return false
	}
	if text == ":" && parent != nil {
		if _, ok := parent.(parser.ILiteralEntryContext); ok {
			return false
		}
	}
	return true
}

func filterText(text string) string {
	return strings.TrimPrefix(text, "$")
}

func CalculateSmartReplaceRange(e antlr.Tree, start, stop int) *Filter {
	// If we are in relationship type or label, and we have error nodes in there.
	// This means that we typed in just ':' and Antlr consumed other tokens in element.
	// In this case replace only ':'.
	_, ok1 := e.(*parser.RelationshipTypeContext)
	_, ok2 := e.(*parser.NodeLabelContext)
	if ok1 || ok2 {
		if ast.HasErrorNode(e) {
			return &Filter{":", start, start}
		}
	}
	return nil
}

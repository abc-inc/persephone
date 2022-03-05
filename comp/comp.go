package comp

import (
	"regexp"
	"strings"

	"github.com/abc-inc/merovingian/db/neo4j"
	"github.com/abc-inc/merovingian/lang"
	"github.com/abc-inc/merovingian/parser"
	"github.com/abc-inc/merovingian/types"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

var KEYWORD_ITEMS []Item

func init() {
	for _, kw := range lang.Keywords {
		KEYWORD_ITEMS = append(KEYWORD_ITEMS,
			Item{Type: types.Keyword, View: kw, Content: kw, Postfix: nil})
	}
}

type Comp interface {
	CalculateItems(ts types.Type, query string) []Item
	Complete(ts []types.Type, query antlr.Tree)
}

type AutoCompletion struct {
	QueryBased  *QueryBased
	SchemaBased *SchemaBased
}

func NewAutoCompletion(schema neo4j.Schema) *AutoCompletion {
	a := &AutoCompletion{}
	a.UpdateSchema(schema)
	return a
}

func (a AutoCompletion) GetItems(types []types.Type, query antlr.Tree, filter string) (items []Item) {
	// TODO check if the original implementation is case-insensitive
	text := filter // strings.ToLower(filter)
	filteredText := filterText(text)

	if a.QueryBased != nil {
		items = append(items, a.QueryBased.Complete(types, query)...)
	}
	if a.SchemaBased != nil {
		items = append(items, a.SchemaBased.Complete(types, query)...)
	}

	if len(filteredText) > 0 {
		return fuzzySearch(items, filteredText, func(i Item) string { return i.View })
	}
	if len(text) > 0 {
		return fuzzySearch(items, text, func(i Item) string { return i.View })
	}
	return items
}

func (a *AutoCompletion) UpdateSchema(schema neo4j.Schema) {
	a.SchemaBased = NewSchemaBased(schema)
}

func (a *AutoCompletion) UpdateReferenceProviders(refProvs map[string]Provider) {
	a.QueryBased = NewQueryBased(refProvs)
}

// ShouldBeReplaced defines whether element should be replaced or not.
func ShouldBeReplaced(element antlr.Tree) bool {
	if element == nil {
		return false
	}

	text := element.(antlr.ParseTree).GetText()
	parent := element.GetParent()

	// If element is whitespace
	if ok, err := regexp.MatchString(`^\s+$`, text); err != nil && ok {
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
		if _, ok := parent.(parser.LiteralEntryContext); ok {
			return false
		}
	}
	return true
}

func filterText(text string) string {
	return strings.TrimPrefix(text, "$")
}

func CalculateSmartReplaceRange(element antlr.Tree, start, stop int) Filter {
	// If we are in relationship type or label, and we have error nodes in there.
	// This means that we typed in just ':' and Antlr consumed other tokens in element.
	// In this case replace only ':'.
	_, ok1 := element.(parser.RelationshipTypeContext)
	_, ok2 := element.(parser.NodeLabelContext)
	if ok1 || ok2 {
		//if ast.HasErrorNode(element) {
		//	return Filter{":", start, stop}
		//}
	}
	return Filter{}
}

func fuzzySearch(items []Item, query string, keyEx func(Item) string) (res []Item) {
	for _, i := range items {
		key := keyEx(i)
		if strings.Contains(key, query) {
			res = append(res, i)
		}
	}
	return
}

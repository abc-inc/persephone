package comp

import (
	"regexp"
	"strings"

	"github.com/abc-inc/persephone/ast"
	"github.com/abc-inc/persephone/fuzzy"
	"github.com/abc-inc/persephone/lang"
	"github.com/abc-inc/persephone/graph"
	"github.com/abc-inc/persephone/parser"
	"github.com/abc-inc/persephone/ref"
	"github.com/abc-inc/persephone/types"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

var KeywordItems []Item

func init() {
	for _, kw := range lang.Keywords {
		KeywordItems = append(KeywordItems,
			Item{Type: types.Keyword, View: kw, Content: kw})
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

func NewAutoCompletion(schema graph.Schema) *AutoCompletion {
	a := &AutoCompletion{}
	a.UpdateSchema(schema)
	return a
}

func (a AutoCompletion) GetItems(types []types.Data, query antlr.Tree, filter string) (items []Item) {
	text := strings.ToLower(filter)
	filteredText := filterText(text)

	if a.QueryBased != nil {
		items = append(items, a.QueryBased.Complete(types, query)...)
	}
	if a.SchemaBased != nil {
		items = append(items, a.SchemaBased.Complete(types, query)...)
	}

	if len(filteredText) > 0 {
		return fuzzy.Search(items, filteredText, func(i Item) string { return i.View })
	}
	if len(text) > 0 {
		return fuzzy.Search(items, text, func(i Item) string { return i.View })
	}
	return items
}

func (a *AutoCompletion) UpdateSchema(schema graph.Schema) {
	a.SchemaBased = NewSchemaBased(schema)
}

func (a *AutoCompletion) UpdateReferenceProviders(refProvs map[string]ref.Provider) {
	a.QueryBased = NewQueryBased(refProvs)
}

// ShouldBeReplaced defines whether element should be replaced or not.
func ShouldBeReplaced(element antlr.Tree) bool {
	if element == nil {
		return false
	}

	text := element.(antlr.ParseTree).GetText()
	parent := ast.GetParent(element)

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

func CalculateSmartReplaceRange(element antlr.Tree, start, stop int) *Filter {
	// If we are in relationship type or label, and we have error nodes in there.
	// This means that we typed in just ':' and Antlr consumed other tokens in element.
	// In this case replace only ':'.
	_, ok1 := element.(*parser.RelationshipTypeContext)
	_, ok2 := element.(*parser.NodeLabelContext)
	if ok1 || ok2 {
		if ast.HasErrorNode(element) {
			return &Filter{":", start, start}
		}
	}
	return nil
}

package comp

import (
	"os"
	"regexp"
	"sort"
	"strings"
	"unicode"

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

func NewAutoCompletion(schema neo4j.Schema) *AutoCompletion {
	a := &AutoCompletion{}
	a.UpdateSchema(schema)
	return a
}

func (a AutoCompletion) GetItems(types []TypeData, query antlr.Tree, filter string) (items []Item) {
	text := strings.ToLower(filter)
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
	if query == "" {
		return items
	}
	queryHasSlashes := strings.ContainsRune(query, '/')
	query = strings.ReplaceAll(query, " ", "")
	return filter(items, query, queryHasSlashes, keyEx)
}

type ScoredCandidate struct {
	item  Item
	score float64
}

func filter(items []Item, query string, queryHasSlashes bool, keyEx func(Item) string) (candidates []Item) {
	var scoredCandidates []ScoredCandidate
	for _, i := range items {
		s := keyEx(i)
		if s == "" {
			continue
		}

		score := Score(s, query)
		if !queryHasSlashes {
			score = basenameScore(s, query, score)
		}
		if score > 0 {
			scoredCandidates = append(scoredCandidates, ScoredCandidate{i, score})
		}
	}

	sort.Slice(scoredCandidates, func(i, j int) bool {
		return scoredCandidates[i].score > scoredCandidates[j].score
	})
	for _, sc := range scoredCandidates {
		candidates = append(candidates, sc.item)
	}
	return
}

func basenameScore(s, query string, score float64) float64 {
	index := len(s) - 1
	for s[index] == os.PathSeparator {
		index--
	}

	var base string
	slashCount := 0
	lastChar := index
	for index >= 0 {
		if s[index] == os.PathSeparator {
			slashCount++
			if base == "" {
				base = s[index+1 : lastChar+1]
			}
		} else if index == 0 {
			if lastChar < len(s)-1 {
				if base == "" {
					base = s[0 : lastChar+1]
				}
			} else {
				if base == "" {
					base = s
				}
			}
		}
		index--
	}

	if base == s {
		score *= 2
	} else if base != "" {
		score += Score(base, query)
	}
	segmentCount := slashCount + 1
	depth := max(1, 10-segmentCount)
	score *= float64(depth) * 0.01
	return score
}

func Score(s, query string) float64 {
	if s == query {
		return 1
	}
	if queryIsLastPathSegment(s, query) {
		return 1
	}
	totalCharScore := 0.0
	strLen := len(s)
	indexInString := 0
	for indexInQuery := 0; indexInQuery < len(query); indexInQuery++ {
		c := rune(query[indexInQuery])
		lowerCaseIndex := strings.IndexRune(s, unicode.ToLower(c))
		upperCaseIndex := strings.IndexRune(s, unicode.ToUpper(c))
		minIndex := min(lowerCaseIndex, upperCaseIndex)
		if minIndex == -1 {
			minIndex = max(lowerCaseIndex, upperCaseIndex)
		}
		indexInString = minIndex
		if indexInString == -1 {
			return 0
		}

		charScore := 0.1
		if s[indexInString] == byte(c) {
			charScore += 0.1
		}
		if indexInString == 0 || s[indexInString-1] == os.PathSeparator {
			charScore += 0.8
		} else if c == '-' || c == '_' || c == ' ' {
			charScore += 0.7
		}
		s = s[indexInString+1:]
		totalCharScore += charScore
	}
	queryScore := totalCharScore / float64(len(query))
	return ((queryScore * (float64(len(query)) / float64(strLen))) + queryScore) / float64(2)
}

func queryIsLastPathSegment(s, query string) bool {
	if len(s)>len(query) && s[len(s)-len(query)-1] == os.PathSeparator {
		return strings.LastIndex(s, query) == len(s)-len(query)
	}
	return false
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

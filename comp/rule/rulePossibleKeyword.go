package rule

import (
	"sort"
	"strings"

	"github.com/abc-inc/merovingian/comp"
	"github.com/abc-inc/merovingian/lang"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// If any of the keywords contains element text, return ALL
func rulePossibleKeyword(e antlr.ParseTree) (is []Info) {
	text := strings.ToUpper(e.GetText())
	if pos := sort.SearchStrings(lang.Keywords, text); pos < len(lang.Keywords) && strings.Contains(lang.Keywords[pos], text) {
		for _, t := range comp.All {
			is = append(is, Info{Type: t.Type})
		}
	}
	return is
}

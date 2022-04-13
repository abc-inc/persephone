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

package rule

import (
	"sort"
	"strings"

	"github.com/abc-inc/persephone/lang"
	"github.com/abc-inc/persephone/types"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// If any of the keywords contains element text, return ALL
func rulePossibleKeyword(e antlr.ParseTree) (is []Info) {
	text := strings.ToUpper(e.GetText())
	if pos := sort.SearchStrings(lang.Keywords, text); pos < len(lang.Keywords) && strings.Contains(lang.Keywords[pos], text) {
		for _, t := range types.AllComp {
			is = append(is, Info{Type: t})
		}
	}
	return is
}

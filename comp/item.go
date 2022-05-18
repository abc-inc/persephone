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
	"fmt"

	"github.com/abc-inc/persephone/types"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Item is completion candidate of a certain type.
type Item struct {
	Type    types.Type
	View    string
	Content string
	Postfix string
}

// String returns a text representation of this completion item.
func (i Item) String() string {
	return fmt.Sprintf("%s(%s):%s", i.Type, i.View, i.Content)
}

// LineCol represents the position of a character in a multi-line string.
type LineCol struct {
	Line, Col int
}

// String returns "(line,column)".
func (l LineCol) String() string {
	return fmt.Sprintf("(%d,%d)", l.Line, l.Col)
}

// Range represents the position range of a substring in a multi-line string.
type Range struct {
	From, To LineCol
}

// String returns "[(fromLine,fromColumn),(toLine,toColumn)]".
func (r Range) String() string {
	return fmt.Sprintf("[%d,%d]", r.From, r.To)
}

// Filter is used limit the replacement range in a longer input string.
type Filter struct {
	FilterText string
	Start      int
	Stop       int
}

// Result holds the completion candidates and their insert position.
type Result struct {
	Items []Item
	Range Range
}

// Info lists all completion types, which are allowed at a certain position.
type Info struct {
	Element antlr.Tree
	Query   antlr.Tree
	Found   bool
	Types   []types.Data
}

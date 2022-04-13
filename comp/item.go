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
	_ "github.com/abc-inc/persephone/types"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type Item struct {
	Type    types.Type `json:"type"`
	View    string     `json:"view"`
	Content string     `json:"content"`
	Postfix string     `json:"postfix"`
}

func (i Item) String() string {
	return fmt.Sprintf("%s(%s):%s", i.Type, i.View, i.Content)
}

type LineCol struct {
	Line, Col int
}

func (l LineCol) String() string {
	return fmt.Sprintf("(%d,%d)", l.Line, l.Col)
}

type Range struct {
	From, To LineCol
}

func (r Range) String() string {
	return fmt.Sprintf("[%d,%d]", r.From, r.To)
}

type Filter struct {
	FilterText string
	Start      int
	Stop       int
}

type Result struct {
	Items []Item
	Range Range
}

type Info struct {
	Element antlr.Tree
	Query   antlr.Tree
	Found   bool
	Types   []types.Data
}

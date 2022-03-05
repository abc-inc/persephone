package comp

import (
	"fmt"

	"github.com/abc-inc/merovingian/types"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type Item struct {
	Type    types.Type  `json:"type"`
	View    string      `json:"view"`
	Content string      `json:"content"`
	Postfix interface{} `json:"postfix"`
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

type ComplInfo struct {
	Element antlr.Tree
	Query   antlr.Tree
	Found   bool
	Types   []types.Type
}


type TypeData struct {
	Name              string
	Path              []string
	FilterLastElement bool
}
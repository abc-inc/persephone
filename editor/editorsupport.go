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

package editor

import (
	"reflect"

	"github.com/abc-inc/persephone/ast"
	"github.com/abc-inc/persephone/comp"
	"github.com/abc-inc/persephone/lang"
	"github.com/abc-inc/persephone/ref"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/gschauer/cypher2go/v4/parser"
)

type Editor struct {
	schema  comp.Metadata
	Input   string
	posConv PosConv

	ParseTree           antlr.ParseTree
	ParseErrors         []SynErr
	referencesProviders map[string]ref.Provider
	completion          comp.AutoCompletion
	statements          []parser.CypherPartContext
}

// NewEditor initializes a new Editor with various autocompletion capabilities.
func NewEditor(input string) *Editor {
	e := &Editor{}
	e.completion = *comp.NewAutoCompletion(comp.Metadata{})
	e.Update(input)
	return e
}

// Update parses the given input, records various ParseTree elements and updates
// the reference providers to provide more appropriate completion results.
func (e *Editor) Update(input string) {
	e.posConv = *NewPosConv(input)

	e.Input = input
	parseTree, referencesListener, errorListener, referencesProviders := Parse(input)
	e.ParseTree = parseTree

	e.ParseErrors = errorListener.errors

	e.statements = referencesListener.Statements
	e.referencesProviders = referencesProviders

	e.completion.UpdateReferenceProviders(e.referencesProviders)
}

// SetSchema updates the metadata used for providing schema-specific completion.
func (e *Editor) SetSchema(schema comp.Metadata) {
	e.schema = schema
	e.completion.UpdateSchema(e.schema)
}

// GetElement returns the ParseTree at the given line and column.
func (e Editor) GetElement(line, column int) antlr.Tree {
	abs := e.posConv.ToAbsolute(line, column)
	return getElement(e.ParseTree, abs)
}

// getElement returns the ParseTree at the given position in the input string.
func getElement(pt antlr.Tree, abs int) antlr.Tree {
	pos := ast.GetPosition(pt)
	if pos != nil && (abs < pos.GetStart() || abs > pos.GetStop()) {
		return nil
	}

	c := pt.GetChildCount()
	if c == 0 && pos != nil {
		return pt
	}

	for _, c := range pt.GetChildren() {
		if e := getElement(c, abs); e != nil {
			return e
		}
	}

	if pos != nil {
		return pt
	}
	return nil
}

func (e Editor) GetReferences(line, column int) []antlr.ParserRuleContext {
	el := ast.FindAnyParent(e.GetElement(line, column), lang.SymbolicContexts)
	if el == nil {
		return nil
	}

	var query antlr.Tree
	typ := reflect.TypeOf(el).Elem().Name()
	if typ == lang.VariableContext {
		query = ast.FindAnyParent(el, []string{lang.QueryContext})
	}

	text := el.(antlr.ParseTree).GetText()
	if query == nil {
		return e.referencesProviders[typ].GetReferences(text, nil)
	}
	return e.referencesProviders[typ].GetReferences(text, query.(*parser.CypherQueryContext))
}

func (e Editor) GetCompletionInfo(line, column int) comp.Info {
	el := e.GetElementForCompletion(line, column)
	query := ast.FindAnyParent(el, []string{lang.QueryContext})
	info := comp.GetTypes(el)
	return comp.Info{
		Element: el,
		Query:   query,
		Found:   info.Found,
		Types:   info.Types,
	}
}

func (e Editor) GetElementForCompletion(line, column int) antlr.Tree {
	el := e.GetElement(line, column)
	if p := ast.FindAnyParent(el, lang.CompletionCandidates); p != nil {
		return p
	}
	return el
}

func (e Editor) GetCompletion(line, column int, doFilter bool) comp.Result {
	info := e.GetCompletionInfo(line, column)
	if !info.Found && column > 0 {
		if prevInfo := e.GetCompletionInfo(line, column-1); prevInfo.Found {
			info = prevInfo
		}
	}

	element, query, found, types := info.Element, info.Query, info.Found, info.Types

	replFrom := comp.LineCol{Line: line, Col: column}
	replTo := comp.LineCol{Line: line, Col: column}
	var filter string

	shouldBeReplaced := comp.ShouldBeReplaced(element)
	if found && shouldBeReplaced {
		// There are number of situations where we need to be smarter than default behavior
		pos := ast.GetPosition(element)
		smartReplaceRange := comp.CalculateSmartReplaceRange(element, pos.GetStart(), pos.GetStop())
		if smartReplaceRange != nil {
			replFrom.Line, replFrom.Col = e.posConv.ToRelative(smartReplaceRange.Start)
			replTo.Line, replTo.Col = e.posConv.ToRelative(smartReplaceRange.Stop + 1)

			if smartReplaceRange.FilterText != "" {
				filter = smartReplaceRange.FilterText
			}
		} else {
			replFrom.Line, replFrom.Col = e.posConv.ToRelative(pos.GetStart())
			replTo.Line, replTo.Col = e.posConv.ToRelative(pos.GetStop() + 1)
		}
	}

	if filter == "" {
		if doFilter && found && shouldBeReplaced {
			if e, ok := element.(interface{ GetText() string }); ok {
				filter = e.GetText()
			} else {
				panic(reflect.TypeOf(element).String() + " does not have text")
			}
		}
	}

	items := e.completion.GetItems(types, query, filter)
	return comp.Result{
		Items: items,
		Range: comp.Range{From: replFrom, To: replTo},
	}
}

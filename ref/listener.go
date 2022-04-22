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

package ref

import (
	"github.com/abc-inc/persephone/lang"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/gschauer/cypher2go/v4/parser"
)

type Listener struct {
	Queries            []parser.CypherQueryContext
	QueriesAndCommands []antlr.ParserRuleContext
	Statements         []parser.CypherPartContext
	Raw                []antlr.ParserRuleContext
	Indexes            map[string]*Index
	InConsoleCommand   bool
	parser.BaseCypherListener
}

func NewRefListener() *Listener {
	is := map[string]*Index{}
	for _, ctx := range lang.SymbolicContexts {
		is[ctx] = NewIndex()
	}
	return &Listener{
		Indexes: is,
	}
}

func (l *Listener) EnterRaw(ctx antlr.ParserRuleContext) {
	l.Raw = append(l.Raw, ctx)
}

func (l *Listener) ExitRaw(ctx antlr.ParserRuleContext) {
	if len(l.Raw) == 0 {
		l.Raw = append(l.Raw, ctx)
	}
}

func (l *Listener) EnterCypherPart(ctx *parser.CypherPartContext) {
	l.Statements = append(l.Statements, *ctx)
}

func (l *Listener) ExitCypherPart(ctx *parser.CypherPartContext) {
	if len(l.Statements) == 0 {
		l.Statements = append(l.Statements, *ctx)
	}
}

func (l *Listener) EnterCypherConsoleCommand(ctx *parser.CypherConsoleCommandContext) {
	l.QueriesAndCommands = append(l.QueriesAndCommands, ctx)
	for _, i := range l.Indexes {
		i.AddQuery()
	}
	l.InConsoleCommand = true
}

func (l *Listener) ExitCypherConsoleCommand(ctx *parser.CypherConsoleCommandContext) {
	l.InConsoleCommand = false
}

func (l *Listener) EnterCypherQuery(ctx *parser.CypherQueryContext) {
	l.Queries = append(l.Queries, *ctx)
	l.QueriesAndCommands = append(l.QueriesAndCommands, ctx)
	for _, i := range l.Indexes {
		i.AddQuery()
	}
}

func (l *Listener) ExitVariable(ctx *parser.VariableContext) {
	if l.InConsoleCommand {
		return
	}
	l.Indexes[lang.VariableContext].AddVariable(ctx)
}

func (l *Listener) ExitLabelName(ctx *parser.LabelNameContext) {
	if l.InConsoleCommand {
		return
	}
	l.Indexes[lang.LabelNameContext].Add(ctx, true)
}

func (l *Listener) ExitRelTypeName(ctx *parser.RelTypeNameContext) {
	if l.InConsoleCommand {
		return
	}
	l.Indexes[lang.RelationshipTypeNameContext].Add(ctx, true)
}

func (l *Listener) ExitPropertyKeyName(ctx *parser.PropertyKeyNameContext) {
	if l.InConsoleCommand {
		return
	}
	l.Indexes[lang.PropertyKeyNameContext].Add(ctx, true)
}

func (l *Listener) ExitParameterName(ctx *parser.ParameterNameContext) {
	if l.InConsoleCommand {
		return
	}
	l.Indexes[lang.ParameterNameContext].Add(ctx, true)
}

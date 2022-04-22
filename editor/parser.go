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
	"github.com/abc-inc/persephone/lang"
	"github.com/abc-inc/persephone/ref"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/gschauer/cypher2go/v4/parser"
)

func Parse(input string) (antlr.ParseTree, *ref.Listener, *ErrorListener, map[string]ref.Provider) {
	refListener := ref.NewRefListener()
	errListener := &ErrorListener{}
	chars := antlr.NewInputStream(input)
	lexer := parser.NewCypherLexer(chars)
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(errListener)
	tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	p := parser.NewCypherParser(tokens)
	p.BuildParseTrees = true
	p.RemoveErrorListeners()
	p.AddErrorListener(errListener)
	p.AddParseListener(refListener)
	parseTree := p.Cypher()
	queries, indexes := refListener.Queries, refListener.Indexes
	refProvs := make(map[string]ref.Provider, len(lang.SymbolicContexts))
	for _, ctx := range lang.SymbolicContexts {
		refProvs[ctx] = *ref.NewProvider(queries, indexes[ctx])
	}
	return parseTree, refListener, errListener, refProvs
}

// Copyright 2022 The Persephone authors
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
	"reflect"

	"github.com/abc-inc/persephone/ast"
	"github.com/abc-inc/persephone/lang"
	"github.com/abc-inc/persephone/types"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/gschauer/cypher2go/v4/parser"
)

// ruleConsoleCommandSubcommands checks if we are in a console command, and not
// in console command name, and then returns the path.
func ruleConsoleCommandSubcommands(e antlr.ParseTree) []Info {
	pt := ast.GetParent(e)
	consCmd := ast.FindParent(pt, reflect.TypeOf(parser.CypherConsoleCommandContext{}))
	isAtTheEnd := false
	if consCmd == nil {
		// We are not in console command.
		// But maybe we are on a space at the end of console command?
		// If first child of parent contains console command
		// and second child is our current element
		// then we are at the space at the end of console command
		child1 := ast.FindChild(pt.GetChild(0), lang.ConsoleCommandContext)
		if child1 != nil && pt.GetChildCount() > 1 && pt.GetChild(1) == e {
			consCmd = child1
			isAtTheEnd = true
		} else {
			return nil
		}
	}

	// Find current parameter or space
	currentElement := ast.FindParent(e, reflect.TypeOf(parser.CypherConsoleCommandParametersContext{}))
	if currentElement == nil {
		currentElement = e
	}

	var path []string
	currentElementInParameter := false

	// Iterate over parameters, and stop when we found current one.
	for _, child := range consCmd.GetChildren() {
		if ctx, ok := child.(*parser.CypherConsoleCommandNameContext); ok {
			path = append(path, ctx.GetText())
		}
		if _, ok := child.(*parser.CypherConsoleCommandParametersContext); ok {
			for _, paramChild := range child.GetChildren() {
				if ctx, ok := paramChild.(*parser.CypherConsoleCommandParameterContext); ok {
					path = append(path, ctx.GetText())
					currentElementInParameter = true
				} else {
					currentElementInParameter = false
				}
				if paramChild == currentElement {
					break
				}
			}
		}
	}

	// If we are at the end of console command, nothing to filter.
	filterLastElement := false
	if isAtTheEnd {
		filterLastElement = false
	} else {
		// If we are in parameter, filter, otherwise not
		filterLastElement = currentElementInParameter
	}

	return []Info{{types.ConsoleCommandSubCommand, path, filterLastElement}}
}

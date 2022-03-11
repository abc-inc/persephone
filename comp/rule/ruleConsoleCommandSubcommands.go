package rule

import (
	"reflect"

	"github.com/abc-inc/merovingian/ast"
	"github.com/abc-inc/merovingian/lang"
	"github.com/abc-inc/merovingian/parser"
	"github.com/abc-inc/merovingian/types"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// If we are in console command, and not in console command name, return path
func ruleConsoleCommandSubcommands(e antlr.ParseTree) []Info {
	consCmd := ast.FindParent(e.GetParent(), reflect.TypeOf(parser.CypherConsoleCommandContext{}))
	isAtTheEnd := false
	if consCmd == nil {
		// We are not in console command. But maybe we are on a space at the end of console command?
		// If first child of parent contains console command
		// and second child is our current element
		// then we are at the space at the end of console command
		parent := ast.GetParent(e)
		child1 := ast.FindChild(parent.GetChild(0), lang.CONSOLE_COMMAND_CONTEXT)
		if child1 != nil && parent.GetChildCount() > 1 && parent.GetChild(1) == e {
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

	return []Info{{
		Type:  types.ConsoleCommandSubCommand,
		Path:  path,
		Found: filterLastElement,
	},
	}
}

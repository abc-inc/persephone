package autocompletion

import (
	"testing"

	"github.com/abc-inc/merovingian/comp"
	"github.com/abc-inc/merovingian/types"
)

func TestTypesYieldsConsoleCommandTypeAtTheCommandColon(t *testing.T) {
	checkCompletionTypes(t, "▼:", true, []types.Type{types.ConsoleCommandName})
}

func TestTypesYieldsConsoleCommandTypeAtTheCommandName(t *testing.T) {
	checkCompletionTypes(t, ":▼p", true, []types.Type{types.ConsoleCommandName})
}

func TestTypesYieldsConsoleCommandAtSubCommand(t *testing.T) {
	checkCompletionTypes(t, ":help▼ ", true, []types.Type{types.ConsoleCommandSubCommand})
}

func TestTypesYieldsConsoleCommandAtSubCommandPartly(t *testing.T) {
	checkCompletionTypes(t, ":help m▼a", true, []types.Type{types.ConsoleCommandSubCommand})
}

func TestTypesYieldsConsoleCommandAtSubCommandAtSubCommand(t *testing.T) {
	checkCompletionTypes(t, ":server user▼ ", true, []types.Type{types.ConsoleCommandSubCommand})
}

func TestTypesYieldsConsoleCommandAtSubCommandAtSubCommandPartly(t *testing.T) {
	checkCompletionTypes(t, ":server user l▼i", true, []types.Type{types.ConsoleCommandSubCommand})
}

func TestWithoutFiltersYieldsCommandNamesIfColon(t *testing.T) {
	expected := comp.Result{
		Items: []comp.Item{
			{Type: types.ConsoleCommandName, View: ":clear", Content: ":clear"},
			{Type: types.ConsoleCommandName, View: ":play", Content: ":play"},
			{Type: types.ConsoleCommandName, View: ":help", Content: ":help", Postfix: "helpdesc"},
			{Type: types.ConsoleCommandName, View: ":server", Content: ":server"},
		},
		Range: comp.Range{
			From: comp.LineCol{Line: 1, Col: 0},
			To:   comp.LineCol{Line: 1, Col: 1},
		},
	}
	checkCompletion(t, ":▼", expected, false)
}

func TestWithoutFiltersYieldsCommandNamesIfHalfWritten(t *testing.T) {
	expected := comp.Result{
		Items: []comp.Item{
			{Type: types.ConsoleCommandName, View: ":clear", Content: ":clear"},
			{Type: types.ConsoleCommandName, View: ":play", Content: ":play"},
			{Type: types.ConsoleCommandName, View: ":help", Content: ":help", Postfix: "helpdesc"},
			{Type: types.ConsoleCommandName, View: ":server", Content: ":server"},
		},
		Range: comp.Range{
			From: comp.LineCol{Line: 1, Col: 0},
			To:   comp.LineCol{Line: 1, Col: 3},
		},
	}
	checkCompletion(t, ":▼pl", expected, false)
	checkCompletion(t, ":pl▼", expected, false)
}

func TestWithFiltersYieldsCommandNamesIfHalfWritten(t *testing.T) {
	expected := comp.Result{
		Items: []comp.Item{
			{Type: types.ConsoleCommandName, View: ":play", Content: ":play"},
		},
		Range: comp.Range{
			From: comp.LineCol{Line: 1, Col: 0},
			To:   comp.LineCol{Line: 1, Col: 3},
		},
	}
	checkCompletion(t, ":▼pl", expected, true)
	checkCompletion(t, ":pl▼", expected, true)
}

func TestWithFiltersYieldsHelpSubCommand(t *testing.T) {
	expected := comp.Result{
		Items: []comp.Item{
			{Type: types.ConsoleCommandSubCommand, View: "match", Content: "match"},
			{Type: types.ConsoleCommandSubCommand, View: "create", Content: "create"},
		},
		Range: comp.Range{
			From: comp.LineCol{Line: 1, Col: 6},
			To:   comp.LineCol{Line: 1, Col: 6},
		},
	}
	checkCompletion(t, ":help ▼", expected, true)
}

func TestWithFiltersYieldsHelpSubCommandPartly(t *testing.T) {
	expected := comp.Result{
		Items: []comp.Item{
			{Type: types.ConsoleCommandSubCommand, View: "match", Content: "match"},
		},
		Range: comp.Range{
			From: comp.LineCol{Line: 1, Col: 6},
			To:   comp.LineCol{Line: 1, Col: 8},
		},
	}
	checkCompletion(t, ":help ma▼", expected, true)
}

func TestWithFiltersYieldsHelpSubCommandSubCommand(t *testing.T) {
	expected := comp.Result{
		Items: []comp.Item{
			{Type: types.ConsoleCommandSubCommand, View: "list", Content: "list", Postfix: "listdesc"},
			{Type: types.ConsoleCommandSubCommand, View: "add", Content: "add"},
		},
		Range: comp.Range{
			From: comp.LineCol{Line: 1, Col: 13},
			To:   comp.LineCol{Line: 1, Col: 13},
		},
	}
	checkCompletion(t, ":server user ▼", expected, true)
}

func TestWithFiltersYieldsSeverSubCommandSubCommandPartly(t *testing.T) {
	expected := comp.Result{
		Items: []comp.Item{
			{Type: types.ConsoleCommandSubCommand, View: "list", Content: "list", Postfix: "listdesc"},
		},
		Range: comp.Range{
			From: comp.LineCol{Line: 1, Col: 13},
			To:   comp.LineCol{Line: 1, Col: 15},
		},
	}
	checkCompletion(t, ":server user li▼", expected, true)
}

func TestWithFiltersYieldsSeverSubCommandSubCommandNoSubCommand(t *testing.T) {
	expected := comp.Result{
		Range: comp.Range{
			From: comp.LineCol{Line: 1, Col: 18},
			To:   comp.LineCol{Line: 1, Col: 18},
		},
	}
	checkCompletion(t, ":server user list ▼", expected, true)
}

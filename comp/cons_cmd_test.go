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

package comp_test

import (
	"testing"

	. "github.com/abc-inc/persephone/comp"
	"github.com/abc-inc/persephone/types"
)

func TestTypesYieldsConsoleCommandTypeAtTheCommandColon(t *testing.T) {
	checkCompletionTypes(t, "▼:", true, []types.Type{types.ConsoleCommandName})
}

func TestTypesYieldsConsoleCommandTypeAtTheCommandName(t *testing.T) {
	checkCompletionTypes(t, ":▼p", true, []types.Type{types.ConsoleCommandName})
}

func TestTypesYieldsConsoleCommandAtSubCommand(t *testing.T) {
	checkCompletionTypesInfo(t, ":help▼ ", Info{Found: true, Types: []types.Data{{Type: types.ConsoleCommandSubCommand, Path: []string{":help"}}}})
}

func TestTypesYieldsConsoleCommandAtSubCommandPartly(t *testing.T) {
	checkCompletionTypesInfo(t, ":help m▼a",
		Info{Found: true, Types: []types.Data{{Type: types.ConsoleCommandSubCommand, Path: []string{":help", "ma"}, FilterLastElement: true}}})
}

func TestTypesYieldsConsoleCommandAtSubCommandAtSubCommand(t *testing.T) {
	checkCompletionTypesInfo(t, ":server user▼ ",
		Info{Found: true, Types: []types.Data{{Type: types.ConsoleCommandSubCommand, Path: []string{":server", "user"}}}})
}

func TestTypesYieldsConsoleCommandAtSubCommandAtSubCommandPartly(t *testing.T) {
	checkCompletionTypesInfo(t, ":server user l▼i",
		Info{Found: true, Types: []types.Data{{Type: types.ConsoleCommandSubCommand, Path: []string{":server", "user", "li"}, FilterLastElement: true}}})
}

func TestWithoutFiltersYieldsCommandNamesIfColon(t *testing.T) {
	expected := Result{
		Items: []Item{
			{Type: types.ConsoleCommandName, View: ":clear", Content: ":clear"},
			{Type: types.ConsoleCommandName, View: ":play", Content: ":play"},
			{Type: types.ConsoleCommandName, View: ":help", Content: ":help", Postfix: "helpdesc"},
			{Type: types.ConsoleCommandName, View: ":server", Content: ":server"},
		},
		Range: Range{
			From: LineCol{Line: 1, Col: 0},
			To:   LineCol{Line: 1, Col: 1},
		},
	}
	checkCompletion(t, ":▼", expected, false)
}

func TestWithoutFiltersYieldsCommandNamesIfHalfWritten(t *testing.T) {
	expected := Result{
		Items: []Item{
			{Type: types.ConsoleCommandName, View: ":clear", Content: ":clear"},
			{Type: types.ConsoleCommandName, View: ":play", Content: ":play"},
			{Type: types.ConsoleCommandName, View: ":help", Content: ":help", Postfix: "helpdesc"},
			{Type: types.ConsoleCommandName, View: ":server", Content: ":server"},
		},
		Range: Range{
			From: LineCol{Line: 1, Col: 0},
			To:   LineCol{Line: 1, Col: 3},
		},
	}
	checkCompletion(t, ":▼pl", expected, false)
	checkCompletion(t, ":pl▼", expected, false)
}

func TestWithFiltersYieldsCommandNamesIfHalfWritten(t *testing.T) {
	expected := Result{
		Items: []Item{
			{Type: types.ConsoleCommandName, View: ":play", Content: ":play"},
		},
		Range: Range{
			From: LineCol{Line: 1, Col: 0},
			To:   LineCol{Line: 1, Col: 3},
		},
	}
	checkCompletion(t, ":▼pl", expected, true)
	checkCompletion(t, ":pl▼", expected, true)
}

func TestWithFiltersYieldsHelpSubCommand(t *testing.T) {
	expected := Result{
		Items: []Item{
			{Type: types.ConsoleCommandSubCommand, View: "match", Content: "match"},
			{Type: types.ConsoleCommandSubCommand, View: "create", Content: "create"},
		},
		Range: Range{
			From: LineCol{Line: 1, Col: 6},
			To:   LineCol{Line: 1, Col: 6},
		},
	}
	checkCompletion(t, ":help ▼", expected, true)
}

func TestWithFiltersYieldsHelpSubCommandPartly(t *testing.T) {
	expected := Result{
		Items: []Item{
			{Type: types.ConsoleCommandSubCommand, View: "match", Content: "match"},
		},
		Range: Range{
			From: LineCol{Line: 1, Col: 6},
			To:   LineCol{Line: 1, Col: 8},
		},
	}
	checkCompletion(t, ":help ma▼", expected, true)
}

func TestWithFiltersYieldsHelpSubCommandSubCommand(t *testing.T) {
	expected := Result{
		Items: []Item{
			{Type: types.ConsoleCommandSubCommand, View: "list", Content: "list", Postfix: "listdesc"},
			{Type: types.ConsoleCommandSubCommand, View: "add", Content: "add"},
		},
		Range: Range{
			From: LineCol{Line: 1, Col: 13},
			To:   LineCol{Line: 1, Col: 13},
		},
	}
	checkCompletion(t, ":server user ▼", expected, true)
}

func TestWithFiltersYieldsSeverSubCommandSubCommandPartly(t *testing.T) {
	expected := Result{
		Items: []Item{
			{Type: types.ConsoleCommandSubCommand, View: "list", Content: "list", Postfix: "listdesc"},
		},
		Range: Range{
			From: LineCol{Line: 1, Col: 13},
			To:   LineCol{Line: 1, Col: 15},
		},
	}
	checkCompletion(t, ":server user li▼", expected, true)
}

func TestWithFiltersYieldsSeverSubCommandSubCommandNoSubCommand(t *testing.T) {
	expected := Result{
		Range: Range{
			From: LineCol{Line: 1, Col: 18},
			To:   LineCol{Line: 1, Col: 18},
		},
	}
	checkCompletion(t, ":server user list ▼", expected, true)
}

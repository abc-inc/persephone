package autocompletion

import (
	"testing"

	"github.com/abc-inc/merovingian/types"
)

func TestTypesYieldsConsoleCommandTypeAtTheCommandColon(t *testing.T) {
	checkCompletionTypes(t, "▼:", true, []types.Type{types.ConsoleCommandName})
}

func TestTypesYieldsConsoleCommandTypeAtTheCommandName(t *testing.T) {
	checkCompletionTypes(t, ":▼p", true, []types.Type{types.ConsoleCommandName})
}

func TestTypesYieldsConsoleCommandAtSubCommand(t *testing.T) {
	checkCompletionTypes(t, ":help▼", true, []types.Type{types.ConsoleCommandSubCommand})
}

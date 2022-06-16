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

package cmd

import (
	"bytes"
	"fmt"
	"regexp"
	"sort"
	"strings"

	"github.com/abc-inc/persephone/cmd/persephone/cmd/cmdutil"
	"github.com/abc-inc/persephone/cmd/persephone/cmd/help"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func rootUsageFunc(command *cobra.Command) error {
	command.Printf("Usage:  %s", command.UseLine())

	subcommands := command.Commands()
	if len(subcommands) > 0 {
		command.Print("\n\nAvailable commands:\n")
		for _, c := range subcommands {
			if c.Hidden {
				continue
			}
			command.Printf("  %s\n", c.Name())
		}
		return nil
	}

	flagUsages := command.LocalFlags().FlagUsages()
	if flagUsages != "" {
		command.Println("\n\nFlags:")
		command.Print(indent(dedent(flagUsages), "  "))
	}
	return nil
}

var hasFailed bool

// HasFailed signals that the main process should exit with non-zero status
func HasFailed() bool {
	return hasFailed
}

// Display helpful error message in case subcommand name was mistyped.
// This matches Cobra's behavior for root command, which Cobra
// confusingly doesn't apply to nested commands.
func nestedSuggestFunc(command *cobra.Command, arg string) {
	command.Printf("unknown command %q for %q\n", arg, command.CommandPath())

	var candidates []string
	if arg == "help" {
		candidates = []string{"--help"}
	} else {
		if command.SuggestionsMinimumDistance <= 0 {
			command.SuggestionsMinimumDistance = 2
		}
		candidates = command.SuggestionsFor(arg)
	}

	if len(candidates) > 0 {
		command.Print("\nDid you mean this?\n")
		for _, c := range candidates {
			command.Printf("\t%s\n", c)
		}
	}

	command.Print("\n")
	_ = rootUsageFunc(command)
}

func isRootCmd(command *cobra.Command) bool {
	return command != nil && !command.HasParent()
}

func rootHelpFunc(f *cmdutil.Factory, cmd *cobra.Command, args []string) {
	if isRootCmd(cmd.Parent()) && len(args) >= 2 && args[1] != "--help" && args[1] != "-h" {
		nestedSuggestFunc(cmd, args[1])
		hasFailed = true
		return
	}

	namePadding := 12
	coreCommands := []string{}
	actionsCommands := []string{}
	additionalCommands := []string{}
	for _, c := range cmd.Commands() {
		if c.Short == "" {
			continue
		}
		if c.Hidden {
			continue
		}

		s := rpad(c.Name()+":", namePadding) + c.Short
		if _, ok := c.Annotations["IsCore"]; ok {
			coreCommands = append(coreCommands, s)
		} else if _, ok := c.Annotations["IsActions"]; ok {
			actionsCommands = append(actionsCommands, s)
		} else {
			additionalCommands = append(additionalCommands, s)
		}
	}

	// If there are no core commands, assume everything is a core command
	if len(coreCommands) == 0 {
		coreCommands = additionalCommands
		additionalCommands = []string{}
	}

	type helpEntry struct {
		Title string
		Body  string
	}

	longText := cmd.Long
	if longText == "" {
		longText = cmd.Short
	}
	if longText != "" && cmd.LocalFlags().Lookup("jq") != nil {
		longText = strings.TrimRight(longText, "\n") +
			"\n\nFor more information about output formatting flags, see `persephone help formatting`."
	}

	helpEntries := []helpEntry{}
	if longText != "" {
		helpEntries = append(helpEntries, helpEntry{"", longText})
	}
	helpEntries = append(helpEntries, helpEntry{"USAGE", cmd.UseLine()})
	if len(coreCommands) > 0 {
		helpEntries = append(helpEntries, helpEntry{"CORE COMMANDS", strings.Join(coreCommands, "\n")})
	}
	if len(actionsCommands) > 0 {
		helpEntries = append(helpEntries, helpEntry{"ACTIONS COMMANDS", strings.Join(actionsCommands, "\n")})
	}
	if len(additionalCommands) > 0 {
		helpEntries = append(helpEntries, helpEntry{"ADDITIONAL COMMANDS", strings.Join(additionalCommands, "\n")})
	}

	if isRootCmd(cmd) {
		var helpTopics []string
		if c := findCommand(cmd, "actions"); c != nil {
			helpTopics = append(helpTopics, rpad(c.Name()+":", namePadding)+c.Short)
		}
		for _, topic := range help.Topics {
			helpTopics = append(helpTopics, rpad(topic.File+":", namePadding)+topic.Desc)
		}
		sort.Strings(helpTopics)
		helpEntries = append(helpEntries, helpEntry{"HELP TOPICS", strings.Join(helpTopics, "\n")})
	}

	flagUsages := cmd.LocalFlags().FlagUsages()
	if flagUsages != "" {
		helpEntries = append(helpEntries, helpEntry{"FLAGS", dedent(flagUsages)})
	}
	inheritedFlagUsages := cmd.InheritedFlags().FlagUsages()
	if inheritedFlagUsages != "" {
		helpEntries = append(helpEntries, helpEntry{"INHERITED FLAGS", dedent(inheritedFlagUsages)})
	}
	if _, ok := cmd.Annotations["help:arguments"]; ok {
		helpEntries = append(helpEntries, helpEntry{"ARGUMENTS", cmd.Annotations["help:arguments"]})
	}
	if cmd.Example != "" {
		helpEntries = append(helpEntries, helpEntry{"EXAMPLES", cmd.Example})
	}
	if _, ok := cmd.Annotations["help:environment"]; ok {
		helpEntries = append(helpEntries, helpEntry{"ENVIRONMENT VARIABLES", cmd.Annotations["help:environment"]})
	}
	helpEntries = append(helpEntries, helpEntry{"LEARN MORE", `
Use 'persephone <command> <subcommand> --help' for more information about a command.`})
	if _, ok := cmd.Annotations["help:feedback"]; ok {
		helpEntries = append(helpEntries, helpEntry{"FEEDBACK", cmd.Annotations["help:feedback"]})
	}

	for _, e := range helpEntries {
		if e.Title != "" {
			// If there is a title, add indentation to each line in the body
			cmd.Println(color.New(color.Bold).Sprint(e.Title))
			cmd.Println(indent(strings.Trim(e.Body, "\r\n"), "  "))
		} else {
			// If there is no title print the body as is
			cmd.Println(e.Body)
		}
	}
}

func findCommand(cmd *cobra.Command, name string) *cobra.Command {
	for _, c := range cmd.Commands() {
		if c.Name() == name {
			return c
		}
	}
	return nil
}

// rpad adds padding to the right of a string.
func rpad(s string, padding int) string {
	template := fmt.Sprintf("%%-%ds ", padding)
	return fmt.Sprintf(template, s)
}

func dedent(s string) string {
	lines := strings.Split(s, "\n")
	minIndent := -1

	for _, l := range lines {
		if len(l) == 0 {
			continue
		}

		indent := len(l) - len(strings.TrimLeft(l, " "))
		if minIndent == -1 || indent < minIndent {
			minIndent = indent
		}
	}

	if minIndent <= 0 {
		return s
	}

	var buf bytes.Buffer
	for _, l := range lines {
		fmt.Fprintln(&buf, strings.TrimPrefix(l, strings.Repeat(" ", minIndent)))
	}
	return strings.TrimSuffix(buf.String(), "\n")
}

var lineRE = regexp.MustCompile(`(?m)^`)

func indent(s, indent string) string {
	if len(strings.TrimSpace(s)) == 0 {
		return s
	}
	return lineRE.ReplaceAllLiteralString(s, indent)
}

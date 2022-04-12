package cmd

import (
	"strings"

	"github.com/spf13/cobra"
)

const (
	Offline = "offline"
)

func Annotate(keys ...string) map[string]string {
	m := make(map[string]string, len(keys))
	for _, k := range keys {
		m[k] = "true"
	}
	return m
}

func FQCmdName(cmd *cobra.Command) string {
	if cmd == nil || cmd.Parent() == nil {
		return ""
	}
	return strings.TrimPrefix(FQCmdName(cmd.Parent())+" "+cmd.Name(), " ")
}

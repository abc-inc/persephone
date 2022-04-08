package cmd

import (
	"errors"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/abc-inc/browser"
	"github.com/abc-inc/persephone/console"
	"github.com/spf13/cobra"
)

var errTmplOpen = errors.New("cannot open template")

var TemplateCmd = &cobra.Command{
	Use:         ":template",
	Short:       "Define a template",
	Args:        cobra.MinimumNArgs(1),
	Annotations: Annotate(Offline),
	Run:         func(cmd *cobra.Command, args []string) { templateCmd(cmd, args) },
}

func init() {
	TemplateCmd.Flags().StringP("template", "t", "", "template")

	TemplateCmd.AddCommand(&cobra.Command{
		Use:   "edit",
		Short: "Open the template in the default editor",
		Args:  cobra.ExactArgs(1),
		Run:   func(cmd *cobra.Command, args []string) { TemplateEdit(args[0]) },
	})
}

func templateCmd(cmd *cobra.Command, args []string) {
	Template(args...)
}

func Template(args ...string) {
	if len(args) == 1 {
		if t, ok := console.Tmpls[args[0]]; !ok {
			console.Writeln(fmt.Errorf("template %s does not exist", args[0]))
		} else {
			console.Writeln(t.Root.String())
		}
		return
	}

	text := strings.Join(args[1:], " ")
	if _, err := console.SetTemplate(args[0], text); err != nil {
		console.Writeln(err)
	}
}

func TemplateEdit(path string) {
	if !browser.Open(filepath.Join(console.TmplDir, filepath.Base(path))) {
		console.Writeln(errTmplOpen)
	}
}

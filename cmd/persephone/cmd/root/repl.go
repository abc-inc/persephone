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

package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/abc-inc/go-data-neo4j/graph"
	"github.com/abc-inc/go-data-neo4j/meta"
	"github.com/abc-inc/persephone/comp"
	"github.com/abc-inc/persephone/config"
	"github.com/abc-inc/persephone/console"
	"github.com/abc-inc/persephone/console/repl"
	"github.com/abc-inc/persephone/editor"
	"github.com/abc-inc/persephone/internal"
	"github.com/abc-inc/persephone/types"
	"github.com/alecthomas/chroma/quick"
	"github.com/c-bata/go-prompt"
	"github.com/mattn/go-isatty"
	"github.com/muesli/termenv"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var p *prompt.Prompt

func runRepl(cfg config.Config, cmd *cobra.Command) {
	e := editor.NewEditor("")
	e.SetSchema(loadSchema(cmd))

	f := filepath.Join(internal.Must(os.UserCacheDir()), "persephone", "history")
	hist := repl.GetHistory()
	defer func() {
		if err := hist.Save(f); err != nil {
			console.WriteErr(err)
		}
	}()

	var opts []prompt.Option
	opts = append(opts,
		prompt.OptionPrefix(""),
		prompt.OptionCompletionWordSeparator(" "),
		prompt.OptionLivePrefix(func() (prefix string, useLivePrefix bool) {
			if graph.GetConn().DBName == "" {
				return "Disconnected>", true
			}
			return fmt.Sprintf("%s@%s> ", graph.GetConn().Username(), graph.GetConn().DBName), len(lines) == 0
		}),
		prompt.OptionInput(inputCallback(cfg)),
		prompt.OptionPrefixTextColor(prompt.Cyan),
		prompt.OptionHistory(hist.Entries()),
		prompt.OptionSetExitCheckerOnInput(func(in string, breakLine bool) bool {
			return breakLine && in == "exit"
		}),
	)
	opts = append(opts, colorOpts(cfg)...)

	p = prompt.New(func(cyp string) { console.WriteErr(executor(cyp, cmd, hist)) },
		func(d prompt.Document) []prompt.Suggest { return completer(e, d) },
		opts...,
	)
	p.Run()
}

func completer(e *editor.Editor, document prompt.Document) (ss []prompt.Suggest) {
	stmt := strings.TrimLeft(document.TextBeforeCursor(), " ")
	if stmt == "exit" || stmt == ":exit" {
		return nil
	}
	if stmt == "" || strings.IndexRune(");'\"", rune(stmt[len(stmt)-1])) >= 0 {
		return nil
	}

	buf := strings.Join(lines, "\n")
	buf += "\n" + stmt
	buf = strings.TrimPrefix(buf, "\n")

	var res comp.Result
	if strings.HasPrefix(stmt, ":") && strings.IndexByte(stmt, ' ') > 0 {
		res = compConsCmd(stmt)
	} else {
		e.Update(buf)
		line, col := editor.NewPosConv(buf).ToRelative(len(buf))
		res = e.GetCompletion(line, col, true)
	}

	for _, i := range res.Items {
		if stmt == "" && (i.Type == types.Variable || i.Type == types.PropertyKey) {
			continue
		}
		if strings.HasPrefix(i.View, "apoc.") && !strings.Contains(stmt, "apoc.") {
			continue
		}
		if i.View == strings.Trim(i.Content, "`") {
			ss = append(ss, prompt.Suggest{Text: i.View})
		} else {
			ss = append(ss, prompt.Suggest{Text: i.View, Description: i.Content})
		}
	}

	sep := " "
	start := res.Range.From.Col - 1
	if start >= 0 && start < len(document.CurrentLine()) {
		sep = document.CurrentLine()[start : start+1]
	}
	internal.MustNoErr(prompt.OptionCompletionWordSeparator(sep)(p))
	return ss
}

func compConsCmd(stmt string) (res comp.Result) {
	items := listItems(stmt)
	for _, p := range items {
		it := comp.Item{View: p.View, Content: p.Content}
		res.Items = append(res.Items, it)
	}

	start := strings.LastIndex(stmt, "/") + 1
	if start == 0 {
		start = strings.IndexByte(stmt, ' ') + 1
	}
	res.Range = comp.Range{
		From: comp.LineCol{Line: 0, Col: start},
		To:   comp.LineCol{Line: 0, Col: len(stmt)},
	}
	return
}

func listItems(stmt string) []repl.Item {
	args, parts := "", strings.SplitN(stmt, " ", 3)
	for i := len(parts) - 1; i >= 0; i-- {
		if cmdComp := compByConsCmd[strings.Join(parts[:i], " ")]; cmdComp != nil {
			if len(parts) > i {
				args = strings.Join(parts[i:], " ")
			}
			return cmdComp(args)
		}
	}
	return nil
}

func loadSchema(cmd *cobra.Command) comp.Metadata {
	md, err := meta.FetchMetadata(graph.GetConn())
	if err != nil {
		log.Fatal().Err(err).Send()
	}

	ls := make([]string, len(md.Nodes))
	for i, n := range md.Nodes {
		ls[i] = n.String()
	}

	ts := make([]string, len(md.Rels))
	for i, r := range md.Rels {
		ts[i] = r.String()
	}

	ccs := make([]comp.Cmd, len(cmd.Root().Commands()))
	for i, c := range cmd.Root().Commands() {
		ccs[i] = comp.Cmd{Name: c.Name(), Desc: strings.TrimPrefix(c.Name(), ":")}
	}

	schema := comp.Metadata{
		Schema: meta.Schema{
			Labels:   ls,
			RelTypes: ts,
			PropKeys: md.Props,
			Funcs:    md.Funcs,
			Procs:    md.Procs,
		},
		Params:  nil,
		ConCmds: ccs,
	}
	return schema
}

func colorOpts(cfg config.Config) (os []prompt.Option) {
	optByCol := map[string]func(prompt.Color) prompt.Option{
		"prefixTextColor":              prompt.OptionPrefixTextColor,
		"prefixBGColor":                prompt.OptionPrefixBackgroundColor,
		"inputTextColor":               prompt.OptionInputTextColor,
		"inputBGColor":                 prompt.OptionInputBGColor,
		"previewSuggestionTextColor":   prompt.OptionPreviewSuggestionTextColor,
		"previewSuggestionBGColor":     prompt.OptionPreviewSuggestionBGColor,
		"suggestionTextColor":          prompt.OptionSuggestionTextColor,
		"suggestionBGColor":            prompt.OptionSuggestionBGColor,
		"selectedSuggestionTextColor":  prompt.OptionSelectedSuggestionTextColor,
		"selectedSuggestionBGColor":    prompt.OptionSelectedSuggestionBGColor,
		"descriptionTextColor":         prompt.OptionDescriptionTextColor,
		"descriptionBGColor":           prompt.OptionDescriptionBGColor,
		"selectedDescriptionTextColor": prompt.OptionSelectedDescriptionTextColor,
		"selectedDescriptionBGColor":   prompt.OptionSelectedDescriptionBGColor,
		"scrollbarThumbColor":          prompt.OptionScrollbarThumbColor,
		"scrollbarBGColor":             prompt.OptionScrollbarBGColor,
	}

	colByName := map[string]prompt.Color{
		"Default":   prompt.DefaultColor,
		"Black":     prompt.Black,
		"DarkRed":   prompt.DarkRed,
		"DarkGreen": prompt.DarkGreen,
		"Brown":     prompt.Brown,
		"DarkBlue":  prompt.DarkBlue,
		"Purple":    prompt.Purple,
		"Cyan":      prompt.Cyan,
		"LightGray": prompt.LightGray,
		"DarkGray":  prompt.DarkGray,
		"Red":       prompt.Red,
		"Green":     prompt.Green,
		"Yellow":    prompt.Yellow,
		"Blue":      prompt.Blue,
		"Fuchsia":   prompt.Fuchsia,
		"Turquoise": prompt.Turquoise,
		"White":     prompt.White,
	}

	for key, opt := range optByCol {
		if val := cfg.Get(key, "").(string); val != "" {
			if col, ok := colByName[val]; ok {
				log.Debug().Str("property", key).Str("color", val).Msg("Configuring color schema")
				os = append(os, opt(col))
			}
		}
	}
	return
}

func inputCallback(cfg config.Config) func(string) string {
	if !isatty.IsTerminal(os.Stdout.Fd()) {
		return func(s string) string { return s }
	}

	fmtByProfile := map[termenv.Profile]string{
		termenv.TrueColor: "terminal16m",
		termenv.ANSI256:   "terminal256",
		termenv.ANSI:      "terminal16",
		termenv.Ascii:     "noop",
	}
	f := fmtByProfile[termenv.ColorProfile()]

	return func(s string) string {
		style := cfg.Get("colors", "").(string)
		if style == "" {
			return s
		}

		buf := strings.Builder{}
		if err := quick.Highlight(&buf, s, "Cypher", f, style); err != nil {
			return s
		}
		return buf.String()
	}
}

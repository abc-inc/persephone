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

	"github.com/abc-inc/persephone/comp"
	"github.com/abc-inc/persephone/console"
	"github.com/abc-inc/persephone/console/repl"
	"github.com/abc-inc/persephone/editor"
	"github.com/abc-inc/persephone/graph"
	"github.com/abc-inc/persephone/internal"
	"github.com/abc-inc/persephone/types"
	"github.com/c-bata/go-prompt"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var p *prompt.Prompt

func runRepl(cmd *cobra.Command) {
	e := editor.NewEditor("")
	e.SetSchema(loadSchema(cmd))

	f := filepath.Join(internal.Must(os.UserCacheDir()), "persephone", "history")
	hist := repl.GetHistory()
	defer func() {
		if err := hist.Save(f); err != nil {
			console.WriteErr(err)
		}
	}()

	p = prompt.New(func(cyp string) { console.WriteErr(executor(cyp, cmd, hist)) },
		func(d prompt.Document) []prompt.Suggest { return completer(e, d) },
		prompt.OptionSetExitCheckerOnInput(func(in string, breakLine bool) bool {
			return breakLine && in == "exit"
		}), prompt.OptionPrefix(""),
		prompt.OptionPrefixTextColor(prompt.Cyan),
		prompt.OptionCompletionWordSeparator(" "),
		prompt.OptionHistory(hist.Entries()),
		prompt.OptionLivePrefix(func() (prefix string, useLivePrefix bool) {
			if graph.GetConn().DBName == "" {
				return "Disconnected>", true
			}
			return fmt.Sprintf("%s@%s> ", graph.GetConn().Username(), graph.GetConn().DBName), len(lines) == 0
		}),
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

func loadSchema(cmd *cobra.Command) graph.Schema {
	md, err := graph.GetConn().Metadata()
	if err != nil {
		log.Fatal().Err(err).Send()
	}

	ls := make([]string, len(md.Nodes))
	var pkeys []string
	for i, e := range md.Nodes {
		ls[i] = e.String()
		for _, p := range e.Properties {
			pkeys = append(pkeys, p)
		}
	}
	if len(pkeys) == 0 {
		pkeys = append(pkeys, md.Props...)
	}

	ts := make([]string, len(md.Rels))
	for i, r := range md.Rels {
		ts[i] = r.Type
		for p := range r.Properties {
			pkeys = append(pkeys, p)
		}
	}

	ccs := make([]graph.Cmd, len(cmd.Root().Commands()))
	for i, c := range cmd.Root().Commands() {
		ccs[i] = graph.Cmd{Name: c.Name(), Desc: strings.TrimPrefix(c.Name(), ":")}
	}

	schema := graph.Schema{
		Labels:   ls,
		RelTypes: ts,
		PropKeys: pkeys,
		Funcs:    md.Funcs,
		Procs:    md.Procs,
		ConCmds:  ccs,
	}
	return schema
}

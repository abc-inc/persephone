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

package console

import (
	"os"
	"reflect"

	"github.com/abc-inc/gutenfmt/formatter"
	"github.com/abc-inc/gutenfmt/gfmt"
	"github.com/abc-inc/gutenfmt/meta"
	"github.com/mattn/go-isatty"
	"github.com/rs/zerolog/log"
)

type FormatInfo struct {
	Format string
	Sep    string
}

type FmtChangeListener func(i FormatInfo)

var listeners []FmtChangeListener
var info FormatInfo
var w gfmt.Writer

func init() {
	meta.Resolve = meta.TagResolver{TagName: "view"}.Lookup
}

func OnFormatChange(ls ...FmtChangeListener) {
	listeners = append(listeners, ls...)
}

// ChangeFmt creates a new Writer.
func ChangeFmt(f string) {
	sepsByType := map[string]string{"csv": ",", "text": " ", "tsv": "\t"}

	switch f {
	case "":
		w = gfmt.NewAutoJSON(os.Stdout)
	case "auto":
		if isatty.IsTerminal(os.Stdout.Fd()) {
			ChangeFmt("table")
		} else {
			ChangeFmt("json")
		}
		return
	case "csv":
		w = gfmt.NewText(os.Stdout)
		w.(*gfmt.Text).Sep = sepsByType[f]
	case "json":
		w = gfmt.NewJSON(os.Stdout)
	case "jsonc":
		w = gfmt.NewPrettyJSON(os.Stdout)
	case "raw":
		w = gfmt.NewJSON(os.Stdout)
	case "rawc":
		w = gfmt.NewPrettyJSON(os.Stdout)
	case "table":
		w = gfmt.NewTab(os.Stdout)
	case "text":
		w = gfmt.NewText(os.Stdout)
		w.(*gfmt.Text).Sep = sepsByType[f]
	case "tsv":
		w = gfmt.NewText(os.Stdout)
		w.(*gfmt.Text).Sep = sepsByType[f]
	case "yaml":
		w = gfmt.NewYAML(os.Stdout)
	case "yamlc":
		w = gfmt.NewPrettyYAML(os.Stdout)
	default:
		log.Error().Str("format", f).Msg("Unsupported format")
		return
	}

	log.Debug().Str("format", f).Msg("Change output")
	info = FormatInfo{Format: f, Sep: sepsByType[f]}
	for _, l := range listeners {
		l(info)
	}
}

func SetFormatter(i any, f formatter.Func) {
	if t, ok := w.(*gfmt.Text); ok {
		t.Formatter.SetFormatterFunc(reflect.TypeOf(i).String(), f)
	} else if t, ok := w.(*gfmt.Tab); ok {
		t.Formatter.SetFormatterFunc(reflect.TypeOf(i).String(), f)
	}
}

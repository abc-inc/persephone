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

package console

import (
	"os"
	"reflect"

	"github.com/abc-inc/gutenfmt/formatter"
	"github.com/abc-inc/gutenfmt/gfmt"
	"github.com/abc-inc/persephone/event"
	"github.com/mattn/go-isatty"
	"github.com/rs/zerolog/log"
)

var fmtName string
var w gfmt.Writer

func ChangeFmt(f string) {
	sepsByType := map[string]string{"csv": ",", "text": " ", "tsv": "\t"}

	switch f {
	case "":
		w = gfmt.NewAutoJSON(os.Stdout)
	case "auto":
		if isatty.IsTerminal(os.Stdout.Fd()) {
			ChangeFmt("table")
		} else {
			ChangeFmt("")
		}
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

	fmtName = f
	event.Publish(event.FormatEvent{Format: f, Sep: sepsByType[f]})
}

func FormatName() string {
	return fmtName
}

func SetFormatter(i interface{}, f formatter.Func) {
	if t, ok := w.(*gfmt.Text); ok {
		t.Formatter.SetFormatterFunc(reflect.TypeOf(i).String(), f)
	} else if t, ok := w.(*gfmt.Tab); ok {
		t.Formatter.SetFormatterFunc(reflect.TypeOf(i).String(), f)
	}
}

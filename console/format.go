package console

import (
	"os"
	"reflect"
	"unicode"

	"github.com/abc-inc/gutenfmt/formatter"
	"github.com/abc-inc/gutenfmt/gfmt"
	"github.com/abc-inc/persephone/event"
	"github.com/fatih/color"
	"github.com/mattn/go-isatty"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
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

func Writeln(i interface{}) {
	switch err := i.(type) {
	case *neo4j.Neo4jError:
		color.Red(err.Msg)
		return
	case error:
		msg := err.Error()
		r := []rune(msg[0:1])
		r[0] = unicode.ToUpper(r[0])
		color.Red(string(r) + msg[1:])
		return
	}

	_, err := w.Write(i)
	if err != nil {
		log.Fatal().Err(err).Send()
	}
	_, _ = w.Write("\n")
}

func SetFormatter(i interface{}, f formatter.Func) {
	if t, ok := w.(*gfmt.Text); ok {
		t.Formatter.SetFormatterFunc(reflect.TypeOf(i).String(), f)
	} else if t, ok := w.(*gfmt.Tab); ok {
		t.Formatter.SetFormatterFunc(reflect.TypeOf(i).String(), f)
	}
}

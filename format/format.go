package format

import (
	"errors"
	"fmt"
	"log"
	"os"
	"reflect"
	"unicode"

	"github.com/abc-inc/gutenfmt/formatter"
	"github.com/abc-inc/gutenfmt/gfmt"
	"github.com/abc-inc/persephone/event"
	"github.com/fatih/color"
	"github.com/mattn/go-isatty"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j/dbtype"
)

var w gfmt.Writer

func init() {
	if isatty.IsTerminal(os.Stdout.Fd()) && isatty.IsTerminal(os.Stdin.Fd()) {
		Change("table")
	} else {
		Change("")
	}
}

func Change(f string) {
	sepsByType := map[string]string{"csv": ",", "text": " ", "tsv": "\t"}

	switch f {
	case "":
		w = gfmt.NewAutoJSON(os.Stdout)
	case "csv":
		w = gfmt.NewText(os.Stdout)
		w.(*gfmt.Text).Sep = sepsByType[f]
	case "json":
		w = gfmt.NewJSON(os.Stdout)
	case "jsonc":
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
		Writeln(errors.New(fmt.Sprintf("unsupported format '%s'", f)))
		return
	}

	event.Publish(event.FormatEvent{Format: f, Sep: sepsByType[f]})
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
		log.Fatalln(err)
	}
	_, _ = w.Write("\n")
}

func SetFormatter(i interface{}, f formatter.Func) {
	if t, ok := w.(*gfmt.Text); ok {
		t.Formatter.SetFormatterFunc(reflect.TypeOf(i).Name(), f)
	}
}

func MapValues(vs *neo4j.Record) (m map[string]interface{}) {
	m = make(map[string]interface{})
	for i, v := range vs.Values {
		k := vs.Keys[i]
		switch t := v.(type) {
		case dbtype.Node:
			for pk, pv := range t.Props {
				m[k+"."+pk] = pv
			}
		case dbtype.Relationship:
			for pk, pv := range t.Props {
				m[k+"."+pk] = pv
			}
		default:
			panic("not implemented yet: " + reflect.TypeOf(v).Name())
		}
	}
	return m
}

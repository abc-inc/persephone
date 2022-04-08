package console

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/abc-inc/gutenfmt/formatter"
	"github.com/abc-inc/persephone/event"
	"github.com/abc-inc/persephone/graph"
	"github.com/abc-inc/persephone/internal"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

const TmplExt = ".tmpl"

var TmplDir = filepath.Join(internal.Must(os.UserConfigDir()), "persephone", "templates")
var Tmpls = make(map[string]*template.Template)

func init() {
	fs := internal.Must(filepath.Glob(filepath.Join(TmplDir, "*"+TmplExt)))
	for _, f := range fs {
		if t, err := template.ParseFiles(f); err != nil {
			Writeln(err)
		} else {
			Tmpls[filepath.Base(f)] = t
		}
	}

	event.Subscribe(event.FormatEvent{}, func(e event.FormatEvent) {
		SetFormatter(&neo4j.Record{}, FormatTemplate)
	})
}

func SetTemplate(name, text string) (tmpl *template.Template, err error) {
	tmpl = template.New(name)
	for n := range Tmpls {
		if strings.HasSuffix(n, TmplExt) {
			tmpl.Funcs(map[string]any{strings.TrimSuffix(n, TmplExt): func(_ string) string { return "" }})
		}
	}

	if tmpl, err = tmpl.Parse(text); err != nil {
		return nil, err
	}
	Tmpls[name] = tmpl
	if strings.HasSuffix(name, TmplExt) {
		Writeln("Saving template '" + name + "'")
		err = os.WriteFile(filepath.Join(TmplDir, name), []byte(text), 0644)
	}
	return
}

func FormatTemplate(i interface{}) (string, error) {
	rec := i.(*neo4j.Record)
	return format(rec)
}

func format(rec *neo4j.Record) (string, error) {
	m := graph.MapValues(rec)
	j, err := json.Marshal(m)
	if err != nil {
		return "", err
	}
	if err = json.Unmarshal(j, &m); err != nil {
		return "", err
	}

	var ret = ""
	for i := range rec.Keys {
		m2 := m.Values[i].(map[string]interface{})
		if l, ok := m2["@label"]; ok {
			s, err := apply((l.(string) + TmplExt), m2)
			if err != nil {
				return s, err
			}
			ret += s
		}
	}
	return ret, nil
}

func apply(name string, m map[string]interface{}) (string, error) {
	tmpl, ok := Tmpls[name]
	if !ok {
		return "", formatter.ErrUnsupported
	}

	for f, t := range Tmpls {
		n := strings.TrimSuffix(f, TmplExt)
		if f == n {
			continue
		}
		tmpl.Funcs(map[string]any{n: func(v string) (string, error) {
			props, ok := m[v].(map[string]interface{})
			if !ok {
				return "", fmt.Errorf("cannot apply template '%s' because variable '%s' does not exist", n, v)
			}
			s := strings.Builder{}
			if err := t.Execute(&s, props); err != nil {
				return "", err
			}
			return s.String(), nil
		}})
	}

	out := strings.Builder{}
	err := tmpl.Execute(&out, m)
	return out.String(), err
}

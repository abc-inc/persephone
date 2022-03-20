package playground

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"io"
	"reflect"
	"strings"
	"text/template"

	"github.com/abc-inc/gutenfmt/gfmt"
	"github.com/abc-inc/persephone/graph"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j/dbtype"
)

func FormatTemplate(text string, rec graph.Record) (string, error) {
	m := mapValues(rec)
	j, err := json.Marshal(m)
	if err != nil {
		return "", err
	}
	if err = json.Unmarshal(j, &m); err != nil {
		return "", err
	}

	hash := sha256.Sum256([]byte(text))
	n := hex.EncodeToString(hash[:])
	tmpl := template.Must(template.New(n).Parse(text))

	out := strings.Builder{}
	err = tmpl.Execute(&out, m)
	return out.String(), err
}

func mapValues(vs graph.Record) (m map[string]interface{}) {
	m = make(map[string]interface{})
	for k, v := range vs {
		switch t := v.(type) {
		case string:
			m[k] = t
		case dbtype.Node:
			for pk, pv := range t.Props {
				m[k+"."+pk] = pv
			}
		// case dbtype.Relationship:
		default:
			panic("not implemented yet: " + reflect.TypeOf(v).Name())
		}
	}
	return m
}

func NewWriter(fmt string, w io.Writer) gfmt.Writer {
	var f gfmt.Writer
	switch strings.ToLower(fmt) {
	case "":
		f = gfmt.NewAutoJSON(w)
	case "json":
		f = gfmt.NewJSON(w)
	case "jsonc":
		f = gfmt.NewPrettyJSON(w)
	case "table":
		f = gfmt.NewTab(w)
	case "text":
		f = gfmt.NewText(w)
		f.(*gfmt.Text).Sep = "="
	case "tsv":
		f = gfmt.NewText(w)
		f.(*gfmt.Text).Sep = "\t"
	case "yaml":
		f = gfmt.NewYAML(w)
	case "yamlc":
		f = gfmt.NewPrettyYAML(w)
	default:
		f = gfmt.NewText(w)
	}
	return f
}

package playground

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"reflect"
	"strings"
	"text/template"

	"github.com/abc-inc/persephone/graph"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j/dbtype"
)

func FormatTemplate(text string, rec graph.Record) (string, error) {
	m := MapValues(rec)
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

func MapValues(vs graph.Record) (m map[string]interface{}) {
	m = make(map[string]interface{})
	for k, v := range vs.Values {
		switch t := v.(type) {
		case string:
			m[k] = t
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

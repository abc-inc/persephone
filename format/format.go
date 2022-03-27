package format

import (
	"os"
	"reflect"

	"github.com/abc-inc/gutenfmt/gfmt"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j/dbtype"
)

var w gfmt.Writer = gfmt.NewTab(os.Stdout)

func Writeln(i interface{}) {
	_, _ = w.Write(i)
	_, _ = w.Write("\n")
}

func mapValues(vs *neo4j.Record) (m map[string]interface{}) {
	m = make(map[string]interface{})
	for i, v := range vs.Values {
		k := vs.Keys[i]
		switch t := v.(type) {
		case dbtype.Node:
			for pk, pv := range t.Props {
				m[k+"."+pk] = pv
			}
		case dbtype.Relationship:
		default:
			panic("not implemented yet: " + reflect.TypeOf(v).Name())
		}
	}
	return m
}

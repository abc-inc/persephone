package graph

import (
	"reflect"
	"time"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j/dbtype"
)

type Request struct {
	Query  string
	Params map[string]interface{}
}

type ValueExtractor func(key string) (interface{}, bool)

type RecordExtractor func(keys []string, rse ValueExtractor) *neo4j.Record

type RowMapper[T any] func(rec *neo4j.Record) T

func NewSingleColumnRowMapper[T any]() RowMapper[T] {
	return func(rec *neo4j.Record) T {
		return rec.Values[0].(T)
	}
}

func NewMapRowMapper() RowMapper[map[string]interface{}] {
	return func(rec *neo4j.Record) (m map[string]interface{}) {
		m = make(map[string]interface{})
		for i, k := range rec.Keys {
			if _, ok := rec.Values[i].(dbtype.Node); ok {
				m2 := MapValues(rec)
				for i, k2 := range m2.Keys {
					m[k2] = m2.Values[i]
				}
			} else if _, ok := rec.Values[i].(dbtype.Relationship); ok {
				m2 := MapValues(rec)
				for i, k2 := range m2.Keys {
					m[k2] = m2.Values[i]
				}
			} else {
				m[k] = rec.Values[i]
			}
		}
		return
	}
}

func NewResultRowMapper() RowMapper[Result] {
	return func(rec *neo4j.Record) (m Result) {
		for i, k := range rec.Keys {
			if n, ok := rec.Values[i].(neo4j.Node); ok {
				m2 := mapNode(n)
				m.Add(k, m2)
			} else if r, ok := rec.Values[i].(dbtype.Relationship); ok {
				m2 := mapRel(r)
				m.Add(k, m2)
			} else {
				m.Add(k, rec.Values[i])
			}
		}
		return
	}
}

func NewRawResultRowMapper() RowMapper[map[string]interface{}] {
	return func(rec *neo4j.Record) map[string]interface{} {
		m := make(map[string]interface{})
		for i, k := range rec.Keys {
			m[k] = rec.Values[i]
		}
		return m
	}
}

// MapValues maps Neo4j values to serializable Go equivalents.
//
// https://neo4j.com/docs/go-manual/current/cypher-workflow/#go-driver-type-mapping
func MapValues(vs *neo4j.Record) (m Result) {
	for i, k := range vs.Keys {
		v := vs.Values[i]
		// TODO: implement missing types
		switch t := v.(type) {
		// case []interface{}:
		// case map[string]interface{}:
		case bool:
			m.Add(k, t)
		case int64:
			m.Add(k, t)
		case float64:
			m.Add(k, t)
		case string:
			m.Add(k, t)
		case []byte:
			m.Add(k, t)
		case neo4j.Date:
			m.Add(k, t.Time())
		case neo4j.OffsetTime:
			m.Add(k, t.Time())
		case time.Time:
			m.Add(k, t)
		case neo4j.LocalDateTime:
			m.Add(k, t.Time())
		case neo4j.Duration:
			m.Add(k, (time.Duration(t.Days)*24*time.Hour)+(time.Duration(t.Seconds)*time.Second)+time.Duration(t.Nanos))
		case neo4j.Point2D:
			m.Add(k, t)
		case neo4j.Point3D:
			m.Add(k, t)
		case neo4j.Node:
			m.Add(k, mapNode(t))
		case neo4j.Relationship:
			m.Add(k, mapRel(t))
		case neo4j.Path:
			p := []map[string]interface{}{}
			for _, n := range t.Nodes {
				p = append(p, mapNode(n))
			}
			for _, r := range t.Relationships {
				p = append(p, mapRel(r))
			}
			m.Add(k, p)
		default:
			panic("not implemented yet: " + reflect.TypeOf(v).Name())
		}
	}
	return m
}

func mapNode(n neo4j.Node) map[string]interface{} {
	m := make(map[string]interface{})
	for pk, pv := range n.Props {
		m[pk] = pv
	}
	m["@id"] = n.Id
	m["@labels"] = n.Labels
	m["@label"] = n.Labels[0]
	return m
}

func mapRel(r neo4j.Relationship) map[string]interface{} {
	n := make(map[string]interface{})
	for pk, pv := range r.Props {
		n[pk] = pv
	}
	n["@id"] = r.Id
	n["@startId"] = r.StartId
	n["@endId"] = r.EndId
	n["@type"] = r.Type
	return n
}

func NewStructRowMapper[S any]() RowMapper[S] {
	mrm := NewMapRowMapper()
	return func(rec *neo4j.Record) (s S) {
		m := mrm(rec)
		return fillStruct[S](m)
	}
}

func fillStruct[S any](data map[string]interface{}) (result S) {
	t := reflect.ValueOf(result).Elem()
	for k, v := range data {
		val := t.FieldByName(k)
		val.Set(reflect.ValueOf(v))
	}
	return
}

func fillStructResult[S any](data Result) (result S) {
	t := reflect.ValueOf(result).Elem()
	for i, k := range data.Keys {
		val := t.FieldByName(k)
		val.Set(reflect.ValueOf(data.Values[i]))
	}
	return
}

func NewRecordRowMapper() RowMapper[*neo4j.Record] {
	return func(rec *neo4j.Record) *neo4j.Record {
		return rec
	}
}

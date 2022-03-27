package graph

import (
	"reflect"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

const (
	ParamFormat   = "f"
	ParamQuery    = "q"
	ParamTemplate = "t"
)

type Request struct {
	Query    string
	Format   string
	Template string
	Params   map[string]interface{}
}

type Record struct {
	Keys   []string
	Values map[string]interface{}
}

func NewRecord() *Record {
	return &Record{Values: make(map[string]interface{})}
}

func (r *Record) Add(k string, v interface{}) {
	r.Keys = append(r.Keys, k)
	r.Values[k] = v
}

type Result []Record

type ValueExtractor func(key string) (interface{}, bool)

type RecordExtractor func(keys []string, rse ValueExtractor) Record

type RowMapper[T any] func(rec *neo4j.Record) T

func NewSingleColumnRowMapper[T any]() RowMapper[T] {
	return func(rec *neo4j.Record) T {
		return rec.GetByIndex(0).(T)
	}
}

func NewMapRowMapper() RowMapper[map[string]interface{}] {
	return func(rec *neo4j.Record) (m map[string]interface{}) {
		for i, k := range rec.Keys {
			m[k] = rec.GetByIndex(i)
		}
		return
	}
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

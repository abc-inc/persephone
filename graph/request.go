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

package graph

import (
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j/dbtype"
)

type Request struct {
	Query  string
	Params map[string]interface{}
}

type RowMapper[T any] func(rec *neo4j.Record) T

func NewSingleColumnRowMapper[T any]() RowMapper[T] {
	return func(rec *neo4j.Record) T {
		return rec.Values[0].(T)
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

func mapNode(n neo4j.Node) map[string]interface{} {
	m := make(map[string]interface{})
	for pk, pv := range n.Props {
		m[pk] = pv
	}
	m[Id] = n.Id
	m[Labels] = n.Labels
	m[Label] = n.Labels[0]
	return m
}

func mapRel(r neo4j.Relationship) map[string]interface{} {
	n := make(map[string]interface{})
	for pk, pv := range r.Props {
		n[pk] = pv
	}
	n[Id] = r.Id
	n[StartId] = r.StartId
	n[EndId] = r.EndId
	n[Type] = r.Type
	return n
}

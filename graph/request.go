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

import "github.com/neo4j/neo4j-go-driver/v4/neo4j"

// Request is a Cypher query and bind parameters.
type Request struct {
	Query  string
	Params map[string]interface{}
}

// String returns the Cypher query.
func (r Request) String() string {
	return r.Query
}

// Mapper is used by TypedTemplate for mapping results on a per-record basis.
// Implementations of this type perform the actual work of mapping each Record
// to a type, but don't need to worry about error handling. Errors will be
// handled by the calling TypedTemplate.
type Mapper[T any] func(rec *neo4j.Record) T

// NewSingleValueMapper creates a new Mapper that converts a single column into
// a single result value per record. The type of the result value for each
// record can be specified. The value for the single column will be extracted
// from the Record and cast to the specified target type.
func NewSingleValueMapper[T any](idx int) Mapper[T] {
	return func(rec *neo4j.Record) T {
		return rec.Values[idx].(T)
	}
}

// NewResultMapper creates a new Mapper that converts Records to Results.
// For each Node and Relationship, its properties are extracted into a map.
// Additionally, IDs, labels and types are added so that specific modifications
// can be applied by the caller.
func NewResultMapper() Mapper[Result] {
	return func(rec *neo4j.Record) (m Result) {
		for i, k := range rec.Keys {
			if n, ok := rec.Values[i].(neo4j.Node); ok {
				m2 := mapNode(n)
				m.Add(k, m2)
			} else if r, ok := rec.Values[i].(neo4j.Relationship); ok {
				m2 := mapRel(r)
				m.Add(k, m2)
			} else {
				m.Add(k, rec.Values[i])
			}
		}
		return
	}
}

// NewRawResultMapper creates a new Mapper that extracts all columns to
// key-value pairs as they are returned.
func NewRawResultMapper() Mapper[map[string]interface{}] {
	return func(rec *neo4j.Record) map[string]interface{} {
		m := make(map[string]interface{})
		for i, k := range rec.Keys {
			m[k] = rec.Values[i]
		}
		return m
	}
}

// mapNode extracts all properties from the given Node.
// Additionally, it adds "ID", "Labels" and the primary label as "Label".
func mapNode(n neo4j.Node) map[string]interface{} {
	m := make(map[string]interface{})
	for pk, pv := range n.Props {
		m[pk] = pv
	}
	m[ID] = n.Id
	m[Labels] = n.Labels
	m[Label] = n.Labels[0]
	return m
}

// mapRel extracts all properties from the given Relationship.
// Additionally, it adds "ID", "StartID", "EndID" and "Type".
func mapRel(r neo4j.Relationship) map[string]interface{} {
	n := make(map[string]interface{})
	for pk, pv := range r.Props {
		n[pk] = pv
	}
	n[ID] = r.Id
	n[StartID] = r.StartId
	n[EndID] = r.EndId
	n[Type] = r.Type
	return n
}

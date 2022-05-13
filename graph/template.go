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
	"errors"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

// ErrEmpty indicates that a query returned no result.
var ErrEmpty = errors.New("empty")

// ErrMultiple indicates that a query returned more Records than expected.
var ErrMultiple = errors.New("multiple")

// Template simplifies the use of Neo4j and helps to avoid common errors.
// It executes core Neo4j workflow, leaving application code to provide Cypher
// and extract results. Template executes Cypher queries or updates, initiating
// iteration over Results and catching errors. Callers need only to implement
// callback functions, giving them a clearly defined contract.
// All Neo4j operations performed are logged at debug level, using the Logger.
type Template struct {
	conn *Conn
}

// TypedTemplate is an instantiated Template for a specific type.
type TypedTemplate[T any] struct {
	Template
}

// NewTemplate creates a new Template with the given Conn.
func NewTemplate(conn *Conn) *Template {
	return &Template{conn}
}

// NewTypedTemplate creates a new TypedTemplate with the given Conn.
func NewTypedTemplate[T any](conn *Conn) *TypedTemplate[T] {
	return &TypedTemplate[T]{Template: *NewTemplate(conn)}
}

// Query executes the given Cypher with list of parameters to bind to the query,
// mapping each record to a value via a RowMapper. If there is no Transaction
// on this Session, then an explicit transaction is started and committed
// afterwards.
func (t TypedTemplate[T]) Query(r Request, m Mapper[T]) (
	list []T, summary neo4j.ResultSummary, err error) {

	tx, created, err := t.conn.GetTransaction()
	if err != nil {
		return nil, summary, err
	} else if created {
		defer func(tx neo4j.Transaction) {
			_, _ = t.conn.Rollback()
		}(tx)
	}

	res, err := tx.Run(r.Query, r.Params)
	if err != nil {
		return nil, nil, err
	}

	for res.Next() {
		list = append(list, m(res.Record()))
	}
	summary, _ = res.Consume()

	if created {
		_, err = t.conn.Commit()
	}
	return list, summary, err
}

// QuerySingle is like Query, but maps a exactly one result record to a value
// via a Mapper. If the query does not return exactly one record, an error is
// returned.
func (t TypedTemplate[T]) QuerySingle(
	cyp string, params map[string]any, m Mapper[T]) (val T, err error) {

	tx, created, err := t.conn.GetTransaction()
	if err != nil {
		return val, err
	} else if created {
		defer func(conn *Conn) {
			_, _ = conn.Rollback()
		}(t.conn)
	}

	res, err := tx.Run(cyp, params)
	if err != nil {
		return val, err
	} else if !res.Next() {
		return val, ErrEmpty
	}

	val = m(res.Record())
	if res.Next() {
		return val, ErrMultiple
	}

	if created {
		_, err = t.conn.Commit()
	}
	return val, err
}

package graph

import (
	"errors"

	"github.com/abc-inc/persephone/format"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

var errEmpty = errors.New("empty")
var errMultiple = errors.New("multiple")
var errTxCommit = errors.New("cannot commit transaction - maybe it timed out?")
var errTxRollback = errors.New("cannot rollback transaction - maybe it timed out?")

type Template struct {
	conn *Conn
}

type TypedTemplate[T any] struct {
	Template
}

func NewTemplate(conn *Conn) *Template {
	return &Template{conn}
}

func NewTypedTemplate[T any](conn *Conn) *TypedTemplate[T] {
	return &TypedTemplate[T]{Template: *NewTemplate(conn)}
}

func (t TypedTemplate[T]) Query(cyp string, args map[string]interface{}, rm RowMapper[T]) (
	list []T, summary neo4j.ResultSummary, err error) {

	tx, created, err := t.conn.GetTransaction()
	if err != nil {
		return nil, summary, err
	} else if created {
		defer func(tx neo4j.Transaction) {
			if _, err := t.conn.Commit(); err != nil {
				format.Writeln(errTxRollback)
			}
		}(tx)
	}

	res, err := tx.Run(cyp, args)
	if err != nil {
		return nil, nil, err
	}

	for res.Next() {
		list = append(list, rm(res.Record()))
	}
	summary, _ = res.Consume()

	if created {
		if _, err := t.conn.Commit(); err != nil {
			format.Writeln(errTxCommit)
		}
	}
	return list, summary, err
}

func (t TypedTemplate[T]) QuerySingle(cyp string, args map[string]interface{}, rm RowMapper[T]) (val T, err error) {
	tx, created, err := t.conn.GetTransaction()
	if err != nil {
		return val, err
	} else if created {
		defer func(conn *Conn) {
			if _, err := conn.Rollback(); err != nil {
				format.Writeln(errTxRollback)
			}
		}(t.conn)
	}

	res, err := tx.Run(cyp, args)
	if !res.Next() {
		return val, errEmpty
	}

	val = rm(res.Record())
	if res.Next() {
		return val, errMultiple
	}

	if created {
		if _, err := t.conn.Commit(); err != nil {
			format.Writeln(errTxCommit)
		}
	}
	return val, nil
}

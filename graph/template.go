package graph

import (
	"errors"
)

var errEmpty = errors.New("empty")
var errMultiple = errors.New("multiple")

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

func (t TypedTemplate[T]) Query(cyp string, args map[string]interface{}, rm RowMapper[T]) (list []T, err error) {
	tx, created, err := t.conn.GetTransaction()
	if err != nil {
		return nil, err
	} else if created {
		defer tx.Close()
	}

	res, err := tx.Run(cyp, args)
	if err != nil {
		return nil, err
	}

	for res.Next() {
		list = append(list, rm(res.Record()))
	}
	return list, err
}

func (t TypedTemplate[T]) QuerySingle(cyp string, args map[string]interface{}, rm RowMapper[T]) (val T, err error) {
	tx, created, err := t.conn.GetTransaction()
	if err != nil {
		return val, err
	} else if created {
		defer tx.Close()
	}

	res, err := tx.Run(cyp, args)
	if !res.Next() {
		return val, errEmpty
	}

	val = rm(res.Record())
	if res.Next() {
		return val, errMultiple
	}
	return val, nil
}

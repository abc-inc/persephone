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
	"sort"

	"github.com/abc-inc/persephone/internal"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

const systemDB = "system"

var defConn *Conn

// Conn represents a database connection, which can open multiple Sessions.
type Conn struct {
	Driver neo4j.Driver
	user   string
	auth   neo4j.AuthToken
	DBName string
	Tx     neo4j.Transaction
	Params map[string]interface{}
}

// IsConnected returns whether the database connection is established.
func IsConnected() bool {
	return defConn != nil && defConn.DBName != ""
}

// GetConn returns the default connection, regardless of it's connection state.
// It panics if there is no connection.
func GetConn() *Conn {
	if defConn == nil {
		panic("Not connected to Neo4j")
	}
	return defConn
}

// NewConn creates a new Neo4j Driver and returns the new Conn.
func NewConn(addr string, user string, auth neo4j.AuthToken, dbName string) *Conn {
	d := internal.Must(neo4j.NewDriver(addr, auth, func(config *neo4j.Config) {
		config.UserAgent = "persephone (" + neo4j.UserAgent + ")"
	}))

	conn := &Conn{
		Driver: d,
		user:   user,
		auth:   auth,
		DBName: dbName,
		Params: make(map[string]interface{}),
	}
	internal.MustNoErr(conn.UseDB(dbName))

	defConn = conn
	return conn
}

// Close the driver and all underlaying connections.
func (c *Conn) Close() (err error) {
	if c.Driver != nil {
		if err = c.Driver.Close(); err == nil {
			c.Driver, c.Tx = nil, nil
			c.Params = make(map[string]interface{})
			c.DBName = ""
		}
	}
	return err
}

// Session creates a new Session.
func (c Conn) Session() neo4j.Session {
	cfg := neo4j.SessionConfig{DatabaseName: c.DBName}
	return c.Driver.NewSession(cfg)
}

// GetTransaction returns the current Transaction ocr creates a new one.
func (c *Conn) GetTransaction() (tx neo4j.Transaction, created bool, err error) {
	if c.Tx == nil {
		c.Tx, err = c.Session().BeginTransaction()
		created = true
	}
	return c.Tx, created, err
}

// Commit commits the current Transaction.
// If there is no active Transaction, false is returned.
func (c *Conn) Commit() (done bool, err error) {
	if c.Tx != nil {
		err = c.Tx.Commit()
		c.Tx, done = nil, err != nil
	}
	return
}

// Rollback rolls back the current Transaction.
// If there is no active Transaction, false is returned.
func (c *Conn) Rollback() (done bool, err error) {
	if c.Tx != nil {
		err = c.Tx.Rollback()
		c.Tx, done = nil, err != nil
	}
	return
}

func (c *Conn) UseDB(dbName string) (err error) {
	if _, err = c.Rollback(); err != nil {
		return err
	}

	currDBName := c.DBName
	c.DBName = dbName
	_, err = c.Session().ReadTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		return tx.Run("CALL db.ping()", nil)
	})
	var nerr *neo4j.Neo4jError
	if err != nil && errors.As(err, &nerr) {
		if nerr.Title() == "CredentialsExpired" && dbName == systemDB {
			return nil
		}
	} else if err != nil {
		c.DBName = currDBName
	}
	return err
}

// Metadata retrieves schema information like labels, relationships, properties,
// functions and procedures.
func (c Conn) Metadata() (Metadata, error) {
	m := Metadata{}
	if c.DBName == systemDB {
		return m, nil
	}

	funcs, err := c.listFuncs("CALL dbms.functions() YIELD name, signature RETURN name, signature ORDER BY toLower(name)")
	if err != nil {
		return m, err
	}
	m.Funcs = funcs

	funcs, err = c.listFuncs("CALL dbms.procedures() YIELD name, signature RETURN name, signature ORDER BY toLower(name)")
	if err != nil {
		return m, err
	}
	m.Procs = funcs

	idx := sort.Search(len(m.Procs), func(i int) bool {
		return m.Procs[i].Name >= "apoc.meta.schema"
	})

	hasApoc := idx < len(m.Procs) && m.Procs[idx].Name == "apoc.meta.schema"
	if !hasApoc {
		return fallback(c, m)
	}
	return apocMetaGraph(c, m), nil
}

func (c Conn) listFuncs(cyp string) (funcs []Func, err error) {
	res, err := c.Session().Run(cyp, nil)
	if err != nil {
		return nil, err
	}

	for res.Next() {
		f := Func{
			Name: res.Record().Values[0].(string),
			Sig:  res.Record().Values[1].(string),
		}
		funcs = append(funcs, f)
	}
	return
}

func (c *Conn) Username() string {
	if c.user == "" {
		u, err := NewTypedTemplate[string](c).QuerySingle(
			"CALL dbms.showCurrentUser()", nil, NewSingleColumnRowMapper[string]())
		c.user = internal.Must(u, err)
	}
	return c.user
}

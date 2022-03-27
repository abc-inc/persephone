package graph

import (
	"sort"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"github.com/rs/zerolog"
)

var defConn *Conn

type Conn struct {
	Logger zerolog.Logger
	Driver neo4j.Driver
	Tx     neo4j.Transaction
	DBName string
	Params map[string]interface{}
}

func GetConn() *Conn {
	if defConn == nil {
		panic("No connection.")
	}
	return defConn
}

func NewConn(d neo4j.Driver, dbName string) *Conn {
	l := zerolog.New(zerolog.NewConsoleWriter())
	conn := &Conn{l, d, nil, dbName, make(map[string]interface{})}
	if defConn == nil {
		defConn = conn
	}
	return conn
}

func (c *Conn) Close() (err error) {
	if c.Driver != nil {
		if err = c.Driver.Close(); err == nil {
			c.Driver, c.Tx = nil, nil
			c.Params = make(map[string]interface{})
		}
	}
	return err
}

func (c Conn) Session() neo4j.Session {
	cfg := neo4j.SessionConfig{DatabaseName: c.DBName}
	return c.Driver.NewSession(cfg)
}

func (c Conn) Exec(r Request, m RecordExtractor) (Result, error) {
	c.Logger.Info().
		Str("query", r.Query).
		Str("format", r.Format).
		Str("template", r.Template).
		Interface("params", r.Params).
		Msg("Executing query")

	res, err := c.Session().Run(r.Query, r.Params)
	if err != nil {
		return nil, err
	}

	recs := Result{}
	for res.Next() {
		getValue := func(key string) (interface{}, bool) {
			return res.Record().Get(key)
		}
		row := m(res.Record().Keys, getValue)
		recs = append(recs, row)
	}
	return recs, nil
}

func (c *Conn) GetTransaction() (tx neo4j.Transaction, created bool, err error) {
	if c.Tx == nil {
		c.Tx, err = c.Session().BeginTransaction()
		created = true
	}
	return c.Tx, created, err
}

func (c *Conn) Commit() (err error) {
	if c.Tx != nil {
		err = c.Tx.Commit()
		c.Tx = nil
	}
	return
}

func (c *Conn) Rollback() (err error) {
	if c.Tx != nil {
		err = c.Tx.Rollback()
		c.Tx = nil
	}
	return
}

func (c Conn) Metadata() (Metadata, error) {
	m := Metadata{}
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
		return fallback(c, m), nil
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

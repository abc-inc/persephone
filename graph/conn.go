package graph

import (
	"errors"
	"sort"

	"github.com/abc-inc/persephone/internal"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"github.com/rs/zerolog"
)

const systemDB = "system"

var defConn *Conn

type Conn struct {
	Logger zerolog.Logger
	Driver neo4j.Driver
	user   string
	auth   neo4j.AuthToken
	DBName string
	Tx     neo4j.Transaction
	Params map[string]interface{}
}

func IsConnected() bool {
	return defConn != nil && defConn.DBName != ""
}

func GetConn() *Conn {
	if defConn == nil {
		panic("Not connected to Neo4j")
	}
	return defConn
}

func NewConn(addr string, user string, auth neo4j.AuthToken, dbName string) *Conn {
	l := zerolog.New(zerolog.NewConsoleWriter())
	d := internal.Must(neo4j.NewDriver(addr, auth, func(config *neo4j.Config) {
		config.UserAgent = "persephone (" + neo4j.UserAgent + ")"
	}))

	conn := &Conn{
		Logger: l,
		Driver: d,
		user:   user,
		auth:   auth,
		DBName: dbName,
		Params: make(map[string]interface{}),
	}
	internal.MustNoErr(conn.UseDB(dbName))

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
			c.DBName = ""
		}
	}
	return err
}

func (c Conn) Session() neo4j.Session {
	cfg := neo4j.SessionConfig{DatabaseName: c.DBName}
	return c.Driver.NewSession(cfg)
}

func (c Conn) Exec(r Request, m RecordExtractor) (Result, neo4j.ResultSummary, error) {
	c.Logger.Info().
		Str("query", r.Query).
		Str("format", r.Format).
		Str("template", r.Template).
		Interface("params", r.Params).
		Msg("Executing query")

	res, err := c.Session().Run(r.Query, r.Params)
	if err != nil {
		return nil, nil, err
	}

	recs := Result{}
	for res.Next() {
		getValue := func(key string) (interface{}, bool) {
			return res.Record().Get(key)
		}
		row := m(res.Record().Keys, getValue)
		recs = append(recs, row)
	}

	summary, err := res.Consume()
	return recs, summary, err
}

func (c *Conn) GetTransaction() (tx neo4j.Transaction, created bool, err error) {
	if c.Tx == nil {
		c.Tx, err = c.Session().BeginTransaction()
		created = true
	}
	return c.Tx, created, err
}

func (c *Conn) Commit() (done bool, err error) {
	if c.Tx != nil {
		err = c.Tx.Commit()
		c.Tx, done = nil, err != nil
	}
	return
}

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

func (c *Conn) Username() string {
	if c.user == "" {
		u, err := NewTypedTemplate[string](c).QuerySingle(
			"CALL dbms.showCurrentUser()", nil, NewSingleColumnRowMapper[string]())
		c.user = internal.Must(u, err)
	}
	return c.user
}

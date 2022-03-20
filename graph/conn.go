package graph

import (
	"sort"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"github.com/rs/zerolog"
)

type Conn struct {
	logger zerolog.Logger
	driver neo4j.Driver
	DBName string
}

func NewConn(d neo4j.Driver) *Conn {
	l := zerolog.New(zerolog.NewConsoleWriter())
	return &Conn{l, d, "neo4j"}
}

func (c Conn) Close() error {
	return c.driver.Close()
}

func (c Conn) Session() neo4j.Session {
	config := neo4j.SessionConfig{DatabaseName: c.DBName}
	return c.driver.NewSession(config)
}

func (c Conn) Exec(r Request, m RecordExtractor) (Result, error) {
	c.logger.Info().
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

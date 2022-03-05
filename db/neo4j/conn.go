package neo4j

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/abc-inc/merovingian/db"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"github.com/rs/zerolog"
)

type Conn struct {
	logger zerolog.Logger
	driver neo4j.Driver
}

func NewConn(d neo4j.Driver) *Conn {
	l := zerolog.New(zerolog.NewConsoleWriter())
	return &Conn{l, d}
}

func (c Conn) Close() error {
	return c.driver.Close()
}

func (c Conn) Session() neo4j.Session {
	return c.driver.NewSession(neo4j.SessionConfig{})
}

func (c Conn) Exec(r db.Request, m db.RecordExtractor) (db.Result, error) {
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

	recs := db.Result{}
	for res.Next() {
		getValue := func(key string) (interface{}, bool) {
			return res.Record().Get(key)
		}
		row := m(res.Record().Keys, getValue)
		recs = append(recs, row)
	}
	return recs, nil
}

func (c Conn) Metadata() (es []db.Entity, err error) {
	// res, err := c.Session().Run("CALL db.labels()", nil)
	res, err := c.Session().Run("CALL apoc.meta.schema", nil)
	if err != nil {
		return nil, err
	}

	// for res.Next() {
	// l, _ := res.Record().Get(res.Record().Keys[0])
	// es = append(es, db.Entity{Name: l.(string)})
	// }
	for res.Next() {
		m := res.Record().Values[0].(map[string]interface{})
		j, err := json.Marshal(m)
		if err != nil {
			panic(err)
		}
		var nor NodeOrRel
		err = json.Unmarshal(j, &nor)
		if err != nil {
			panic(err)
		}
		fmt.Println(strings.Repeat("%", 80))
		fmt.Println(nor)
		fmt.Println(strings.Repeat("%", 80))
		fmt.Println(string(j))
		fmt.Println(strings.Repeat("%", 80))
		j, err = json.Marshal(m)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(j))
		fmt.Println(strings.Repeat("%", 80))

		for n := range m {
			ps := m[n].(map[string]interface{})
			fmt.Println(n, ps["type"], ps["labels"], ps["types"], ps["properties"])
		}

		// es = append(es, db.Entity{Name: l.(string)})
	}
	return es, nil
}

type NodeOrRel interface {
}

type Node struct {
	Count      int `json:"count"`
	Relationships map[string]RelProperty
	Type string `json:"type"`
	Properties struct {
	} `json:"properties"`
	Labels    []string `json:"labels"`
	NodeOrRel
}

type Relationship struct {
	Count      int    `json:"count"`
	Type       string `json:"type"`
	Properties map[string]NodeRelProperty `json:"properties"`
	NodeOrRel
}

type RelProperty struct {
	Count      int `json:"count"`
	Properties map[string]NodeRelProperty `json:"properties"`
	Direction string   `json:"direction"`
	Labels    []string `json:"labels"`

}
type NodeProperty struct {
	Existence bool   `json:"existence"`
	Type      string `json:"type"`
	Indexed   bool   `json:"indexed"`
	Unique    bool   `json:"unique"`
}

type NodeRelProperty struct {
Existence bool   `json:"existence"`
Type      string `json:"type"`
Array     bool   `json:"array"`
}

type MetaSchema map[string]NodeOrRel

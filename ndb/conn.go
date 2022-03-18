package ndb

import (
	"encoding/json"
	"fmt"
	"strings"

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

func (c Conn) Metadata() (ns []Node, rs []Relationship, err error) {
	// res, err := c.Session().Run("CALL db.labels()", nil)
	res, err := c.Session().Run("CALL apoc.meta.schema", nil)
	if err != nil {
		return nil, nil, err
	}

	// for res.Next() {
	// l, _ := res.Record().Get(res.Record().Keys[0])
	// es = append(es, ndb.Entity{Name: l.(string)})
	// }
	for res.Next() {
		m := res.Record().Values[0].(map[string]interface{})
		j, err := json.Marshal(m)
		if err != nil {
			panic(err)
		}
		var nor Placeholder
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

			var pkeys []string
			for p := range ps["properties"].(map[string]interface{}) {
				pkeys = append(pkeys, p)
			}

			if ps["type"] == "node" {
				ns = append(ns, Node{
					Count:         ps["count"].(int64),
					Relationships: nil,
					Type:          ps["type"].(string),
					Properties:    pkeys,
					Labels:        []string{n},
				})
			} else {
				rs = append(rs, Relationship{
					Count:      ps["count"].(int64),
					Type:       ps["type"].(string),
					Properties: nil,
				})
			}
		}

		// es = append(es, ndb.Entity{Name: l.(string)})
	}
	return ns, rs, nil
}

type Placeholder interface {
}

type NodeOrRel interface {
	String() string
}

type Node struct {
	Count         int64 `json:"count"`
	Relationships map[string]RelProperty
	Type          string   `json:"type"`
	Properties    []string `json:"properties"`
	Labels        []string `json:"labels"`
}

func (n Node) String() string {
	return n.Labels[0]
}

var _ NodeOrRel = (*Node)(nil)

type Relationship struct {
	Count      int64                      `json:"count"`
	Type       string                     `json:"type"`
	Properties map[string]NodeRelProperty `json:"properties"`
}

type RelProperty struct {
	Count      int                        `json:"count"`
	Properties map[string]NodeRelProperty `json:"properties"`
	Direction  string                     `json:"direction"`
	Labels     []string                   `json:"labels"`
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

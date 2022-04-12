package graph

import (
	"strings"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j/db"
)

const (
	Id      = "@id"
	Label   = "@label"
	Labels  = "@labels"
	Type    = "@type"
	StartId = "@startId"
	EndId   = "endId"
)

type Result db.Record

func (r *Result) Add(key string, val interface{}) {
	r.Keys = append(r.Keys, key)
	r.Values = append(r.Values, val)
}

func (r *Result) Index(key string) int {
	for i, k := range r.Keys {
		if k == key {
			return i
		}
	}
	return -1
}

func (r *Result) Value(key string) (interface{}, bool) {
	idx := r.Index(key)
	if idx < 0 {
		return nil, false
	}
	return r.Values[idx], true
}

type Node struct {
	Count         int64 `json:"count"`
	Relationships map[string]RelProperty
	Type          string   `json:"type"`
	Properties    []string `json:"properties"`
	Labels        []string `json:"labels"`
}

func (n Node) String() string {
	return ":" + strings.Join(n.Labels, ":")
}

type Relationship struct {
	Count      int64                      `json:"count"`
	Type       string                     `json:"type"`
	Properties map[string]NodeRelProperty `json:"properties"`
}

func (r Relationship) String() string {
	return r.Type
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

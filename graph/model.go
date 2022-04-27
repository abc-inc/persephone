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
	"fmt"
	"strings"

	"github.com/dustin/go-humanize/english"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j/db"
)

const (
	ID      = "@id"
	Label   = "@label"
	Labels  = "@labels"
	Type    = "@type"
	StartID = "@startId"
	EndID   = "@endId"
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

// Node describes the metamodel of a kind of nodes.
type Node struct {
	Count         int64 `json:"count"`
	Relationships map[string]RelProperty
	Type          string   `json:"type"`
	Properties    []string `json:"properties"`
	Labels        []string `json:"labels"`
}

// String returns the labels of the Node.
func (n Node) String() string {
	return ":" + strings.Join(n.Labels, ":")
}

// Relationship describes the metamodel of a kind of relationships.
type Relationship struct {
	Count      int64                      `json:"count"`
	Type       string                     `json:"type"`
	Properties map[string]NodeRelProperty `json:"properties"`
}

// String returns the relationship type.
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

type StmtType neo4j.StatementType

func (s StmtType) String() string {
	return []string{"UNKNOWN", "READ_ONLY", "READ_WRITE", "WRITE_ONLY", "SCHEMA_WRITE"}[s]
}

func (s StmtType) MarshalJSON() ([]byte, error) {
	return []byte("\"" + s.String() + "\""), nil
}

func (s StmtType) MarshalYAML() (interface{}, error) {
	return s.String(), nil
}

type PlanStats struct {
	Plan      string   `json:"plan" yaml:"plan" table:"Plan"`
	Statement StmtType `json:"queryType" yaml:"queryType" table:"Statement"`
	Version   string   `json:"version" yaml:"version" table:"Version"`
	Planner   string   `json:"planner" yaml:"planner" table:"Planner"`
	Runtime   string   `json:"runtime" yaml:"runtime" table:"Runtime"`
	Time      int64    `json:"time" yaml:"time" table:"Time"`
	DBHits    int64    `json:"dbHits" yaml:"dbHits" table:"DB Hits"`
	Rows      int64    `json:"rows" yaml:"rows" table:"rows"`
	Memory    int64    `json:"memory" yaml:"memory" table:"Memory (Bytes)"`
}

func (p PlanStats) String() string {
	return fmt.Sprintf("%s (%d ms, %d rows, %d DB hits)", p.Plan, p.Time, p.Rows, p.DBHits)
}

type PlanOp struct {
	Op          string    `json:"operatorType" yaml:"operatorType" table:"Operator"`
	Details     string    `json:"details,omitempty" yaml:"details,omitempty" table:"Details,omitempty"`
	RowsEst     int64     `json:"estimatedRows" yaml:"estimatedRows" table:"Estimated Rows"`
	Rows        int64     `json:"rows" yaml:"rows" table:"Rows"`
	DBHits      int64     `json:"dbHits" yaml:"dbHits" table:"DB Hits"`
	Memory      int64     `json:"memory" yaml:"memory" table:"Memory (Bytes)"`
	CacheHits   int64     `json:"pageCacheHits" yaml:"pageCacheHits" table:"Cache Hits"`
	CacheMisses int64     `json:"pageCacheMisses" yaml:"pageCacheMisses" table:"Cache Misses"`
	Order       string    `json:"order,omitempty" yaml:"order,omitempty" table:"Ordered by,omitempty"`
	Children    []*PlanOp `json:"children,omitempty" yaml:"children,omitempty" table:"-"`
}

func (p PlanOp) String() string {
	return fmt.Sprintf("%s (%d %s, %d DB %s)", p.Op,
		p.Rows, english.PluralWord(int(p.Rows), "row", ""),
		p.DBHits, english.PluralWord(int(p.DBHits), "hit", ""))
}

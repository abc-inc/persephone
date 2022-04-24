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
	"strings"

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

type PlanOp struct {
	Op      string   `json:"Operator"`
	Details []string `json:"Details"`
	RowsEst int64    `json:"Estimated Rows"`
	Rows    int64    `json:"Rows"`
	DBHits  int64    `json:"DB Hits"`
	Cache   string   `json:"Page Cache Hits/Misses"`
}

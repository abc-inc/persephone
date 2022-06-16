// Copyright 2022 The Persephone authors
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

package console

import (
	"fmt"
	"math"
	"strings"

	"github.com/abc-inc/go-data-neo4j/plan"
	"github.com/abc-inc/gutenfmt/gfmt"
	"github.com/abc-inc/persephone/internal"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"github.com/xlab/treeprint"
)

// simplePlan contains all the methods Plan and ProfiledPlan have in common.
type simplePlan interface {
	Operator() string
	Arguments() map[string]interface{}
	Identifiers() []string
}

// writePlan writes statistics and the detailed execution plan to the console.
func writePlan(n int, sum neo4j.ResultSummary, sp simplePlan) {
	t := treeprint.New()
	var p []*plan.Op
	if ep, ok := sp.(neo4j.Plan); ok {
		p = formatPlan(t, ep,
			func(p neo4j.Plan) []neo4j.Plan { return p.Children() })
	} else {
		p = formatPlan(t, sp.(neo4j.ProfiledPlan),
			func(p neo4j.ProfiledPlan) []neo4j.ProfiledPlan { return p.Children() })
	}

	planStats := stats(n, sum, sp)
	for _, op := range p {
		planStats.DBHits += op.DBHits
	}

	lines := strings.Split(t.String(), "\n")
	if _, ok := w.(*gfmt.JSON); ok {
		p = []*plan.Op{p[0]}
	} else if _, ok := w.(*gfmt.YAML); ok {
		p = []*plan.Op{p[0]}
	} else {
		for i := range p {
			p[i].Op = lines[i+1]
			p[i].Children = nil
		}
	}
	Write([]plan.Stats{planStats})
	fmt.Println()
	Write(p)
	fmt.Println()
}

// stats extracts the statistics of the execution plan.
func stats(n int, sum neo4j.ResultSummary, sp simplePlan) plan.Stats {
	typ := "EXPLAIN"
	if _, ok := sp.(neo4j.ProfiledPlan); ok {
		typ = "PROFILE"
	}
	planStats := plan.Stats{
		Plan:      typ,
		Statement: plan.StmtType(sum.StatementType()),
		Version:   sp.Arguments()["version"].(string),
		Planner:   sp.Arguments()["planner"].(string),
		Runtime:   sp.Arguments()["runtime"].(string),
		Time:      (sum.ResultAvailableAfter() + sum.ResultConsumedAfter()).Milliseconds(),
		DBHits:    0,
		Rows:      int64(n),
		Memory:    internal.NilToZero[int64](sp.Arguments()["GlobalMemory"]),
	}
	return planStats
}

// formatPlan traverses the execution plan and initializes a data structure
// that holds all operations and metadata.
func formatPlan[SP simplePlan](t treeprint.Tree, sp SP, children func(SP) []SP) []*plan.Op {
	if len(children(sp)) == 0 {
		n := toNode(sp)
		t.AddNode(n.Op)
		return []*plan.Op{n}
	}

	var res []*plan.Op
	n := toNode(sp)
	res = append(res, n)
	br := t.AddBranch(n.Op)

	for _, c := range children(sp) {
		sp := formatPlan[SP](br, c, children)
		n.Children = append(n.Children, sp[0])
		res = append(res, sp...)
	}
	return res
}

// toNode extracts the metadata from the plan entry.
func toNode(p simplePlan) *plan.Op {
	rows, dbHits, cHits, cMiss := int64(0), int64(0), int64(0), int64(0)
	if pp, ok := p.(neo4j.ProfiledPlan); ok {
		rows, dbHits, cHits, cMiss = pp.Records(), pp.DbHits(), pp.PageCacheHits(), pp.PageCacheMisses()
	}
	op, _, _ := strings.Cut(p.Operator(), "@")
	return &plan.Op{
		Op:          op,
		Details:     internal.NilToZero[string](p.Arguments()["Details"]),
		RowsEst:     int64(math.Round(p.Arguments()["EstimatedRows"].(float64))),
		Rows:        rows,
		DBHits:      dbHits,
		Memory:      internal.NilToZero[int64](p.Arguments()["Memory"]),
		CacheHits:   cHits,
		CacheMisses: cMiss,
		Order:       internal.NilToZero[string](p.Arguments()["Order"]),
	}
}

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

package comp_test

import (
	"testing"

	"github.com/abc-inc/persephone/comp"
	"github.com/abc-inc/persephone/types"
)

func TestReturnProcedureNameType(t *testing.T) {
	checkCompletionTypesInfo(t, "call ▼db.proc()",
		comp.Info{Found: true, Types: []types.Data{{Type: types.ProcedureName, FilterLastElement: false}}})
}

func TestReturnProcedureNameTypeIfOnlyCallPresent(t *testing.T) {
	checkCompletionTypesInfo(t, "call▼ ",
		comp.Info{Found: true, Types: []types.Data{{Type: types.ProcedureName, FilterLastElement: false}}})
}

func TestReturnOutputAtTheBeginningOfYield(t *testing.T) {
	checkCompletionTypesInfo(t, "call db.proc() yield▼ ",
		comp.Info{Found: true, Types: []types.Data{{Type: types.ProcedureOutput, Path: []string{"db.proc"}, FilterLastElement: false}}})
}

func TestReturnOutputBeforeTheFirstTypedSymbol(t *testing.T) {
	checkCompletionTypesInfo(t, "call db.proc() yield ▼a",
		comp.Info{Found: true, Types: []types.Data{{Type: types.ProcedureOutput, Path: []string{"db.proc"}, FilterLastElement: false}}})
}

func TestReturnOutputAtTheBeginningOfSecondOutput(t *testing.T) {
	checkCompletionTypesInfo(t, "call db.proc() yield a as b,▼ ",
		comp.Info{Found: true, Types: []types.Data{{Type: types.ProcedureOutput, Path: []string{"db.proc"}, FilterLastElement: false}}})
}

func TestWithoutFiltersYieldsProcedureNameList(t *testing.T) {
	exp := comp.Result{
		Range: comp.Range{
			From: comp.LineCol{Line: 1, Col: 5},
			To:   comp.LineCol{Line: 1, Col: 8},
		},
		Items: []comp.Item{
			{Type: types.ProcedureName, View: "db.indexes", Content: "db.indexes", Postfix: "()"},
			{Type: types.ProcedureName, View: "org.neo4j.graph.traverse", Content: "org.neo4j.graph.traverse", Postfix: "expression"},
		},
	}

	checkCompletion(t, "call ▼d.p", exp, false)
	checkCompletion(t, "call d▼.p", exp, false)
	checkCompletion(t, "call d.▼p", exp, false)
	checkCompletion(t, "call d.p▼", exp, false)
}

func TestWithoutFiltersYieldsProcedureNameListIfOnlyCallPresent(t *testing.T) {
	exp := comp.Result{
		Range: comp.Range{
			From: comp.LineCol{Line: 1, Col: 5},
			To:   comp.LineCol{Line: 1, Col: 5},
		},
		Items: []comp.Item{
			{Type: types.ProcedureName, View: "db.indexes", Content: "db.indexes", Postfix: "()"},
			{Type: types.ProcedureName, View: "org.neo4j.graph.traverse", Content: "org.neo4j.graph.traverse", Postfix: "expression"},
		},
	}

	checkCompletion(t, "call ▼", exp, false)
}

func TestWithFiltersYieldsProcedureNameList(t *testing.T) {
	exp := comp.Result{
		Range: comp.Range{
			From: comp.LineCol{Line: 1, Col: 5},
			To:   comp.LineCol{Line: 1, Col: 9},
		},
		Items: []comp.Item{
			{Type: types.ProcedureName, View: "db.indexes", Content: "db.indexes", Postfix: "()"},
		},
	}

	checkCompletion(t, "call ▼db.i", exp, true)
	checkCompletion(t, "call d▼b.i", exp, true)
	checkCompletion(t, "call db▼.i", exp, true)
	checkCompletion(t, "call db.▼i", exp, true)
	checkCompletion(t, "call db.i▼", exp, true)
}

func TestWithFiltersYieldsAllProcedureReturnItems(t *testing.T) {
	exp := comp.Result{
		Range: comp.Range{
			From: comp.LineCol{Line: 1, Col: 24},
			To:   comp.LineCol{Line: 1, Col: 24},
		},
		Items: []comp.Item{
			{Type: types.ProcedureOutput, View: "description", Content: "description", Postfix: " :: STRING?"},
			{Type: types.ProcedureOutput, View: "state", Content: "state", Postfix: " :: STRING?"},
			{Type: types.ProcedureOutput, View: "type", Content: "type", Postfix: " :: STRING?"},
		},
	}

	checkCompletion(t, "call db.indexes() yield ▼", exp, true)
}

func TestWithFiltersYieldsAllProcedureReturnItemsAfterFirstSymbolIsTyped(t *testing.T) {
	exp := comp.Result{
		Range: comp.Range{
			From: comp.LineCol{Line: 1, Col: 24},
			To:   comp.LineCol{Line: 1, Col: 25},
		},
		Items: []comp.Item{
			{Type: types.ProcedureOutput, View: "state", Content: "state", Postfix: " :: STRING?"},
		},
	}

	checkCompletion(t, "call db.indexes() yield a▼", exp, true)
}

func TestWithFiltersYieldsAllProcedureReturnItemsAfterExpression(t *testing.T) {
	exp := comp.Result{
		Range: comp.Range{
			From: comp.LineCol{Line: 1, Col: 32},
			To:   comp.LineCol{Line: 1, Col: 32},
		},
		Items: []comp.Item{
			{Type: types.ProcedureOutput, View: "description", Content: "description", Postfix: " :: STRING?"},
			{Type: types.ProcedureOutput, View: "state", Content: "state", Postfix: " :: STRING?"},
			{Type: types.ProcedureOutput, View: "type", Content: "type", Postfix: " :: STRING?"},
		},
	}

	checkCompletion(t, "call db.indexes() yield a as b, ▼", exp, true)
}

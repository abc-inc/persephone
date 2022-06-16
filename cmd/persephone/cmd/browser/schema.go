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

package cmd

import (
	"github.com/abc-inc/go-data-neo4j/graph"
	"github.com/abc-inc/persephone/cmd/persephone/cmd/cmdutil"
	"github.com/abc-inc/persephone/console"
	"github.com/abc-inc/persephone/internal"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"github.com/spf13/cobra"
)

type index struct {
	Name          string   `json:"Index Name" view:"Index Name" yaml:"Index Name"`
	Type          string   `json:"Type" view:"Type" yaml:"Type"`
	Uniqueness    string   `json:"Uniqueness" view:"Uniqueness" yaml:"Uniqueness"`
	EntityType    string   `json:"EntityType" view:"EntityType" yaml:"EntityType"`
	LabelsOrTypes []string `json:"LabelsOrTypes" view:"LabelsOrTypes" yaml:"LabelsOrTypes"`
	Properties    []string `json:"Properties" view:"Properties" yaml:"Properties"`
	State         string   `json:"State" view:"State" yaml:"State"`
}

func (i index) String() string {
	return i.Name
}

func NewCmdSchema(f *cmdutil.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use:   ":schema",
		Short: "Show information about database schema indexes and constraints",
		Run:   func(cmd *cobra.Command, args []string) { Schema() },
	}

	return cmd
}

func Schema() {
	cyp := "CALL db.indexes() YIELD name AS `Index Name`, type AS Type, uniqueness AS Uniqueness, " +
		"entityType AS EntityType, labelsOrTypes AS LabelsOrTypes, properties AS Properties, state AS State " +
		"RETURN `Index Name`, Type, Uniqueness, EntityType, LabelsOrTypes, Properties, State " +
		"ORDER BY `Index Name`;"

	t := graph.NewTemplate[index](graph.GetConn())
	r := graph.Request{Query: cyp}
	idxs, _ := internal.MustTuple(t.Query(r, func(rec *neo4j.Record) index {
		ls := internal.MustOk(rec.Get("LabelsOrTypes")).([]any)
		ps := internal.MustOk(rec.Get("Properties")).([]any)

		return index{
			Name:          internal.MustOk(rec.Get("Index Name")).(string),
			Type:          internal.MustOk(rec.Get("Type")).(string),
			Uniqueness:    internal.MustOk(rec.Get("Uniqueness")).(string),
			EntityType:    internal.MustOk(rec.Get("Index Name")).(string),
			LabelsOrTypes: internal.ReSlice[string](ls),
			Properties:    internal.ReSlice[string](ps),
			State:         internal.MustOk(rec.Get("State")).(string),
		}
	}))

	console.Write(idxs)
}

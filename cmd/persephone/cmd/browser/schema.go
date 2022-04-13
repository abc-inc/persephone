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

package cmd

import (
	"github.com/abc-inc/persephone/console"
	"github.com/abc-inc/persephone/graph"
	. "github.com/abc-inc/persephone/internal"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"github.com/spf13/cobra"
)

type index struct {
	Name          string   `json:"Index Name"`
	Type          string   `json:"Type"`
	Uniqueness    string   `json:"Uniqueness"`
	EntityType    string   `json:"EntityType"`
	LabelsOrTypes []string `json:"LabelsOrTypes"`
	Properties    []string `json:"Properties"`
	State         string   `json:"State"`
}

func (i index) String() string {
	return i.Name
}

var SchemaCmd = &cobra.Command{
	Use:   ":schema",
	Short: "Show information about database schema indexes and constraints",
	Run:   func(cmd *cobra.Command, args []string) { Schema() },
}

func Schema() {
	cyp := "CALL db.indexes() YIELD name AS `Index Name`, type AS Type, uniqueness AS Uniqueness, " +
		"entityType AS EntityType, labelsOrTypes AS LabelsOrTypes, properties AS Properties, state AS State " +
		"RETURN `Index Name`, Type, Uniqueness, EntityType, LabelsOrTypes, Properties, State " +
		"ORDER BY `Index Name`;"

	t := graph.NewTypedTemplate[index](graph.GetConn())
	idxs, _ := MustTuple(t.Query(cyp, nil, func(rec *neo4j.Record) index {
		ls := MustOk(rec.Get("LabelsOrTypes")).([]interface{})
		ps := MustOk(rec.Get("Properties")).([]interface{})

		return index{
			Name:          MustOk(rec.Get("Index Name")).(string),
			Type:          MustOk(rec.Get("Type")).(string),
			Uniqueness:    MustOk(rec.Get("Uniqueness")).(string),
			EntityType:    MustOk(rec.Get("Index Name")).(string),
			LabelsOrTypes: ReSlice[string](ls),
			Properties:    ReSlice[string](ps),
			State:         MustOk(rec.Get("State")).(string),
		}
	}))

	console.Write(idxs)
}

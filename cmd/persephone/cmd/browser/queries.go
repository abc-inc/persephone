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
	"encoding/json"
	"time"

	"github.com/abc-inc/persephone/cmd/persephone/cmd/cmdutil"
	"github.com/abc-inc/persephone/console"
	"github.com/abc-inc/persephone/graph"
	"github.com/abc-inc/persephone/internal"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"github.com/spf13/cobra"
)

type query struct {
	DBURI   string        `json:"Database URI" table:"Database URI"`
	User    string        `json:"User" table:"User"`
	Query   string        `json:"Query" table:"Query"`
	Params  any           `json:"Params" table:"Params"`
	Meta    any           `json:"Meta" table:"Meta"`
	Elapsed time.Duration `json:"Elapsed time" table:"Elapsed time"`
}

func (q query) String() string {
	return q.Query
}

func NewCmdQueries(f *cmdutil.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use:   ":queries",
		Short: "List your servers and clusters running queries",
		Run:   func(cmd *cobra.Command, args []string) { Queries() },
	}

	return cmd
}

func Queries() {
	t := graph.NewTypedTemplate[query](graph.GetConn())
	r := graph.Request{Query: "CALL dbms.listQueries()"}
	qs, _, err := t.Query(r, func(rec *neo4j.Record) query {
		return query{
			DBURI:   "neo4j://" + internal.MustOk(rec.Get("requestUri")).(string),
			User:    internal.MustOk(rec.Get("username")).(string),
			Meta:    internal.MustOk(rec.Get("metaData")),
			Query:   internal.MustOk(rec.Get("query")).(string),
			Params:  string(internal.Must(json.Marshal(internal.MustOk(rec.Get("parameters"))))),
			Elapsed: time.Duration(internal.MustOk(rec.Get("elapsedTimeMillis")).(int64)),
		}
	})

	if err != nil {
		console.WriteErr(err)
	} else {
		console.Write(qs)
	}
}

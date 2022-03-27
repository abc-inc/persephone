package cmd

import (
	"github.com/abc-inc/persephone/format"
	"github.com/abc-inc/persephone/graph"
	. "github.com/abc-inc/persephone/internal"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"github.com/spf13/cobra"
)

type query struct {
	QueryId    string
	Username   string
	MetaData   interface{}
	Query      string
	Parameters interface{}
	StartTime  neo4j.Date
}

var QueriesCmd = &cobra.Command{
	Use:   ":queries",
	Short: "List your servers and clusters running queries",
	Run:   queriesCmd,
}

func queriesCmd(cmd *cobra.Command, args []string) {
	t := graph.NewTypedTemplate[query](graph.GetConn())
	qs := Must(t.Query("CALL dbms.listQueries()", nil, func(rec *neo4j.Record) query {
		return query{
			QueryId:    MustOk(rec.Get("queryId")).(string),
			Username:   MustOk(rec.Get("username")).(string),
			MetaData:   MustOk(rec.Get("metaData")),
			Query:      MustOk(rec.Get("query")).(string),
			Parameters: MustOk(rec.Get("parameters")),
			StartTime:  MustOk(rec.Get("startTime")).(neo4j.Date),
		}
	}))

	format.Writeln(qs)
}

package cmd

import (
	"github.com/abc-inc/persephone/format"
	"github.com/abc-inc/persephone/graph"
	. "github.com/abc-inc/persephone/internal"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"github.com/spf13/cobra"
)

type dbInfo struct {
	Name    string
	Address string
	Role    string
	Status  string
	Default bool
	Error   string
}

var SysinfoCmd = &cobra.Command{
	Use:   ":sysinfo",
	Short: "Print system information",
	Run:   sysinfoCmd,
}

func sysinfoCmd(cmd *cobra.Command, args []string) {
	t := graph.NewTypedTemplate[dbInfo](graph.GetConn())
	dbs := Must(t.Query("SHOW DATABASES", nil, func(rec *neo4j.Record) dbInfo {
		return dbInfo{
			Name:    MustOk(rec.Get("name")).(string),
			Address: MustOk(rec.Get("address")).(string),
			Role:    MustOk(rec.Get("role")).(string),
			Status:  MustOk(rec.Get("requestedStatus")).(string),
			Default: MustOk(rec.Get("default")).(bool),
			Error:   MustOk(rec.Get("error")).(string),
		}
	}))

	format.Writeln(dbs)
}

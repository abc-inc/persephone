package cmd

import (
	"strings"
	"sync"

	cmd "github.com/abc-inc/persephone/cmd/persephone/cmd/browser"
	"github.com/abc-inc/persephone/console"
	"github.com/abc-inc/persephone/console/repl"
	"github.com/abc-inc/persephone/graph"
	"github.com/spf13/cobra"
)

var dbsCache sync.Once

var UseCmd = &cobra.Command{
	Use:   ":use database",
	Short: "Set the active database",
	Long:  "Set the active database that transactions are executed on",
	Run:   func(cmd *cobra.Command, args []string) { Use(args[0]) },
}

func Use(dbName string) {
	if err := graph.GetConn().UseDB(dbName); err != nil {
		console.WriteErr(err)
	}
}

func UseCompFunc() func(string) []repl.Item {
	var dbs []cmd.DBInfo
	return func(s string) (its []repl.Item) {
		dbsCache.Do(func() {
			dbs = cmd.ListDBs()
		})

		for _, db := range dbs {
			if strings.HasPrefix(db.Name, s) || strings.Contains(db.Address, s) {
				its = append(its, repl.Item{View: db.Name, Content: db.Address})
			}
		}
		return
	}
}

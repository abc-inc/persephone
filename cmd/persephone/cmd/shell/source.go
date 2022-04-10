package cmd

import (
	"bufio"
	"os"
	"strings"

	"github.com/abc-inc/persephone/console"
	"github.com/abc-inc/persephone/graph"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var SourceCmd = &cobra.Command{
	Use:   ":source [filename]",
	Short: "Execute Cypher statements from a file",
	Args:  cobra.ExactArgs(1),
	Run:   func(cmd *cobra.Command, args []string) { Source(args[0]) },
}

func Source(path string) {
	f, err := os.Open(path)
	if err != nil {
		console.Writeln(err)
		return
	}

	defer func(f *os.File) {
		_ = f.Close()
	}(f)

	tx, created, err := graph.GetConn().GetTransaction()
	if err != nil {
		console.Writeln(err)
		return
	} else if created {
		defer func(tx neo4j.Transaction) {
			_, _ = graph.GetConn().Rollback()
		}(tx)
	}

	sc := bufio.NewScanner(f)
	for sc.Scan() {
		if strings.HasPrefix(sc.Text(), ":") {
			continue
		}

		log.Debug().Str("statement", sc.Text()).Fields(graph.GetConn().Params).Msg("Executing")
		if result, err := tx.Run(sc.Text(), graph.GetConn().Params); err != nil {
			console.Writeln(err)
			return
		} else {
			for result.Next() {
				console.Writeln(result.Record())
			}
		}
	}

	_, _ = graph.GetConn().Commit()
}

package cmd

import (
	"bufio"
	"os"
	"strings"

	"github.com/abc-inc/persephone/format"
	"github.com/abc-inc/persephone/graph"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
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
		format.Writeln(err)
		return
	}

	defer func(f *os.File) {
		_ = f.Close()
	}(f)

	tx, created, err := graph.GetConn().GetTransaction()
	if err != nil {
		format.Writeln(err)
		return
	} else if created {
		defer func(tx neo4j.Transaction) {
			_, _ = graph.GetConn().Rollback()
		}(tx)
	}

	s := bufio.NewScanner(f)
	for s.Scan() {
		if strings.HasPrefix(s.Text(), ":") {
			continue
		}

		format.Writeln(s.Text())
		if result, err := tx.Run(s.Text(), nil); err != nil {
			format.Writeln(err)
			return
		} else if _, err := result.Consume(); err != nil {
			format.Writeln(err)
			return
		}
	}

	_, _ = graph.GetConn().Commit()
}

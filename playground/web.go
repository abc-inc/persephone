package playground

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"os"
	"text/template"

	"github.com/abc-inc/gutenfmt/formatter"
	"github.com/abc-inc/persephone/console"
	"github.com/abc-inc/persephone/format"
	"github.com/abc-inc/persephone/graph"
	"github.com/dustin/go-humanize/english"
)

func Foo(req graph.Request) error {
	t := graph.NewTypedTemplate[map[string]interface{}](graph.GetConn())

	sp := console.NewSpinner()
	sp.Start()
	list, summary, err := t.Query(req.Query, req.Params, graph.NewMapRowMapper())
	sp.Stop()
	if err != nil {
		return err
	}

	format.Writeln(list)
	fmt.Fprintf(os.Stderr, "\n%d %s, ready to start consuming query after %s, results consumed after another %s\n",
		len(list), english.PluralWord(len(list), "row", "rows"),
		summary.ResultAvailableAfter(), summary.ResultConsumedAfter())

	return err

	tStr := `
	   {{range $a := .}}
	   {{index $a "id"}}: {{index $a "name"}}
	   {{end}}`
	tmpl := template.Must(template.New("main").Parse(tStr))
	j, err := json.Marshal(list)
	jj := []map[string]interface{}{}
	err = json.Unmarshal(j, &jj)
	if err != nil {
		fmt.Println("ERROR", err)
		return err
	}
	str, err := formatter.FromTemplate(tmpl).Format(jj)
	if err != nil {
		return err
	}

	fmt.Println(jj, str, err)
	format.Writeln(list)
	return nil
}

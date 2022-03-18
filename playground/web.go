package playground

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"io"
	"text/template"

	"github.com/abc-inc/gutenfmt/formatter"
	"github.com/abc-inc/merovingian/ndb"
)

func Foo(w io.Writer, s ndb.Connector, req ndb.Request) error {
	rse := func(keys []string, rse ndb.ValueExtractor) ndb.Record {
		m := map[string]interface{}{}
		for _, k := range keys {
			m[k], _ = rse(k)
		}
		return m
	}

	res, err := s.Exec(req, rse)
	if err != nil {
		fmt.Println(err)
		return err
	}

	t := `
{{range $a := .}}
{{index $a "id"}}: {{index $a "name"}}
{{end}}`
	tmpl := template.Must(template.New("main").Parse(t))
	j, err := json.Marshal(res)
	jj := []map[string]interface{}{}
	err = json.Unmarshal(j, &jj)
	if err != nil {
		fmt.Println("ERROR", err)
		return err
	}
	str, err := formatter.FromTemplate(tmpl).Format(jj)
	fmt.Println(jj, str, err)

	gfmt := NewWriter(req.Format, w)
	_, err = gfmt.Write(res)
	return err
}

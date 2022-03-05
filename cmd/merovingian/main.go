package main

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"text/template"
	"time"

	"github.com/abc-inc/merovingian/_deprecated"
	"github.com/briandowns/spinner"
	"github.com/c-bata/go-prompt"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j/db"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j/dbtype"
)

func main() {
	dbUri := "neo4j://localhost:7687"
	driver, err := neo4j.NewDriver(dbUri, neo4j.BasicAuth("neo4j", "root", ""))
	if err != nil {
		panic(err)
	}
	// Handle driver lifetime based on your application lifetime requirements  driver's lifetime is usually
	// bound by the application lifetime, which usually implies one driver instance per application
	defer driver.Close()

	sess := driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})
	defer sess.Close()
	res, err := sess.Run("CALL db.labels()", nil)
	if err != nil {
		panic(err)
	}

	labels := []prompt.Suggest{}
	for res.Next() {
		labels = append(labels, prompt.Suggest{Text: res.Record().Values[0].(string)})
	}

	p := _deprecated.NewPrompt()
	p.Register(_deprecated.Node, func(d prompt.Document) []prompt.Suggest {
		return labels
	})
	p.Run()

	fmt.Println("You selected " + p.Input())
	return

	item, err := insertItem(driver)
	if err != nil {
		panic(err)
	}
	n := mapValues(item.(*db.Record))
	json, err := json.Marshal(n)
	fmt.Println("[" + string(json) + "]")
	return

	tmpl := template.Must(template.New("").Parse("{{index . \"r.path\"}}\t{{index . \"f.gradleVersion\"}}"))
	err = tmpl.Execute(os.Stderr, n)
	if err != nil {
		panic(err)
	}
	return

	s := spinner.New(spinner.CharSets[14], 100*time.Millisecond) // Build our new spinner
	s.Start()                                                    // Start the spinner
	time.Sleep(4 * time.Millisecond)                             // Run for some time to simulate work
	s.Stop()
	return

}

func mapValues(vs *neo4j.Record) (m map[string]interface{}) {
	m = make(map[string]interface{})
	for i, v := range vs.Values {
		k := vs.Keys[i]
		switch t := v.(type) {
		case dbtype.Node:
			for pk, pv := range t.Props {
				m[k+"."+pk] = pv
			}
		case dbtype.Relationship:
		default:
			panic("not implemented yet: " + reflect.TypeOf(v).Name())
		}
	}
	return m
}

func insertItem(driver neo4j.Driver) (interface{}, error) {
	session := driver.NewSession(neo4j.SessionConfig{})
	defer session.Close()
	result, err := session.WriteTransaction(createItemFn)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func createItemFn(tx neo4j.Transaction) (interface{}, error) {
	records, err := tx.Run("MATCH (r:Repository)-[r_f]-(f:File) RETURN r, r_f, f LIMIT 1", map[string]interface{}{
		"id":   1,
		"name": "Item 1",
	})
	// In face of driver native errors, make sure to return them directly.
	// Depending on the error, the driver may try to execute the function again.
	if err != nil {
		return nil, err
	}
	record, err := records.Single()
	if err != nil {
		return nil, err
	}
	// You can also retrieve values by name, with e.g. `id, found := record.Get("n.id")`
	return record, nil
}

type Item struct {
	Id   int64
	Name string
}

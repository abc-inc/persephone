package graph

import (
	"log"
	"sort"
)

type Metadata struct {
	Nodes []Node
	Rels  []Relationship
	Funcs []Func
	Procs []Func
	Props []string
}

func apocMetaGraph(c Conn, m Metadata) Metadata {
	res, err := c.Session().Run("CALL apoc.meta.schema", nil)
	if err != nil {
		return m
	}
	rec, err := res.Single()
	if err != nil {
		return m
	}

	var present interface{}
	allProps := make(map[string]interface{}, 0)
	kvs := rec.Values[0].(map[string]interface{})
	for k, v := range kvs {
		kv := v.(map[string]interface{})

		var pkeys []string
		for p := range kv["properties"].(map[string]interface{}) {
			pkeys = append(pkeys, p)
			allProps[p] = present
		}

		if kv["type"] == "node" {
			m.Nodes = append(m.Nodes, Node{
				Count:         kv["count"].(int64),
				Relationships: nil,
				Type:          kv["type"].(string),
				Properties:    pkeys,
				Labels:        []string{k},
			})
		} else {
			m.Rels = append(m.Rels, Relationship{
				Count:      kv["count"].(int64),
				Type:       kv["type"].(string),
				Properties: nil,
			})
		}
	}

	for k := range allProps {
		m.Props = append(m.Props, k)
	}
	sort.Strings(m.Props)

	return m
}

func fallback(c Conn, m Metadata) Metadata {
	res, err := c.Session().Run("CALL db.labels() YIELD labels RETURN labels ORDER BY labels", nil)
	if err != nil {
		log.Println("Cannot retrieve metadata.")
		return m
	}
	for res.Next() {
		l := res.Record().Values[0].(string)
		m.Nodes = append(m.Nodes, Node{Labels: []string{l}})
	}

	res, _ = c.Session().Run("CALL db.relationshipTypes() YIELD relationshipType RETURN relationshipType ORDER BY relationshipType", nil)
	for res.Next() {
		l := res.Record().Values[0].(string)
		m.Nodes = append(m.Nodes, Node{Labels: []string{l}})
	}

	res, _ = c.Session().Run("CALL db.propertyKeys() YIELD propertyKey RETURN propertyKey ORDER BY propertyKey", nil)
	for res.Next() {
		l := res.Record().Values[0].(string)
		m.Props = append(m.Props, l)
	}

	return m
}

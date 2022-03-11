package repl

import (
	"github.com/abc-inc/merovingian/comp"
	"github.com/abc-inc/merovingian/cypher"
	"github.com/abc-inc/merovingian/db/neo4j"
)

func Foo(q string) (is []comp.Item) {
	e := cypher.NewEditorSupport(q)
	e.SetSchema(neo4j.Schema{
		Labels:   []string{"Node", "Other"},
		RelTypes: []string{"RELATED_TO"},
		PropKeys: nil,
		Funcs: []neo4j.Func{{
			Name:     "count",
			Sig:      "expression :: any",
			RetItems: nil,
		}},
		Procs: []neo4j.Func{{
			Name:     "count",
			Sig:      "expression :: any",
			RetItems: nil,
		}},
		ConCmds: nil,
		Params:  nil,
	})
	e.Update(q)

	//fmt.Println(e.GetCompletion(0, 0, true).Items)
	//fmt.Println(e.GetCompletion(1, strings.Index(q, " ")+1, true).Items)
	//fmt.Println(e.GetCompletion(1, strings.Index(q, "TCH"), true).Items)

	//fmt.Println(e.GetCompletion(1, strings.Index(q, "n:"), true).Items)
	//fmt.Println(e.GetCompletion(1, strings.Index(q, "Node"), true).Items)
	//fmt.Println(e.GetCompletion(1, strings.Index(q, "n,"), true).Items)
	//fmt.Println(e.GetCompletion(1, strings.Index(q, "cou"), true).Items)
	//fmt.Println(e.GetCompletion(1, strings.Index(q, "(o) AS cnt"), true).Items)
	//fmt.Println(e.GetCompletion(1, strings.Index(q, "n."), true).Items)
	//fmt.Println(e.GetCompletion(1, strings.Index(q, "n.")+2, true).Items)
	//fmt.Println(e.GetCompletion(1, strings.LastIndex(q, "cnt"), true).Items)
	pc := cypher.NewPosConv(q)
	x, y := pc.ToRelative(len(q))
	its := e.GetCompletion(x, y, true)
	return its.Items
}

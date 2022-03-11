package comp

import (
	"fmt"

	"github.com/abc-inc/merovingian/db/neo4j"
	"github.com/abc-inc/merovingian/lang"
	"github.com/abc-inc/merovingian/types"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type ItemProvider func(schema neo4j.Schema, typeData TypeData) []Item

var providers = make(map[types.Type]ItemProvider)

func init() {
	providers[types.ProcedureOutput] = func(schema neo4j.Schema, typeData TypeData) []Item {
		for _, p := range schema.Procs {
			if p.Name == typeData.Path[0] && len(p.RetItems) > 0 {
				its := make([]Item, len(p.RetItems))
				for i, ri := range p.RetItems {
					its[i] = Item{
						Type:    types.ProcedureOutput,
						View:    ri.Name,
						Content: ri.Sig,
						Postfix: " :: " + ri.Sig,
					}
				}
				return its
			}
		}
		return nil
	}
	providers[types.ConsoleCommandSubCommand] = func(schema neo4j.Schema, typeData TypeData) []Item {
		filterLastElement, path := typeData.FilterLastElement, typeData.Path
		length := len(path)
		if filterLastElement {
			length--
		}
		currentLevel := schema.ConCmds
		for i := 0; i < length; i++ {
			ok := false
			for _, foo := range currentLevel {
				if foo.Name == path[i] {
					currentLevel = foo.SubCmds
					ok = true
					break
				}
			}
			if !ok {
				return nil
			}
		}

		its := make([]Item, len(currentLevel))
		for i, cmd := range currentLevel {
			its[i] = Item{
				Type:    types.ConsoleCommandSubCommand,
				View:    cmd.Name,
				Content: cmd.Name,
				Postfix: cmd.Desc,
			}
		}
		return its
	}
}

type SchemaBased struct {
	Schema neo4j.Schema
	Cache
}

func NewSchemaBased(schema neo4j.Schema) *SchemaBased {
	s := &SchemaBased{Schema: schema}
	s.Map = make(map[types.Type][]Item)
	s.Map[types.Keyword] = KEYWORD_ITEMS
	s.Map[types.Label] = mapItems(schema.Labels, types.Label, strF, escF, nilF)
	s.Map[types.RelationshipType] = mapItems(schema.RelTypes, types.RelationshipType, strF, escF, nilF)
	s.Map[types.PropertyKey] = mapItems(schema.PropKeys, types.PropertyKey, strF, escF, nilF)
	s.Map[types.FunctionName] = mapItemsStruct(schema.Funcs, types.FunctionName,
		func(n neo4j.Func) string { return n.Name }, func(n neo4j.Func) string { return lang.EscapeCypher(n.Name) }, func(n neo4j.Func) string { return n.Sig })
	s.Map[types.ProcedureName] = mapItemsStruct(schema.Procs, types.ProcedureName,
		func(n neo4j.Func) string { return n.Name }, func(n neo4j.Func) string { return n.Name }, func(n neo4j.Func) string { return n.Sig })
	s.Map[types.ConsoleCommandName] = mapItemsCmd(schema.ConCmds, types.ConsoleCommandName,
		func(n neo4j.Cmd) string { return n.Name }, func(n neo4j.Cmd) string { return n.Name }, func(n neo4j.Cmd) string { return n.Desc })
	s.Map[types.Parameter] = mapItems(schema.Params, types.Parameter, strF, strF, nilF)
	return s
}

func (s SchemaBased) CalculateItems(t TypeData, query antlr.Tree) (is []Item) {
	if p, ok := providers[t.Type]; ok {
		is = append(is, p(s.Schema, t)...)
	}
	return is
}

func (s SchemaBased) Complete(ts []TypeData, query antlr.Tree) (is []Item) {
	if len(ts) == 0 {
		return nil
	}

	for _, t := range ts {
		if items := s.Map[t.Type]; items != nil {
			is = append(is, items...)
		} else {
			is = append(is, s.CalculateItems(t, query)...)
		}
	}

	return
}

func strF(s interface{}) string {
	if str, ok := s.(string); ok {
		return str
	} else if str, ok := s.(fmt.Stringer); ok {
		return str.String()
	} else {
		panic(s)
	}
}

func escF(i interface{}) string { return lang.EscapeCypher(i.(string)) }

func nilF(_ interface{}) string {
	return ""
}

func mapItems(ns []string, typ types.Type,
	viewFunc func(interface{}) string,
	contFunc func(interface{}) string,
	pfFunc func(interface{}) string) (its []Item) {

	for _, n := range ns {
		its = append(its, Item{
			Type:    typ,
			View:    viewFunc(n),
			Content: contFunc(n),
			Postfix: pfFunc(n),
		})
	}
	return
}

func mapItemsStruct(ns []neo4j.Func, typ types.Type,
	viewFunc func(neo4j.Func) string,
	contFunc func(neo4j.Func) string,
	pfFunc func(neo4j.Func) string) (its []Item) {

	for _, n := range ns {
		its = append(its, Item{
			Type:    typ,
			View:    viewFunc(n),
			Content: contFunc(n),
			Postfix: pfFunc(n),
		})
	}
	return
}

func mapItemsCmd(ns []neo4j.Cmd, typ types.Type,
	viewFunc func(neo4j.Cmd) string,
	contFunc func(neo4j.Cmd) string,
	pfFunc func(cmd neo4j.Cmd) string) (its []Item) {

	for _, n := range ns {
		its = append(its, Item{
			Type:    typ,
			View:    viewFunc(n),
			Content: contFunc(n),
			Postfix: pfFunc(n),
		})
	}
	return
}

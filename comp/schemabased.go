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
			if p.Name == typeData.Name && len(p.RetItems) > 0 {
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
			for _, foo := range currentLevel {
				if foo.Name == path[i] {
					currentLevel = foo.SubCmds
				} else {
					return nil
				}
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
	s.Map[types.FunctionName] = mapItemsStruct(schema.Funcs, types.FunctionName, strF, strF, nilF)
	//s.Map[types.ProcedureName] =
	//s.Map[types.ConsoleCommandName] =
	s.Map[types.Parameter] = mapItems(schema.Params, types.Parameter, strF, strF, nilF)
	return s
}

func (s SchemaBased) CalculateItems(ts ComplInfo, query antlr.Tree) (is []Item) {
	typeData := TypeData{
		Name:              "",
		Path:              nil,
		FilterLastElement: false,
	}
	for _, t := range ts.Types {
		if p, ok := providers[t]; ok {
			is = append(is, p(s.Schema, typeData)...)
		}
	}
	return is
}

func (s SchemaBased) Complete(ts []types.Type, query antlr.Tree) (is []Item) {
	if len(ts) == 0 {
		return nil
	}

	for _, t := range ts {
		if items := s.Map[t]; items != nil {
			is = append(is, items...)
		}
	}
	is = append(is, s.CalculateItems(ComplInfo{
		Element: query,
		Query:   query,
		Found:   false,
		Types:   ts,
	}, query)...)

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

func nilF(i interface{}) interface{} {
	return nil
}

func mapItems(ns []string, typ types.Type,
	viewFunc func(interface{}) string,
	contFunc func(interface{}) string,
	pfFunc func(interface{}) interface{}) (its []Item) {

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

func mapItemsStruct(ns []interface{}, typ types.Type,
	viewFunc func(interface{}) string,
	contFunc func(interface{}) string,
	pfFunc func(interface{}) interface{}) (its []Item) {

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

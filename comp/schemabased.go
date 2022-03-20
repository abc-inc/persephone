package comp

import (
	"fmt"

	"github.com/abc-inc/persephone/lang"
	"github.com/abc-inc/persephone/ndb"
	"github.com/abc-inc/persephone/types"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type ItemProvider func(schema ndb.Schema, typeData types.Data) []Item

var providers = make(map[types.Type]ItemProvider)

func init() {
	providers[types.ProcedureOutput] = func(schema ndb.Schema, typeData types.Data) []Item {
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
	providers[types.ConsoleCommandSubCommand] = func(schema ndb.Schema, typeData types.Data) []Item {
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
	Schema ndb.Schema
	cache  map[types.Type][]Item
}

func NewSchemaBased(schema ndb.Schema) *SchemaBased {
	s := &SchemaBased{Schema: schema}
	s.cache = make(map[types.Type][]Item)
	s.cache[types.Keyword] = KeywordItems
	s.cache[types.Label] = mapItems(schema.Labels, types.Label, strF, escF, nilF)
	s.cache[types.RelationshipType] = mapItems(schema.RelTypes, types.RelationshipType, strF, escF, nilF)
	s.cache[types.PropertyKey] = mapItems(schema.PropKeys, types.PropertyKey, strF, escF, nilF)
	s.cache[types.FunctionName] = mapItemsStruct(schema.Funcs, types.FunctionName,
		func(n ndb.Func) string { return n.Name }, func(n ndb.Func) string { return lang.EscapeCypher(n.Name) }, func(n ndb.Func) string { return n.Sig })
	s.cache[types.ProcedureName] = mapItemsStruct(schema.Procs, types.ProcedureName,
		func(n ndb.Func) string { return n.Name }, func(n ndb.Func) string { return n.Name }, func(n ndb.Func) string { return n.Sig })
	s.cache[types.ConsoleCommandName] = mapItemsCmd(schema.ConCmds, types.ConsoleCommandName,
		func(n ndb.Cmd) string { return n.Name }, func(n ndb.Cmd) string { return n.Name }, func(n ndb.Cmd) string { return n.Desc })
	s.cache[types.Parameter] = mapItems(schema.Params, types.Parameter, strF, strF, nilF)
	return s
}

func (s SchemaBased) CalculateItems(t types.Data, query antlr.Tree) (is []Item) {
	if p, ok := providers[t.Type]; ok {
		is = append(is, p(s.Schema, t)...)
	}
	return is
}

func (s SchemaBased) Complete(ts []types.Data, query antlr.Tree) (is []Item) {
	if len(ts) == 0 {
		return nil
	}

	for _, t := range ts {
		if items := s.cache[t.Type]; items != nil {
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

func mapItemsStruct(ns []ndb.Func, typ types.Type,
	viewFunc func(ndb.Func) string,
	contFunc func(ndb.Func) string,
	pfFunc func(ndb.Func) string) (its []Item) {

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

func mapItemsCmd(ns []ndb.Cmd, typ types.Type,
	viewFunc func(ndb.Cmd) string,
	contFunc func(ndb.Cmd) string,
	pfFunc func(cmd ndb.Cmd) string) (its []Item) {

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

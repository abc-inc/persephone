// Copyright 2022 The Persephone authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package comp

import (
	"fmt"

	"github.com/abc-inc/persephone/lang"
	"github.com/abc-inc/persephone/types"
	"github.com/abc-inc/roland/meta"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Cmd represents a console command.
type Cmd struct {
	Name    string
	Desc    string
	SubCmds []Cmd
}

// String returns command name.
func (c Cmd) String() string {
	return c.Name
}

// Metadata represents holds schema information and other completion metadata.
type Metadata struct {
	meta.Schema
	Params  []string
	ConCmds []Cmd
}

type ItemProvider func(schema Metadata, typeData types.Data) []Item

var providers = make(map[types.Type]ItemProvider)

func init() {
	providers[types.ProcedureOutput] = func(schema Metadata, typeData types.Data) []Item {
		for _, p := range schema.Procs {
			if p.Name == typeData.Path[0] && len(p.RetItems) > 0 {
				its := make([]Item, len(p.RetItems))
				for i, ri := range p.RetItems {
					its[i] = Item{
						Type:    types.ProcedureOutput,
						View:    ri.Name,
						Content: ri.Name,
						Postfix: " :: " + ri.Sig,
					}
				}
				return its
			}
		}
		return nil
	}
	providers[types.ConsoleCommandSubCommand] = func(schema Metadata, typeData types.Data) []Item {
		filterLastElement, path := typeData.FilterLastElement, typeData.Path
		length := len(path)
		if filterLastElement {
			length--
		}
		currLvlCmds := schema.ConCmds
		for i := 0; i < length; i++ {
			ok := false
			for _, currLvlCmd := range currLvlCmds {
				if currLvlCmd.Name == path[i] {
					currLvlCmds = currLvlCmd.SubCmds
					ok = true
					break
				}
			}
			if !ok {
				return nil
			}
		}

		its := make([]Item, len(currLvlCmds))
		for i, cmd := range currLvlCmds {
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
	Schema Metadata
	cache  map[types.Type][]Item
}

var _ Comp = (*SchemaBased)(nil)

func NewSchemaBased(schema Metadata) *SchemaBased {
	s := &SchemaBased{Schema: schema}
	s.cache = make(map[types.Type][]Item)
	s.cache[types.Keyword] = KeywordItems()
	s.cache[types.Label] = mapItems(schema.Labels, types.Label, strF, escF, nilF)
	s.cache[types.RelationshipType] = mapItems(schema.RelTypes, types.RelationshipType, strF, escF, nilF)
	s.cache[types.PropertyKey] = mapItems(schema.PropKeys, types.PropertyKey, strF, escF, nilF)
	s.cache[types.FunctionName] = mapItemsStruct(schema.Funcs, types.FunctionName,
		func(n meta.Func) string { return n.Name }, func(n meta.Func) string { return lang.EscapeCypher(n.Name) }, func(n meta.Func) string { return n.Sig })
	s.cache[types.ProcedureName] = mapItemsStruct(schema.Procs, types.ProcedureName,
		func(n meta.Func) string { return n.Name }, func(n meta.Func) string { return n.Name }, func(n meta.Func) string { return n.Sig })
	s.cache[types.ConsoleCommandName] = mapItemsCmd(schema.ConCmds, types.ConsoleCommandName,
		func(n Cmd) string { return n.Name }, func(n Cmd) string { return n.Name }, func(n Cmd) string { return n.Desc })
	s.cache[types.Parameter] = mapItems(schema.Params, types.Parameter, strF, strF, nilF)
	return s
}

func (s SchemaBased) CalculateItems(t types.Data, _ antlr.Tree) (its []Item) {
	if p, ok := providers[t.Type]; ok {
		its = append(its, p(s.Schema, t)...)
	}
	return its
}

func (s SchemaBased) Complete(ts []types.Data, query antlr.Tree) (its []Item) {
	if len(ts) == 0 {
		return nil
	}

	for _, t := range ts {
		if items := s.cache[t.Type]; items != nil {
			its = append(its, items...)
		} else {
			its = append(its, s.CalculateItems(t, query)...)
		}
	}

	return
}

func strF(s any) string {
	if str, ok := s.(string); ok {
		return str
	} else if str, ok := s.(fmt.Stringer); ok {
		return str.String()
	} else {
		panic(s)
	}
}

func escF(i any) string { return lang.EscapeCypher(i.(string)) }

func nilF(_ any) string {
	return ""
}

func mapItems(ns []string, typ types.Type,
	viewFunc func(any) string,
	contFunc func(any) string,
	pfFunc func(any) string) (its []Item) {

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

func mapItemsStruct(ns []meta.Func, typ types.Type,
	viewFunc func(meta.Func) string,
	contFunc func(meta.Func) string,
	pfFunc func(meta.Func) string) (its []Item) {

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

func mapItemsCmd(ns []Cmd, typ types.Type,
	viewFunc func(Cmd) string,
	contFunc func(Cmd) string,
	pfFunc func(Cmd) string) (its []Item) {

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

// Copyright 2022 The persephone authors
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

package ast

import (
	"reflect"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Pos provides position of Tokens in a ParseTree.
type Pos interface {
	// GetStart returns the position where the Token begins.
	GetStart() int
	// GetStop returns the position where the Token ends.
	GetStop() int
}

// posStruct holds a position range.
type posStruct struct {
	Start, Stop int
}

// GetStart returns the position where the Token begins.
func (p posStruct) GetStart() int {
	return p.Start
}

// GetStop returns the position where the Token ends.
func (p posStruct) GetStop() int {
	return p.Stop
}

// FindParent returns the closest parent RuleContext matching a certain type.
// If pt is of type t, it is returned directly.
func FindParent(pt antlr.Tree, t reflect.Type) antlr.Tree {
	if pt == nil || reflect.TypeOf(pt).Elem() == t {
		return pt
	}
	e := pt.GetParent()
	for e != nil {
		if reflect.TypeOf(e).Elem() == t {
			break
		}
		for _, c := range e.GetChildren() {
			if reflect.TypeOf(c).Elem() == t {
				return c
			}
		}
		e = e.GetParent()
	}
	return e
}

// FindAnyParent returns the closest parent RuleContext matching any of the
// given types.
func FindAnyParent(pt antlr.Tree, types []string) antlr.Tree {
	el := pt
	for el != nil {
		for _, t := range types {
			if t == reflect.TypeOf(el).Elem().Name() {
				return el
			}
		}
		el = GetParent(el)
	}
	return el
}

// FindChild performs a depth-first search traversal and returns the first child
// of a certain type.
func FindChild(element antlr.Tree, typ string) antlr.Tree {
	if element == nil {
		return nil
	}
	if reflect.TypeOf(element).Elem().Name() == typ {
		return element
	}
	for _, c := range element.GetChildren() {
		if result := FindChild(c, typ); result != nil {
			return result
		}
	}
	return nil
}

// GetPosition returns the position of the given Token or Symbol.
func GetPosition(el antlr.Tree) Pos {
	if el == nil {
		return nil
	}

	var start, stop, symbol antlr.Token
	if o, ok := el.(interface{ GetStart() antlr.Token }); ok {
		start = o.GetStart()
	}
	if o, ok := el.(interface{ GetStop() antlr.Token }); ok {
		stop = o.GetStop()
	}
	if o, ok := el.(interface{ GetSymbol() antlr.Token }); ok {
		symbol = o.GetSymbol()
	}

	if symbol != nil {
		return posStruct{symbol.GetStart(), symbol.GetStop()}
	} else if start != nil && stop != nil {
		return posStruct{start.GetStart(), stop.GetStop()}
	}
	return nil
}

// HasErrorNode checks whether the given ParseTree contains an ErrorNode.
func HasErrorNode(element antlr.Tree) bool {
	if element == nil {
		return false
	}
	if _, ok := element.(antlr.ErrorNode); ok {
		return true
	}
	for _, c := range element.GetChildren() {
		if HasErrorNode(c) {
			return true
		}
	}
	return false
}

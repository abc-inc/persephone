package ast

import (
	"reflect"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type Pos interface {
	GetStart() int
	GetStop() int
}

type posStruct struct {
	Start, Stop int
}

func (p posStruct) GetStart() int {
	return p.Start
}

func (p posStruct) GetStop() int {
	return p.Stop
}

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

func FindAnyParent(pt antlr.Tree, types []string) antlr.Tree {
	el := pt
	for el != nil {
		for _, t := range types {
			if t == reflect.TypeOf(el).Elem().Name() {
				return el
			}
			//for _, c := range el.GetChildren() {
			//	if t == reflect.TypeOf(c).Elem().Name() {
			//		return c
			//	}
			//}
		}
		el = GetParent(el)
	}
	return el
}

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

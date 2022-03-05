package ast

import (
	"reflect"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type Pos interface {
	GetStart() int
	GetStop() int
}

type PosStruct struct {
	Start, Stop int
}

func (p PosStruct) GetStart() int {
	return p.Start
}

func (p PosStruct) GetStop() int {
	return p.Stop
}

func FindParent(pt antlr.Tree, t reflect.Type) antlr.Tree {
	e := pt
	for e != nil {
		if reflect.TypeOf(e) == t {
			break
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
		}
		el = el.GetParent()
	}
	return el
}

func FindChild(element antlr.Tree, typ string) antlr.Tree {
	if element == nil {
		return nil
	}
	if reflect.TypeOf(element).Name() == typ {
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
		return PosStruct{symbol.GetStart(), symbol.GetStop()}
	} else if start != nil && stop != nil {
		return PosStruct{start.GetStart(), stop.GetStop()}
	}
	return nil
}

func HasErrorNode(element antlr.Tree) bool {
	if element == nil {
		return false
	}
	// TODO: implement
	return false
}

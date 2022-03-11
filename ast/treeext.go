package ast

import (
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

func GetParent(e antlr.Tree) antlr.Tree {
	//fmt.Println(reflect.TypeOf(e).Elem(), reflect.TypeOf(e.GetParent()).Elem())
	e = e.GetParent()
	if _, ok := e.(*antlr.BaseParserRuleContext); ok && e.GetParent() != nil {
		if e.GetParent().GetChildCount() == 0 {
			fmt.Println(e.GetParent())
		}
		e = e.GetParent().GetChild(0)
		//fmt.Println(reflect.TypeOf(e).Elem(), reflect.TypeOf(p).Elem())
	}
	return e
}
package ast

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

func GetParent(e antlr.Tree) antlr.Tree {
	pt := e.GetParent()
	if _, ok := pt.(*antlr.BaseParserRuleContext); ok && pt.GetParent() != nil {
		if pt.GetParent().GetChildCount() == 1 {
			return pt.GetParent().GetChild(0)
		}
		eStart := pt.(*antlr.BaseParserRuleContext).GetStart()
		eStop := pt.(*antlr.BaseParserRuleContext).GetStop()
		for _, c := range pt.GetParent().GetChildren() {
			if _, ok := c.(antlr.ParserRuleContext); ok {
				cStart := c.(antlr.ParserRuleContext).GetStart()
				cStop := c.(antlr.ParserRuleContext).GetStop()
				if cStart.GetStart() >= eStart.GetStart() && cStop.GetStop() <= eStop.GetStop() {
					return c
				}
			}
		}
		pt = pt.GetParent().GetChild(0)
	}
	return pt
}

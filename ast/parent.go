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
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// GetParent returns the parent RuleContext in the ParseTree.
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
				if cStop == nil {
					// In rare cases e.g., if there's not even a single valid keyword,
					// there is no stop token yet.
					cStop = cStart
				}
				if cStart.GetStart() >= eStart.GetStart() && cStop.GetStop() <= eStop.GetStop() {
					return c
				}
			}
		}
		pt = pt.GetParent().GetChild(0)
	}
	return pt
}

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

package rule

import (
	"reflect"

	"github.com/abc-inc/persephone/ast"
	"github.com/abc-inc/persephone/lang"
	"github.com/abc-inc/persephone/types"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

var childToParent = map[string]types.Type{
	lang.VariableContext:                      types.Variable,
	lang.ParameterNameContext:                 types.Parameter,
	lang.PropertyKeyNameContext:               types.PropertyKey,
	lang.FunctionNameContext:                  types.FunctionName,
	lang.ProcedureNameContext:                 types.ProcedureName,
	lang.NodeLabelContext:                     types.Label,
	lang.RelationshipTypeContext:              types.RelationshipType,
	lang.RelationshipTypeOptionalColonContext: types.RelationshipType,
	lang.ConsoleCommandNameContext:            types.ConsoleCommandName,
	lang.NodeLabelsContext:                    types.Label,
	lang.RelationshipTypesContext:             types.RelationshipType,
}

// ruleSpecificParent checks that the ParseTree element is inside a specific
// parent context.
func ruleSpecificParent(e antlr.ParseTree) []Info {
	ctxNames := make([]string, len(childToParent))
	for name := range childToParent {
		ctxNames = append(ctxNames, name)
	}

	parent := ast.FindAnyParent(e, ctxNames)
	if parent != nil {
		if t, ok := childToParent[reflect.TypeOf(parent).Elem().Name()]; !ok {
			panic(t)
		} else {
			return []Info{{Type: t}}
		}
	}
	return nil
}

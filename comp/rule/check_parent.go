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

package rule

import (
	"reflect"

	"github.com/abc-inc/persephone/ast"
	"github.com/abc-inc/persephone/lang"
	"github.com/abc-inc/persephone/types"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

var childToParent = map[string]types.Type{
	lang.VARIABLE_CONTEXT:                         types.Variable,
	lang.PARAMETER_NAME_CONTEXT:                   types.Parameter,
	lang.PROPERTY_KEY_NAME_CONTEXT:                types.PropertyKey,
	lang.FUNCTION_NAME_CONTEXT:                    types.FunctionName,
	lang.PROCEDURE_NAME_CONTEXT:                   types.ProcedureName,
	lang.NODE_LABEL_CONTEXT:                       types.Label,
	lang.RELATIONSHIP_TYPE_CONTEXT:                types.RelationshipType,
	lang.RELATIONSHIP_TYPE_OPTIONAL_COLON_CONTEXT: types.RelationshipType,
	lang.CONSOLE_COMMAND_NAME_CONTEXT:             types.ConsoleCommandName,
	lang.NODE_LABELS_CONTEXT:                      types.Label,
	lang.RELATIONSHIP_TYPES_CONTEXT:               types.RelationshipType,
}

func ruleSpecificParent(e antlr.ParseTree) []Info {
	var ctxNames []string
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

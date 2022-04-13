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

package lang

import (
	"reflect"

	"github.com/abc-inc/persephone/parser"
)

var VARIABLE_CONTEXT = reflect.TypeOf(parser.VariableContext{}).Name()
var LABEL_NAME_CONTEXT = reflect.TypeOf(parser.LabelNameContext{}).Name()
var RELATIONSHIP_TYPE_NAME_CONTEXT = reflect.TypeOf(parser.RelTypeNameContext{}).Name()
var PROPERTY_KEY_NAME_CONTEXT = reflect.TypeOf(parser.PropertyKeyNameContext{}).Name()
var PARAMETER_NAME_CONTEXT = reflect.TypeOf(parser.ParameterNameContext{}).Name()
var PARAMETER_CONTEXT = reflect.TypeOf(parser.ParameterContext{}).Name()
var FUNCTION_NAME_CONTEXT = reflect.TypeOf(parser.FunctionInvocationBodyContext{}).Name()
var PROCEDURE_NAME_CONTEXT = reflect.TypeOf(parser.ProcedureInvocationBodyContext{}).Name()
var CONSOLE_COMMAND_NAME_CONTEXT = reflect.TypeOf(parser.CypherConsoleCommandNameContext{}).Name()
var CONSOLE_COMMAND_CONTEXT = reflect.TypeOf(parser.CypherConsoleCommandContext{}).Name()
var CONSOLE_COMMAND_PARAMETERS_CONTEXT = reflect.TypeOf(parser.CypherConsoleCommandParametersContext{}).Name()
var CONSOLE_COMMAND_PARAMETER_CONTEXT = reflect.TypeOf(parser.CypherConsoleCommandParameterContext{}).Name()
var CONSOLE_COMMAND_SUBCOMMAND_CONTEXT = reflect.TypeOf(parser.SubCommandContext{}).Name()
var CONSOLE_COMMAND_PATH_CONTEXT = reflect.TypeOf(parser.CommandPathContext{}).Name()
var PROCEDURE_OUTPUT_CONTEXT = reflect.TypeOf(parser.ProcedureOutputContext{}).Name()
var PROCEDURE_RESULTS_CONTEXT = reflect.TypeOf(parser.ProcedureResultsContext{}).Name()

var ALL_FUNCTION_NAME_CONTEXT = reflect.TypeOf(parser.AllFunctionNameContext{}).Name()
var ANY_FUNCTION_NAME_CONTEXT = reflect.TypeOf(parser.AnyFunctionNameContext{}).Name()
var SINGLE_FUNCTION_NAME_CONTEXT = reflect.TypeOf(parser.SingleFunctionNameContext{}).Name()
var NONE_FUNCTION_NAME_CONTEXT = reflect.TypeOf(parser.NoneFunctionNameContext{}).Name()
var EXTRACT_FUNCTION_NAME_CONTEXT = reflect.TypeOf(parser.ExtractFunctionNameContext{}).Name()
var REDUCE_FUNCTION_NAME_CONTEXT = reflect.TypeOf(parser.ReduceFunctionNameContext{}).Name()
var SHORTEST_PATH_FUNCTION_NAME_CONTEXT = reflect.TypeOf(parser.ShortestPathFunctionNameContext{}).Name()
var ALL_SHORTEST_PATH_FUNCTION_NAME_CONTEXT = reflect.TypeOf(parser.AllShortestPathFunctionNameContext{}).Name()
var FILTER_FUNCTION_NAME_CONTEXT = reflect.TypeOf(parser.FilterFunctionNameContext{}).Name()
var EXISTS_FUNCTION_NAME_CONTEXT = reflect.TypeOf(parser.ExistsFunctionNameContext{}).Name()

var CALL_CONTEXT = reflect.TypeOf(parser.CallContext{}).Name()
var EXPRESSION_CONTEXT = reflect.TypeOf(parser.ExpressionContext{}).Name()
var PATTERN_ELEMENT_CONTEXT = reflect.TypeOf(parser.PatternElementContext{}).Name()
var NODE_PATTERN_CONTEXT = reflect.TypeOf(parser.NodePatternContext{}).Name()
var NODE_LABEL_CONTEXT = reflect.TypeOf(parser.NodeLabelContext{}).Name()
var NODE_LABELS_CONTEXT = reflect.TypeOf(parser.NodeLabelsContext{}).Name()
var RELATIONSHIP_TYPE_CONTEXT = reflect.TypeOf(parser.RelationshipTypeContext{}).Name()
var RELATIONSHIP_TYPE_OPTIONAL_COLON_CONTEXT = reflect.TypeOf(parser.RelationshipTypeOptionalColonContext{}).Name()
var RELATIONSHIP_TYPES_CONTEXT = reflect.TypeOf(parser.RelationshipTypesContext{}).Name()
var RELATIONSHIP_PATTERN_CONTEXT = reflect.TypeOf(parser.RelationshipPatternContext{}).Name()
var PROPERTY_LOOKUP_CONTEXT = reflect.TypeOf(parser.PropertyLookupContext{}).Name()
var MAP_LITERAL_CONTEXT = reflect.TypeOf(parser.MapLiteralContext{}).Name()
var PROPERTIES_CONTEXT = reflect.TypeOf(parser.PropertiesContext{}).Name()
var MAP_LITERAL_ENTRY = reflect.TypeOf(parser.LiteralEntryContext{}).Name()
var STRING_LITERAL_CONTEXT = reflect.TypeOf(parser.StringLiteralContext{}).Name()
var ATOM_CONTEXT = reflect.TypeOf(parser.AtomContext{}).Name()

var QUERY_CONTEXT = reflect.TypeOf(parser.CypherQueryContext{}).Name()
var SYMBOLIC_NAME_CONTEXT = reflect.TypeOf(parser.SymbolicNameContext{}).Name()

var CompletionCandidates = []string{
	STRING_LITERAL_CONTEXT,
	VARIABLE_CONTEXT,
	PROCEDURE_NAME_CONTEXT,
	FUNCTION_NAME_CONTEXT,
	CONSOLE_COMMAND_NAME_CONTEXT,
	NODE_LABEL_CONTEXT,
	RELATIONSHIP_TYPE_CONTEXT,
	RELATIONSHIP_TYPE_OPTIONAL_COLON_CONTEXT,
}

var SymbolicContexts = []string{
	VARIABLE_CONTEXT,
	LABEL_NAME_CONTEXT,
	RELATIONSHIP_TYPE_NAME_CONTEXT,
	PROPERTY_KEY_NAME_CONTEXT,
	PARAMETER_NAME_CONTEXT,
}

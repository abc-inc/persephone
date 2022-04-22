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

	"github.com/gschauer/cypher2go/v4/parser"
)

var VariableContext = reflect.TypeOf(parser.VariableContext{}).Name()
var LabelNameContext = reflect.TypeOf(parser.LabelNameContext{}).Name()
var RelationshipTypeNameContext = reflect.TypeOf(parser.RelTypeNameContext{}).Name()
var PropertyKeyNameContext = reflect.TypeOf(parser.PropertyKeyNameContext{}).Name()
var ParameterNameContext = reflect.TypeOf(parser.ParameterNameContext{}).Name()
var ParameterContext = reflect.TypeOf(parser.ParameterContext{}).Name()
var FunctionNameContext = reflect.TypeOf(parser.FunctionInvocationBodyContext{}).Name()
var ProcedureNameContext = reflect.TypeOf(parser.ProcedureInvocationBodyContext{}).Name()
var ConsoleCommandNameContext = reflect.TypeOf(parser.CypherConsoleCommandNameContext{}).Name()
var ConsoleCommandContext = reflect.TypeOf(parser.CypherConsoleCommandContext{}).Name()
var ConsoleCommandParametersContext = reflect.TypeOf(parser.CypherConsoleCommandParametersContext{}).Name()
var ConsoleCommandParameterContext = reflect.TypeOf(parser.CypherConsoleCommandParameterContext{}).Name()
var ConsoleCommandSubcommandContext = reflect.TypeOf(parser.SubCommandContext{}).Name()
var ConsoleCommandPathContext = reflect.TypeOf(parser.CommandPathContext{}).Name()
var ProcedureOutputContext = reflect.TypeOf(parser.ProcedureOutputContext{}).Name()
var ProcedureResultsContext = reflect.TypeOf(parser.ProcedureResultsContext{}).Name()

var AllFunctionNameContext = reflect.TypeOf(parser.AllFunctionNameContext{}).Name()
var AnyFunctionNameContext = reflect.TypeOf(parser.AnyFunctionNameContext{}).Name()
var SingleFunctionNameContext = reflect.TypeOf(parser.SingleFunctionNameContext{}).Name()
var NoneFunctionNameContext = reflect.TypeOf(parser.NoneFunctionNameContext{}).Name()
var ExtractFunctionNameContext = reflect.TypeOf(parser.ExtractFunctionNameContext{}).Name()
var ReduceFunctionNameContext = reflect.TypeOf(parser.ReduceFunctionNameContext{}).Name()
var ShortestPathFunctionNameContext = reflect.TypeOf(parser.ShortestPathFunctionNameContext{}).Name()
var AllShortestPathFunctionNameContext = reflect.TypeOf(parser.AllShortestPathFunctionNameContext{}).Name()
var FilterFunctionNameContext = reflect.TypeOf(parser.FilterFunctionNameContext{}).Name()
var ExistsFunctionNameContext = reflect.TypeOf(parser.ExistsFunctionNameContext{}).Name()

var CallContext = reflect.TypeOf(parser.CallContext{}).Name()
var ExpressionContext = reflect.TypeOf(parser.ExpressionContext{}).Name()
var PatternElementContext = reflect.TypeOf(parser.PatternElementContext{}).Name()
var NodePatternContext = reflect.TypeOf(parser.NodePatternContext{}).Name()
var NodeLabelContext = reflect.TypeOf(parser.NodeLabelContext{}).Name()
var NodeLabelsContext = reflect.TypeOf(parser.NodeLabelsContext{}).Name()
var RelationshipTypeContext = reflect.TypeOf(parser.RelationshipTypeContext{}).Name()
var RelationshipTypeOptionalColonContext = reflect.TypeOf(parser.RelationshipTypeOptionalColonContext{}).Name()
var RelationshipTypesContext = reflect.TypeOf(parser.RelationshipTypesContext{}).Name()
var RelationshipPatternContext = reflect.TypeOf(parser.RelationshipPatternContext{}).Name()
var PropertyLookupContext = reflect.TypeOf(parser.PropertyLookupContext{}).Name()
var MapLiteralContext = reflect.TypeOf(parser.MapLiteralContext{}).Name()
var PropertiesContext = reflect.TypeOf(parser.PropertiesContext{}).Name()
var MapLiteralEntry = reflect.TypeOf(parser.LiteralEntryContext{}).Name()
var StringLiteralContext = reflect.TypeOf(parser.StringLiteralContext{}).Name()
var AtomContext = reflect.TypeOf(parser.AtomContext{}).Name()

var QueryContext = reflect.TypeOf(parser.CypherQueryContext{}).Name()
var SymbolicNameContext = reflect.TypeOf(parser.SymbolicNameContext{}).Name()

var CompletionCandidates = []string{
	StringLiteralContext,
	VariableContext,
	ProcedureNameContext,
	FunctionNameContext,
	ConsoleCommandNameContext,
	NodeLabelContext,
	RelationshipTypeContext,
	RelationshipTypeOptionalColonContext,
}

var SymbolicContexts = []string{
	VariableContext,
	LabelNameContext,
	RelationshipTypeNameContext,
	PropertyKeyNameContext,
	ParameterNameContext,
}

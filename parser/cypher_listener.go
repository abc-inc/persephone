// Generated from cypher-editor-support/src/_generated/Cypher.g4 by ANTLR 4.7.

package parser // Cypher

import "github.com/antlr/antlr4/runtime/Go/antlr"

// CypherListener is a complete listener for a parse tree produced by CypherParser.
type CypherListener interface {
	antlr.ParseTreeListener

	// EnterCypher is called when entering the cypher production.
	EnterCypher(c *CypherContext)

	// EnterCypherPart is called when entering the cypherPart production.
	EnterCypherPart(c *CypherPartContext)

	// EnterCypherConsoleCommand is called when entering the cypherConsoleCommand production.
	EnterCypherConsoleCommand(c *CypherConsoleCommandContext)

	// EnterCypherConsoleCommandName is called when entering the cypherConsoleCommandName production.
	EnterCypherConsoleCommandName(c *CypherConsoleCommandNameContext)

	// EnterCypherConsoleCommandParameters is called when entering the cypherConsoleCommandParameters production.
	EnterCypherConsoleCommandParameters(c *CypherConsoleCommandParametersContext)

	// EnterCypherConsoleCommandParameter is called when entering the cypherConsoleCommandParameter production.
	EnterCypherConsoleCommandParameter(c *CypherConsoleCommandParameterContext)

	// EnterArrowExpression is called when entering the arrowExpression production.
	EnterArrowExpression(c *ArrowExpressionContext)

	// EnterUrl is called when entering the url production.
	EnterUrl(c *UrlContext)

	// EnterUri is called when entering the uri production.
	EnterUri(c *UriContext)

	// EnterScheme is called when entering the scheme production.
	EnterScheme(c *SchemeContext)

	// EnterHost is called when entering the host production.
	EnterHost(c *HostContext)

	// EnterHostname is called when entering the hostname production.
	EnterHostname(c *HostnameContext)

	// EnterHostnumber is called when entering the hostnumber production.
	EnterHostnumber(c *HostnumberContext)

	// EnterPort is called when entering the port production.
	EnterPort(c *PortContext)

	// EnterPath is called when entering the path production.
	EnterPath(c *PathContext)

	// EnterUser is called when entering the user production.
	EnterUser(c *UserContext)

	// EnterLogin is called when entering the login production.
	EnterLogin(c *LoginContext)

	// EnterPassword is called when entering the password production.
	EnterPassword(c *PasswordContext)

	// EnterFrag is called when entering the frag production.
	EnterFrag(c *FragContext)

	// EnterUrlQuery is called when entering the urlQuery production.
	EnterUrlQuery(c *UrlQueryContext)

	// EnterSearch is called when entering the search production.
	EnterSearch(c *SearchContext)

	// EnterSearchparameter is called when entering the searchparameter production.
	EnterSearchparameter(c *SearchparameterContext)

	// EnterStr is called when entering the str production.
	EnterStr(c *StrContext)

	// EnterUrlDigits is called when entering the urlDigits production.
	EnterUrlDigits(c *UrlDigitsContext)

	// EnterJson is called when entering the json production.
	EnterJson(c *JsonContext)

	// EnterObj is called when entering the obj production.
	EnterObj(c *ObjContext)

	// EnterPair is called when entering the pair production.
	EnterPair(c *PairContext)

	// EnterArray is called when entering the array production.
	EnterArray(c *ArrayContext)

	// EnterValue is called when entering the value production.
	EnterValue(c *ValueContext)

	// EnterKeyValueLiteral is called when entering the keyValueLiteral production.
	EnterKeyValueLiteral(c *KeyValueLiteralContext)

	// EnterCommandPath is called when entering the commandPath production.
	EnterCommandPath(c *CommandPathContext)

	// EnterSubCommand is called when entering the subCommand production.
	EnterSubCommand(c *SubCommandContext)

	// EnterCypherQuery is called when entering the cypherQuery production.
	EnterCypherQuery(c *CypherQueryContext)

	// EnterQueryOptions is called when entering the queryOptions production.
	EnterQueryOptions(c *QueryOptionsContext)

	// EnterAnyCypherOption is called when entering the anyCypherOption production.
	EnterAnyCypherOption(c *AnyCypherOptionContext)

	// EnterCypherOption is called when entering the cypherOption production.
	EnterCypherOption(c *CypherOptionContext)

	// EnterVersionNumber is called when entering the versionNumber production.
	EnterVersionNumber(c *VersionNumberContext)

	// EnterExplain is called when entering the explain production.
	EnterExplain(c *ExplainContext)

	// EnterProfile is called when entering the profile production.
	EnterProfile(c *ProfileContext)

	// EnterConfigurationOption is called when entering the configurationOption production.
	EnterConfigurationOption(c *ConfigurationOptionContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterQuery is called when entering the query production.
	EnterQuery(c *QueryContext)

	// EnterRegularQuery is called when entering the regularQuery production.
	EnterRegularQuery(c *RegularQueryContext)

	// EnterBulkImportQuery is called when entering the bulkImportQuery production.
	EnterBulkImportQuery(c *BulkImportQueryContext)

	// EnterSingleQuery is called when entering the singleQuery production.
	EnterSingleQuery(c *SingleQueryContext)

	// EnterPeriodicCommitHint is called when entering the periodicCommitHint production.
	EnterPeriodicCommitHint(c *PeriodicCommitHintContext)

	// EnterLoadCSVQuery is called when entering the loadCSVQuery production.
	EnterLoadCSVQuery(c *LoadCSVQueryContext)

	// EnterUnion is called when entering the union production.
	EnterUnion(c *UnionContext)

	// EnterClause is called when entering the clause production.
	EnterClause(c *ClauseContext)

	// EnterCommand is called when entering the command production.
	EnterCommand(c *CommandContext)

	// EnterSystemCommand is called when entering the systemCommand production.
	EnterSystemCommand(c *SystemCommandContext)

	// EnterMultidatabaseCommand is called when entering the multidatabaseCommand production.
	EnterMultidatabaseCommand(c *MultidatabaseCommandContext)

	// EnterUserCommand is called when entering the userCommand production.
	EnterUserCommand(c *UserCommandContext)

	// EnterPrivilegeCommand is called when entering the privilegeCommand production.
	EnterPrivilegeCommand(c *PrivilegeCommandContext)

	// EnterShowRoles is called when entering the showRoles production.
	EnterShowRoles(c *ShowRolesContext)

	// EnterCreateRole is called when entering the createRole production.
	EnterCreateRole(c *CreateRoleContext)

	// EnterCopyRole is called when entering the copyRole production.
	EnterCopyRole(c *CopyRoleContext)

	// EnterDropRole is called when entering the dropRole production.
	EnterDropRole(c *DropRoleContext)

	// EnterShowUsers is called when entering the showUsers production.
	EnterShowUsers(c *ShowUsersContext)

	// EnterCreateUser is called when entering the createUser production.
	EnterCreateUser(c *CreateUserContext)

	// EnterDropUser is called when entering the dropUser production.
	EnterDropUser(c *DropUserContext)

	// EnterAlterUser is called when entering the alterUser production.
	EnterAlterUser(c *AlterUserContext)

	// EnterShowPrivileges is called when entering the showPrivileges production.
	EnterShowPrivileges(c *ShowPrivilegesContext)

	// EnterGrantPrivilege is called when entering the grantPrivilege production.
	EnterGrantPrivilege(c *GrantPrivilegeContext)

	// EnterDenyPrivilege is called when entering the denyPrivilege production.
	EnterDenyPrivilege(c *DenyPrivilegeContext)

	// EnterRevokePrivilege is called when entering the revokePrivilege production.
	EnterRevokePrivilege(c *RevokePrivilegeContext)

	// EnterRevokePart is called when entering the revokePart production.
	EnterRevokePart(c *RevokePartContext)

	// EnterDatabaseScope is called when entering the databaseScope production.
	EnterDatabaseScope(c *DatabaseScopeContext)

	// EnterGraphScope is called when entering the graphScope production.
	EnterGraphScope(c *GraphScopeContext)

	// EnterRoles is called when entering the roles production.
	EnterRoles(c *RolesContext)

	// EnterGrantableGraphPrivileges is called when entering the grantableGraphPrivileges production.
	EnterGrantableGraphPrivileges(c *GrantableGraphPrivilegesContext)

	// EnterRevokeableGraphPrivileges is called when entering the revokeableGraphPrivileges production.
	EnterRevokeableGraphPrivileges(c *RevokeableGraphPrivilegesContext)

	// EnterDatasbasePrivilege is called when entering the datasbasePrivilege production.
	EnterDatasbasePrivilege(c *DatasbasePrivilegeContext)

	// EnterDbmsPrivilege is called when entering the dbmsPrivilege production.
	EnterDbmsPrivilege(c *DbmsPrivilegeContext)

	// EnterElementScope is called when entering the elementScope production.
	EnterElementScope(c *ElementScopeContext)

	// EnterPropertiesList is called when entering the propertiesList production.
	EnterPropertiesList(c *PropertiesListContext)

	// EnterPropertyScope is called when entering the propertyScope production.
	EnterPropertyScope(c *PropertyScopeContext)

	// EnterShowDatabase is called when entering the showDatabase production.
	EnterShowDatabase(c *ShowDatabaseContext)

	// EnterCreateDatabase is called when entering the createDatabase production.
	EnterCreateDatabase(c *CreateDatabaseContext)

	// EnterDropDatabase is called when entering the dropDatabase production.
	EnterDropDatabase(c *DropDatabaseContext)

	// EnterStartDatabase is called when entering the startDatabase production.
	EnterStartDatabase(c *StartDatabaseContext)

	// EnterStopDatabase is called when entering the stopDatabase production.
	EnterStopDatabase(c *StopDatabaseContext)

	// EnterIfNotExists is called when entering the ifNotExists production.
	EnterIfNotExists(c *IfNotExistsContext)

	// EnterIfExists is called when entering the ifExists production.
	EnterIfExists(c *IfExistsContext)

	// EnterOrReplace is called when entering the orReplace production.
	EnterOrReplace(c *OrReplaceContext)

	// EnterSetPassword is called when entering the setPassword production.
	EnterSetPassword(c *SetPasswordContext)

	// EnterPasswordStatus is called when entering the passwordStatus production.
	EnterPasswordStatus(c *PasswordStatusContext)

	// EnterSetStatus is called when entering the setStatus production.
	EnterSetStatus(c *SetStatusContext)

	// EnterUserStatus is called when entering the userStatus production.
	EnterUserStatus(c *UserStatusContext)

	// EnterCreateUniqueConstraint is called when entering the createUniqueConstraint production.
	EnterCreateUniqueConstraint(c *CreateUniqueConstraintContext)

	// EnterCreateNodeKeyConstraint is called when entering the createNodeKeyConstraint production.
	EnterCreateNodeKeyConstraint(c *CreateNodeKeyConstraintContext)

	// EnterCreateNodePropertyExistenceConstraint is called when entering the createNodePropertyExistenceConstraint production.
	EnterCreateNodePropertyExistenceConstraint(c *CreateNodePropertyExistenceConstraintContext)

	// EnterCreateRelationshipPropertyExistenceConstraint is called when entering the createRelationshipPropertyExistenceConstraint production.
	EnterCreateRelationshipPropertyExistenceConstraint(c *CreateRelationshipPropertyExistenceConstraintContext)

	// EnterCreateIndex is called when entering the createIndex production.
	EnterCreateIndex(c *CreateIndexContext)

	// EnterDropUniqueConstraint is called when entering the dropUniqueConstraint production.
	EnterDropUniqueConstraint(c *DropUniqueConstraintContext)

	// EnterDropNodeKeyConstraint is called when entering the dropNodeKeyConstraint production.
	EnterDropNodeKeyConstraint(c *DropNodeKeyConstraintContext)

	// EnterDropNodePropertyExistenceConstraint is called when entering the dropNodePropertyExistenceConstraint production.
	EnterDropNodePropertyExistenceConstraint(c *DropNodePropertyExistenceConstraintContext)

	// EnterDropRelationshipPropertyExistenceConstraint is called when entering the dropRelationshipPropertyExistenceConstraint production.
	EnterDropRelationshipPropertyExistenceConstraint(c *DropRelationshipPropertyExistenceConstraintContext)

	// EnterDropIndex is called when entering the dropIndex production.
	EnterDropIndex(c *DropIndexContext)

	// EnterIndex is called when entering the index production.
	EnterIndex(c *IndexContext)

	// EnterUniqueConstraint is called when entering the uniqueConstraint production.
	EnterUniqueConstraint(c *UniqueConstraintContext)

	// EnterNodeKeyConstraint is called when entering the nodeKeyConstraint production.
	EnterNodeKeyConstraint(c *NodeKeyConstraintContext)

	// EnterNodePropertyExistenceConstraint is called when entering the nodePropertyExistenceConstraint production.
	EnterNodePropertyExistenceConstraint(c *NodePropertyExistenceConstraintContext)

	// EnterRelationshipPropertyExistenceConstraint is called when entering the relationshipPropertyExistenceConstraint production.
	EnterRelationshipPropertyExistenceConstraint(c *RelationshipPropertyExistenceConstraintContext)

	// EnterRelationshipPatternSyntax is called when entering the relationshipPatternSyntax production.
	EnterRelationshipPatternSyntax(c *RelationshipPatternSyntaxContext)

	// EnterLoadCSVClause is called when entering the loadCSVClause production.
	EnterLoadCSVClause(c *LoadCSVClauseContext)

	// EnterMatchClause is called when entering the matchClause production.
	EnterMatchClause(c *MatchClauseContext)

	// EnterUnwindClause is called when entering the unwindClause production.
	EnterUnwindClause(c *UnwindClauseContext)

	// EnterMergeClause is called when entering the mergeClause production.
	EnterMergeClause(c *MergeClauseContext)

	// EnterMergeAction is called when entering the mergeAction production.
	EnterMergeAction(c *MergeActionContext)

	// EnterCreateClause is called when entering the createClause production.
	EnterCreateClause(c *CreateClauseContext)

	// EnterCreateUniqueClause is called when entering the createUniqueClause production.
	EnterCreateUniqueClause(c *CreateUniqueClauseContext)

	// EnterSetClause is called when entering the setClause production.
	EnterSetClause(c *SetClauseContext)

	// EnterSetItem is called when entering the setItem production.
	EnterSetItem(c *SetItemContext)

	// EnterDeleteClause is called when entering the deleteClause production.
	EnterDeleteClause(c *DeleteClauseContext)

	// EnterRemoveClause is called when entering the removeClause production.
	EnterRemoveClause(c *RemoveClauseContext)

	// EnterRemoveItem is called when entering the removeItem production.
	EnterRemoveItem(c *RemoveItemContext)

	// EnterForeachClause is called when entering the foreachClause production.
	EnterForeachClause(c *ForeachClauseContext)

	// EnterWithClause is called when entering the withClause production.
	EnterWithClause(c *WithClauseContext)

	// EnterReturnClause is called when entering the returnClause production.
	EnterReturnClause(c *ReturnClauseContext)

	// EnterReturnBody is called when entering the returnBody production.
	EnterReturnBody(c *ReturnBodyContext)

	// EnterFunction is called when entering the function production.
	EnterFunction(c *FunctionContext)

	// EnterReturnItems is called when entering the returnItems production.
	EnterReturnItems(c *ReturnItemsContext)

	// EnterReturnItem is called when entering the returnItem production.
	EnterReturnItem(c *ReturnItemContext)

	// EnterCall is called when entering the call production.
	EnterCall(c *CallContext)

	// EnterProcedureInvocation is called when entering the procedureInvocation production.
	EnterProcedureInvocation(c *ProcedureInvocationContext)

	// EnterProcedureInvocationBody is called when entering the procedureInvocationBody production.
	EnterProcedureInvocationBody(c *ProcedureInvocationBodyContext)

	// EnterProcedureArguments is called when entering the procedureArguments production.
	EnterProcedureArguments(c *ProcedureArgumentsContext)

	// EnterProcedureResults is called when entering the procedureResults production.
	EnterProcedureResults(c *ProcedureResultsContext)

	// EnterProcedureResult is called when entering the procedureResult production.
	EnterProcedureResult(c *ProcedureResultContext)

	// EnterAliasedProcedureResult is called when entering the aliasedProcedureResult production.
	EnterAliasedProcedureResult(c *AliasedProcedureResultContext)

	// EnterSimpleProcedureResult is called when entering the simpleProcedureResult production.
	EnterSimpleProcedureResult(c *SimpleProcedureResultContext)

	// EnterProcedureOutput is called when entering the procedureOutput production.
	EnterProcedureOutput(c *ProcedureOutputContext)

	// EnterOrder is called when entering the order production.
	EnterOrder(c *OrderContext)

	// EnterSkip is called when entering the skip production.
	EnterSkip(c *SkipContext)

	// EnterLimit is called when entering the limit production.
	EnterLimit(c *LimitContext)

	// EnterSortItem is called when entering the sortItem production.
	EnterSortItem(c *SortItemContext)

	// EnterHint is called when entering the hint production.
	EnterHint(c *HintContext)

	// EnterStartClause is called when entering the startClause production.
	EnterStartClause(c *StartClauseContext)

	// EnterStartPoint is called when entering the startPoint production.
	EnterStartPoint(c *StartPointContext)

	// EnterLookup is called when entering the lookup production.
	EnterLookup(c *LookupContext)

	// EnterNodeLookup is called when entering the nodeLookup production.
	EnterNodeLookup(c *NodeLookupContext)

	// EnterRelationshipLookup is called when entering the relationshipLookup production.
	EnterRelationshipLookup(c *RelationshipLookupContext)

	// EnterIdentifiedIndexLookup is called when entering the identifiedIndexLookup production.
	EnterIdentifiedIndexLookup(c *IdentifiedIndexLookupContext)

	// EnterIndexQuery is called when entering the indexQuery production.
	EnterIndexQuery(c *IndexQueryContext)

	// EnterIdLookup is called when entering the idLookup production.
	EnterIdLookup(c *IdLookupContext)

	// EnterLiteralIds is called when entering the literalIds production.
	EnterLiteralIds(c *LiteralIdsContext)

	// EnterWhere is called when entering the where production.
	EnterWhere(c *WhereContext)

	// EnterPattern is called when entering the pattern production.
	EnterPattern(c *PatternContext)

	// EnterPatternPart is called when entering the patternPart production.
	EnterPatternPart(c *PatternPartContext)

	// EnterAnonymousPatternPart is called when entering the anonymousPatternPart production.
	EnterAnonymousPatternPart(c *AnonymousPatternPartContext)

	// EnterPatternElement is called when entering the patternElement production.
	EnterPatternElement(c *PatternElementContext)

	// EnterNodePattern is called when entering the nodePattern production.
	EnterNodePattern(c *NodePatternContext)

	// EnterPatternElementChain is called when entering the patternElementChain production.
	EnterPatternElementChain(c *PatternElementChainContext)

	// EnterRelationshipPattern is called when entering the relationshipPattern production.
	EnterRelationshipPattern(c *RelationshipPatternContext)

	// EnterRelationshipPatternStart is called when entering the relationshipPatternStart production.
	EnterRelationshipPatternStart(c *RelationshipPatternStartContext)

	// EnterRelationshipPatternEnd is called when entering the relationshipPatternEnd production.
	EnterRelationshipPatternEnd(c *RelationshipPatternEndContext)

	// EnterRelationshipDetail is called when entering the relationshipDetail production.
	EnterRelationshipDetail(c *RelationshipDetailContext)

	// EnterProperties is called when entering the properties production.
	EnterProperties(c *PropertiesContext)

	// EnterRelType is called when entering the relType production.
	EnterRelType(c *RelTypeContext)

	// EnterRelationshipTypes is called when entering the relationshipTypes production.
	EnterRelationshipTypes(c *RelationshipTypesContext)

	// EnterRelationshipType is called when entering the relationshipType production.
	EnterRelationshipType(c *RelationshipTypeContext)

	// EnterRelationshipTypeOptionalColon is called when entering the relationshipTypeOptionalColon production.
	EnterRelationshipTypeOptionalColon(c *RelationshipTypeOptionalColonContext)

	// EnterNodeLabels is called when entering the nodeLabels production.
	EnterNodeLabels(c *NodeLabelsContext)

	// EnterNodeLabel is called when entering the nodeLabel production.
	EnterNodeLabel(c *NodeLabelContext)

	// EnterRangeLiteral is called when entering the rangeLiteral production.
	EnterRangeLiteral(c *RangeLiteralContext)

	// EnterLabelName is called when entering the labelName production.
	EnterLabelName(c *LabelNameContext)

	// EnterRelTypeName is called when entering the relTypeName production.
	EnterRelTypeName(c *RelTypeNameContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterOrExpression is called when entering the orExpression production.
	EnterOrExpression(c *OrExpressionContext)

	// EnterXorExpression is called when entering the xorExpression production.
	EnterXorExpression(c *XorExpressionContext)

	// EnterAndExpression is called when entering the andExpression production.
	EnterAndExpression(c *AndExpressionContext)

	// EnterNotExpression is called when entering the notExpression production.
	EnterNotExpression(c *NotExpressionContext)

	// EnterComparisonExpression is called when entering the comparisonExpression production.
	EnterComparisonExpression(c *ComparisonExpressionContext)

	// EnterAddOrSubtractExpression is called when entering the addOrSubtractExpression production.
	EnterAddOrSubtractExpression(c *AddOrSubtractExpressionContext)

	// EnterMultiplyDivideModuloExpression is called when entering the multiplyDivideModuloExpression production.
	EnterMultiplyDivideModuloExpression(c *MultiplyDivideModuloExpressionContext)

	// EnterPowerOfExpression is called when entering the powerOfExpression production.
	EnterPowerOfExpression(c *PowerOfExpressionContext)

	// EnterUnaryAddOrSubtractExpression is called when entering the unaryAddOrSubtractExpression production.
	EnterUnaryAddOrSubtractExpression(c *UnaryAddOrSubtractExpressionContext)

	// EnterStringListNullOperatorExpression is called when entering the stringListNullOperatorExpression production.
	EnterStringListNullOperatorExpression(c *StringListNullOperatorExpressionContext)

	// EnterPropertyOrLabelsExpression is called when entering the propertyOrLabelsExpression production.
	EnterPropertyOrLabelsExpression(c *PropertyOrLabelsExpressionContext)

	// EnterFilterFunction is called when entering the filterFunction production.
	EnterFilterFunction(c *FilterFunctionContext)

	// EnterFilterFunctionName is called when entering the filterFunctionName production.
	EnterFilterFunctionName(c *FilterFunctionNameContext)

	// EnterExistsFunction is called when entering the existsFunction production.
	EnterExistsFunction(c *ExistsFunctionContext)

	// EnterExistsFunctionName is called when entering the existsFunctionName production.
	EnterExistsFunctionName(c *ExistsFunctionNameContext)

	// EnterAllFunction is called when entering the allFunction production.
	EnterAllFunction(c *AllFunctionContext)

	// EnterAllFunctionName is called when entering the allFunctionName production.
	EnterAllFunctionName(c *AllFunctionNameContext)

	// EnterAnyFunction is called when entering the anyFunction production.
	EnterAnyFunction(c *AnyFunctionContext)

	// EnterAnyFunctionName is called when entering the anyFunctionName production.
	EnterAnyFunctionName(c *AnyFunctionNameContext)

	// EnterNoneFunction is called when entering the noneFunction production.
	EnterNoneFunction(c *NoneFunctionContext)

	// EnterNoneFunctionName is called when entering the noneFunctionName production.
	EnterNoneFunctionName(c *NoneFunctionNameContext)

	// EnterSingleFunction is called when entering the singleFunction production.
	EnterSingleFunction(c *SingleFunctionContext)

	// EnterSingleFunctionName is called when entering the singleFunctionName production.
	EnterSingleFunctionName(c *SingleFunctionNameContext)

	// EnterExtractFunction is called when entering the extractFunction production.
	EnterExtractFunction(c *ExtractFunctionContext)

	// EnterExtractFunctionName is called when entering the extractFunctionName production.
	EnterExtractFunctionName(c *ExtractFunctionNameContext)

	// EnterReduceFunction is called when entering the reduceFunction production.
	EnterReduceFunction(c *ReduceFunctionContext)

	// EnterReduceFunctionName is called when entering the reduceFunctionName production.
	EnterReduceFunctionName(c *ReduceFunctionNameContext)

	// EnterShortestPathPatternFunction is called when entering the shortestPathPatternFunction production.
	EnterShortestPathPatternFunction(c *ShortestPathPatternFunctionContext)

	// EnterShortestPathFunctionName is called when entering the shortestPathFunctionName production.
	EnterShortestPathFunctionName(c *ShortestPathFunctionNameContext)

	// EnterAllShortestPathFunctionName is called when entering the allShortestPathFunctionName production.
	EnterAllShortestPathFunctionName(c *AllShortestPathFunctionNameContext)

	// EnterAtom is called when entering the atom production.
	EnterAtom(c *AtomContext)

	// EnterLiteral is called when entering the literal production.
	EnterLiteral(c *LiteralContext)

	// EnterStringLiteral is called when entering the stringLiteral production.
	EnterStringLiteral(c *StringLiteralContext)

	// EnterBooleanLiteral is called when entering the booleanLiteral production.
	EnterBooleanLiteral(c *BooleanLiteralContext)

	// EnterListLiteral is called when entering the listLiteral production.
	EnterListLiteral(c *ListLiteralContext)

	// EnterPartialComparisonExpression is called when entering the partialComparisonExpression production.
	EnterPartialComparisonExpression(c *PartialComparisonExpressionContext)

	// EnterParenthesizedExpression is called when entering the parenthesizedExpression production.
	EnterParenthesizedExpression(c *ParenthesizedExpressionContext)

	// EnterRelationshipsPattern is called when entering the relationshipsPattern production.
	EnterRelationshipsPattern(c *RelationshipsPatternContext)

	// EnterFilterExpression is called when entering the filterExpression production.
	EnterFilterExpression(c *FilterExpressionContext)

	// EnterIdInColl is called when entering the idInColl production.
	EnterIdInColl(c *IdInCollContext)

	// EnterFunctionInvocation is called when entering the functionInvocation production.
	EnterFunctionInvocation(c *FunctionInvocationContext)

	// EnterFunctionInvocationBody is called when entering the functionInvocationBody production.
	EnterFunctionInvocationBody(c *FunctionInvocationBodyContext)

	// EnterFunctionName is called when entering the functionName production.
	EnterFunctionName(c *FunctionNameContext)

	// EnterProcedureName is called when entering the procedureName production.
	EnterProcedureName(c *ProcedureNameContext)

	// EnterListComprehension is called when entering the listComprehension production.
	EnterListComprehension(c *ListComprehensionContext)

	// EnterPatternComprehension is called when entering the patternComprehension production.
	EnterPatternComprehension(c *PatternComprehensionContext)

	// EnterPropertyLookup is called when entering the propertyLookup production.
	EnterPropertyLookup(c *PropertyLookupContext)

	// EnterCaseExpression is called when entering the caseExpression production.
	EnterCaseExpression(c *CaseExpressionContext)

	// EnterCaseAlternatives is called when entering the caseAlternatives production.
	EnterCaseAlternatives(c *CaseAlternativesContext)

	// EnterVariable is called when entering the variable production.
	EnterVariable(c *VariableContext)

	// EnterNumberLiteral is called when entering the numberLiteral production.
	EnterNumberLiteral(c *NumberLiteralContext)

	// EnterMapLiteral is called when entering the mapLiteral production.
	EnterMapLiteral(c *MapLiteralContext)

	// EnterMapProjection is called when entering the mapProjection production.
	EnterMapProjection(c *MapProjectionContext)

	// EnterMapProjectionVariants is called when entering the mapProjectionVariants production.
	EnterMapProjectionVariants(c *MapProjectionVariantsContext)

	// EnterLiteralEntry is called when entering the literalEntry production.
	EnterLiteralEntry(c *LiteralEntryContext)

	// EnterPropertySelector is called when entering the propertySelector production.
	EnterPropertySelector(c *PropertySelectorContext)

	// EnterVariableSelector is called when entering the variableSelector production.
	EnterVariableSelector(c *VariableSelectorContext)

	// EnterAllPropertiesSelector is called when entering the allPropertiesSelector production.
	EnterAllPropertiesSelector(c *AllPropertiesSelectorContext)

	// EnterParameter is called when entering the parameter production.
	EnterParameter(c *ParameterContext)

	// EnterLegacyParameter is called when entering the legacyParameter production.
	EnterLegacyParameter(c *LegacyParameterContext)

	// EnterDollarParameter is called when entering the dollarParameter production.
	EnterDollarParameter(c *DollarParameterContext)

	// EnterParameterName is called when entering the parameterName production.
	EnterParameterName(c *ParameterNameContext)

	// EnterPropertyExpressions is called when entering the propertyExpressions production.
	EnterPropertyExpressions(c *PropertyExpressionsContext)

	// EnterPropertyExpression is called when entering the propertyExpression production.
	EnterPropertyExpression(c *PropertyExpressionContext)

	// EnterPropertyKeys is called when entering the propertyKeys production.
	EnterPropertyKeys(c *PropertyKeysContext)

	// EnterPropertyKeyName is called when entering the propertyKeyName production.
	EnterPropertyKeyName(c *PropertyKeyNameContext)

	// EnterIntegerLiteral is called when entering the integerLiteral production.
	EnterIntegerLiteral(c *IntegerLiteralContext)

	// EnterDoubleLiteral is called when entering the doubleLiteral production.
	EnterDoubleLiteral(c *DoubleLiteralContext)

	// EnterNamespace is called when entering the namespace production.
	EnterNamespace(c *NamespaceContext)

	// EnterLeftArrowHead is called when entering the leftArrowHead production.
	EnterLeftArrowHead(c *LeftArrowHeadContext)

	// EnterRightArrowHead is called when entering the rightArrowHead production.
	EnterRightArrowHead(c *RightArrowHeadContext)

	// EnterDash is called when entering the dash production.
	EnterDash(c *DashContext)

	// EnterSymbolicName is called when entering the symbolicName production.
	EnterSymbolicName(c *SymbolicNameContext)

	// EnterKeyword is called when entering the keyword production.
	EnterKeyword(c *KeywordContext)

	// ExitCypher is called when exiting the cypher production.
	ExitCypher(c *CypherContext)

	// ExitCypherPart is called when exiting the cypherPart production.
	ExitCypherPart(c *CypherPartContext)

	// ExitCypherConsoleCommand is called when exiting the cypherConsoleCommand production.
	ExitCypherConsoleCommand(c *CypherConsoleCommandContext)

	// ExitCypherConsoleCommandName is called when exiting the cypherConsoleCommandName production.
	ExitCypherConsoleCommandName(c *CypherConsoleCommandNameContext)

	// ExitCypherConsoleCommandParameters is called when exiting the cypherConsoleCommandParameters production.
	ExitCypherConsoleCommandParameters(c *CypherConsoleCommandParametersContext)

	// ExitCypherConsoleCommandParameter is called when exiting the cypherConsoleCommandParameter production.
	ExitCypherConsoleCommandParameter(c *CypherConsoleCommandParameterContext)

	// ExitArrowExpression is called when exiting the arrowExpression production.
	ExitArrowExpression(c *ArrowExpressionContext)

	// ExitUrl is called when exiting the url production.
	ExitUrl(c *UrlContext)

	// ExitUri is called when exiting the uri production.
	ExitUri(c *UriContext)

	// ExitScheme is called when exiting the scheme production.
	ExitScheme(c *SchemeContext)

	// ExitHost is called when exiting the host production.
	ExitHost(c *HostContext)

	// ExitHostname is called when exiting the hostname production.
	ExitHostname(c *HostnameContext)

	// ExitHostnumber is called when exiting the hostnumber production.
	ExitHostnumber(c *HostnumberContext)

	// ExitPort is called when exiting the port production.
	ExitPort(c *PortContext)

	// ExitPath is called when exiting the path production.
	ExitPath(c *PathContext)

	// ExitUser is called when exiting the user production.
	ExitUser(c *UserContext)

	// ExitLogin is called when exiting the login production.
	ExitLogin(c *LoginContext)

	// ExitPassword is called when exiting the password production.
	ExitPassword(c *PasswordContext)

	// ExitFrag is called when exiting the frag production.
	ExitFrag(c *FragContext)

	// ExitUrlQuery is called when exiting the urlQuery production.
	ExitUrlQuery(c *UrlQueryContext)

	// ExitSearch is called when exiting the search production.
	ExitSearch(c *SearchContext)

	// ExitSearchparameter is called when exiting the searchparameter production.
	ExitSearchparameter(c *SearchparameterContext)

	// ExitStr is called when exiting the str production.
	ExitStr(c *StrContext)

	// ExitUrlDigits is called when exiting the urlDigits production.
	ExitUrlDigits(c *UrlDigitsContext)

	// ExitJson is called when exiting the json production.
	ExitJson(c *JsonContext)

	// ExitObj is called when exiting the obj production.
	ExitObj(c *ObjContext)

	// ExitPair is called when exiting the pair production.
	ExitPair(c *PairContext)

	// ExitArray is called when exiting the array production.
	ExitArray(c *ArrayContext)

	// ExitValue is called when exiting the value production.
	ExitValue(c *ValueContext)

	// ExitKeyValueLiteral is called when exiting the keyValueLiteral production.
	ExitKeyValueLiteral(c *KeyValueLiteralContext)

	// ExitCommandPath is called when exiting the commandPath production.
	ExitCommandPath(c *CommandPathContext)

	// ExitSubCommand is called when exiting the subCommand production.
	ExitSubCommand(c *SubCommandContext)

	// ExitCypherQuery is called when exiting the cypherQuery production.
	ExitCypherQuery(c *CypherQueryContext)

	// ExitQueryOptions is called when exiting the queryOptions production.
	ExitQueryOptions(c *QueryOptionsContext)

	// ExitAnyCypherOption is called when exiting the anyCypherOption production.
	ExitAnyCypherOption(c *AnyCypherOptionContext)

	// ExitCypherOption is called when exiting the cypherOption production.
	ExitCypherOption(c *CypherOptionContext)

	// ExitVersionNumber is called when exiting the versionNumber production.
	ExitVersionNumber(c *VersionNumberContext)

	// ExitExplain is called when exiting the explain production.
	ExitExplain(c *ExplainContext)

	// ExitProfile is called when exiting the profile production.
	ExitProfile(c *ProfileContext)

	// ExitConfigurationOption is called when exiting the configurationOption production.
	ExitConfigurationOption(c *ConfigurationOptionContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitQuery is called when exiting the query production.
	ExitQuery(c *QueryContext)

	// ExitRegularQuery is called when exiting the regularQuery production.
	ExitRegularQuery(c *RegularQueryContext)

	// ExitBulkImportQuery is called when exiting the bulkImportQuery production.
	ExitBulkImportQuery(c *BulkImportQueryContext)

	// ExitSingleQuery is called when exiting the singleQuery production.
	ExitSingleQuery(c *SingleQueryContext)

	// ExitPeriodicCommitHint is called when exiting the periodicCommitHint production.
	ExitPeriodicCommitHint(c *PeriodicCommitHintContext)

	// ExitLoadCSVQuery is called when exiting the loadCSVQuery production.
	ExitLoadCSVQuery(c *LoadCSVQueryContext)

	// ExitUnion is called when exiting the union production.
	ExitUnion(c *UnionContext)

	// ExitClause is called when exiting the clause production.
	ExitClause(c *ClauseContext)

	// ExitCommand is called when exiting the command production.
	ExitCommand(c *CommandContext)

	// ExitSystemCommand is called when exiting the systemCommand production.
	ExitSystemCommand(c *SystemCommandContext)

	// ExitMultidatabaseCommand is called when exiting the multidatabaseCommand production.
	ExitMultidatabaseCommand(c *MultidatabaseCommandContext)

	// ExitUserCommand is called when exiting the userCommand production.
	ExitUserCommand(c *UserCommandContext)

	// ExitPrivilegeCommand is called when exiting the privilegeCommand production.
	ExitPrivilegeCommand(c *PrivilegeCommandContext)

	// ExitShowRoles is called when exiting the showRoles production.
	ExitShowRoles(c *ShowRolesContext)

	// ExitCreateRole is called when exiting the createRole production.
	ExitCreateRole(c *CreateRoleContext)

	// ExitCopyRole is called when exiting the copyRole production.
	ExitCopyRole(c *CopyRoleContext)

	// ExitDropRole is called when exiting the dropRole production.
	ExitDropRole(c *DropRoleContext)

	// ExitShowUsers is called when exiting the showUsers production.
	ExitShowUsers(c *ShowUsersContext)

	// ExitCreateUser is called when exiting the createUser production.
	ExitCreateUser(c *CreateUserContext)

	// ExitDropUser is called when exiting the dropUser production.
	ExitDropUser(c *DropUserContext)

	// ExitAlterUser is called when exiting the alterUser production.
	ExitAlterUser(c *AlterUserContext)

	// ExitShowPrivileges is called when exiting the showPrivileges production.
	ExitShowPrivileges(c *ShowPrivilegesContext)

	// ExitGrantPrivilege is called when exiting the grantPrivilege production.
	ExitGrantPrivilege(c *GrantPrivilegeContext)

	// ExitDenyPrivilege is called when exiting the denyPrivilege production.
	ExitDenyPrivilege(c *DenyPrivilegeContext)

	// ExitRevokePrivilege is called when exiting the revokePrivilege production.
	ExitRevokePrivilege(c *RevokePrivilegeContext)

	// ExitRevokePart is called when exiting the revokePart production.
	ExitRevokePart(c *RevokePartContext)

	// ExitDatabaseScope is called when exiting the databaseScope production.
	ExitDatabaseScope(c *DatabaseScopeContext)

	// ExitGraphScope is called when exiting the graphScope production.
	ExitGraphScope(c *GraphScopeContext)

	// ExitRoles is called when exiting the roles production.
	ExitRoles(c *RolesContext)

	// ExitGrantableGraphPrivileges is called when exiting the grantableGraphPrivileges production.
	ExitGrantableGraphPrivileges(c *GrantableGraphPrivilegesContext)

	// ExitRevokeableGraphPrivileges is called when exiting the revokeableGraphPrivileges production.
	ExitRevokeableGraphPrivileges(c *RevokeableGraphPrivilegesContext)

	// ExitDatasbasePrivilege is called when exiting the datasbasePrivilege production.
	ExitDatasbasePrivilege(c *DatasbasePrivilegeContext)

	// ExitDbmsPrivilege is called when exiting the dbmsPrivilege production.
	ExitDbmsPrivilege(c *DbmsPrivilegeContext)

	// ExitElementScope is called when exiting the elementScope production.
	ExitElementScope(c *ElementScopeContext)

	// ExitPropertiesList is called when exiting the propertiesList production.
	ExitPropertiesList(c *PropertiesListContext)

	// ExitPropertyScope is called when exiting the propertyScope production.
	ExitPropertyScope(c *PropertyScopeContext)

	// ExitShowDatabase is called when exiting the showDatabase production.
	ExitShowDatabase(c *ShowDatabaseContext)

	// ExitCreateDatabase is called when exiting the createDatabase production.
	ExitCreateDatabase(c *CreateDatabaseContext)

	// ExitDropDatabase is called when exiting the dropDatabase production.
	ExitDropDatabase(c *DropDatabaseContext)

	// ExitStartDatabase is called when exiting the startDatabase production.
	ExitStartDatabase(c *StartDatabaseContext)

	// ExitStopDatabase is called when exiting the stopDatabase production.
	ExitStopDatabase(c *StopDatabaseContext)

	// ExitIfNotExists is called when exiting the ifNotExists production.
	ExitIfNotExists(c *IfNotExistsContext)

	// ExitIfExists is called when exiting the ifExists production.
	ExitIfExists(c *IfExistsContext)

	// ExitOrReplace is called when exiting the orReplace production.
	ExitOrReplace(c *OrReplaceContext)

	// ExitSetPassword is called when exiting the setPassword production.
	ExitSetPassword(c *SetPasswordContext)

	// ExitPasswordStatus is called when exiting the passwordStatus production.
	ExitPasswordStatus(c *PasswordStatusContext)

	// ExitSetStatus is called when exiting the setStatus production.
	ExitSetStatus(c *SetStatusContext)

	// ExitUserStatus is called when exiting the userStatus production.
	ExitUserStatus(c *UserStatusContext)

	// ExitCreateUniqueConstraint is called when exiting the createUniqueConstraint production.
	ExitCreateUniqueConstraint(c *CreateUniqueConstraintContext)

	// ExitCreateNodeKeyConstraint is called when exiting the createNodeKeyConstraint production.
	ExitCreateNodeKeyConstraint(c *CreateNodeKeyConstraintContext)

	// ExitCreateNodePropertyExistenceConstraint is called when exiting the createNodePropertyExistenceConstraint production.
	ExitCreateNodePropertyExistenceConstraint(c *CreateNodePropertyExistenceConstraintContext)

	// ExitCreateRelationshipPropertyExistenceConstraint is called when exiting the createRelationshipPropertyExistenceConstraint production.
	ExitCreateRelationshipPropertyExistenceConstraint(c *CreateRelationshipPropertyExistenceConstraintContext)

	// ExitCreateIndex is called when exiting the createIndex production.
	ExitCreateIndex(c *CreateIndexContext)

	// ExitDropUniqueConstraint is called when exiting the dropUniqueConstraint production.
	ExitDropUniqueConstraint(c *DropUniqueConstraintContext)

	// ExitDropNodeKeyConstraint is called when exiting the dropNodeKeyConstraint production.
	ExitDropNodeKeyConstraint(c *DropNodeKeyConstraintContext)

	// ExitDropNodePropertyExistenceConstraint is called when exiting the dropNodePropertyExistenceConstraint production.
	ExitDropNodePropertyExistenceConstraint(c *DropNodePropertyExistenceConstraintContext)

	// ExitDropRelationshipPropertyExistenceConstraint is called when exiting the dropRelationshipPropertyExistenceConstraint production.
	ExitDropRelationshipPropertyExistenceConstraint(c *DropRelationshipPropertyExistenceConstraintContext)

	// ExitDropIndex is called when exiting the dropIndex production.
	ExitDropIndex(c *DropIndexContext)

	// ExitIndex is called when exiting the index production.
	ExitIndex(c *IndexContext)

	// ExitUniqueConstraint is called when exiting the uniqueConstraint production.
	ExitUniqueConstraint(c *UniqueConstraintContext)

	// ExitNodeKeyConstraint is called when exiting the nodeKeyConstraint production.
	ExitNodeKeyConstraint(c *NodeKeyConstraintContext)

	// ExitNodePropertyExistenceConstraint is called when exiting the nodePropertyExistenceConstraint production.
	ExitNodePropertyExistenceConstraint(c *NodePropertyExistenceConstraintContext)

	// ExitRelationshipPropertyExistenceConstraint is called when exiting the relationshipPropertyExistenceConstraint production.
	ExitRelationshipPropertyExistenceConstraint(c *RelationshipPropertyExistenceConstraintContext)

	// ExitRelationshipPatternSyntax is called when exiting the relationshipPatternSyntax production.
	ExitRelationshipPatternSyntax(c *RelationshipPatternSyntaxContext)

	// ExitLoadCSVClause is called when exiting the loadCSVClause production.
	ExitLoadCSVClause(c *LoadCSVClauseContext)

	// ExitMatchClause is called when exiting the matchClause production.
	ExitMatchClause(c *MatchClauseContext)

	// ExitUnwindClause is called when exiting the unwindClause production.
	ExitUnwindClause(c *UnwindClauseContext)

	// ExitMergeClause is called when exiting the mergeClause production.
	ExitMergeClause(c *MergeClauseContext)

	// ExitMergeAction is called when exiting the mergeAction production.
	ExitMergeAction(c *MergeActionContext)

	// ExitCreateClause is called when exiting the createClause production.
	ExitCreateClause(c *CreateClauseContext)

	// ExitCreateUniqueClause is called when exiting the createUniqueClause production.
	ExitCreateUniqueClause(c *CreateUniqueClauseContext)

	// ExitSetClause is called when exiting the setClause production.
	ExitSetClause(c *SetClauseContext)

	// ExitSetItem is called when exiting the setItem production.
	ExitSetItem(c *SetItemContext)

	// ExitDeleteClause is called when exiting the deleteClause production.
	ExitDeleteClause(c *DeleteClauseContext)

	// ExitRemoveClause is called when exiting the removeClause production.
	ExitRemoveClause(c *RemoveClauseContext)

	// ExitRemoveItem is called when exiting the removeItem production.
	ExitRemoveItem(c *RemoveItemContext)

	// ExitForeachClause is called when exiting the foreachClause production.
	ExitForeachClause(c *ForeachClauseContext)

	// ExitWithClause is called when exiting the withClause production.
	ExitWithClause(c *WithClauseContext)

	// ExitReturnClause is called when exiting the returnClause production.
	ExitReturnClause(c *ReturnClauseContext)

	// ExitReturnBody is called when exiting the returnBody production.
	ExitReturnBody(c *ReturnBodyContext)

	// ExitFunction is called when exiting the function production.
	ExitFunction(c *FunctionContext)

	// ExitReturnItems is called when exiting the returnItems production.
	ExitReturnItems(c *ReturnItemsContext)

	// ExitReturnItem is called when exiting the returnItem production.
	ExitReturnItem(c *ReturnItemContext)

	// ExitCall is called when exiting the call production.
	ExitCall(c *CallContext)

	// ExitProcedureInvocation is called when exiting the procedureInvocation production.
	ExitProcedureInvocation(c *ProcedureInvocationContext)

	// ExitProcedureInvocationBody is called when exiting the procedureInvocationBody production.
	ExitProcedureInvocationBody(c *ProcedureInvocationBodyContext)

	// ExitProcedureArguments is called when exiting the procedureArguments production.
	ExitProcedureArguments(c *ProcedureArgumentsContext)

	// ExitProcedureResults is called when exiting the procedureResults production.
	ExitProcedureResults(c *ProcedureResultsContext)

	// ExitProcedureResult is called when exiting the procedureResult production.
	ExitProcedureResult(c *ProcedureResultContext)

	// ExitAliasedProcedureResult is called when exiting the aliasedProcedureResult production.
	ExitAliasedProcedureResult(c *AliasedProcedureResultContext)

	// ExitSimpleProcedureResult is called when exiting the simpleProcedureResult production.
	ExitSimpleProcedureResult(c *SimpleProcedureResultContext)

	// ExitProcedureOutput is called when exiting the procedureOutput production.
	ExitProcedureOutput(c *ProcedureOutputContext)

	// ExitOrder is called when exiting the order production.
	ExitOrder(c *OrderContext)

	// ExitSkip is called when exiting the skip production.
	ExitSkip(c *SkipContext)

	// ExitLimit is called when exiting the limit production.
	ExitLimit(c *LimitContext)

	// ExitSortItem is called when exiting the sortItem production.
	ExitSortItem(c *SortItemContext)

	// ExitHint is called when exiting the hint production.
	ExitHint(c *HintContext)

	// ExitStartClause is called when exiting the startClause production.
	ExitStartClause(c *StartClauseContext)

	// ExitStartPoint is called when exiting the startPoint production.
	ExitStartPoint(c *StartPointContext)

	// ExitLookup is called when exiting the lookup production.
	ExitLookup(c *LookupContext)

	// ExitNodeLookup is called when exiting the nodeLookup production.
	ExitNodeLookup(c *NodeLookupContext)

	// ExitRelationshipLookup is called when exiting the relationshipLookup production.
	ExitRelationshipLookup(c *RelationshipLookupContext)

	// ExitIdentifiedIndexLookup is called when exiting the identifiedIndexLookup production.
	ExitIdentifiedIndexLookup(c *IdentifiedIndexLookupContext)

	// ExitIndexQuery is called when exiting the indexQuery production.
	ExitIndexQuery(c *IndexQueryContext)

	// ExitIdLookup is called when exiting the idLookup production.
	ExitIdLookup(c *IdLookupContext)

	// ExitLiteralIds is called when exiting the literalIds production.
	ExitLiteralIds(c *LiteralIdsContext)

	// ExitWhere is called when exiting the where production.
	ExitWhere(c *WhereContext)

	// ExitPattern is called when exiting the pattern production.
	ExitPattern(c *PatternContext)

	// ExitPatternPart is called when exiting the patternPart production.
	ExitPatternPart(c *PatternPartContext)

	// ExitAnonymousPatternPart is called when exiting the anonymousPatternPart production.
	ExitAnonymousPatternPart(c *AnonymousPatternPartContext)

	// ExitPatternElement is called when exiting the patternElement production.
	ExitPatternElement(c *PatternElementContext)

	// ExitNodePattern is called when exiting the nodePattern production.
	ExitNodePattern(c *NodePatternContext)

	// ExitPatternElementChain is called when exiting the patternElementChain production.
	ExitPatternElementChain(c *PatternElementChainContext)

	// ExitRelationshipPattern is called when exiting the relationshipPattern production.
	ExitRelationshipPattern(c *RelationshipPatternContext)

	// ExitRelationshipPatternStart is called when exiting the relationshipPatternStart production.
	ExitRelationshipPatternStart(c *RelationshipPatternStartContext)

	// ExitRelationshipPatternEnd is called when exiting the relationshipPatternEnd production.
	ExitRelationshipPatternEnd(c *RelationshipPatternEndContext)

	// ExitRelationshipDetail is called when exiting the relationshipDetail production.
	ExitRelationshipDetail(c *RelationshipDetailContext)

	// ExitProperties is called when exiting the properties production.
	ExitProperties(c *PropertiesContext)

	// ExitRelType is called when exiting the relType production.
	ExitRelType(c *RelTypeContext)

	// ExitRelationshipTypes is called when exiting the relationshipTypes production.
	ExitRelationshipTypes(c *RelationshipTypesContext)

	// ExitRelationshipType is called when exiting the relationshipType production.
	ExitRelationshipType(c *RelationshipTypeContext)

	// ExitRelationshipTypeOptionalColon is called when exiting the relationshipTypeOptionalColon production.
	ExitRelationshipTypeOptionalColon(c *RelationshipTypeOptionalColonContext)

	// ExitNodeLabels is called when exiting the nodeLabels production.
	ExitNodeLabels(c *NodeLabelsContext)

	// ExitNodeLabel is called when exiting the nodeLabel production.
	ExitNodeLabel(c *NodeLabelContext)

	// ExitRangeLiteral is called when exiting the rangeLiteral production.
	ExitRangeLiteral(c *RangeLiteralContext)

	// ExitLabelName is called when exiting the labelName production.
	ExitLabelName(c *LabelNameContext)

	// ExitRelTypeName is called when exiting the relTypeName production.
	ExitRelTypeName(c *RelTypeNameContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitOrExpression is called when exiting the orExpression production.
	ExitOrExpression(c *OrExpressionContext)

	// ExitXorExpression is called when exiting the xorExpression production.
	ExitXorExpression(c *XorExpressionContext)

	// ExitAndExpression is called when exiting the andExpression production.
	ExitAndExpression(c *AndExpressionContext)

	// ExitNotExpression is called when exiting the notExpression production.
	ExitNotExpression(c *NotExpressionContext)

	// ExitComparisonExpression is called when exiting the comparisonExpression production.
	ExitComparisonExpression(c *ComparisonExpressionContext)

	// ExitAddOrSubtractExpression is called when exiting the addOrSubtractExpression production.
	ExitAddOrSubtractExpression(c *AddOrSubtractExpressionContext)

	// ExitMultiplyDivideModuloExpression is called when exiting the multiplyDivideModuloExpression production.
	ExitMultiplyDivideModuloExpression(c *MultiplyDivideModuloExpressionContext)

	// ExitPowerOfExpression is called when exiting the powerOfExpression production.
	ExitPowerOfExpression(c *PowerOfExpressionContext)

	// ExitUnaryAddOrSubtractExpression is called when exiting the unaryAddOrSubtractExpression production.
	ExitUnaryAddOrSubtractExpression(c *UnaryAddOrSubtractExpressionContext)

	// ExitStringListNullOperatorExpression is called when exiting the stringListNullOperatorExpression production.
	ExitStringListNullOperatorExpression(c *StringListNullOperatorExpressionContext)

	// ExitPropertyOrLabelsExpression is called when exiting the propertyOrLabelsExpression production.
	ExitPropertyOrLabelsExpression(c *PropertyOrLabelsExpressionContext)

	// ExitFilterFunction is called when exiting the filterFunction production.
	ExitFilterFunction(c *FilterFunctionContext)

	// ExitFilterFunctionName is called when exiting the filterFunctionName production.
	ExitFilterFunctionName(c *FilterFunctionNameContext)

	// ExitExistsFunction is called when exiting the existsFunction production.
	ExitExistsFunction(c *ExistsFunctionContext)

	// ExitExistsFunctionName is called when exiting the existsFunctionName production.
	ExitExistsFunctionName(c *ExistsFunctionNameContext)

	// ExitAllFunction is called when exiting the allFunction production.
	ExitAllFunction(c *AllFunctionContext)

	// ExitAllFunctionName is called when exiting the allFunctionName production.
	ExitAllFunctionName(c *AllFunctionNameContext)

	// ExitAnyFunction is called when exiting the anyFunction production.
	ExitAnyFunction(c *AnyFunctionContext)

	// ExitAnyFunctionName is called when exiting the anyFunctionName production.
	ExitAnyFunctionName(c *AnyFunctionNameContext)

	// ExitNoneFunction is called when exiting the noneFunction production.
	ExitNoneFunction(c *NoneFunctionContext)

	// ExitNoneFunctionName is called when exiting the noneFunctionName production.
	ExitNoneFunctionName(c *NoneFunctionNameContext)

	// ExitSingleFunction is called when exiting the singleFunction production.
	ExitSingleFunction(c *SingleFunctionContext)

	// ExitSingleFunctionName is called when exiting the singleFunctionName production.
	ExitSingleFunctionName(c *SingleFunctionNameContext)

	// ExitExtractFunction is called when exiting the extractFunction production.
	ExitExtractFunction(c *ExtractFunctionContext)

	// ExitExtractFunctionName is called when exiting the extractFunctionName production.
	ExitExtractFunctionName(c *ExtractFunctionNameContext)

	// ExitReduceFunction is called when exiting the reduceFunction production.
	ExitReduceFunction(c *ReduceFunctionContext)

	// ExitReduceFunctionName is called when exiting the reduceFunctionName production.
	ExitReduceFunctionName(c *ReduceFunctionNameContext)

	// ExitShortestPathPatternFunction is called when exiting the shortestPathPatternFunction production.
	ExitShortestPathPatternFunction(c *ShortestPathPatternFunctionContext)

	// ExitShortestPathFunctionName is called when exiting the shortestPathFunctionName production.
	ExitShortestPathFunctionName(c *ShortestPathFunctionNameContext)

	// ExitAllShortestPathFunctionName is called when exiting the allShortestPathFunctionName production.
	ExitAllShortestPathFunctionName(c *AllShortestPathFunctionNameContext)

	// ExitAtom is called when exiting the atom production.
	ExitAtom(c *AtomContext)

	// ExitLiteral is called when exiting the literal production.
	ExitLiteral(c *LiteralContext)

	// ExitStringLiteral is called when exiting the stringLiteral production.
	ExitStringLiteral(c *StringLiteralContext)

	// ExitBooleanLiteral is called when exiting the booleanLiteral production.
	ExitBooleanLiteral(c *BooleanLiteralContext)

	// ExitListLiteral is called when exiting the listLiteral production.
	ExitListLiteral(c *ListLiteralContext)

	// ExitPartialComparisonExpression is called when exiting the partialComparisonExpression production.
	ExitPartialComparisonExpression(c *PartialComparisonExpressionContext)

	// ExitParenthesizedExpression is called when exiting the parenthesizedExpression production.
	ExitParenthesizedExpression(c *ParenthesizedExpressionContext)

	// ExitRelationshipsPattern is called when exiting the relationshipsPattern production.
	ExitRelationshipsPattern(c *RelationshipsPatternContext)

	// ExitFilterExpression is called when exiting the filterExpression production.
	ExitFilterExpression(c *FilterExpressionContext)

	// ExitIdInColl is called when exiting the idInColl production.
	ExitIdInColl(c *IdInCollContext)

	// ExitFunctionInvocation is called when exiting the functionInvocation production.
	ExitFunctionInvocation(c *FunctionInvocationContext)

	// ExitFunctionInvocationBody is called when exiting the functionInvocationBody production.
	ExitFunctionInvocationBody(c *FunctionInvocationBodyContext)

	// ExitFunctionName is called when exiting the functionName production.
	ExitFunctionName(c *FunctionNameContext)

	// ExitProcedureName is called when exiting the procedureName production.
	ExitProcedureName(c *ProcedureNameContext)

	// ExitListComprehension is called when exiting the listComprehension production.
	ExitListComprehension(c *ListComprehensionContext)

	// ExitPatternComprehension is called when exiting the patternComprehension production.
	ExitPatternComprehension(c *PatternComprehensionContext)

	// ExitPropertyLookup is called when exiting the propertyLookup production.
	ExitPropertyLookup(c *PropertyLookupContext)

	// ExitCaseExpression is called when exiting the caseExpression production.
	ExitCaseExpression(c *CaseExpressionContext)

	// ExitCaseAlternatives is called when exiting the caseAlternatives production.
	ExitCaseAlternatives(c *CaseAlternativesContext)

	// ExitVariable is called when exiting the variable production.
	ExitVariable(c *VariableContext)

	// ExitNumberLiteral is called when exiting the numberLiteral production.
	ExitNumberLiteral(c *NumberLiteralContext)

	// ExitMapLiteral is called when exiting the mapLiteral production.
	ExitMapLiteral(c *MapLiteralContext)

	// ExitMapProjection is called when exiting the mapProjection production.
	ExitMapProjection(c *MapProjectionContext)

	// ExitMapProjectionVariants is called when exiting the mapProjectionVariants production.
	ExitMapProjectionVariants(c *MapProjectionVariantsContext)

	// ExitLiteralEntry is called when exiting the literalEntry production.
	ExitLiteralEntry(c *LiteralEntryContext)

	// ExitPropertySelector is called when exiting the propertySelector production.
	ExitPropertySelector(c *PropertySelectorContext)

	// ExitVariableSelector is called when exiting the variableSelector production.
	ExitVariableSelector(c *VariableSelectorContext)

	// ExitAllPropertiesSelector is called when exiting the allPropertiesSelector production.
	ExitAllPropertiesSelector(c *AllPropertiesSelectorContext)

	// ExitParameter is called when exiting the parameter production.
	ExitParameter(c *ParameterContext)

	// ExitLegacyParameter is called when exiting the legacyParameter production.
	ExitLegacyParameter(c *LegacyParameterContext)

	// ExitDollarParameter is called when exiting the dollarParameter production.
	ExitDollarParameter(c *DollarParameterContext)

	// ExitParameterName is called when exiting the parameterName production.
	ExitParameterName(c *ParameterNameContext)

	// ExitPropertyExpressions is called when exiting the propertyExpressions production.
	ExitPropertyExpressions(c *PropertyExpressionsContext)

	// ExitPropertyExpression is called when exiting the propertyExpression production.
	ExitPropertyExpression(c *PropertyExpressionContext)

	// ExitPropertyKeys is called when exiting the propertyKeys production.
	ExitPropertyKeys(c *PropertyKeysContext)

	// ExitPropertyKeyName is called when exiting the propertyKeyName production.
	ExitPropertyKeyName(c *PropertyKeyNameContext)

	// ExitIntegerLiteral is called when exiting the integerLiteral production.
	ExitIntegerLiteral(c *IntegerLiteralContext)

	// ExitDoubleLiteral is called when exiting the doubleLiteral production.
	ExitDoubleLiteral(c *DoubleLiteralContext)

	// ExitNamespace is called when exiting the namespace production.
	ExitNamespace(c *NamespaceContext)

	// ExitLeftArrowHead is called when exiting the leftArrowHead production.
	ExitLeftArrowHead(c *LeftArrowHeadContext)

	// ExitRightArrowHead is called when exiting the rightArrowHead production.
	ExitRightArrowHead(c *RightArrowHeadContext)

	// ExitDash is called when exiting the dash production.
	ExitDash(c *DashContext)

	// ExitSymbolicName is called when exiting the symbolicName production.
	ExitSymbolicName(c *SymbolicNameContext)

	// ExitKeyword is called when exiting the keyword production.
	ExitKeyword(c *KeywordContext)
}

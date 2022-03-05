// Generated from cypher-editor-support/src/_generated/Cypher.g4 by ANTLR 4.7.

package parser // Cypher

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseCypherListener is a complete listener for a parse tree produced by CypherParser.
type BaseCypherListener struct{}

var _ CypherListener = &BaseCypherListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseCypherListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseCypherListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseCypherListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseCypherListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterCypher is called when production cypher is entered.
func (s *BaseCypherListener) EnterCypher(ctx *CypherContext) {}

// ExitCypher is called when production cypher is exited.
func (s *BaseCypherListener) ExitCypher(ctx *CypherContext) {}

// EnterCypherPart is called when production cypherPart is entered.
func (s *BaseCypherListener) EnterCypherPart(ctx *CypherPartContext) {}

// ExitCypherPart is called when production cypherPart is exited.
func (s *BaseCypherListener) ExitCypherPart(ctx *CypherPartContext) {}

// EnterCypherConsoleCommand is called when production cypherConsoleCommand is entered.
func (s *BaseCypherListener) EnterCypherConsoleCommand(ctx *CypherConsoleCommandContext) {}

// ExitCypherConsoleCommand is called when production cypherConsoleCommand is exited.
func (s *BaseCypherListener) ExitCypherConsoleCommand(ctx *CypherConsoleCommandContext) {}

// EnterCypherConsoleCommandName is called when production cypherConsoleCommandName is entered.
func (s *BaseCypherListener) EnterCypherConsoleCommandName(ctx *CypherConsoleCommandNameContext) {}

// ExitCypherConsoleCommandName is called when production cypherConsoleCommandName is exited.
func (s *BaseCypherListener) ExitCypherConsoleCommandName(ctx *CypherConsoleCommandNameContext) {}

// EnterCypherConsoleCommandParameters is called when production cypherConsoleCommandParameters is entered.
func (s *BaseCypherListener) EnterCypherConsoleCommandParameters(ctx *CypherConsoleCommandParametersContext) {
}

// ExitCypherConsoleCommandParameters is called when production cypherConsoleCommandParameters is exited.
func (s *BaseCypherListener) ExitCypherConsoleCommandParameters(ctx *CypherConsoleCommandParametersContext) {
}

// EnterCypherConsoleCommandParameter is called when production cypherConsoleCommandParameter is entered.
func (s *BaseCypherListener) EnterCypherConsoleCommandParameter(ctx *CypherConsoleCommandParameterContext) {
}

// ExitCypherConsoleCommandParameter is called when production cypherConsoleCommandParameter is exited.
func (s *BaseCypherListener) ExitCypherConsoleCommandParameter(ctx *CypherConsoleCommandParameterContext) {
}

// EnterArrowExpression is called when production arrowExpression is entered.
func (s *BaseCypherListener) EnterArrowExpression(ctx *ArrowExpressionContext) {}

// ExitArrowExpression is called when production arrowExpression is exited.
func (s *BaseCypherListener) ExitArrowExpression(ctx *ArrowExpressionContext) {}

// EnterUrl is called when production url is entered.
func (s *BaseCypherListener) EnterUrl(ctx *UrlContext) {}

// ExitUrl is called when production url is exited.
func (s *BaseCypherListener) ExitUrl(ctx *UrlContext) {}

// EnterUri is called when production uri is entered.
func (s *BaseCypherListener) EnterUri(ctx *UriContext) {}

// ExitUri is called when production uri is exited.
func (s *BaseCypherListener) ExitUri(ctx *UriContext) {}

// EnterScheme is called when production scheme is entered.
func (s *BaseCypherListener) EnterScheme(ctx *SchemeContext) {}

// ExitScheme is called when production scheme is exited.
func (s *BaseCypherListener) ExitScheme(ctx *SchemeContext) {}

// EnterHost is called when production host is entered.
func (s *BaseCypherListener) EnterHost(ctx *HostContext) {}

// ExitHost is called when production host is exited.
func (s *BaseCypherListener) ExitHost(ctx *HostContext) {}

// EnterHostname is called when production hostname is entered.
func (s *BaseCypherListener) EnterHostname(ctx *HostnameContext) {}

// ExitHostname is called when production hostname is exited.
func (s *BaseCypherListener) ExitHostname(ctx *HostnameContext) {}

// EnterHostnumber is called when production hostnumber is entered.
func (s *BaseCypherListener) EnterHostnumber(ctx *HostnumberContext) {}

// ExitHostnumber is called when production hostnumber is exited.
func (s *BaseCypherListener) ExitHostnumber(ctx *HostnumberContext) {}

// EnterPort is called when production port is entered.
func (s *BaseCypherListener) EnterPort(ctx *PortContext) {}

// ExitPort is called when production port is exited.
func (s *BaseCypherListener) ExitPort(ctx *PortContext) {}

// EnterPath is called when production path is entered.
func (s *BaseCypherListener) EnterPath(ctx *PathContext) {}

// ExitPath is called when production path is exited.
func (s *BaseCypherListener) ExitPath(ctx *PathContext) {}

// EnterUser is called when production user is entered.
func (s *BaseCypherListener) EnterUser(ctx *UserContext) {}

// ExitUser is called when production user is exited.
func (s *BaseCypherListener) ExitUser(ctx *UserContext) {}

// EnterLogin is called when production login is entered.
func (s *BaseCypherListener) EnterLogin(ctx *LoginContext) {}

// ExitLogin is called when production login is exited.
func (s *BaseCypherListener) ExitLogin(ctx *LoginContext) {}

// EnterPassword is called when production password is entered.
func (s *BaseCypherListener) EnterPassword(ctx *PasswordContext) {}

// ExitPassword is called when production password is exited.
func (s *BaseCypherListener) ExitPassword(ctx *PasswordContext) {}

// EnterFrag is called when production frag is entered.
func (s *BaseCypherListener) EnterFrag(ctx *FragContext) {}

// ExitFrag is called when production frag is exited.
func (s *BaseCypherListener) ExitFrag(ctx *FragContext) {}

// EnterUrlQuery is called when production urlQuery is entered.
func (s *BaseCypherListener) EnterUrlQuery(ctx *UrlQueryContext) {}

// ExitUrlQuery is called when production urlQuery is exited.
func (s *BaseCypherListener) ExitUrlQuery(ctx *UrlQueryContext) {}

// EnterSearch is called when production search is entered.
func (s *BaseCypherListener) EnterSearch(ctx *SearchContext) {}

// ExitSearch is called when production search is exited.
func (s *BaseCypherListener) ExitSearch(ctx *SearchContext) {}

// EnterSearchparameter is called when production searchparameter is entered.
func (s *BaseCypherListener) EnterSearchparameter(ctx *SearchparameterContext) {}

// ExitSearchparameter is called when production searchparameter is exited.
func (s *BaseCypherListener) ExitSearchparameter(ctx *SearchparameterContext) {}

// EnterStr is called when production str is entered.
func (s *BaseCypherListener) EnterStr(ctx *StrContext) {}

// ExitStr is called when production str is exited.
func (s *BaseCypherListener) ExitStr(ctx *StrContext) {}

// EnterUrlDigits is called when production urlDigits is entered.
func (s *BaseCypherListener) EnterUrlDigits(ctx *UrlDigitsContext) {}

// ExitUrlDigits is called when production urlDigits is exited.
func (s *BaseCypherListener) ExitUrlDigits(ctx *UrlDigitsContext) {}

// EnterJson is called when production json is entered.
func (s *BaseCypherListener) EnterJson(ctx *JsonContext) {}

// ExitJson is called when production json is exited.
func (s *BaseCypherListener) ExitJson(ctx *JsonContext) {}

// EnterObj is called when production obj is entered.
func (s *BaseCypherListener) EnterObj(ctx *ObjContext) {}

// ExitObj is called when production obj is exited.
func (s *BaseCypherListener) ExitObj(ctx *ObjContext) {}

// EnterPair is called when production pair is entered.
func (s *BaseCypherListener) EnterPair(ctx *PairContext) {}

// ExitPair is called when production pair is exited.
func (s *BaseCypherListener) ExitPair(ctx *PairContext) {}

// EnterArray is called when production array is entered.
func (s *BaseCypherListener) EnterArray(ctx *ArrayContext) {}

// ExitArray is called when production array is exited.
func (s *BaseCypherListener) ExitArray(ctx *ArrayContext) {}

// EnterValue is called when production value is entered.
func (s *BaseCypherListener) EnterValue(ctx *ValueContext) {}

// ExitValue is called when production value is exited.
func (s *BaseCypherListener) ExitValue(ctx *ValueContext) {}

// EnterKeyValueLiteral is called when production keyValueLiteral is entered.
func (s *BaseCypherListener) EnterKeyValueLiteral(ctx *KeyValueLiteralContext) {}

// ExitKeyValueLiteral is called when production keyValueLiteral is exited.
func (s *BaseCypherListener) ExitKeyValueLiteral(ctx *KeyValueLiteralContext) {}

// EnterCommandPath is called when production commandPath is entered.
func (s *BaseCypherListener) EnterCommandPath(ctx *CommandPathContext) {}

// ExitCommandPath is called when production commandPath is exited.
func (s *BaseCypherListener) ExitCommandPath(ctx *CommandPathContext) {}

// EnterSubCommand is called when production subCommand is entered.
func (s *BaseCypherListener) EnterSubCommand(ctx *SubCommandContext) {}

// ExitSubCommand is called when production subCommand is exited.
func (s *BaseCypherListener) ExitSubCommand(ctx *SubCommandContext) {}

// EnterCypherQuery is called when production cypherQuery is entered.
func (s *BaseCypherListener) EnterCypherQuery(ctx *CypherQueryContext) {}

// ExitCypherQuery is called when production cypherQuery is exited.
func (s *BaseCypherListener) ExitCypherQuery(ctx *CypherQueryContext) {}

// EnterQueryOptions is called when production queryOptions is entered.
func (s *BaseCypherListener) EnterQueryOptions(ctx *QueryOptionsContext) {}

// ExitQueryOptions is called when production queryOptions is exited.
func (s *BaseCypherListener) ExitQueryOptions(ctx *QueryOptionsContext) {}

// EnterAnyCypherOption is called when production anyCypherOption is entered.
func (s *BaseCypherListener) EnterAnyCypherOption(ctx *AnyCypherOptionContext) {}

// ExitAnyCypherOption is called when production anyCypherOption is exited.
func (s *BaseCypherListener) ExitAnyCypherOption(ctx *AnyCypherOptionContext) {}

// EnterCypherOption is called when production cypherOption is entered.
func (s *BaseCypherListener) EnterCypherOption(ctx *CypherOptionContext) {}

// ExitCypherOption is called when production cypherOption is exited.
func (s *BaseCypherListener) ExitCypherOption(ctx *CypherOptionContext) {}

// EnterVersionNumber is called when production versionNumber is entered.
func (s *BaseCypherListener) EnterVersionNumber(ctx *VersionNumberContext) {}

// ExitVersionNumber is called when production versionNumber is exited.
func (s *BaseCypherListener) ExitVersionNumber(ctx *VersionNumberContext) {}

// EnterExplain is called when production explain is entered.
func (s *BaseCypherListener) EnterExplain(ctx *ExplainContext) {}

// ExitExplain is called when production explain is exited.
func (s *BaseCypherListener) ExitExplain(ctx *ExplainContext) {}

// EnterProfile is called when production profile is entered.
func (s *BaseCypherListener) EnterProfile(ctx *ProfileContext) {}

// ExitProfile is called when production profile is exited.
func (s *BaseCypherListener) ExitProfile(ctx *ProfileContext) {}

// EnterConfigurationOption is called when production configurationOption is entered.
func (s *BaseCypherListener) EnterConfigurationOption(ctx *ConfigurationOptionContext) {}

// ExitConfigurationOption is called when production configurationOption is exited.
func (s *BaseCypherListener) ExitConfigurationOption(ctx *ConfigurationOptionContext) {}

// EnterStatement is called when production statement is entered.
func (s *BaseCypherListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BaseCypherListener) ExitStatement(ctx *StatementContext) {}

// EnterQuery is called when production query is entered.
func (s *BaseCypherListener) EnterQuery(ctx *QueryContext) {}

// ExitQuery is called when production query is exited.
func (s *BaseCypherListener) ExitQuery(ctx *QueryContext) {}

// EnterRegularQuery is called when production regularQuery is entered.
func (s *BaseCypherListener) EnterRegularQuery(ctx *RegularQueryContext) {}

// ExitRegularQuery is called when production regularQuery is exited.
func (s *BaseCypherListener) ExitRegularQuery(ctx *RegularQueryContext) {}

// EnterBulkImportQuery is called when production bulkImportQuery is entered.
func (s *BaseCypherListener) EnterBulkImportQuery(ctx *BulkImportQueryContext) {}

// ExitBulkImportQuery is called when production bulkImportQuery is exited.
func (s *BaseCypherListener) ExitBulkImportQuery(ctx *BulkImportQueryContext) {}

// EnterSingleQuery is called when production singleQuery is entered.
func (s *BaseCypherListener) EnterSingleQuery(ctx *SingleQueryContext) {}

// ExitSingleQuery is called when production singleQuery is exited.
func (s *BaseCypherListener) ExitSingleQuery(ctx *SingleQueryContext) {}

// EnterPeriodicCommitHint is called when production periodicCommitHint is entered.
func (s *BaseCypherListener) EnterPeriodicCommitHint(ctx *PeriodicCommitHintContext) {}

// ExitPeriodicCommitHint is called when production periodicCommitHint is exited.
func (s *BaseCypherListener) ExitPeriodicCommitHint(ctx *PeriodicCommitHintContext) {}

// EnterLoadCSVQuery is called when production loadCSVQuery is entered.
func (s *BaseCypherListener) EnterLoadCSVQuery(ctx *LoadCSVQueryContext) {}

// ExitLoadCSVQuery is called when production loadCSVQuery is exited.
func (s *BaseCypherListener) ExitLoadCSVQuery(ctx *LoadCSVQueryContext) {}

// EnterUnion is called when production union is entered.
func (s *BaseCypherListener) EnterUnion(ctx *UnionContext) {}

// ExitUnion is called when production union is exited.
func (s *BaseCypherListener) ExitUnion(ctx *UnionContext) {}

// EnterClause is called when production clause is entered.
func (s *BaseCypherListener) EnterClause(ctx *ClauseContext) {}

// ExitClause is called when production clause is exited.
func (s *BaseCypherListener) ExitClause(ctx *ClauseContext) {}

// EnterCommand is called when production command is entered.
func (s *BaseCypherListener) EnterCommand(ctx *CommandContext) {}

// ExitCommand is called when production command is exited.
func (s *BaseCypherListener) ExitCommand(ctx *CommandContext) {}

// EnterSystemCommand is called when production systemCommand is entered.
func (s *BaseCypherListener) EnterSystemCommand(ctx *SystemCommandContext) {}

// ExitSystemCommand is called when production systemCommand is exited.
func (s *BaseCypherListener) ExitSystemCommand(ctx *SystemCommandContext) {}

// EnterMultidatabaseCommand is called when production multidatabaseCommand is entered.
func (s *BaseCypherListener) EnterMultidatabaseCommand(ctx *MultidatabaseCommandContext) {}

// ExitMultidatabaseCommand is called when production multidatabaseCommand is exited.
func (s *BaseCypherListener) ExitMultidatabaseCommand(ctx *MultidatabaseCommandContext) {}

// EnterUserCommand is called when production userCommand is entered.
func (s *BaseCypherListener) EnterUserCommand(ctx *UserCommandContext) {}

// ExitUserCommand is called when production userCommand is exited.
func (s *BaseCypherListener) ExitUserCommand(ctx *UserCommandContext) {}

// EnterPrivilegeCommand is called when production privilegeCommand is entered.
func (s *BaseCypherListener) EnterPrivilegeCommand(ctx *PrivilegeCommandContext) {}

// ExitPrivilegeCommand is called when production privilegeCommand is exited.
func (s *BaseCypherListener) ExitPrivilegeCommand(ctx *PrivilegeCommandContext) {}

// EnterShowRoles is called when production showRoles is entered.
func (s *BaseCypherListener) EnterShowRoles(ctx *ShowRolesContext) {}

// ExitShowRoles is called when production showRoles is exited.
func (s *BaseCypherListener) ExitShowRoles(ctx *ShowRolesContext) {}

// EnterCreateRole is called when production createRole is entered.
func (s *BaseCypherListener) EnterCreateRole(ctx *CreateRoleContext) {}

// ExitCreateRole is called when production createRole is exited.
func (s *BaseCypherListener) ExitCreateRole(ctx *CreateRoleContext) {}

// EnterCopyRole is called when production copyRole is entered.
func (s *BaseCypherListener) EnterCopyRole(ctx *CopyRoleContext) {}

// ExitCopyRole is called when production copyRole is exited.
func (s *BaseCypherListener) ExitCopyRole(ctx *CopyRoleContext) {}

// EnterDropRole is called when production dropRole is entered.
func (s *BaseCypherListener) EnterDropRole(ctx *DropRoleContext) {}

// ExitDropRole is called when production dropRole is exited.
func (s *BaseCypherListener) ExitDropRole(ctx *DropRoleContext) {}

// EnterShowUsers is called when production showUsers is entered.
func (s *BaseCypherListener) EnterShowUsers(ctx *ShowUsersContext) {}

// ExitShowUsers is called when production showUsers is exited.
func (s *BaseCypherListener) ExitShowUsers(ctx *ShowUsersContext) {}

// EnterCreateUser is called when production createUser is entered.
func (s *BaseCypherListener) EnterCreateUser(ctx *CreateUserContext) {}

// ExitCreateUser is called when production createUser is exited.
func (s *BaseCypherListener) ExitCreateUser(ctx *CreateUserContext) {}

// EnterDropUser is called when production dropUser is entered.
func (s *BaseCypherListener) EnterDropUser(ctx *DropUserContext) {}

// ExitDropUser is called when production dropUser is exited.
func (s *BaseCypherListener) ExitDropUser(ctx *DropUserContext) {}

// EnterAlterUser is called when production alterUser is entered.
func (s *BaseCypherListener) EnterAlterUser(ctx *AlterUserContext) {}

// ExitAlterUser is called when production alterUser is exited.
func (s *BaseCypherListener) ExitAlterUser(ctx *AlterUserContext) {}

// EnterShowPrivileges is called when production showPrivileges is entered.
func (s *BaseCypherListener) EnterShowPrivileges(ctx *ShowPrivilegesContext) {}

// ExitShowPrivileges is called when production showPrivileges is exited.
func (s *BaseCypherListener) ExitShowPrivileges(ctx *ShowPrivilegesContext) {}

// EnterGrantPrivilege is called when production grantPrivilege is entered.
func (s *BaseCypherListener) EnterGrantPrivilege(ctx *GrantPrivilegeContext) {}

// ExitGrantPrivilege is called when production grantPrivilege is exited.
func (s *BaseCypherListener) ExitGrantPrivilege(ctx *GrantPrivilegeContext) {}

// EnterDenyPrivilege is called when production denyPrivilege is entered.
func (s *BaseCypherListener) EnterDenyPrivilege(ctx *DenyPrivilegeContext) {}

// ExitDenyPrivilege is called when production denyPrivilege is exited.
func (s *BaseCypherListener) ExitDenyPrivilege(ctx *DenyPrivilegeContext) {}

// EnterRevokePrivilege is called when production revokePrivilege is entered.
func (s *BaseCypherListener) EnterRevokePrivilege(ctx *RevokePrivilegeContext) {}

// ExitRevokePrivilege is called when production revokePrivilege is exited.
func (s *BaseCypherListener) ExitRevokePrivilege(ctx *RevokePrivilegeContext) {}

// EnterRevokePart is called when production revokePart is entered.
func (s *BaseCypherListener) EnterRevokePart(ctx *RevokePartContext) {}

// ExitRevokePart is called when production revokePart is exited.
func (s *BaseCypherListener) ExitRevokePart(ctx *RevokePartContext) {}

// EnterDatabaseScope is called when production databaseScope is entered.
func (s *BaseCypherListener) EnterDatabaseScope(ctx *DatabaseScopeContext) {}

// ExitDatabaseScope is called when production databaseScope is exited.
func (s *BaseCypherListener) ExitDatabaseScope(ctx *DatabaseScopeContext) {}

// EnterGraphScope is called when production graphScope is entered.
func (s *BaseCypherListener) EnterGraphScope(ctx *GraphScopeContext) {}

// ExitGraphScope is called when production graphScope is exited.
func (s *BaseCypherListener) ExitGraphScope(ctx *GraphScopeContext) {}

// EnterRoles is called when production roles is entered.
func (s *BaseCypherListener) EnterRoles(ctx *RolesContext) {}

// ExitRoles is called when production roles is exited.
func (s *BaseCypherListener) ExitRoles(ctx *RolesContext) {}

// EnterGrantableGraphPrivileges is called when production grantableGraphPrivileges is entered.
func (s *BaseCypherListener) EnterGrantableGraphPrivileges(ctx *GrantableGraphPrivilegesContext) {}

// ExitGrantableGraphPrivileges is called when production grantableGraphPrivileges is exited.
func (s *BaseCypherListener) ExitGrantableGraphPrivileges(ctx *GrantableGraphPrivilegesContext) {}

// EnterRevokeableGraphPrivileges is called when production revokeableGraphPrivileges is entered.
func (s *BaseCypherListener) EnterRevokeableGraphPrivileges(ctx *RevokeableGraphPrivilegesContext) {}

// ExitRevokeableGraphPrivileges is called when production revokeableGraphPrivileges is exited.
func (s *BaseCypherListener) ExitRevokeableGraphPrivileges(ctx *RevokeableGraphPrivilegesContext) {}

// EnterDatasbasePrivilege is called when production datasbasePrivilege is entered.
func (s *BaseCypherListener) EnterDatasbasePrivilege(ctx *DatasbasePrivilegeContext) {}

// ExitDatasbasePrivilege is called when production datasbasePrivilege is exited.
func (s *BaseCypherListener) ExitDatasbasePrivilege(ctx *DatasbasePrivilegeContext) {}

// EnterDbmsPrivilege is called when production dbmsPrivilege is entered.
func (s *BaseCypherListener) EnterDbmsPrivilege(ctx *DbmsPrivilegeContext) {}

// ExitDbmsPrivilege is called when production dbmsPrivilege is exited.
func (s *BaseCypherListener) ExitDbmsPrivilege(ctx *DbmsPrivilegeContext) {}

// EnterElementScope is called when production elementScope is entered.
func (s *BaseCypherListener) EnterElementScope(ctx *ElementScopeContext) {}

// ExitElementScope is called when production elementScope is exited.
func (s *BaseCypherListener) ExitElementScope(ctx *ElementScopeContext) {}

// EnterPropertiesList is called when production propertiesList is entered.
func (s *BaseCypherListener) EnterPropertiesList(ctx *PropertiesListContext) {}

// ExitPropertiesList is called when production propertiesList is exited.
func (s *BaseCypherListener) ExitPropertiesList(ctx *PropertiesListContext) {}

// EnterPropertyScope is called when production propertyScope is entered.
func (s *BaseCypherListener) EnterPropertyScope(ctx *PropertyScopeContext) {}

// ExitPropertyScope is called when production propertyScope is exited.
func (s *BaseCypherListener) ExitPropertyScope(ctx *PropertyScopeContext) {}

// EnterShowDatabase is called when production showDatabase is entered.
func (s *BaseCypherListener) EnterShowDatabase(ctx *ShowDatabaseContext) {}

// ExitShowDatabase is called when production showDatabase is exited.
func (s *BaseCypherListener) ExitShowDatabase(ctx *ShowDatabaseContext) {}

// EnterCreateDatabase is called when production createDatabase is entered.
func (s *BaseCypherListener) EnterCreateDatabase(ctx *CreateDatabaseContext) {}

// ExitCreateDatabase is called when production createDatabase is exited.
func (s *BaseCypherListener) ExitCreateDatabase(ctx *CreateDatabaseContext) {}

// EnterDropDatabase is called when production dropDatabase is entered.
func (s *BaseCypherListener) EnterDropDatabase(ctx *DropDatabaseContext) {}

// ExitDropDatabase is called when production dropDatabase is exited.
func (s *BaseCypherListener) ExitDropDatabase(ctx *DropDatabaseContext) {}

// EnterStartDatabase is called when production startDatabase is entered.
func (s *BaseCypherListener) EnterStartDatabase(ctx *StartDatabaseContext) {}

// ExitStartDatabase is called when production startDatabase is exited.
func (s *BaseCypherListener) ExitStartDatabase(ctx *StartDatabaseContext) {}

// EnterStopDatabase is called when production stopDatabase is entered.
func (s *BaseCypherListener) EnterStopDatabase(ctx *StopDatabaseContext) {}

// ExitStopDatabase is called when production stopDatabase is exited.
func (s *BaseCypherListener) ExitStopDatabase(ctx *StopDatabaseContext) {}

// EnterIfNotExists is called when production ifNotExists is entered.
func (s *BaseCypherListener) EnterIfNotExists(ctx *IfNotExistsContext) {}

// ExitIfNotExists is called when production ifNotExists is exited.
func (s *BaseCypherListener) ExitIfNotExists(ctx *IfNotExistsContext) {}

// EnterIfExists is called when production ifExists is entered.
func (s *BaseCypherListener) EnterIfExists(ctx *IfExistsContext) {}

// ExitIfExists is called when production ifExists is exited.
func (s *BaseCypherListener) ExitIfExists(ctx *IfExistsContext) {}

// EnterOrReplace is called when production orReplace is entered.
func (s *BaseCypherListener) EnterOrReplace(ctx *OrReplaceContext) {}

// ExitOrReplace is called when production orReplace is exited.
func (s *BaseCypherListener) ExitOrReplace(ctx *OrReplaceContext) {}

// EnterSetPassword is called when production setPassword is entered.
func (s *BaseCypherListener) EnterSetPassword(ctx *SetPasswordContext) {}

// ExitSetPassword is called when production setPassword is exited.
func (s *BaseCypherListener) ExitSetPassword(ctx *SetPasswordContext) {}

// EnterPasswordStatus is called when production passwordStatus is entered.
func (s *BaseCypherListener) EnterPasswordStatus(ctx *PasswordStatusContext) {}

// ExitPasswordStatus is called when production passwordStatus is exited.
func (s *BaseCypherListener) ExitPasswordStatus(ctx *PasswordStatusContext) {}

// EnterSetStatus is called when production setStatus is entered.
func (s *BaseCypherListener) EnterSetStatus(ctx *SetStatusContext) {}

// ExitSetStatus is called when production setStatus is exited.
func (s *BaseCypherListener) ExitSetStatus(ctx *SetStatusContext) {}

// EnterUserStatus is called when production userStatus is entered.
func (s *BaseCypherListener) EnterUserStatus(ctx *UserStatusContext) {}

// ExitUserStatus is called when production userStatus is exited.
func (s *BaseCypherListener) ExitUserStatus(ctx *UserStatusContext) {}

// EnterCreateUniqueConstraint is called when production createUniqueConstraint is entered.
func (s *BaseCypherListener) EnterCreateUniqueConstraint(ctx *CreateUniqueConstraintContext) {}

// ExitCreateUniqueConstraint is called when production createUniqueConstraint is exited.
func (s *BaseCypherListener) ExitCreateUniqueConstraint(ctx *CreateUniqueConstraintContext) {}

// EnterCreateNodeKeyConstraint is called when production createNodeKeyConstraint is entered.
func (s *BaseCypherListener) EnterCreateNodeKeyConstraint(ctx *CreateNodeKeyConstraintContext) {}

// ExitCreateNodeKeyConstraint is called when production createNodeKeyConstraint is exited.
func (s *BaseCypherListener) ExitCreateNodeKeyConstraint(ctx *CreateNodeKeyConstraintContext) {}

// EnterCreateNodePropertyExistenceConstraint is called when production createNodePropertyExistenceConstraint is entered.
func (s *BaseCypherListener) EnterCreateNodePropertyExistenceConstraint(ctx *CreateNodePropertyExistenceConstraintContext) {
}

// ExitCreateNodePropertyExistenceConstraint is called when production createNodePropertyExistenceConstraint is exited.
func (s *BaseCypherListener) ExitCreateNodePropertyExistenceConstraint(ctx *CreateNodePropertyExistenceConstraintContext) {
}

// EnterCreateRelationshipPropertyExistenceConstraint is called when production createRelationshipPropertyExistenceConstraint is entered.
func (s *BaseCypherListener) EnterCreateRelationshipPropertyExistenceConstraint(ctx *CreateRelationshipPropertyExistenceConstraintContext) {
}

// ExitCreateRelationshipPropertyExistenceConstraint is called when production createRelationshipPropertyExistenceConstraint is exited.
func (s *BaseCypherListener) ExitCreateRelationshipPropertyExistenceConstraint(ctx *CreateRelationshipPropertyExistenceConstraintContext) {
}

// EnterCreateIndex is called when production createIndex is entered.
func (s *BaseCypherListener) EnterCreateIndex(ctx *CreateIndexContext) {}

// ExitCreateIndex is called when production createIndex is exited.
func (s *BaseCypherListener) ExitCreateIndex(ctx *CreateIndexContext) {}

// EnterDropUniqueConstraint is called when production dropUniqueConstraint is entered.
func (s *BaseCypherListener) EnterDropUniqueConstraint(ctx *DropUniqueConstraintContext) {}

// ExitDropUniqueConstraint is called when production dropUniqueConstraint is exited.
func (s *BaseCypherListener) ExitDropUniqueConstraint(ctx *DropUniqueConstraintContext) {}

// EnterDropNodeKeyConstraint is called when production dropNodeKeyConstraint is entered.
func (s *BaseCypherListener) EnterDropNodeKeyConstraint(ctx *DropNodeKeyConstraintContext) {}

// ExitDropNodeKeyConstraint is called when production dropNodeKeyConstraint is exited.
func (s *BaseCypherListener) ExitDropNodeKeyConstraint(ctx *DropNodeKeyConstraintContext) {}

// EnterDropNodePropertyExistenceConstraint is called when production dropNodePropertyExistenceConstraint is entered.
func (s *BaseCypherListener) EnterDropNodePropertyExistenceConstraint(ctx *DropNodePropertyExistenceConstraintContext) {
}

// ExitDropNodePropertyExistenceConstraint is called when production dropNodePropertyExistenceConstraint is exited.
func (s *BaseCypherListener) ExitDropNodePropertyExistenceConstraint(ctx *DropNodePropertyExistenceConstraintContext) {
}

// EnterDropRelationshipPropertyExistenceConstraint is called when production dropRelationshipPropertyExistenceConstraint is entered.
func (s *BaseCypherListener) EnterDropRelationshipPropertyExistenceConstraint(ctx *DropRelationshipPropertyExistenceConstraintContext) {
}

// ExitDropRelationshipPropertyExistenceConstraint is called when production dropRelationshipPropertyExistenceConstraint is exited.
func (s *BaseCypherListener) ExitDropRelationshipPropertyExistenceConstraint(ctx *DropRelationshipPropertyExistenceConstraintContext) {
}

// EnterDropIndex is called when production dropIndex is entered.
func (s *BaseCypherListener) EnterDropIndex(ctx *DropIndexContext) {}

// ExitDropIndex is called when production dropIndex is exited.
func (s *BaseCypherListener) ExitDropIndex(ctx *DropIndexContext) {}

// EnterIndex is called when production index is entered.
func (s *BaseCypherListener) EnterIndex(ctx *IndexContext) {}

// ExitIndex is called when production index is exited.
func (s *BaseCypherListener) ExitIndex(ctx *IndexContext) {}

// EnterUniqueConstraint is called when production uniqueConstraint is entered.
func (s *BaseCypherListener) EnterUniqueConstraint(ctx *UniqueConstraintContext) {}

// ExitUniqueConstraint is called when production uniqueConstraint is exited.
func (s *BaseCypherListener) ExitUniqueConstraint(ctx *UniqueConstraintContext) {}

// EnterNodeKeyConstraint is called when production nodeKeyConstraint is entered.
func (s *BaseCypherListener) EnterNodeKeyConstraint(ctx *NodeKeyConstraintContext) {}

// ExitNodeKeyConstraint is called when production nodeKeyConstraint is exited.
func (s *BaseCypherListener) ExitNodeKeyConstraint(ctx *NodeKeyConstraintContext) {}

// EnterNodePropertyExistenceConstraint is called when production nodePropertyExistenceConstraint is entered.
func (s *BaseCypherListener) EnterNodePropertyExistenceConstraint(ctx *NodePropertyExistenceConstraintContext) {
}

// ExitNodePropertyExistenceConstraint is called when production nodePropertyExistenceConstraint is exited.
func (s *BaseCypherListener) ExitNodePropertyExistenceConstraint(ctx *NodePropertyExistenceConstraintContext) {
}

// EnterRelationshipPropertyExistenceConstraint is called when production relationshipPropertyExistenceConstraint is entered.
func (s *BaseCypherListener) EnterRelationshipPropertyExistenceConstraint(ctx *RelationshipPropertyExistenceConstraintContext) {
}

// ExitRelationshipPropertyExistenceConstraint is called when production relationshipPropertyExistenceConstraint is exited.
func (s *BaseCypherListener) ExitRelationshipPropertyExistenceConstraint(ctx *RelationshipPropertyExistenceConstraintContext) {
}

// EnterRelationshipPatternSyntax is called when production relationshipPatternSyntax is entered.
func (s *BaseCypherListener) EnterRelationshipPatternSyntax(ctx *RelationshipPatternSyntaxContext) {}

// ExitRelationshipPatternSyntax is called when production relationshipPatternSyntax is exited.
func (s *BaseCypherListener) ExitRelationshipPatternSyntax(ctx *RelationshipPatternSyntaxContext) {}

// EnterLoadCSVClause is called when production loadCSVClause is entered.
func (s *BaseCypherListener) EnterLoadCSVClause(ctx *LoadCSVClauseContext) {}

// ExitLoadCSVClause is called when production loadCSVClause is exited.
func (s *BaseCypherListener) ExitLoadCSVClause(ctx *LoadCSVClauseContext) {}

// EnterMatchClause is called when production matchClause is entered.
func (s *BaseCypherListener) EnterMatchClause(ctx *MatchClauseContext) {}

// ExitMatchClause is called when production matchClause is exited.
func (s *BaseCypherListener) ExitMatchClause(ctx *MatchClauseContext) {}

// EnterUnwindClause is called when production unwindClause is entered.
func (s *BaseCypherListener) EnterUnwindClause(ctx *UnwindClauseContext) {}

// ExitUnwindClause is called when production unwindClause is exited.
func (s *BaseCypherListener) ExitUnwindClause(ctx *UnwindClauseContext) {}

// EnterMergeClause is called when production mergeClause is entered.
func (s *BaseCypherListener) EnterMergeClause(ctx *MergeClauseContext) {}

// ExitMergeClause is called when production mergeClause is exited.
func (s *BaseCypherListener) ExitMergeClause(ctx *MergeClauseContext) {}

// EnterMergeAction is called when production mergeAction is entered.
func (s *BaseCypherListener) EnterMergeAction(ctx *MergeActionContext) {}

// ExitMergeAction is called when production mergeAction is exited.
func (s *BaseCypherListener) ExitMergeAction(ctx *MergeActionContext) {}

// EnterCreateClause is called when production createClause is entered.
func (s *BaseCypherListener) EnterCreateClause(ctx *CreateClauseContext) {}

// ExitCreateClause is called when production createClause is exited.
func (s *BaseCypherListener) ExitCreateClause(ctx *CreateClauseContext) {}

// EnterCreateUniqueClause is called when production createUniqueClause is entered.
func (s *BaseCypherListener) EnterCreateUniqueClause(ctx *CreateUniqueClauseContext) {}

// ExitCreateUniqueClause is called when production createUniqueClause is exited.
func (s *BaseCypherListener) ExitCreateUniqueClause(ctx *CreateUniqueClauseContext) {}

// EnterSetClause is called when production setClause is entered.
func (s *BaseCypherListener) EnterSetClause(ctx *SetClauseContext) {}

// ExitSetClause is called when production setClause is exited.
func (s *BaseCypherListener) ExitSetClause(ctx *SetClauseContext) {}

// EnterSetItem is called when production setItem is entered.
func (s *BaseCypherListener) EnterSetItem(ctx *SetItemContext) {}

// ExitSetItem is called when production setItem is exited.
func (s *BaseCypherListener) ExitSetItem(ctx *SetItemContext) {}

// EnterDeleteClause is called when production deleteClause is entered.
func (s *BaseCypherListener) EnterDeleteClause(ctx *DeleteClauseContext) {}

// ExitDeleteClause is called when production deleteClause is exited.
func (s *BaseCypherListener) ExitDeleteClause(ctx *DeleteClauseContext) {}

// EnterRemoveClause is called when production removeClause is entered.
func (s *BaseCypherListener) EnterRemoveClause(ctx *RemoveClauseContext) {}

// ExitRemoveClause is called when production removeClause is exited.
func (s *BaseCypherListener) ExitRemoveClause(ctx *RemoveClauseContext) {}

// EnterRemoveItem is called when production removeItem is entered.
func (s *BaseCypherListener) EnterRemoveItem(ctx *RemoveItemContext) {}

// ExitRemoveItem is called when production removeItem is exited.
func (s *BaseCypherListener) ExitRemoveItem(ctx *RemoveItemContext) {}

// EnterForeachClause is called when production foreachClause is entered.
func (s *BaseCypherListener) EnterForeachClause(ctx *ForeachClauseContext) {}

// ExitForeachClause is called when production foreachClause is exited.
func (s *BaseCypherListener) ExitForeachClause(ctx *ForeachClauseContext) {}

// EnterWithClause is called when production withClause is entered.
func (s *BaseCypherListener) EnterWithClause(ctx *WithClauseContext) {}

// ExitWithClause is called when production withClause is exited.
func (s *BaseCypherListener) ExitWithClause(ctx *WithClauseContext) {}

// EnterReturnClause is called when production returnClause is entered.
func (s *BaseCypherListener) EnterReturnClause(ctx *ReturnClauseContext) {}

// ExitReturnClause is called when production returnClause is exited.
func (s *BaseCypherListener) ExitReturnClause(ctx *ReturnClauseContext) {}

// EnterReturnBody is called when production returnBody is entered.
func (s *BaseCypherListener) EnterReturnBody(ctx *ReturnBodyContext) {}

// ExitReturnBody is called when production returnBody is exited.
func (s *BaseCypherListener) ExitReturnBody(ctx *ReturnBodyContext) {}

// EnterFunction is called when production function is entered.
func (s *BaseCypherListener) EnterFunction(ctx *FunctionContext) {}

// ExitFunction is called when production function is exited.
func (s *BaseCypherListener) ExitFunction(ctx *FunctionContext) {}

// EnterReturnItems is called when production returnItems is entered.
func (s *BaseCypherListener) EnterReturnItems(ctx *ReturnItemsContext) {}

// ExitReturnItems is called when production returnItems is exited.
func (s *BaseCypherListener) ExitReturnItems(ctx *ReturnItemsContext) {}

// EnterReturnItem is called when production returnItem is entered.
func (s *BaseCypherListener) EnterReturnItem(ctx *ReturnItemContext) {}

// ExitReturnItem is called when production returnItem is exited.
func (s *BaseCypherListener) ExitReturnItem(ctx *ReturnItemContext) {}

// EnterCall is called when production call is entered.
func (s *BaseCypherListener) EnterCall(ctx *CallContext) {}

// ExitCall is called when production call is exited.
func (s *BaseCypherListener) ExitCall(ctx *CallContext) {}

// EnterProcedureInvocation is called when production procedureInvocation is entered.
func (s *BaseCypherListener) EnterProcedureInvocation(ctx *ProcedureInvocationContext) {}

// ExitProcedureInvocation is called when production procedureInvocation is exited.
func (s *BaseCypherListener) ExitProcedureInvocation(ctx *ProcedureInvocationContext) {}

// EnterProcedureInvocationBody is called when production procedureInvocationBody is entered.
func (s *BaseCypherListener) EnterProcedureInvocationBody(ctx *ProcedureInvocationBodyContext) {}

// ExitProcedureInvocationBody is called when production procedureInvocationBody is exited.
func (s *BaseCypherListener) ExitProcedureInvocationBody(ctx *ProcedureInvocationBodyContext) {}

// EnterProcedureArguments is called when production procedureArguments is entered.
func (s *BaseCypherListener) EnterProcedureArguments(ctx *ProcedureArgumentsContext) {}

// ExitProcedureArguments is called when production procedureArguments is exited.
func (s *BaseCypherListener) ExitProcedureArguments(ctx *ProcedureArgumentsContext) {}

// EnterProcedureResults is called when production procedureResults is entered.
func (s *BaseCypherListener) EnterProcedureResults(ctx *ProcedureResultsContext) {}

// ExitProcedureResults is called when production procedureResults is exited.
func (s *BaseCypherListener) ExitProcedureResults(ctx *ProcedureResultsContext) {}

// EnterProcedureResult is called when production procedureResult is entered.
func (s *BaseCypherListener) EnterProcedureResult(ctx *ProcedureResultContext) {}

// ExitProcedureResult is called when production procedureResult is exited.
func (s *BaseCypherListener) ExitProcedureResult(ctx *ProcedureResultContext) {}

// EnterAliasedProcedureResult is called when production aliasedProcedureResult is entered.
func (s *BaseCypherListener) EnterAliasedProcedureResult(ctx *AliasedProcedureResultContext) {}

// ExitAliasedProcedureResult is called when production aliasedProcedureResult is exited.
func (s *BaseCypherListener) ExitAliasedProcedureResult(ctx *AliasedProcedureResultContext) {}

// EnterSimpleProcedureResult is called when production simpleProcedureResult is entered.
func (s *BaseCypherListener) EnterSimpleProcedureResult(ctx *SimpleProcedureResultContext) {}

// ExitSimpleProcedureResult is called when production simpleProcedureResult is exited.
func (s *BaseCypherListener) ExitSimpleProcedureResult(ctx *SimpleProcedureResultContext) {}

// EnterProcedureOutput is called when production procedureOutput is entered.
func (s *BaseCypherListener) EnterProcedureOutput(ctx *ProcedureOutputContext) {}

// ExitProcedureOutput is called when production procedureOutput is exited.
func (s *BaseCypherListener) ExitProcedureOutput(ctx *ProcedureOutputContext) {}

// EnterOrder is called when production order is entered.
func (s *BaseCypherListener) EnterOrder(ctx *OrderContext) {}

// ExitOrder is called when production order is exited.
func (s *BaseCypherListener) ExitOrder(ctx *OrderContext) {}

// EnterSkip is called when production skip is entered.
func (s *BaseCypherListener) EnterSkip(ctx *SkipContext) {}

// ExitSkip is called when production skip is exited.
func (s *BaseCypherListener) ExitSkip(ctx *SkipContext) {}

// EnterLimit is called when production limit is entered.
func (s *BaseCypherListener) EnterLimit(ctx *LimitContext) {}

// ExitLimit is called when production limit is exited.
func (s *BaseCypherListener) ExitLimit(ctx *LimitContext) {}

// EnterSortItem is called when production sortItem is entered.
func (s *BaseCypherListener) EnterSortItem(ctx *SortItemContext) {}

// ExitSortItem is called when production sortItem is exited.
func (s *BaseCypherListener) ExitSortItem(ctx *SortItemContext) {}

// EnterHint is called when production hint is entered.
func (s *BaseCypherListener) EnterHint(ctx *HintContext) {}

// ExitHint is called when production hint is exited.
func (s *BaseCypherListener) ExitHint(ctx *HintContext) {}

// EnterStartClause is called when production startClause is entered.
func (s *BaseCypherListener) EnterStartClause(ctx *StartClauseContext) {}

// ExitStartClause is called when production startClause is exited.
func (s *BaseCypherListener) ExitStartClause(ctx *StartClauseContext) {}

// EnterStartPoint is called when production startPoint is entered.
func (s *BaseCypherListener) EnterStartPoint(ctx *StartPointContext) {}

// ExitStartPoint is called when production startPoint is exited.
func (s *BaseCypherListener) ExitStartPoint(ctx *StartPointContext) {}

// EnterLookup is called when production lookup is entered.
func (s *BaseCypherListener) EnterLookup(ctx *LookupContext) {}

// ExitLookup is called when production lookup is exited.
func (s *BaseCypherListener) ExitLookup(ctx *LookupContext) {}

// EnterNodeLookup is called when production nodeLookup is entered.
func (s *BaseCypherListener) EnterNodeLookup(ctx *NodeLookupContext) {}

// ExitNodeLookup is called when production nodeLookup is exited.
func (s *BaseCypherListener) ExitNodeLookup(ctx *NodeLookupContext) {}

// EnterRelationshipLookup is called when production relationshipLookup is entered.
func (s *BaseCypherListener) EnterRelationshipLookup(ctx *RelationshipLookupContext) {}

// ExitRelationshipLookup is called when production relationshipLookup is exited.
func (s *BaseCypherListener) ExitRelationshipLookup(ctx *RelationshipLookupContext) {}

// EnterIdentifiedIndexLookup is called when production identifiedIndexLookup is entered.
func (s *BaseCypherListener) EnterIdentifiedIndexLookup(ctx *IdentifiedIndexLookupContext) {}

// ExitIdentifiedIndexLookup is called when production identifiedIndexLookup is exited.
func (s *BaseCypherListener) ExitIdentifiedIndexLookup(ctx *IdentifiedIndexLookupContext) {}

// EnterIndexQuery is called when production indexQuery is entered.
func (s *BaseCypherListener) EnterIndexQuery(ctx *IndexQueryContext) {}

// ExitIndexQuery is called when production indexQuery is exited.
func (s *BaseCypherListener) ExitIndexQuery(ctx *IndexQueryContext) {}

// EnterIdLookup is called when production idLookup is entered.
func (s *BaseCypherListener) EnterIdLookup(ctx *IdLookupContext) {}

// ExitIdLookup is called when production idLookup is exited.
func (s *BaseCypherListener) ExitIdLookup(ctx *IdLookupContext) {}

// EnterLiteralIds is called when production literalIds is entered.
func (s *BaseCypherListener) EnterLiteralIds(ctx *LiteralIdsContext) {}

// ExitLiteralIds is called when production literalIds is exited.
func (s *BaseCypherListener) ExitLiteralIds(ctx *LiteralIdsContext) {}

// EnterWhere is called when production where is entered.
func (s *BaseCypherListener) EnterWhere(ctx *WhereContext) {}

// ExitWhere is called when production where is exited.
func (s *BaseCypherListener) ExitWhere(ctx *WhereContext) {}

// EnterPattern is called when production pattern is entered.
func (s *BaseCypherListener) EnterPattern(ctx *PatternContext) {}

// ExitPattern is called when production pattern is exited.
func (s *BaseCypherListener) ExitPattern(ctx *PatternContext) {}

// EnterPatternPart is called when production patternPart is entered.
func (s *BaseCypherListener) EnterPatternPart(ctx *PatternPartContext) {}

// ExitPatternPart is called when production patternPart is exited.
func (s *BaseCypherListener) ExitPatternPart(ctx *PatternPartContext) {}

// EnterAnonymousPatternPart is called when production anonymousPatternPart is entered.
func (s *BaseCypherListener) EnterAnonymousPatternPart(ctx *AnonymousPatternPartContext) {}

// ExitAnonymousPatternPart is called when production anonymousPatternPart is exited.
func (s *BaseCypherListener) ExitAnonymousPatternPart(ctx *AnonymousPatternPartContext) {}

// EnterPatternElement is called when production patternElement is entered.
func (s *BaseCypherListener) EnterPatternElement(ctx *PatternElementContext) {}

// ExitPatternElement is called when production patternElement is exited.
func (s *BaseCypherListener) ExitPatternElement(ctx *PatternElementContext) {}

// EnterNodePattern is called when production nodePattern is entered.
func (s *BaseCypherListener) EnterNodePattern(ctx *NodePatternContext) {}

// ExitNodePattern is called when production nodePattern is exited.
func (s *BaseCypherListener) ExitNodePattern(ctx *NodePatternContext) {}

// EnterPatternElementChain is called when production patternElementChain is entered.
func (s *BaseCypherListener) EnterPatternElementChain(ctx *PatternElementChainContext) {}

// ExitPatternElementChain is called when production patternElementChain is exited.
func (s *BaseCypherListener) ExitPatternElementChain(ctx *PatternElementChainContext) {}

// EnterRelationshipPattern is called when production relationshipPattern is entered.
func (s *BaseCypherListener) EnterRelationshipPattern(ctx *RelationshipPatternContext) {}

// ExitRelationshipPattern is called when production relationshipPattern is exited.
func (s *BaseCypherListener) ExitRelationshipPattern(ctx *RelationshipPatternContext) {}

// EnterRelationshipPatternStart is called when production relationshipPatternStart is entered.
func (s *BaseCypherListener) EnterRelationshipPatternStart(ctx *RelationshipPatternStartContext) {}

// ExitRelationshipPatternStart is called when production relationshipPatternStart is exited.
func (s *BaseCypherListener) ExitRelationshipPatternStart(ctx *RelationshipPatternStartContext) {}

// EnterRelationshipPatternEnd is called when production relationshipPatternEnd is entered.
func (s *BaseCypherListener) EnterRelationshipPatternEnd(ctx *RelationshipPatternEndContext) {}

// ExitRelationshipPatternEnd is called when production relationshipPatternEnd is exited.
func (s *BaseCypherListener) ExitRelationshipPatternEnd(ctx *RelationshipPatternEndContext) {}

// EnterRelationshipDetail is called when production relationshipDetail is entered.
func (s *BaseCypherListener) EnterRelationshipDetail(ctx *RelationshipDetailContext) {}

// ExitRelationshipDetail is called when production relationshipDetail is exited.
func (s *BaseCypherListener) ExitRelationshipDetail(ctx *RelationshipDetailContext) {}

// EnterProperties is called when production properties is entered.
func (s *BaseCypherListener) EnterProperties(ctx *PropertiesContext) {}

// ExitProperties is called when production properties is exited.
func (s *BaseCypherListener) ExitProperties(ctx *PropertiesContext) {}

// EnterRelType is called when production relType is entered.
func (s *BaseCypherListener) EnterRelType(ctx *RelTypeContext) {}

// ExitRelType is called when production relType is exited.
func (s *BaseCypherListener) ExitRelType(ctx *RelTypeContext) {}

// EnterRelationshipTypes is called when production relationshipTypes is entered.
func (s *BaseCypherListener) EnterRelationshipTypes(ctx *RelationshipTypesContext) {}

// ExitRelationshipTypes is called when production relationshipTypes is exited.
func (s *BaseCypherListener) ExitRelationshipTypes(ctx *RelationshipTypesContext) {}

// EnterRelationshipType is called when production relationshipType is entered.
func (s *BaseCypherListener) EnterRelationshipType(ctx *RelationshipTypeContext) {}

// ExitRelationshipType is called when production relationshipType is exited.
func (s *BaseCypherListener) ExitRelationshipType(ctx *RelationshipTypeContext) {}

// EnterRelationshipTypeOptionalColon is called when production relationshipTypeOptionalColon is entered.
func (s *BaseCypherListener) EnterRelationshipTypeOptionalColon(ctx *RelationshipTypeOptionalColonContext) {
}

// ExitRelationshipTypeOptionalColon is called when production relationshipTypeOptionalColon is exited.
func (s *BaseCypherListener) ExitRelationshipTypeOptionalColon(ctx *RelationshipTypeOptionalColonContext) {
}

// EnterNodeLabels is called when production nodeLabels is entered.
func (s *BaseCypherListener) EnterNodeLabels(ctx *NodeLabelsContext) {}

// ExitNodeLabels is called when production nodeLabels is exited.
func (s *BaseCypherListener) ExitNodeLabels(ctx *NodeLabelsContext) {}

// EnterNodeLabel is called when production nodeLabel is entered.
func (s *BaseCypherListener) EnterNodeLabel(ctx *NodeLabelContext) {}

// ExitNodeLabel is called when production nodeLabel is exited.
func (s *BaseCypherListener) ExitNodeLabel(ctx *NodeLabelContext) {}

// EnterRangeLiteral is called when production rangeLiteral is entered.
func (s *BaseCypherListener) EnterRangeLiteral(ctx *RangeLiteralContext) {}

// ExitRangeLiteral is called when production rangeLiteral is exited.
func (s *BaseCypherListener) ExitRangeLiteral(ctx *RangeLiteralContext) {}

// EnterLabelName is called when production labelName is entered.
func (s *BaseCypherListener) EnterLabelName(ctx *LabelNameContext) {}

// ExitLabelName is called when production labelName is exited.
func (s *BaseCypherListener) ExitLabelName(ctx *LabelNameContext) {}

// EnterRelTypeName is called when production relTypeName is entered.
func (s *BaseCypherListener) EnterRelTypeName(ctx *RelTypeNameContext) {}

// ExitRelTypeName is called when production relTypeName is exited.
func (s *BaseCypherListener) ExitRelTypeName(ctx *RelTypeNameContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseCypherListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseCypherListener) ExitExpression(ctx *ExpressionContext) {}

// EnterOrExpression is called when production orExpression is entered.
func (s *BaseCypherListener) EnterOrExpression(ctx *OrExpressionContext) {}

// ExitOrExpression is called when production orExpression is exited.
func (s *BaseCypherListener) ExitOrExpression(ctx *OrExpressionContext) {}

// EnterXorExpression is called when production xorExpression is entered.
func (s *BaseCypherListener) EnterXorExpression(ctx *XorExpressionContext) {}

// ExitXorExpression is called when production xorExpression is exited.
func (s *BaseCypherListener) ExitXorExpression(ctx *XorExpressionContext) {}

// EnterAndExpression is called when production andExpression is entered.
func (s *BaseCypherListener) EnterAndExpression(ctx *AndExpressionContext) {}

// ExitAndExpression is called when production andExpression is exited.
func (s *BaseCypherListener) ExitAndExpression(ctx *AndExpressionContext) {}

// EnterNotExpression is called when production notExpression is entered.
func (s *BaseCypherListener) EnterNotExpression(ctx *NotExpressionContext) {}

// ExitNotExpression is called when production notExpression is exited.
func (s *BaseCypherListener) ExitNotExpression(ctx *NotExpressionContext) {}

// EnterComparisonExpression is called when production comparisonExpression is entered.
func (s *BaseCypherListener) EnterComparisonExpression(ctx *ComparisonExpressionContext) {}

// ExitComparisonExpression is called when production comparisonExpression is exited.
func (s *BaseCypherListener) ExitComparisonExpression(ctx *ComparisonExpressionContext) {}

// EnterAddOrSubtractExpression is called when production addOrSubtractExpression is entered.
func (s *BaseCypherListener) EnterAddOrSubtractExpression(ctx *AddOrSubtractExpressionContext) {}

// ExitAddOrSubtractExpression is called when production addOrSubtractExpression is exited.
func (s *BaseCypherListener) ExitAddOrSubtractExpression(ctx *AddOrSubtractExpressionContext) {}

// EnterMultiplyDivideModuloExpression is called when production multiplyDivideModuloExpression is entered.
func (s *BaseCypherListener) EnterMultiplyDivideModuloExpression(ctx *MultiplyDivideModuloExpressionContext) {
}

// ExitMultiplyDivideModuloExpression is called when production multiplyDivideModuloExpression is exited.
func (s *BaseCypherListener) ExitMultiplyDivideModuloExpression(ctx *MultiplyDivideModuloExpressionContext) {
}

// EnterPowerOfExpression is called when production powerOfExpression is entered.
func (s *BaseCypherListener) EnterPowerOfExpression(ctx *PowerOfExpressionContext) {}

// ExitPowerOfExpression is called when production powerOfExpression is exited.
func (s *BaseCypherListener) ExitPowerOfExpression(ctx *PowerOfExpressionContext) {}

// EnterUnaryAddOrSubtractExpression is called when production unaryAddOrSubtractExpression is entered.
func (s *BaseCypherListener) EnterUnaryAddOrSubtractExpression(ctx *UnaryAddOrSubtractExpressionContext) {
}

// ExitUnaryAddOrSubtractExpression is called when production unaryAddOrSubtractExpression is exited.
func (s *BaseCypherListener) ExitUnaryAddOrSubtractExpression(ctx *UnaryAddOrSubtractExpressionContext) {
}

// EnterStringListNullOperatorExpression is called when production stringListNullOperatorExpression is entered.
func (s *BaseCypherListener) EnterStringListNullOperatorExpression(ctx *StringListNullOperatorExpressionContext) {
}

// ExitStringListNullOperatorExpression is called when production stringListNullOperatorExpression is exited.
func (s *BaseCypherListener) ExitStringListNullOperatorExpression(ctx *StringListNullOperatorExpressionContext) {
}

// EnterPropertyOrLabelsExpression is called when production propertyOrLabelsExpression is entered.
func (s *BaseCypherListener) EnterPropertyOrLabelsExpression(ctx *PropertyOrLabelsExpressionContext) {
}

// ExitPropertyOrLabelsExpression is called when production propertyOrLabelsExpression is exited.
func (s *BaseCypherListener) ExitPropertyOrLabelsExpression(ctx *PropertyOrLabelsExpressionContext) {}

// EnterFilterFunction is called when production filterFunction is entered.
func (s *BaseCypherListener) EnterFilterFunction(ctx *FilterFunctionContext) {}

// ExitFilterFunction is called when production filterFunction is exited.
func (s *BaseCypherListener) ExitFilterFunction(ctx *FilterFunctionContext) {}

// EnterFilterFunctionName is called when production filterFunctionName is entered.
func (s *BaseCypherListener) EnterFilterFunctionName(ctx *FilterFunctionNameContext) {}

// ExitFilterFunctionName is called when production filterFunctionName is exited.
func (s *BaseCypherListener) ExitFilterFunctionName(ctx *FilterFunctionNameContext) {}

// EnterExistsFunction is called when production existsFunction is entered.
func (s *BaseCypherListener) EnterExistsFunction(ctx *ExistsFunctionContext) {}

// ExitExistsFunction is called when production existsFunction is exited.
func (s *BaseCypherListener) ExitExistsFunction(ctx *ExistsFunctionContext) {}

// EnterExistsFunctionName is called when production existsFunctionName is entered.
func (s *BaseCypherListener) EnterExistsFunctionName(ctx *ExistsFunctionNameContext) {}

// ExitExistsFunctionName is called when production existsFunctionName is exited.
func (s *BaseCypherListener) ExitExistsFunctionName(ctx *ExistsFunctionNameContext) {}

// EnterAllFunction is called when production allFunction is entered.
func (s *BaseCypherListener) EnterAllFunction(ctx *AllFunctionContext) {}

// ExitAllFunction is called when production allFunction is exited.
func (s *BaseCypherListener) ExitAllFunction(ctx *AllFunctionContext) {}

// EnterAllFunctionName is called when production allFunctionName is entered.
func (s *BaseCypherListener) EnterAllFunctionName(ctx *AllFunctionNameContext) {}

// ExitAllFunctionName is called when production allFunctionName is exited.
func (s *BaseCypherListener) ExitAllFunctionName(ctx *AllFunctionNameContext) {}

// EnterAnyFunction is called when production anyFunction is entered.
func (s *BaseCypherListener) EnterAnyFunction(ctx *AnyFunctionContext) {}

// ExitAnyFunction is called when production anyFunction is exited.
func (s *BaseCypherListener) ExitAnyFunction(ctx *AnyFunctionContext) {}

// EnterAnyFunctionName is called when production anyFunctionName is entered.
func (s *BaseCypherListener) EnterAnyFunctionName(ctx *AnyFunctionNameContext) {}

// ExitAnyFunctionName is called when production anyFunctionName is exited.
func (s *BaseCypherListener) ExitAnyFunctionName(ctx *AnyFunctionNameContext) {}

// EnterNoneFunction is called when production noneFunction is entered.
func (s *BaseCypherListener) EnterNoneFunction(ctx *NoneFunctionContext) {}

// ExitNoneFunction is called when production noneFunction is exited.
func (s *BaseCypherListener) ExitNoneFunction(ctx *NoneFunctionContext) {}

// EnterNoneFunctionName is called when production noneFunctionName is entered.
func (s *BaseCypherListener) EnterNoneFunctionName(ctx *NoneFunctionNameContext) {}

// ExitNoneFunctionName is called when production noneFunctionName is exited.
func (s *BaseCypherListener) ExitNoneFunctionName(ctx *NoneFunctionNameContext) {}

// EnterSingleFunction is called when production singleFunction is entered.
func (s *BaseCypherListener) EnterSingleFunction(ctx *SingleFunctionContext) {}

// ExitSingleFunction is called when production singleFunction is exited.
func (s *BaseCypherListener) ExitSingleFunction(ctx *SingleFunctionContext) {}

// EnterSingleFunctionName is called when production singleFunctionName is entered.
func (s *BaseCypherListener) EnterSingleFunctionName(ctx *SingleFunctionNameContext) {}

// ExitSingleFunctionName is called when production singleFunctionName is exited.
func (s *BaseCypherListener) ExitSingleFunctionName(ctx *SingleFunctionNameContext) {}

// EnterExtractFunction is called when production extractFunction is entered.
func (s *BaseCypherListener) EnterExtractFunction(ctx *ExtractFunctionContext) {}

// ExitExtractFunction is called when production extractFunction is exited.
func (s *BaseCypherListener) ExitExtractFunction(ctx *ExtractFunctionContext) {}

// EnterExtractFunctionName is called when production extractFunctionName is entered.
func (s *BaseCypherListener) EnterExtractFunctionName(ctx *ExtractFunctionNameContext) {}

// ExitExtractFunctionName is called when production extractFunctionName is exited.
func (s *BaseCypherListener) ExitExtractFunctionName(ctx *ExtractFunctionNameContext) {}

// EnterReduceFunction is called when production reduceFunction is entered.
func (s *BaseCypherListener) EnterReduceFunction(ctx *ReduceFunctionContext) {}

// ExitReduceFunction is called when production reduceFunction is exited.
func (s *BaseCypherListener) ExitReduceFunction(ctx *ReduceFunctionContext) {}

// EnterReduceFunctionName is called when production reduceFunctionName is entered.
func (s *BaseCypherListener) EnterReduceFunctionName(ctx *ReduceFunctionNameContext) {}

// ExitReduceFunctionName is called when production reduceFunctionName is exited.
func (s *BaseCypherListener) ExitReduceFunctionName(ctx *ReduceFunctionNameContext) {}

// EnterShortestPathPatternFunction is called when production shortestPathPatternFunction is entered.
func (s *BaseCypherListener) EnterShortestPathPatternFunction(ctx *ShortestPathPatternFunctionContext) {
}

// ExitShortestPathPatternFunction is called when production shortestPathPatternFunction is exited.
func (s *BaseCypherListener) ExitShortestPathPatternFunction(ctx *ShortestPathPatternFunctionContext) {
}

// EnterShortestPathFunctionName is called when production shortestPathFunctionName is entered.
func (s *BaseCypherListener) EnterShortestPathFunctionName(ctx *ShortestPathFunctionNameContext) {}

// ExitShortestPathFunctionName is called when production shortestPathFunctionName is exited.
func (s *BaseCypherListener) ExitShortestPathFunctionName(ctx *ShortestPathFunctionNameContext) {}

// EnterAllShortestPathFunctionName is called when production allShortestPathFunctionName is entered.
func (s *BaseCypherListener) EnterAllShortestPathFunctionName(ctx *AllShortestPathFunctionNameContext) {
}

// ExitAllShortestPathFunctionName is called when production allShortestPathFunctionName is exited.
func (s *BaseCypherListener) ExitAllShortestPathFunctionName(ctx *AllShortestPathFunctionNameContext) {
}

// EnterAtom is called when production atom is entered.
func (s *BaseCypherListener) EnterAtom(ctx *AtomContext) {}

// ExitAtom is called when production atom is exited.
func (s *BaseCypherListener) ExitAtom(ctx *AtomContext) {}

// EnterLiteral is called when production literal is entered.
func (s *BaseCypherListener) EnterLiteral(ctx *LiteralContext) {}

// ExitLiteral is called when production literal is exited.
func (s *BaseCypherListener) ExitLiteral(ctx *LiteralContext) {}

// EnterStringLiteral is called when production stringLiteral is entered.
func (s *BaseCypherListener) EnterStringLiteral(ctx *StringLiteralContext) {}

// ExitStringLiteral is called when production stringLiteral is exited.
func (s *BaseCypherListener) ExitStringLiteral(ctx *StringLiteralContext) {}

// EnterBooleanLiteral is called when production booleanLiteral is entered.
func (s *BaseCypherListener) EnterBooleanLiteral(ctx *BooleanLiteralContext) {}

// ExitBooleanLiteral is called when production booleanLiteral is exited.
func (s *BaseCypherListener) ExitBooleanLiteral(ctx *BooleanLiteralContext) {}

// EnterListLiteral is called when production listLiteral is entered.
func (s *BaseCypherListener) EnterListLiteral(ctx *ListLiteralContext) {}

// ExitListLiteral is called when production listLiteral is exited.
func (s *BaseCypherListener) ExitListLiteral(ctx *ListLiteralContext) {}

// EnterPartialComparisonExpression is called when production partialComparisonExpression is entered.
func (s *BaseCypherListener) EnterPartialComparisonExpression(ctx *PartialComparisonExpressionContext) {
}

// ExitPartialComparisonExpression is called when production partialComparisonExpression is exited.
func (s *BaseCypherListener) ExitPartialComparisonExpression(ctx *PartialComparisonExpressionContext) {
}

// EnterParenthesizedExpression is called when production parenthesizedExpression is entered.
func (s *BaseCypherListener) EnterParenthesizedExpression(ctx *ParenthesizedExpressionContext) {}

// ExitParenthesizedExpression is called when production parenthesizedExpression is exited.
func (s *BaseCypherListener) ExitParenthesizedExpression(ctx *ParenthesizedExpressionContext) {}

// EnterRelationshipsPattern is called when production relationshipsPattern is entered.
func (s *BaseCypherListener) EnterRelationshipsPattern(ctx *RelationshipsPatternContext) {}

// ExitRelationshipsPattern is called when production relationshipsPattern is exited.
func (s *BaseCypherListener) ExitRelationshipsPattern(ctx *RelationshipsPatternContext) {}

// EnterFilterExpression is called when production filterExpression is entered.
func (s *BaseCypherListener) EnterFilterExpression(ctx *FilterExpressionContext) {}

// ExitFilterExpression is called when production filterExpression is exited.
func (s *BaseCypherListener) ExitFilterExpression(ctx *FilterExpressionContext) {}

// EnterIdInColl is called when production idInColl is entered.
func (s *BaseCypherListener) EnterIdInColl(ctx *IdInCollContext) {}

// ExitIdInColl is called when production idInColl is exited.
func (s *BaseCypherListener) ExitIdInColl(ctx *IdInCollContext) {}

// EnterFunctionInvocation is called when production functionInvocation is entered.
func (s *BaseCypherListener) EnterFunctionInvocation(ctx *FunctionInvocationContext) {}

// ExitFunctionInvocation is called when production functionInvocation is exited.
func (s *BaseCypherListener) ExitFunctionInvocation(ctx *FunctionInvocationContext) {}

// EnterFunctionInvocationBody is called when production functionInvocationBody is entered.
func (s *BaseCypherListener) EnterFunctionInvocationBody(ctx *FunctionInvocationBodyContext) {}

// ExitFunctionInvocationBody is called when production functionInvocationBody is exited.
func (s *BaseCypherListener) ExitFunctionInvocationBody(ctx *FunctionInvocationBodyContext) {}

// EnterFunctionName is called when production functionName is entered.
func (s *BaseCypherListener) EnterFunctionName(ctx *FunctionNameContext) {}

// ExitFunctionName is called when production functionName is exited.
func (s *BaseCypherListener) ExitFunctionName(ctx *FunctionNameContext) {}

// EnterProcedureName is called when production procedureName is entered.
func (s *BaseCypherListener) EnterProcedureName(ctx *ProcedureNameContext) {}

// ExitProcedureName is called when production procedureName is exited.
func (s *BaseCypherListener) ExitProcedureName(ctx *ProcedureNameContext) {}

// EnterListComprehension is called when production listComprehension is entered.
func (s *BaseCypherListener) EnterListComprehension(ctx *ListComprehensionContext) {}

// ExitListComprehension is called when production listComprehension is exited.
func (s *BaseCypherListener) ExitListComprehension(ctx *ListComprehensionContext) {}

// EnterPatternComprehension is called when production patternComprehension is entered.
func (s *BaseCypherListener) EnterPatternComprehension(ctx *PatternComprehensionContext) {}

// ExitPatternComprehension is called when production patternComprehension is exited.
func (s *BaseCypherListener) ExitPatternComprehension(ctx *PatternComprehensionContext) {}

// EnterPropertyLookup is called when production propertyLookup is entered.
func (s *BaseCypherListener) EnterPropertyLookup(ctx *PropertyLookupContext) {}

// ExitPropertyLookup is called when production propertyLookup is exited.
func (s *BaseCypherListener) ExitPropertyLookup(ctx *PropertyLookupContext) {}

// EnterCaseExpression is called when production caseExpression is entered.
func (s *BaseCypherListener) EnterCaseExpression(ctx *CaseExpressionContext) {}

// ExitCaseExpression is called when production caseExpression is exited.
func (s *BaseCypherListener) ExitCaseExpression(ctx *CaseExpressionContext) {}

// EnterCaseAlternatives is called when production caseAlternatives is entered.
func (s *BaseCypherListener) EnterCaseAlternatives(ctx *CaseAlternativesContext) {}

// ExitCaseAlternatives is called when production caseAlternatives is exited.
func (s *BaseCypherListener) ExitCaseAlternatives(ctx *CaseAlternativesContext) {}

// EnterVariable is called when production variable is entered.
func (s *BaseCypherListener) EnterVariable(ctx *VariableContext) {}

// ExitVariable is called when production variable is exited.
func (s *BaseCypherListener) ExitVariable(ctx *VariableContext) {}

// EnterNumberLiteral is called when production numberLiteral is entered.
func (s *BaseCypherListener) EnterNumberLiteral(ctx *NumberLiteralContext) {}

// ExitNumberLiteral is called when production numberLiteral is exited.
func (s *BaseCypherListener) ExitNumberLiteral(ctx *NumberLiteralContext) {}

// EnterMapLiteral is called when production mapLiteral is entered.
func (s *BaseCypherListener) EnterMapLiteral(ctx *MapLiteralContext) {}

// ExitMapLiteral is called when production mapLiteral is exited.
func (s *BaseCypherListener) ExitMapLiteral(ctx *MapLiteralContext) {}

// EnterMapProjection is called when production mapProjection is entered.
func (s *BaseCypherListener) EnterMapProjection(ctx *MapProjectionContext) {}

// ExitMapProjection is called when production mapProjection is exited.
func (s *BaseCypherListener) ExitMapProjection(ctx *MapProjectionContext) {}

// EnterMapProjectionVariants is called when production mapProjectionVariants is entered.
func (s *BaseCypherListener) EnterMapProjectionVariants(ctx *MapProjectionVariantsContext) {}

// ExitMapProjectionVariants is called when production mapProjectionVariants is exited.
func (s *BaseCypherListener) ExitMapProjectionVariants(ctx *MapProjectionVariantsContext) {}

// EnterLiteralEntry is called when production literalEntry is entered.
func (s *BaseCypherListener) EnterLiteralEntry(ctx *LiteralEntryContext) {}

// ExitLiteralEntry is called when production literalEntry is exited.
func (s *BaseCypherListener) ExitLiteralEntry(ctx *LiteralEntryContext) {}

// EnterPropertySelector is called when production propertySelector is entered.
func (s *BaseCypherListener) EnterPropertySelector(ctx *PropertySelectorContext) {}

// ExitPropertySelector is called when production propertySelector is exited.
func (s *BaseCypherListener) ExitPropertySelector(ctx *PropertySelectorContext) {}

// EnterVariableSelector is called when production variableSelector is entered.
func (s *BaseCypherListener) EnterVariableSelector(ctx *VariableSelectorContext) {}

// ExitVariableSelector is called when production variableSelector is exited.
func (s *BaseCypherListener) ExitVariableSelector(ctx *VariableSelectorContext) {}

// EnterAllPropertiesSelector is called when production allPropertiesSelector is entered.
func (s *BaseCypherListener) EnterAllPropertiesSelector(ctx *AllPropertiesSelectorContext) {}

// ExitAllPropertiesSelector is called when production allPropertiesSelector is exited.
func (s *BaseCypherListener) ExitAllPropertiesSelector(ctx *AllPropertiesSelectorContext) {}

// EnterParameter is called when production parameter is entered.
func (s *BaseCypherListener) EnterParameter(ctx *ParameterContext) {}

// ExitParameter is called when production parameter is exited.
func (s *BaseCypherListener) ExitParameter(ctx *ParameterContext) {}

// EnterLegacyParameter is called when production legacyParameter is entered.
func (s *BaseCypherListener) EnterLegacyParameter(ctx *LegacyParameterContext) {}

// ExitLegacyParameter is called when production legacyParameter is exited.
func (s *BaseCypherListener) ExitLegacyParameter(ctx *LegacyParameterContext) {}

// EnterDollarParameter is called when production dollarParameter is entered.
func (s *BaseCypherListener) EnterDollarParameter(ctx *DollarParameterContext) {}

// ExitDollarParameter is called when production dollarParameter is exited.
func (s *BaseCypherListener) ExitDollarParameter(ctx *DollarParameterContext) {}

// EnterParameterName is called when production parameterName is entered.
func (s *BaseCypherListener) EnterParameterName(ctx *ParameterNameContext) {}

// ExitParameterName is called when production parameterName is exited.
func (s *BaseCypherListener) ExitParameterName(ctx *ParameterNameContext) {}

// EnterPropertyExpressions is called when production propertyExpressions is entered.
func (s *BaseCypherListener) EnterPropertyExpressions(ctx *PropertyExpressionsContext) {}

// ExitPropertyExpressions is called when production propertyExpressions is exited.
func (s *BaseCypherListener) ExitPropertyExpressions(ctx *PropertyExpressionsContext) {}

// EnterPropertyExpression is called when production propertyExpression is entered.
func (s *BaseCypherListener) EnterPropertyExpression(ctx *PropertyExpressionContext) {}

// ExitPropertyExpression is called when production propertyExpression is exited.
func (s *BaseCypherListener) ExitPropertyExpression(ctx *PropertyExpressionContext) {}

// EnterPropertyKeys is called when production propertyKeys is entered.
func (s *BaseCypherListener) EnterPropertyKeys(ctx *PropertyKeysContext) {}

// ExitPropertyKeys is called when production propertyKeys is exited.
func (s *BaseCypherListener) ExitPropertyKeys(ctx *PropertyKeysContext) {}

// EnterPropertyKeyName is called when production propertyKeyName is entered.
func (s *BaseCypherListener) EnterPropertyKeyName(ctx *PropertyKeyNameContext) {}

// ExitPropertyKeyName is called when production propertyKeyName is exited.
func (s *BaseCypherListener) ExitPropertyKeyName(ctx *PropertyKeyNameContext) {}

// EnterIntegerLiteral is called when production integerLiteral is entered.
func (s *BaseCypherListener) EnterIntegerLiteral(ctx *IntegerLiteralContext) {}

// ExitIntegerLiteral is called when production integerLiteral is exited.
func (s *BaseCypherListener) ExitIntegerLiteral(ctx *IntegerLiteralContext) {}

// EnterDoubleLiteral is called when production doubleLiteral is entered.
func (s *BaseCypherListener) EnterDoubleLiteral(ctx *DoubleLiteralContext) {}

// ExitDoubleLiteral is called when production doubleLiteral is exited.
func (s *BaseCypherListener) ExitDoubleLiteral(ctx *DoubleLiteralContext) {}

// EnterNamespace is called when production namespace is entered.
func (s *BaseCypherListener) EnterNamespace(ctx *NamespaceContext) {}

// ExitNamespace is called when production namespace is exited.
func (s *BaseCypherListener) ExitNamespace(ctx *NamespaceContext) {}

// EnterLeftArrowHead is called when production leftArrowHead is entered.
func (s *BaseCypherListener) EnterLeftArrowHead(ctx *LeftArrowHeadContext) {}

// ExitLeftArrowHead is called when production leftArrowHead is exited.
func (s *BaseCypherListener) ExitLeftArrowHead(ctx *LeftArrowHeadContext) {}

// EnterRightArrowHead is called when production rightArrowHead is entered.
func (s *BaseCypherListener) EnterRightArrowHead(ctx *RightArrowHeadContext) {}

// ExitRightArrowHead is called when production rightArrowHead is exited.
func (s *BaseCypherListener) ExitRightArrowHead(ctx *RightArrowHeadContext) {}

// EnterDash is called when production dash is entered.
func (s *BaseCypherListener) EnterDash(ctx *DashContext) {}

// ExitDash is called when production dash is exited.
func (s *BaseCypherListener) ExitDash(ctx *DashContext) {}

// EnterSymbolicName is called when production symbolicName is entered.
func (s *BaseCypherListener) EnterSymbolicName(ctx *SymbolicNameContext) {}

// ExitSymbolicName is called when production symbolicName is exited.
func (s *BaseCypherListener) ExitSymbolicName(ctx *SymbolicNameContext) {}

// EnterKeyword is called when production keyword is entered.
func (s *BaseCypherListener) EnterKeyword(ctx *KeywordContext) {}

// ExitKeyword is called when production keyword is exited.
func (s *BaseCypherListener) ExitKeyword(ctx *KeywordContext) {}

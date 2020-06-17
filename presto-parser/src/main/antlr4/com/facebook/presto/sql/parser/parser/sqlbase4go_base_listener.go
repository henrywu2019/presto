// Code generated from SqlBase4Go.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // SqlBase4Go

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseSqlBase4GoListener is a complete listener for a parse tree produced by SqlBase4GoParser.
type BaseSqlBase4GoListener struct{}

var _ SqlBase4GoListener = &BaseSqlBase4GoListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseSqlBase4GoListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseSqlBase4GoListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseSqlBase4GoListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseSqlBase4GoListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterSingleStatement is called when production singleStatement is entered.
func (s *BaseSqlBase4GoListener) EnterSingleStatement(ctx *SingleStatementContext) {}

// ExitSingleStatement is called when production singleStatement is exited.
func (s *BaseSqlBase4GoListener) ExitSingleStatement(ctx *SingleStatementContext) {}

// EnterStandaloneExpression is called when production standaloneExpression is entered.
func (s *BaseSqlBase4GoListener) EnterStandaloneExpression(ctx *StandaloneExpressionContext) {}

// ExitStandaloneExpression is called when production standaloneExpression is exited.
func (s *BaseSqlBase4GoListener) ExitStandaloneExpression(ctx *StandaloneExpressionContext) {}

// EnterStandaloneRoutineBody is called when production standaloneRoutineBody is entered.
func (s *BaseSqlBase4GoListener) EnterStandaloneRoutineBody(ctx *StandaloneRoutineBodyContext) {}

// ExitStandaloneRoutineBody is called when production standaloneRoutineBody is exited.
func (s *BaseSqlBase4GoListener) ExitStandaloneRoutineBody(ctx *StandaloneRoutineBodyContext) {}

// EnterStatementDefault is called when production statementDefault is entered.
func (s *BaseSqlBase4GoListener) EnterStatementDefault(ctx *StatementDefaultContext) {}

// ExitStatementDefault is called when production statementDefault is exited.
func (s *BaseSqlBase4GoListener) ExitStatementDefault(ctx *StatementDefaultContext) {}

// EnterUse is called when production use is entered.
func (s *BaseSqlBase4GoListener) EnterUse(ctx *UseContext) {}

// ExitUse is called when production use is exited.
func (s *BaseSqlBase4GoListener) ExitUse(ctx *UseContext) {}

// EnterCreateSchema is called when production createSchema is entered.
func (s *BaseSqlBase4GoListener) EnterCreateSchema(ctx *CreateSchemaContext) {}

// ExitCreateSchema is called when production createSchema is exited.
func (s *BaseSqlBase4GoListener) ExitCreateSchema(ctx *CreateSchemaContext) {}

// EnterDropSchema is called when production dropSchema is entered.
func (s *BaseSqlBase4GoListener) EnterDropSchema(ctx *DropSchemaContext) {}

// ExitDropSchema is called when production dropSchema is exited.
func (s *BaseSqlBase4GoListener) ExitDropSchema(ctx *DropSchemaContext) {}

// EnterRenameSchema is called when production renameSchema is entered.
func (s *BaseSqlBase4GoListener) EnterRenameSchema(ctx *RenameSchemaContext) {}

// ExitRenameSchema is called when production renameSchema is exited.
func (s *BaseSqlBase4GoListener) ExitRenameSchema(ctx *RenameSchemaContext) {}

// EnterCreateTableAsSelect is called when production createTableAsSelect is entered.
func (s *BaseSqlBase4GoListener) EnterCreateTableAsSelect(ctx *CreateTableAsSelectContext) {}

// ExitCreateTableAsSelect is called when production createTableAsSelect is exited.
func (s *BaseSqlBase4GoListener) ExitCreateTableAsSelect(ctx *CreateTableAsSelectContext) {}

// EnterCreateTable is called when production createTable is entered.
func (s *BaseSqlBase4GoListener) EnterCreateTable(ctx *CreateTableContext) {}

// ExitCreateTable is called when production createTable is exited.
func (s *BaseSqlBase4GoListener) ExitCreateTable(ctx *CreateTableContext) {}

// EnterDropTable is called when production dropTable is entered.
func (s *BaseSqlBase4GoListener) EnterDropTable(ctx *DropTableContext) {}

// ExitDropTable is called when production dropTable is exited.
func (s *BaseSqlBase4GoListener) ExitDropTable(ctx *DropTableContext) {}

// EnterInsertInto is called when production insertInto is entered.
func (s *BaseSqlBase4GoListener) EnterInsertInto(ctx *InsertIntoContext) {}

// ExitInsertInto is called when production insertInto is exited.
func (s *BaseSqlBase4GoListener) ExitInsertInto(ctx *InsertIntoContext) {}

// EnterDelete is called when production delete is entered.
func (s *BaseSqlBase4GoListener) EnterDelete(ctx *DeleteContext) {}

// ExitDelete is called when production delete is exited.
func (s *BaseSqlBase4GoListener) ExitDelete(ctx *DeleteContext) {}

// EnterRenameTable is called when production renameTable is entered.
func (s *BaseSqlBase4GoListener) EnterRenameTable(ctx *RenameTableContext) {}

// ExitRenameTable is called when production renameTable is exited.
func (s *BaseSqlBase4GoListener) ExitRenameTable(ctx *RenameTableContext) {}

// EnterRenameColumn is called when production renameColumn is entered.
func (s *BaseSqlBase4GoListener) EnterRenameColumn(ctx *RenameColumnContext) {}

// ExitRenameColumn is called when production renameColumn is exited.
func (s *BaseSqlBase4GoListener) ExitRenameColumn(ctx *RenameColumnContext) {}

// EnterDropColumn is called when production dropColumn is entered.
func (s *BaseSqlBase4GoListener) EnterDropColumn(ctx *DropColumnContext) {}

// ExitDropColumn is called when production dropColumn is exited.
func (s *BaseSqlBase4GoListener) ExitDropColumn(ctx *DropColumnContext) {}

// EnterAddColumn is called when production addColumn is entered.
func (s *BaseSqlBase4GoListener) EnterAddColumn(ctx *AddColumnContext) {}

// ExitAddColumn is called when production addColumn is exited.
func (s *BaseSqlBase4GoListener) ExitAddColumn(ctx *AddColumnContext) {}

// EnterAnalyze is called when production analyze is entered.
func (s *BaseSqlBase4GoListener) EnterAnalyze(ctx *AnalyzeContext) {}

// ExitAnalyze is called when production analyze is exited.
func (s *BaseSqlBase4GoListener) ExitAnalyze(ctx *AnalyzeContext) {}

// EnterCreateView is called when production createView is entered.
func (s *BaseSqlBase4GoListener) EnterCreateView(ctx *CreateViewContext) {}

// ExitCreateView is called when production createView is exited.
func (s *BaseSqlBase4GoListener) ExitCreateView(ctx *CreateViewContext) {}

// EnterDropView is called when production dropView is entered.
func (s *BaseSqlBase4GoListener) EnterDropView(ctx *DropViewContext) {}

// ExitDropView is called when production dropView is exited.
func (s *BaseSqlBase4GoListener) ExitDropView(ctx *DropViewContext) {}

// EnterCreateFunction is called when production createFunction is entered.
func (s *BaseSqlBase4GoListener) EnterCreateFunction(ctx *CreateFunctionContext) {}

// ExitCreateFunction is called when production createFunction is exited.
func (s *BaseSqlBase4GoListener) ExitCreateFunction(ctx *CreateFunctionContext) {}

// EnterAlterFunction is called when production alterFunction is entered.
func (s *BaseSqlBase4GoListener) EnterAlterFunction(ctx *AlterFunctionContext) {}

// ExitAlterFunction is called when production alterFunction is exited.
func (s *BaseSqlBase4GoListener) ExitAlterFunction(ctx *AlterFunctionContext) {}

// EnterDropFunction is called when production dropFunction is entered.
func (s *BaseSqlBase4GoListener) EnterDropFunction(ctx *DropFunctionContext) {}

// ExitDropFunction is called when production dropFunction is exited.
func (s *BaseSqlBase4GoListener) ExitDropFunction(ctx *DropFunctionContext) {}

// EnterCall is called when production call is entered.
func (s *BaseSqlBase4GoListener) EnterCall(ctx *CallContext) {}

// ExitCall is called when production call is exited.
func (s *BaseSqlBase4GoListener) ExitCall(ctx *CallContext) {}

// EnterCreateRole is called when production createRole is entered.
func (s *BaseSqlBase4GoListener) EnterCreateRole(ctx *CreateRoleContext) {}

// ExitCreateRole is called when production createRole is exited.
func (s *BaseSqlBase4GoListener) ExitCreateRole(ctx *CreateRoleContext) {}

// EnterDropRole is called when production dropRole is entered.
func (s *BaseSqlBase4GoListener) EnterDropRole(ctx *DropRoleContext) {}

// ExitDropRole is called when production dropRole is exited.
func (s *BaseSqlBase4GoListener) ExitDropRole(ctx *DropRoleContext) {}

// EnterGrantRoles is called when production grantRoles is entered.
func (s *BaseSqlBase4GoListener) EnterGrantRoles(ctx *GrantRolesContext) {}

// ExitGrantRoles is called when production grantRoles is exited.
func (s *BaseSqlBase4GoListener) ExitGrantRoles(ctx *GrantRolesContext) {}

// EnterRevokeRoles is called when production revokeRoles is entered.
func (s *BaseSqlBase4GoListener) EnterRevokeRoles(ctx *RevokeRolesContext) {}

// ExitRevokeRoles is called when production revokeRoles is exited.
func (s *BaseSqlBase4GoListener) ExitRevokeRoles(ctx *RevokeRolesContext) {}

// EnterSetRole is called when production setRole is entered.
func (s *BaseSqlBase4GoListener) EnterSetRole(ctx *SetRoleContext) {}

// ExitSetRole is called when production setRole is exited.
func (s *BaseSqlBase4GoListener) ExitSetRole(ctx *SetRoleContext) {}

// EnterGrant is called when production grant is entered.
func (s *BaseSqlBase4GoListener) EnterGrant(ctx *GrantContext) {}

// ExitGrant is called when production grant is exited.
func (s *BaseSqlBase4GoListener) ExitGrant(ctx *GrantContext) {}

// EnterRevoke is called when production revoke is entered.
func (s *BaseSqlBase4GoListener) EnterRevoke(ctx *RevokeContext) {}

// ExitRevoke is called when production revoke is exited.
func (s *BaseSqlBase4GoListener) ExitRevoke(ctx *RevokeContext) {}

// EnterShowGrants is called when production showGrants is entered.
func (s *BaseSqlBase4GoListener) EnterShowGrants(ctx *ShowGrantsContext) {}

// ExitShowGrants is called when production showGrants is exited.
func (s *BaseSqlBase4GoListener) ExitShowGrants(ctx *ShowGrantsContext) {}

// EnterExplain is called when production explain is entered.
func (s *BaseSqlBase4GoListener) EnterExplain(ctx *ExplainContext) {}

// ExitExplain is called when production explain is exited.
func (s *BaseSqlBase4GoListener) ExitExplain(ctx *ExplainContext) {}

// EnterShowCreateTable is called when production showCreateTable is entered.
func (s *BaseSqlBase4GoListener) EnterShowCreateTable(ctx *ShowCreateTableContext) {}

// ExitShowCreateTable is called when production showCreateTable is exited.
func (s *BaseSqlBase4GoListener) ExitShowCreateTable(ctx *ShowCreateTableContext) {}

// EnterShowCreateView is called when production showCreateView is entered.
func (s *BaseSqlBase4GoListener) EnterShowCreateView(ctx *ShowCreateViewContext) {}

// ExitShowCreateView is called when production showCreateView is exited.
func (s *BaseSqlBase4GoListener) ExitShowCreateView(ctx *ShowCreateViewContext) {}

// EnterShowCreateFunction is called when production showCreateFunction is entered.
func (s *BaseSqlBase4GoListener) EnterShowCreateFunction(ctx *ShowCreateFunctionContext) {}

// ExitShowCreateFunction is called when production showCreateFunction is exited.
func (s *BaseSqlBase4GoListener) ExitShowCreateFunction(ctx *ShowCreateFunctionContext) {}

// EnterShowTables is called when production showTables is entered.
func (s *BaseSqlBase4GoListener) EnterShowTables(ctx *ShowTablesContext) {}

// ExitShowTables is called when production showTables is exited.
func (s *BaseSqlBase4GoListener) ExitShowTables(ctx *ShowTablesContext) {}

// EnterShowSchemas is called when production showSchemas is entered.
func (s *BaseSqlBase4GoListener) EnterShowSchemas(ctx *ShowSchemasContext) {}

// ExitShowSchemas is called when production showSchemas is exited.
func (s *BaseSqlBase4GoListener) ExitShowSchemas(ctx *ShowSchemasContext) {}

// EnterShowCatalogs is called when production showCatalogs is entered.
func (s *BaseSqlBase4GoListener) EnterShowCatalogs(ctx *ShowCatalogsContext) {}

// ExitShowCatalogs is called when production showCatalogs is exited.
func (s *BaseSqlBase4GoListener) ExitShowCatalogs(ctx *ShowCatalogsContext) {}

// EnterShowColumns is called when production showColumns is entered.
func (s *BaseSqlBase4GoListener) EnterShowColumns(ctx *ShowColumnsContext) {}

// ExitShowColumns is called when production showColumns is exited.
func (s *BaseSqlBase4GoListener) ExitShowColumns(ctx *ShowColumnsContext) {}

// EnterShowStats is called when production showStats is entered.
func (s *BaseSqlBase4GoListener) EnterShowStats(ctx *ShowStatsContext) {}

// ExitShowStats is called when production showStats is exited.
func (s *BaseSqlBase4GoListener) ExitShowStats(ctx *ShowStatsContext) {}

// EnterShowStatsForQuery is called when production showStatsForQuery is entered.
func (s *BaseSqlBase4GoListener) EnterShowStatsForQuery(ctx *ShowStatsForQueryContext) {}

// ExitShowStatsForQuery is called when production showStatsForQuery is exited.
func (s *BaseSqlBase4GoListener) ExitShowStatsForQuery(ctx *ShowStatsForQueryContext) {}

// EnterShowRoles is called when production showRoles is entered.
func (s *BaseSqlBase4GoListener) EnterShowRoles(ctx *ShowRolesContext) {}

// ExitShowRoles is called when production showRoles is exited.
func (s *BaseSqlBase4GoListener) ExitShowRoles(ctx *ShowRolesContext) {}

// EnterShowRoleGrants is called when production showRoleGrants is entered.
func (s *BaseSqlBase4GoListener) EnterShowRoleGrants(ctx *ShowRoleGrantsContext) {}

// ExitShowRoleGrants is called when production showRoleGrants is exited.
func (s *BaseSqlBase4GoListener) ExitShowRoleGrants(ctx *ShowRoleGrantsContext) {}

// EnterShowFunctions is called when production showFunctions is entered.
func (s *BaseSqlBase4GoListener) EnterShowFunctions(ctx *ShowFunctionsContext) {}

// ExitShowFunctions is called when production showFunctions is exited.
func (s *BaseSqlBase4GoListener) ExitShowFunctions(ctx *ShowFunctionsContext) {}

// EnterShowSession is called when production showSession is entered.
func (s *BaseSqlBase4GoListener) EnterShowSession(ctx *ShowSessionContext) {}

// ExitShowSession is called when production showSession is exited.
func (s *BaseSqlBase4GoListener) ExitShowSession(ctx *ShowSessionContext) {}

// EnterSetSession is called when production setSession is entered.
func (s *BaseSqlBase4GoListener) EnterSetSession(ctx *SetSessionContext) {}

// ExitSetSession is called when production setSession is exited.
func (s *BaseSqlBase4GoListener) ExitSetSession(ctx *SetSessionContext) {}

// EnterResetSession is called when production resetSession is entered.
func (s *BaseSqlBase4GoListener) EnterResetSession(ctx *ResetSessionContext) {}

// ExitResetSession is called when production resetSession is exited.
func (s *BaseSqlBase4GoListener) ExitResetSession(ctx *ResetSessionContext) {}

// EnterStartTransaction is called when production startTransaction is entered.
func (s *BaseSqlBase4GoListener) EnterStartTransaction(ctx *StartTransactionContext) {}

// ExitStartTransaction is called when production startTransaction is exited.
func (s *BaseSqlBase4GoListener) ExitStartTransaction(ctx *StartTransactionContext) {}

// EnterCommit is called when production commit is entered.
func (s *BaseSqlBase4GoListener) EnterCommit(ctx *CommitContext) {}

// ExitCommit is called when production commit is exited.
func (s *BaseSqlBase4GoListener) ExitCommit(ctx *CommitContext) {}

// EnterRollback is called when production rollback is entered.
func (s *BaseSqlBase4GoListener) EnterRollback(ctx *RollbackContext) {}

// ExitRollback is called when production rollback is exited.
func (s *BaseSqlBase4GoListener) ExitRollback(ctx *RollbackContext) {}

// EnterPrepare is called when production prepare is entered.
func (s *BaseSqlBase4GoListener) EnterPrepare(ctx *PrepareContext) {}

// ExitPrepare is called when production prepare is exited.
func (s *BaseSqlBase4GoListener) ExitPrepare(ctx *PrepareContext) {}

// EnterDeallocate is called when production deallocate is entered.
func (s *BaseSqlBase4GoListener) EnterDeallocate(ctx *DeallocateContext) {}

// ExitDeallocate is called when production deallocate is exited.
func (s *BaseSqlBase4GoListener) ExitDeallocate(ctx *DeallocateContext) {}

// EnterExecute is called when production execute is entered.
func (s *BaseSqlBase4GoListener) EnterExecute(ctx *ExecuteContext) {}

// ExitExecute is called when production execute is exited.
func (s *BaseSqlBase4GoListener) ExitExecute(ctx *ExecuteContext) {}

// EnterDescribeInput is called when production describeInput is entered.
func (s *BaseSqlBase4GoListener) EnterDescribeInput(ctx *DescribeInputContext) {}

// ExitDescribeInput is called when production describeInput is exited.
func (s *BaseSqlBase4GoListener) ExitDescribeInput(ctx *DescribeInputContext) {}

// EnterDescribeOutput is called when production describeOutput is entered.
func (s *BaseSqlBase4GoListener) EnterDescribeOutput(ctx *DescribeOutputContext) {}

// ExitDescribeOutput is called when production describeOutput is exited.
func (s *BaseSqlBase4GoListener) ExitDescribeOutput(ctx *DescribeOutputContext) {}

// EnterQuery is called when production query is entered.
func (s *BaseSqlBase4GoListener) EnterQuery(ctx *QueryContext) {}

// ExitQuery is called when production query is exited.
func (s *BaseSqlBase4GoListener) ExitQuery(ctx *QueryContext) {}

// EnterWith is called when production with is entered.
func (s *BaseSqlBase4GoListener) EnterWith(ctx *WithContext) {}

// ExitWith is called when production with is exited.
func (s *BaseSqlBase4GoListener) ExitWith(ctx *WithContext) {}

// EnterTableElement is called when production tableElement is entered.
func (s *BaseSqlBase4GoListener) EnterTableElement(ctx *TableElementContext) {}

// ExitTableElement is called when production tableElement is exited.
func (s *BaseSqlBase4GoListener) ExitTableElement(ctx *TableElementContext) {}

// EnterColumnDefinition is called when production columnDefinition is entered.
func (s *BaseSqlBase4GoListener) EnterColumnDefinition(ctx *ColumnDefinitionContext) {}

// ExitColumnDefinition is called when production columnDefinition is exited.
func (s *BaseSqlBase4GoListener) ExitColumnDefinition(ctx *ColumnDefinitionContext) {}

// EnterLikeClause is called when production likeClause is entered.
func (s *BaseSqlBase4GoListener) EnterLikeClause(ctx *LikeClauseContext) {}

// ExitLikeClause is called when production likeClause is exited.
func (s *BaseSqlBase4GoListener) ExitLikeClause(ctx *LikeClauseContext) {}

// EnterProperties is called when production properties is entered.
func (s *BaseSqlBase4GoListener) EnterProperties(ctx *PropertiesContext) {}

// ExitProperties is called when production properties is exited.
func (s *BaseSqlBase4GoListener) ExitProperties(ctx *PropertiesContext) {}

// EnterProperty is called when production property is entered.
func (s *BaseSqlBase4GoListener) EnterProperty(ctx *PropertyContext) {}

// ExitProperty is called when production property is exited.
func (s *BaseSqlBase4GoListener) ExitProperty(ctx *PropertyContext) {}

// EnterSqlParameterDeclaration is called when production sqlParameterDeclaration is entered.
func (s *BaseSqlBase4GoListener) EnterSqlParameterDeclaration(ctx *SqlParameterDeclarationContext) {}

// ExitSqlParameterDeclaration is called when production sqlParameterDeclaration is exited.
func (s *BaseSqlBase4GoListener) ExitSqlParameterDeclaration(ctx *SqlParameterDeclarationContext) {}

// EnterRoutineCharacteristics is called when production routineCharacteristics is entered.
func (s *BaseSqlBase4GoListener) EnterRoutineCharacteristics(ctx *RoutineCharacteristicsContext) {}

// ExitRoutineCharacteristics is called when production routineCharacteristics is exited.
func (s *BaseSqlBase4GoListener) ExitRoutineCharacteristics(ctx *RoutineCharacteristicsContext) {}

// EnterRoutineCharacteristic is called when production routineCharacteristic is entered.
func (s *BaseSqlBase4GoListener) EnterRoutineCharacteristic(ctx *RoutineCharacteristicContext) {}

// ExitRoutineCharacteristic is called when production routineCharacteristic is exited.
func (s *BaseSqlBase4GoListener) ExitRoutineCharacteristic(ctx *RoutineCharacteristicContext) {}

// EnterAlterRoutineCharacteristics is called when production alterRoutineCharacteristics is entered.
func (s *BaseSqlBase4GoListener) EnterAlterRoutineCharacteristics(ctx *AlterRoutineCharacteristicsContext) {
}

// ExitAlterRoutineCharacteristics is called when production alterRoutineCharacteristics is exited.
func (s *BaseSqlBase4GoListener) ExitAlterRoutineCharacteristics(ctx *AlterRoutineCharacteristicsContext) {
}

// EnterAlterRoutineCharacteristic is called when production alterRoutineCharacteristic is entered.
func (s *BaseSqlBase4GoListener) EnterAlterRoutineCharacteristic(ctx *AlterRoutineCharacteristicContext) {
}

// ExitAlterRoutineCharacteristic is called when production alterRoutineCharacteristic is exited.
func (s *BaseSqlBase4GoListener) ExitAlterRoutineCharacteristic(ctx *AlterRoutineCharacteristicContext) {
}

// EnterRoutineBody is called when production routineBody is entered.
func (s *BaseSqlBase4GoListener) EnterRoutineBody(ctx *RoutineBodyContext) {}

// ExitRoutineBody is called when production routineBody is exited.
func (s *BaseSqlBase4GoListener) ExitRoutineBody(ctx *RoutineBodyContext) {}

// EnterReturnStatement is called when production returnStatement is entered.
func (s *BaseSqlBase4GoListener) EnterReturnStatement(ctx *ReturnStatementContext) {}

// ExitReturnStatement is called when production returnStatement is exited.
func (s *BaseSqlBase4GoListener) ExitReturnStatement(ctx *ReturnStatementContext) {}

// EnterLanguage is called when production language is entered.
func (s *BaseSqlBase4GoListener) EnterLanguage(ctx *LanguageContext) {}

// ExitLanguage is called when production language is exited.
func (s *BaseSqlBase4GoListener) ExitLanguage(ctx *LanguageContext) {}

// EnterDeterminism is called when production determinism is entered.
func (s *BaseSqlBase4GoListener) EnterDeterminism(ctx *DeterminismContext) {}

// ExitDeterminism is called when production determinism is exited.
func (s *BaseSqlBase4GoListener) ExitDeterminism(ctx *DeterminismContext) {}

// EnterNullCallClause is called when production nullCallClause is entered.
func (s *BaseSqlBase4GoListener) EnterNullCallClause(ctx *NullCallClauseContext) {}

// ExitNullCallClause is called when production nullCallClause is exited.
func (s *BaseSqlBase4GoListener) ExitNullCallClause(ctx *NullCallClauseContext) {}

// EnterQueryNoWith is called when production queryNoWith is entered.
func (s *BaseSqlBase4GoListener) EnterQueryNoWith(ctx *QueryNoWithContext) {}

// ExitQueryNoWith is called when production queryNoWith is exited.
func (s *BaseSqlBase4GoListener) ExitQueryNoWith(ctx *QueryNoWithContext) {}

// EnterQueryTermDefault is called when production queryTermDefault is entered.
func (s *BaseSqlBase4GoListener) EnterQueryTermDefault(ctx *QueryTermDefaultContext) {}

// ExitQueryTermDefault is called when production queryTermDefault is exited.
func (s *BaseSqlBase4GoListener) ExitQueryTermDefault(ctx *QueryTermDefaultContext) {}

// EnterSetOperation is called when production setOperation is entered.
func (s *BaseSqlBase4GoListener) EnterSetOperation(ctx *SetOperationContext) {}

// ExitSetOperation is called when production setOperation is exited.
func (s *BaseSqlBase4GoListener) ExitSetOperation(ctx *SetOperationContext) {}

// EnterQueryPrimaryDefault is called when production queryPrimaryDefault is entered.
func (s *BaseSqlBase4GoListener) EnterQueryPrimaryDefault(ctx *QueryPrimaryDefaultContext) {}

// ExitQueryPrimaryDefault is called when production queryPrimaryDefault is exited.
func (s *BaseSqlBase4GoListener) ExitQueryPrimaryDefault(ctx *QueryPrimaryDefaultContext) {}

// EnterTable is called when production table is entered.
func (s *BaseSqlBase4GoListener) EnterTable(ctx *TableContext) {}

// ExitTable is called when production table is exited.
func (s *BaseSqlBase4GoListener) ExitTable(ctx *TableContext) {}

// EnterInlineTable is called when production inlineTable is entered.
func (s *BaseSqlBase4GoListener) EnterInlineTable(ctx *InlineTableContext) {}

// ExitInlineTable is called when production inlineTable is exited.
func (s *BaseSqlBase4GoListener) ExitInlineTable(ctx *InlineTableContext) {}

// EnterSubquery is called when production subquery is entered.
func (s *BaseSqlBase4GoListener) EnterSubquery(ctx *SubqueryContext) {}

// ExitSubquery is called when production subquery is exited.
func (s *BaseSqlBase4GoListener) ExitSubquery(ctx *SubqueryContext) {}

// EnterSortItem is called when production sortItem is entered.
func (s *BaseSqlBase4GoListener) EnterSortItem(ctx *SortItemContext) {}

// ExitSortItem is called when production sortItem is exited.
func (s *BaseSqlBase4GoListener) ExitSortItem(ctx *SortItemContext) {}

// EnterQuerySpecification is called when production querySpecification is entered.
func (s *BaseSqlBase4GoListener) EnterQuerySpecification(ctx *QuerySpecificationContext) {}

// ExitQuerySpecification is called when production querySpecification is exited.
func (s *BaseSqlBase4GoListener) ExitQuerySpecification(ctx *QuerySpecificationContext) {}

// EnterGroupBy is called when production groupBy is entered.
func (s *BaseSqlBase4GoListener) EnterGroupBy(ctx *GroupByContext) {}

// ExitGroupBy is called when production groupBy is exited.
func (s *BaseSqlBase4GoListener) ExitGroupBy(ctx *GroupByContext) {}

// EnterSingleGroupingSet is called when production singleGroupingSet is entered.
func (s *BaseSqlBase4GoListener) EnterSingleGroupingSet(ctx *SingleGroupingSetContext) {}

// ExitSingleGroupingSet is called when production singleGroupingSet is exited.
func (s *BaseSqlBase4GoListener) ExitSingleGroupingSet(ctx *SingleGroupingSetContext) {}

// EnterRollup is called when production rollup is entered.
func (s *BaseSqlBase4GoListener) EnterRollup(ctx *RollupContext) {}

// ExitRollup is called when production rollup is exited.
func (s *BaseSqlBase4GoListener) ExitRollup(ctx *RollupContext) {}

// EnterCube is called when production cube is entered.
func (s *BaseSqlBase4GoListener) EnterCube(ctx *CubeContext) {}

// ExitCube is called when production cube is exited.
func (s *BaseSqlBase4GoListener) ExitCube(ctx *CubeContext) {}

// EnterMultipleGroupingSets is called when production multipleGroupingSets is entered.
func (s *BaseSqlBase4GoListener) EnterMultipleGroupingSets(ctx *MultipleGroupingSetsContext) {}

// ExitMultipleGroupingSets is called when production multipleGroupingSets is exited.
func (s *BaseSqlBase4GoListener) ExitMultipleGroupingSets(ctx *MultipleGroupingSetsContext) {}

// EnterGroupingSet is called when production groupingSet is entered.
func (s *BaseSqlBase4GoListener) EnterGroupingSet(ctx *GroupingSetContext) {}

// ExitGroupingSet is called when production groupingSet is exited.
func (s *BaseSqlBase4GoListener) ExitGroupingSet(ctx *GroupingSetContext) {}

// EnterNamedQuery is called when production namedQuery is entered.
func (s *BaseSqlBase4GoListener) EnterNamedQuery(ctx *NamedQueryContext) {}

// ExitNamedQuery is called when production namedQuery is exited.
func (s *BaseSqlBase4GoListener) ExitNamedQuery(ctx *NamedQueryContext) {}

// EnterSetQuantifier is called when production setQuantifier is entered.
func (s *BaseSqlBase4GoListener) EnterSetQuantifier(ctx *SetQuantifierContext) {}

// ExitSetQuantifier is called when production setQuantifier is exited.
func (s *BaseSqlBase4GoListener) ExitSetQuantifier(ctx *SetQuantifierContext) {}

// EnterSelectSingle is called when production selectSingle is entered.
func (s *BaseSqlBase4GoListener) EnterSelectSingle(ctx *SelectSingleContext) {}

// ExitSelectSingle is called when production selectSingle is exited.
func (s *BaseSqlBase4GoListener) ExitSelectSingle(ctx *SelectSingleContext) {}

// EnterSelectAll is called when production selectAll is entered.
func (s *BaseSqlBase4GoListener) EnterSelectAll(ctx *SelectAllContext) {}

// ExitSelectAll is called when production selectAll is exited.
func (s *BaseSqlBase4GoListener) ExitSelectAll(ctx *SelectAllContext) {}

// EnterRelationDefault is called when production relationDefault is entered.
func (s *BaseSqlBase4GoListener) EnterRelationDefault(ctx *RelationDefaultContext) {}

// ExitRelationDefault is called when production relationDefault is exited.
func (s *BaseSqlBase4GoListener) ExitRelationDefault(ctx *RelationDefaultContext) {}

// EnterJoinRelation is called when production joinRelation is entered.
func (s *BaseSqlBase4GoListener) EnterJoinRelation(ctx *JoinRelationContext) {}

// ExitJoinRelation is called when production joinRelation is exited.
func (s *BaseSqlBase4GoListener) ExitJoinRelation(ctx *JoinRelationContext) {}

// EnterJoinType is called when production joinType is entered.
func (s *BaseSqlBase4GoListener) EnterJoinType(ctx *JoinTypeContext) {}

// ExitJoinType is called when production joinType is exited.
func (s *BaseSqlBase4GoListener) ExitJoinType(ctx *JoinTypeContext) {}

// EnterJoinCriteria is called when production joinCriteria is entered.
func (s *BaseSqlBase4GoListener) EnterJoinCriteria(ctx *JoinCriteriaContext) {}

// ExitJoinCriteria is called when production joinCriteria is exited.
func (s *BaseSqlBase4GoListener) ExitJoinCriteria(ctx *JoinCriteriaContext) {}

// EnterSampledRelation is called when production sampledRelation is entered.
func (s *BaseSqlBase4GoListener) EnterSampledRelation(ctx *SampledRelationContext) {}

// ExitSampledRelation is called when production sampledRelation is exited.
func (s *BaseSqlBase4GoListener) ExitSampledRelation(ctx *SampledRelationContext) {}

// EnterSampleType is called when production sampleType is entered.
func (s *BaseSqlBase4GoListener) EnterSampleType(ctx *SampleTypeContext) {}

// ExitSampleType is called when production sampleType is exited.
func (s *BaseSqlBase4GoListener) ExitSampleType(ctx *SampleTypeContext) {}

// EnterAliasedRelation is called when production aliasedRelation is entered.
func (s *BaseSqlBase4GoListener) EnterAliasedRelation(ctx *AliasedRelationContext) {}

// ExitAliasedRelation is called when production aliasedRelation is exited.
func (s *BaseSqlBase4GoListener) ExitAliasedRelation(ctx *AliasedRelationContext) {}

// EnterColumnAliases is called when production columnAliases is entered.
func (s *BaseSqlBase4GoListener) EnterColumnAliases(ctx *ColumnAliasesContext) {}

// ExitColumnAliases is called when production columnAliases is exited.
func (s *BaseSqlBase4GoListener) ExitColumnAliases(ctx *ColumnAliasesContext) {}

// EnterTableName is called when production tableName is entered.
func (s *BaseSqlBase4GoListener) EnterTableName(ctx *TableNameContext) {}

// ExitTableName is called when production tableName is exited.
func (s *BaseSqlBase4GoListener) ExitTableName(ctx *TableNameContext) {}

// EnterSubqueryRelation is called when production subqueryRelation is entered.
func (s *BaseSqlBase4GoListener) EnterSubqueryRelation(ctx *SubqueryRelationContext) {}

// ExitSubqueryRelation is called when production subqueryRelation is exited.
func (s *BaseSqlBase4GoListener) ExitSubqueryRelation(ctx *SubqueryRelationContext) {}

// EnterUnnest is called when production unnest is entered.
func (s *BaseSqlBase4GoListener) EnterUnnest(ctx *UnnestContext) {}

// ExitUnnest is called when production unnest is exited.
func (s *BaseSqlBase4GoListener) ExitUnnest(ctx *UnnestContext) {}

// EnterLateral is called when production lateral is entered.
func (s *BaseSqlBase4GoListener) EnterLateral(ctx *LateralContext) {}

// ExitLateral is called when production lateral is exited.
func (s *BaseSqlBase4GoListener) ExitLateral(ctx *LateralContext) {}

// EnterParenthesizedRelation is called when production parenthesizedRelation is entered.
func (s *BaseSqlBase4GoListener) EnterParenthesizedRelation(ctx *ParenthesizedRelationContext) {}

// ExitParenthesizedRelation is called when production parenthesizedRelation is exited.
func (s *BaseSqlBase4GoListener) ExitParenthesizedRelation(ctx *ParenthesizedRelationContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseSqlBase4GoListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseSqlBase4GoListener) ExitExpression(ctx *ExpressionContext) {}

// EnterLogicalNot is called when production logicalNot is entered.
func (s *BaseSqlBase4GoListener) EnterLogicalNot(ctx *LogicalNotContext) {}

// ExitLogicalNot is called when production logicalNot is exited.
func (s *BaseSqlBase4GoListener) ExitLogicalNot(ctx *LogicalNotContext) {}

// EnterPredicated is called when production predicated is entered.
func (s *BaseSqlBase4GoListener) EnterPredicated(ctx *PredicatedContext) {}

// ExitPredicated is called when production predicated is exited.
func (s *BaseSqlBase4GoListener) ExitPredicated(ctx *PredicatedContext) {}

// EnterLogicalBinary is called when production logicalBinary is entered.
func (s *BaseSqlBase4GoListener) EnterLogicalBinary(ctx *LogicalBinaryContext) {}

// ExitLogicalBinary is called when production logicalBinary is exited.
func (s *BaseSqlBase4GoListener) ExitLogicalBinary(ctx *LogicalBinaryContext) {}

// EnterComparison is called when production comparison is entered.
func (s *BaseSqlBase4GoListener) EnterComparison(ctx *ComparisonContext) {}

// ExitComparison is called when production comparison is exited.
func (s *BaseSqlBase4GoListener) ExitComparison(ctx *ComparisonContext) {}

// EnterQuantifiedComparison is called when production quantifiedComparison is entered.
func (s *BaseSqlBase4GoListener) EnterQuantifiedComparison(ctx *QuantifiedComparisonContext) {}

// ExitQuantifiedComparison is called when production quantifiedComparison is exited.
func (s *BaseSqlBase4GoListener) ExitQuantifiedComparison(ctx *QuantifiedComparisonContext) {}

// EnterBetween is called when production between is entered.
func (s *BaseSqlBase4GoListener) EnterBetween(ctx *BetweenContext) {}

// ExitBetween is called when production between is exited.
func (s *BaseSqlBase4GoListener) ExitBetween(ctx *BetweenContext) {}

// EnterInList is called when production inList is entered.
func (s *BaseSqlBase4GoListener) EnterInList(ctx *InListContext) {}

// ExitInList is called when production inList is exited.
func (s *BaseSqlBase4GoListener) ExitInList(ctx *InListContext) {}

// EnterInSubquery is called when production inSubquery is entered.
func (s *BaseSqlBase4GoListener) EnterInSubquery(ctx *InSubqueryContext) {}

// ExitInSubquery is called when production inSubquery is exited.
func (s *BaseSqlBase4GoListener) ExitInSubquery(ctx *InSubqueryContext) {}

// EnterLike is called when production like is entered.
func (s *BaseSqlBase4GoListener) EnterLike(ctx *LikeContext) {}

// ExitLike is called when production like is exited.
func (s *BaseSqlBase4GoListener) ExitLike(ctx *LikeContext) {}

// EnterNullPredicate is called when production nullPredicate is entered.
func (s *BaseSqlBase4GoListener) EnterNullPredicate(ctx *NullPredicateContext) {}

// ExitNullPredicate is called when production nullPredicate is exited.
func (s *BaseSqlBase4GoListener) ExitNullPredicate(ctx *NullPredicateContext) {}

// EnterDistinctFrom is called when production distinctFrom is entered.
func (s *BaseSqlBase4GoListener) EnterDistinctFrom(ctx *DistinctFromContext) {}

// ExitDistinctFrom is called when production distinctFrom is exited.
func (s *BaseSqlBase4GoListener) ExitDistinctFrom(ctx *DistinctFromContext) {}

// EnterValueExpressionDefault is called when production valueExpressionDefault is entered.
func (s *BaseSqlBase4GoListener) EnterValueExpressionDefault(ctx *ValueExpressionDefaultContext) {}

// ExitValueExpressionDefault is called when production valueExpressionDefault is exited.
func (s *BaseSqlBase4GoListener) ExitValueExpressionDefault(ctx *ValueExpressionDefaultContext) {}

// EnterConcatenation is called when production concatenation is entered.
func (s *BaseSqlBase4GoListener) EnterConcatenation(ctx *ConcatenationContext) {}

// ExitConcatenation is called when production concatenation is exited.
func (s *BaseSqlBase4GoListener) ExitConcatenation(ctx *ConcatenationContext) {}

// EnterArithmeticBinary is called when production arithmeticBinary is entered.
func (s *BaseSqlBase4GoListener) EnterArithmeticBinary(ctx *ArithmeticBinaryContext) {}

// ExitArithmeticBinary is called when production arithmeticBinary is exited.
func (s *BaseSqlBase4GoListener) ExitArithmeticBinary(ctx *ArithmeticBinaryContext) {}

// EnterArithmeticUnary is called when production arithmeticUnary is entered.
func (s *BaseSqlBase4GoListener) EnterArithmeticUnary(ctx *ArithmeticUnaryContext) {}

// ExitArithmeticUnary is called when production arithmeticUnary is exited.
func (s *BaseSqlBase4GoListener) ExitArithmeticUnary(ctx *ArithmeticUnaryContext) {}

// EnterAtTimeZone is called when production atTimeZone is entered.
func (s *BaseSqlBase4GoListener) EnterAtTimeZone(ctx *AtTimeZoneContext) {}

// ExitAtTimeZone is called when production atTimeZone is exited.
func (s *BaseSqlBase4GoListener) ExitAtTimeZone(ctx *AtTimeZoneContext) {}

// EnterDereference is called when production dereference is entered.
func (s *BaseSqlBase4GoListener) EnterDereference(ctx *DereferenceContext) {}

// ExitDereference is called when production dereference is exited.
func (s *BaseSqlBase4GoListener) ExitDereference(ctx *DereferenceContext) {}

// EnterTypeConstructor is called when production typeConstructor is entered.
func (s *BaseSqlBase4GoListener) EnterTypeConstructor(ctx *TypeConstructorContext) {}

// ExitTypeConstructor is called when production typeConstructor is exited.
func (s *BaseSqlBase4GoListener) ExitTypeConstructor(ctx *TypeConstructorContext) {}

// EnterSpecialDateTimeFunction is called when production specialDateTimeFunction is entered.
func (s *BaseSqlBase4GoListener) EnterSpecialDateTimeFunction(ctx *SpecialDateTimeFunctionContext) {}

// ExitSpecialDateTimeFunction is called when production specialDateTimeFunction is exited.
func (s *BaseSqlBase4GoListener) ExitSpecialDateTimeFunction(ctx *SpecialDateTimeFunctionContext) {}

// EnterSubstring is called when production substring is entered.
func (s *BaseSqlBase4GoListener) EnterSubstring(ctx *SubstringContext) {}

// ExitSubstring is called when production substring is exited.
func (s *BaseSqlBase4GoListener) ExitSubstring(ctx *SubstringContext) {}

// EnterCast is called when production cast is entered.
func (s *BaseSqlBase4GoListener) EnterCast(ctx *CastContext) {}

// ExitCast is called when production cast is exited.
func (s *BaseSqlBase4GoListener) ExitCast(ctx *CastContext) {}

// EnterLambda is called when production lambda is entered.
func (s *BaseSqlBase4GoListener) EnterLambda(ctx *LambdaContext) {}

// ExitLambda is called when production lambda is exited.
func (s *BaseSqlBase4GoListener) ExitLambda(ctx *LambdaContext) {}

// EnterParenthesizedExpression is called when production parenthesizedExpression is entered.
func (s *BaseSqlBase4GoListener) EnterParenthesizedExpression(ctx *ParenthesizedExpressionContext) {}

// ExitParenthesizedExpression is called when production parenthesizedExpression is exited.
func (s *BaseSqlBase4GoListener) ExitParenthesizedExpression(ctx *ParenthesizedExpressionContext) {}

// EnterParameter is called when production parameter is entered.
func (s *BaseSqlBase4GoListener) EnterParameter(ctx *ParameterContext) {}

// ExitParameter is called when production parameter is exited.
func (s *BaseSqlBase4GoListener) ExitParameter(ctx *ParameterContext) {}

// EnterNormalize is called when production normalize is entered.
func (s *BaseSqlBase4GoListener) EnterNormalize(ctx *NormalizeContext) {}

// ExitNormalize is called when production normalize is exited.
func (s *BaseSqlBase4GoListener) ExitNormalize(ctx *NormalizeContext) {}

// EnterIntervalLiteral is called when production intervalLiteral is entered.
func (s *BaseSqlBase4GoListener) EnterIntervalLiteral(ctx *IntervalLiteralContext) {}

// ExitIntervalLiteral is called when production intervalLiteral is exited.
func (s *BaseSqlBase4GoListener) ExitIntervalLiteral(ctx *IntervalLiteralContext) {}

// EnterNumericLiteral is called when production numericLiteral is entered.
func (s *BaseSqlBase4GoListener) EnterNumericLiteral(ctx *NumericLiteralContext) {}

// ExitNumericLiteral is called when production numericLiteral is exited.
func (s *BaseSqlBase4GoListener) ExitNumericLiteral(ctx *NumericLiteralContext) {}

// EnterBooleanLiteral is called when production booleanLiteral is entered.
func (s *BaseSqlBase4GoListener) EnterBooleanLiteral(ctx *BooleanLiteralContext) {}

// ExitBooleanLiteral is called when production booleanLiteral is exited.
func (s *BaseSqlBase4GoListener) ExitBooleanLiteral(ctx *BooleanLiteralContext) {}

// EnterSimpleCase is called when production simpleCase is entered.
func (s *BaseSqlBase4GoListener) EnterSimpleCase(ctx *SimpleCaseContext) {}

// ExitSimpleCase is called when production simpleCase is exited.
func (s *BaseSqlBase4GoListener) ExitSimpleCase(ctx *SimpleCaseContext) {}

// EnterColumnReference is called when production columnReference is entered.
func (s *BaseSqlBase4GoListener) EnterColumnReference(ctx *ColumnReferenceContext) {}

// ExitColumnReference is called when production columnReference is exited.
func (s *BaseSqlBase4GoListener) ExitColumnReference(ctx *ColumnReferenceContext) {}

// EnterNullLiteral is called when production nullLiteral is entered.
func (s *BaseSqlBase4GoListener) EnterNullLiteral(ctx *NullLiteralContext) {}

// ExitNullLiteral is called when production nullLiteral is exited.
func (s *BaseSqlBase4GoListener) ExitNullLiteral(ctx *NullLiteralContext) {}

// EnterRowConstructor is called when production rowConstructor is entered.
func (s *BaseSqlBase4GoListener) EnterRowConstructor(ctx *RowConstructorContext) {}

// ExitRowConstructor is called when production rowConstructor is exited.
func (s *BaseSqlBase4GoListener) ExitRowConstructor(ctx *RowConstructorContext) {}

// EnterSubscript is called when production subscript is entered.
func (s *BaseSqlBase4GoListener) EnterSubscript(ctx *SubscriptContext) {}

// ExitSubscript is called when production subscript is exited.
func (s *BaseSqlBase4GoListener) ExitSubscript(ctx *SubscriptContext) {}

// EnterSubqueryExpression is called when production subqueryExpression is entered.
func (s *BaseSqlBase4GoListener) EnterSubqueryExpression(ctx *SubqueryExpressionContext) {}

// ExitSubqueryExpression is called when production subqueryExpression is exited.
func (s *BaseSqlBase4GoListener) ExitSubqueryExpression(ctx *SubqueryExpressionContext) {}

// EnterBinaryLiteral is called when production binaryLiteral is entered.
func (s *BaseSqlBase4GoListener) EnterBinaryLiteral(ctx *BinaryLiteralContext) {}

// ExitBinaryLiteral is called when production binaryLiteral is exited.
func (s *BaseSqlBase4GoListener) ExitBinaryLiteral(ctx *BinaryLiteralContext) {}

// EnterCurrentUser is called when production currentUser is entered.
func (s *BaseSqlBase4GoListener) EnterCurrentUser(ctx *CurrentUserContext) {}

// ExitCurrentUser is called when production currentUser is exited.
func (s *BaseSqlBase4GoListener) ExitCurrentUser(ctx *CurrentUserContext) {}

// EnterExtract is called when production extract is entered.
func (s *BaseSqlBase4GoListener) EnterExtract(ctx *ExtractContext) {}

// ExitExtract is called when production extract is exited.
func (s *BaseSqlBase4GoListener) ExitExtract(ctx *ExtractContext) {}

// EnterStringLiteral is called when production stringLiteral is entered.
func (s *BaseSqlBase4GoListener) EnterStringLiteral(ctx *StringLiteralContext) {}

// ExitStringLiteral is called when production stringLiteral is exited.
func (s *BaseSqlBase4GoListener) ExitStringLiteral(ctx *StringLiteralContext) {}

// EnterArrayConstructor is called when production arrayConstructor is entered.
func (s *BaseSqlBase4GoListener) EnterArrayConstructor(ctx *ArrayConstructorContext) {}

// ExitArrayConstructor is called when production arrayConstructor is exited.
func (s *BaseSqlBase4GoListener) ExitArrayConstructor(ctx *ArrayConstructorContext) {}

// EnterFunctionCall is called when production functionCall is entered.
func (s *BaseSqlBase4GoListener) EnterFunctionCall(ctx *FunctionCallContext) {}

// ExitFunctionCall is called when production functionCall is exited.
func (s *BaseSqlBase4GoListener) ExitFunctionCall(ctx *FunctionCallContext) {}

// EnterExists is called when production exists is entered.
func (s *BaseSqlBase4GoListener) EnterExists(ctx *ExistsContext) {}

// ExitExists is called when production exists is exited.
func (s *BaseSqlBase4GoListener) ExitExists(ctx *ExistsContext) {}

// EnterPosition is called when production position is entered.
func (s *BaseSqlBase4GoListener) EnterPosition(ctx *PositionContext) {}

// ExitPosition is called when production position is exited.
func (s *BaseSqlBase4GoListener) ExitPosition(ctx *PositionContext) {}

// EnterSearchedCase is called when production searchedCase is entered.
func (s *BaseSqlBase4GoListener) EnterSearchedCase(ctx *SearchedCaseContext) {}

// ExitSearchedCase is called when production searchedCase is exited.
func (s *BaseSqlBase4GoListener) ExitSearchedCase(ctx *SearchedCaseContext) {}

// EnterGroupingOperation is called when production groupingOperation is entered.
func (s *BaseSqlBase4GoListener) EnterGroupingOperation(ctx *GroupingOperationContext) {}

// ExitGroupingOperation is called when production groupingOperation is exited.
func (s *BaseSqlBase4GoListener) ExitGroupingOperation(ctx *GroupingOperationContext) {}

// EnterNullTreatment is called when production nullTreatment is entered.
func (s *BaseSqlBase4GoListener) EnterNullTreatment(ctx *NullTreatmentContext) {}

// ExitNullTreatment is called when production nullTreatment is exited.
func (s *BaseSqlBase4GoListener) ExitNullTreatment(ctx *NullTreatmentContext) {}

// EnterTimeZoneInterval is called when production timeZoneInterval is entered.
func (s *BaseSqlBase4GoListener) EnterTimeZoneInterval(ctx *TimeZoneIntervalContext) {}

// ExitTimeZoneInterval is called when production timeZoneInterval is exited.
func (s *BaseSqlBase4GoListener) ExitTimeZoneInterval(ctx *TimeZoneIntervalContext) {}

// EnterTimeZoneString is called when production timeZoneString is entered.
func (s *BaseSqlBase4GoListener) EnterTimeZoneString(ctx *TimeZoneStringContext) {}

// ExitTimeZoneString is called when production timeZoneString is exited.
func (s *BaseSqlBase4GoListener) ExitTimeZoneString(ctx *TimeZoneStringContext) {}

// EnterComparisonOperator is called when production comparisonOperator is entered.
func (s *BaseSqlBase4GoListener) EnterComparisonOperator(ctx *ComparisonOperatorContext) {}

// ExitComparisonOperator is called when production comparisonOperator is exited.
func (s *BaseSqlBase4GoListener) ExitComparisonOperator(ctx *ComparisonOperatorContext) {}

// EnterComparisonQuantifier is called when production comparisonQuantifier is entered.
func (s *BaseSqlBase4GoListener) EnterComparisonQuantifier(ctx *ComparisonQuantifierContext) {}

// ExitComparisonQuantifier is called when production comparisonQuantifier is exited.
func (s *BaseSqlBase4GoListener) ExitComparisonQuantifier(ctx *ComparisonQuantifierContext) {}

// EnterBooleanValue is called when production booleanValue is entered.
func (s *BaseSqlBase4GoListener) EnterBooleanValue(ctx *BooleanValueContext) {}

// ExitBooleanValue is called when production booleanValue is exited.
func (s *BaseSqlBase4GoListener) ExitBooleanValue(ctx *BooleanValueContext) {}

// EnterInterval is called when production interval is entered.
func (s *BaseSqlBase4GoListener) EnterInterval(ctx *IntervalContext) {}

// ExitInterval is called when production interval is exited.
func (s *BaseSqlBase4GoListener) ExitInterval(ctx *IntervalContext) {}

// EnterIntervalField is called when production intervalField is entered.
func (s *BaseSqlBase4GoListener) EnterIntervalField(ctx *IntervalFieldContext) {}

// ExitIntervalField is called when production intervalField is exited.
func (s *BaseSqlBase4GoListener) ExitIntervalField(ctx *IntervalFieldContext) {}

// EnterNormalForm is called when production normalForm is entered.
func (s *BaseSqlBase4GoListener) EnterNormalForm(ctx *NormalFormContext) {}

// ExitNormalForm is called when production normalForm is exited.
func (s *BaseSqlBase4GoListener) ExitNormalForm(ctx *NormalFormContext) {}

// EnterTypes is called when production types is entered.
func (s *BaseSqlBase4GoListener) EnterTypes(ctx *TypesContext) {}

// ExitTypes is called when production types is exited.
func (s *BaseSqlBase4GoListener) ExitTypes(ctx *TypesContext) {}

// EnterType_r is called when production type_r is entered.
func (s *BaseSqlBase4GoListener) EnterType_r(ctx *Type_rContext) {}

// ExitType_r is called when production type_r is exited.
func (s *BaseSqlBase4GoListener) ExitType_r(ctx *Type_rContext) {}

// EnterTypeParameter is called when production typeParameter is entered.
func (s *BaseSqlBase4GoListener) EnterTypeParameter(ctx *TypeParameterContext) {}

// ExitTypeParameter is called when production typeParameter is exited.
func (s *BaseSqlBase4GoListener) ExitTypeParameter(ctx *TypeParameterContext) {}

// EnterBaseType is called when production baseType is entered.
func (s *BaseSqlBase4GoListener) EnterBaseType(ctx *BaseTypeContext) {}

// ExitBaseType is called when production baseType is exited.
func (s *BaseSqlBase4GoListener) ExitBaseType(ctx *BaseTypeContext) {}

// EnterWhenClause is called when production whenClause is entered.
func (s *BaseSqlBase4GoListener) EnterWhenClause(ctx *WhenClauseContext) {}

// ExitWhenClause is called when production whenClause is exited.
func (s *BaseSqlBase4GoListener) ExitWhenClause(ctx *WhenClauseContext) {}

// EnterFilter is called when production filter is entered.
func (s *BaseSqlBase4GoListener) EnterFilter(ctx *FilterContext) {}

// ExitFilter is called when production filter is exited.
func (s *BaseSqlBase4GoListener) ExitFilter(ctx *FilterContext) {}

// EnterOver is called when production over is entered.
func (s *BaseSqlBase4GoListener) EnterOver(ctx *OverContext) {}

// ExitOver is called when production over is exited.
func (s *BaseSqlBase4GoListener) ExitOver(ctx *OverContext) {}

// EnterWindowFrame is called when production windowFrame is entered.
func (s *BaseSqlBase4GoListener) EnterWindowFrame(ctx *WindowFrameContext) {}

// ExitWindowFrame is called when production windowFrame is exited.
func (s *BaseSqlBase4GoListener) ExitWindowFrame(ctx *WindowFrameContext) {}

// EnterUnboundedFrame is called when production unboundedFrame is entered.
func (s *BaseSqlBase4GoListener) EnterUnboundedFrame(ctx *UnboundedFrameContext) {}

// ExitUnboundedFrame is called when production unboundedFrame is exited.
func (s *BaseSqlBase4GoListener) ExitUnboundedFrame(ctx *UnboundedFrameContext) {}

// EnterCurrentRowBound is called when production currentRowBound is entered.
func (s *BaseSqlBase4GoListener) EnterCurrentRowBound(ctx *CurrentRowBoundContext) {}

// ExitCurrentRowBound is called when production currentRowBound is exited.
func (s *BaseSqlBase4GoListener) ExitCurrentRowBound(ctx *CurrentRowBoundContext) {}

// EnterBoundedFrame is called when production boundedFrame is entered.
func (s *BaseSqlBase4GoListener) EnterBoundedFrame(ctx *BoundedFrameContext) {}

// ExitBoundedFrame is called when production boundedFrame is exited.
func (s *BaseSqlBase4GoListener) ExitBoundedFrame(ctx *BoundedFrameContext) {}

// EnterExplainFormat is called when production explainFormat is entered.
func (s *BaseSqlBase4GoListener) EnterExplainFormat(ctx *ExplainFormatContext) {}

// ExitExplainFormat is called when production explainFormat is exited.
func (s *BaseSqlBase4GoListener) ExitExplainFormat(ctx *ExplainFormatContext) {}

// EnterExplainType is called when production explainType is entered.
func (s *BaseSqlBase4GoListener) EnterExplainType(ctx *ExplainTypeContext) {}

// ExitExplainType is called when production explainType is exited.
func (s *BaseSqlBase4GoListener) ExitExplainType(ctx *ExplainTypeContext) {}

// EnterIsolationLevel is called when production isolationLevel is entered.
func (s *BaseSqlBase4GoListener) EnterIsolationLevel(ctx *IsolationLevelContext) {}

// ExitIsolationLevel is called when production isolationLevel is exited.
func (s *BaseSqlBase4GoListener) ExitIsolationLevel(ctx *IsolationLevelContext) {}

// EnterTransactionAccessMode is called when production transactionAccessMode is entered.
func (s *BaseSqlBase4GoListener) EnterTransactionAccessMode(ctx *TransactionAccessModeContext) {}

// ExitTransactionAccessMode is called when production transactionAccessMode is exited.
func (s *BaseSqlBase4GoListener) ExitTransactionAccessMode(ctx *TransactionAccessModeContext) {}

// EnterReadUncommitted is called when production readUncommitted is entered.
func (s *BaseSqlBase4GoListener) EnterReadUncommitted(ctx *ReadUncommittedContext) {}

// ExitReadUncommitted is called when production readUncommitted is exited.
func (s *BaseSqlBase4GoListener) ExitReadUncommitted(ctx *ReadUncommittedContext) {}

// EnterReadCommitted is called when production readCommitted is entered.
func (s *BaseSqlBase4GoListener) EnterReadCommitted(ctx *ReadCommittedContext) {}

// ExitReadCommitted is called when production readCommitted is exited.
func (s *BaseSqlBase4GoListener) ExitReadCommitted(ctx *ReadCommittedContext) {}

// EnterRepeatableRead is called when production repeatableRead is entered.
func (s *BaseSqlBase4GoListener) EnterRepeatableRead(ctx *RepeatableReadContext) {}

// ExitRepeatableRead is called when production repeatableRead is exited.
func (s *BaseSqlBase4GoListener) ExitRepeatableRead(ctx *RepeatableReadContext) {}

// EnterSerializable is called when production serializable is entered.
func (s *BaseSqlBase4GoListener) EnterSerializable(ctx *SerializableContext) {}

// ExitSerializable is called when production serializable is exited.
func (s *BaseSqlBase4GoListener) ExitSerializable(ctx *SerializableContext) {}

// EnterPositionalArgument is called when production positionalArgument is entered.
func (s *BaseSqlBase4GoListener) EnterPositionalArgument(ctx *PositionalArgumentContext) {}

// ExitPositionalArgument is called when production positionalArgument is exited.
func (s *BaseSqlBase4GoListener) ExitPositionalArgument(ctx *PositionalArgumentContext) {}

// EnterNamedArgument is called when production namedArgument is entered.
func (s *BaseSqlBase4GoListener) EnterNamedArgument(ctx *NamedArgumentContext) {}

// ExitNamedArgument is called when production namedArgument is exited.
func (s *BaseSqlBase4GoListener) ExitNamedArgument(ctx *NamedArgumentContext) {}

// EnterPrivilege is called when production privilege is entered.
func (s *BaseSqlBase4GoListener) EnterPrivilege(ctx *PrivilegeContext) {}

// ExitPrivilege is called when production privilege is exited.
func (s *BaseSqlBase4GoListener) ExitPrivilege(ctx *PrivilegeContext) {}

// EnterQualifiedName is called when production qualifiedName is entered.
func (s *BaseSqlBase4GoListener) EnterQualifiedName(ctx *QualifiedNameContext) {}

// ExitQualifiedName is called when production qualifiedName is exited.
func (s *BaseSqlBase4GoListener) ExitQualifiedName(ctx *QualifiedNameContext) {}

// EnterCurrentUserGrantor is called when production currentUserGrantor is entered.
func (s *BaseSqlBase4GoListener) EnterCurrentUserGrantor(ctx *CurrentUserGrantorContext) {}

// ExitCurrentUserGrantor is called when production currentUserGrantor is exited.
func (s *BaseSqlBase4GoListener) ExitCurrentUserGrantor(ctx *CurrentUserGrantorContext) {}

// EnterCurrentRoleGrantor is called when production currentRoleGrantor is entered.
func (s *BaseSqlBase4GoListener) EnterCurrentRoleGrantor(ctx *CurrentRoleGrantorContext) {}

// ExitCurrentRoleGrantor is called when production currentRoleGrantor is exited.
func (s *BaseSqlBase4GoListener) ExitCurrentRoleGrantor(ctx *CurrentRoleGrantorContext) {}

// EnterSpecifiedPrincipal is called when production specifiedPrincipal is entered.
func (s *BaseSqlBase4GoListener) EnterSpecifiedPrincipal(ctx *SpecifiedPrincipalContext) {}

// ExitSpecifiedPrincipal is called when production specifiedPrincipal is exited.
func (s *BaseSqlBase4GoListener) ExitSpecifiedPrincipal(ctx *SpecifiedPrincipalContext) {}

// EnterUserPrincipal is called when production userPrincipal is entered.
func (s *BaseSqlBase4GoListener) EnterUserPrincipal(ctx *UserPrincipalContext) {}

// ExitUserPrincipal is called when production userPrincipal is exited.
func (s *BaseSqlBase4GoListener) ExitUserPrincipal(ctx *UserPrincipalContext) {}

// EnterRolePrincipal is called when production rolePrincipal is entered.
func (s *BaseSqlBase4GoListener) EnterRolePrincipal(ctx *RolePrincipalContext) {}

// ExitRolePrincipal is called when production rolePrincipal is exited.
func (s *BaseSqlBase4GoListener) ExitRolePrincipal(ctx *RolePrincipalContext) {}

// EnterUnspecifiedPrincipal is called when production unspecifiedPrincipal is entered.
func (s *BaseSqlBase4GoListener) EnterUnspecifiedPrincipal(ctx *UnspecifiedPrincipalContext) {}

// ExitUnspecifiedPrincipal is called when production unspecifiedPrincipal is exited.
func (s *BaseSqlBase4GoListener) ExitUnspecifiedPrincipal(ctx *UnspecifiedPrincipalContext) {}

// EnterRoles is called when production roles is entered.
func (s *BaseSqlBase4GoListener) EnterRoles(ctx *RolesContext) {}

// ExitRoles is called when production roles is exited.
func (s *BaseSqlBase4GoListener) ExitRoles(ctx *RolesContext) {}

// EnterUnquotedIdentifier is called when production unquotedIdentifier is entered.
func (s *BaseSqlBase4GoListener) EnterUnquotedIdentifier(ctx *UnquotedIdentifierContext) {}

// ExitUnquotedIdentifier is called when production unquotedIdentifier is exited.
func (s *BaseSqlBase4GoListener) ExitUnquotedIdentifier(ctx *UnquotedIdentifierContext) {}

// EnterQuotedIdentifier is called when production quotedIdentifier is entered.
func (s *BaseSqlBase4GoListener) EnterQuotedIdentifier(ctx *QuotedIdentifierContext) {}

// ExitQuotedIdentifier is called when production quotedIdentifier is exited.
func (s *BaseSqlBase4GoListener) ExitQuotedIdentifier(ctx *QuotedIdentifierContext) {}

// EnterBackQuotedIdentifier is called when production backQuotedIdentifier is entered.
func (s *BaseSqlBase4GoListener) EnterBackQuotedIdentifier(ctx *BackQuotedIdentifierContext) {}

// ExitBackQuotedIdentifier is called when production backQuotedIdentifier is exited.
func (s *BaseSqlBase4GoListener) ExitBackQuotedIdentifier(ctx *BackQuotedIdentifierContext) {}

// EnterDigitIdentifier is called when production digitIdentifier is entered.
func (s *BaseSqlBase4GoListener) EnterDigitIdentifier(ctx *DigitIdentifierContext) {}

// ExitDigitIdentifier is called when production digitIdentifier is exited.
func (s *BaseSqlBase4GoListener) ExitDigitIdentifier(ctx *DigitIdentifierContext) {}

// EnterDecimalLiteral is called when production decimalLiteral is entered.
func (s *BaseSqlBase4GoListener) EnterDecimalLiteral(ctx *DecimalLiteralContext) {}

// ExitDecimalLiteral is called when production decimalLiteral is exited.
func (s *BaseSqlBase4GoListener) ExitDecimalLiteral(ctx *DecimalLiteralContext) {}

// EnterDoubleLiteral is called when production doubleLiteral is entered.
func (s *BaseSqlBase4GoListener) EnterDoubleLiteral(ctx *DoubleLiteralContext) {}

// ExitDoubleLiteral is called when production doubleLiteral is exited.
func (s *BaseSqlBase4GoListener) ExitDoubleLiteral(ctx *DoubleLiteralContext) {}

// EnterIntegerLiteral is called when production integerLiteral is entered.
func (s *BaseSqlBase4GoListener) EnterIntegerLiteral(ctx *IntegerLiteralContext) {}

// ExitIntegerLiteral is called when production integerLiteral is exited.
func (s *BaseSqlBase4GoListener) ExitIntegerLiteral(ctx *IntegerLiteralContext) {}

// EnterNonReserved is called when production nonReserved is entered.
func (s *BaseSqlBase4GoListener) EnterNonReserved(ctx *NonReservedContext) {}

// ExitNonReserved is called when production nonReserved is exited.
func (s *BaseSqlBase4GoListener) ExitNonReserved(ctx *NonReservedContext) {}

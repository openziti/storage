// Code generated from ZitiQl.g4 by ANTLR 4.7.2. DO NOT EDIT.

package zitiql // ZitiQl
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseZitiQlListener is a complete listener for a parse tree produced by ZitiQlParser.
type BaseZitiQlListener struct{}

var _ ZitiQlListener = &BaseZitiQlListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseZitiQlListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseZitiQlListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseZitiQlListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseZitiQlListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterString_array is called when production string_array is entered.
func (s *BaseZitiQlListener) EnterString_array(ctx *String_arrayContext) {}

// ExitString_array is called when production string_array is exited.
func (s *BaseZitiQlListener) ExitString_array(ctx *String_arrayContext) {}

// EnterNumber_array is called when production number_array is entered.
func (s *BaseZitiQlListener) EnterNumber_array(ctx *Number_arrayContext) {}

// ExitNumber_array is called when production number_array is exited.
func (s *BaseZitiQlListener) ExitNumber_array(ctx *Number_arrayContext) {}

// EnterDatetime_array is called when production datetime_array is entered.
func (s *BaseZitiQlListener) EnterDatetime_array(ctx *Datetime_arrayContext) {}

// ExitDatetime_array is called when production datetime_array is exited.
func (s *BaseZitiQlListener) ExitDatetime_array(ctx *Datetime_arrayContext) {}

// EnterEnd is called when production End is entered.
func (s *BaseZitiQlListener) EnterEnd(ctx *EndContext) {}

// ExitEnd is called when production End is exited.
func (s *BaseZitiQlListener) ExitEnd(ctx *EndContext) {}

// EnterQueryStmt is called when production QueryStmt is entered.
func (s *BaseZitiQlListener) EnterQueryStmt(ctx *QueryStmtContext) {}

// ExitQueryStmt is called when production QueryStmt is exited.
func (s *BaseZitiQlListener) ExitQueryStmt(ctx *QueryStmtContext) {}

// EnterSkipExpr is called when production SkipExpr is entered.
func (s *BaseZitiQlListener) EnterSkipExpr(ctx *SkipExprContext) {}

// ExitSkipExpr is called when production SkipExpr is exited.
func (s *BaseZitiQlListener) ExitSkipExpr(ctx *SkipExprContext) {}

// EnterLimitExpr is called when production LimitExpr is entered.
func (s *BaseZitiQlListener) EnterLimitExpr(ctx *LimitExprContext) {}

// ExitLimitExpr is called when production LimitExpr is exited.
func (s *BaseZitiQlListener) ExitLimitExpr(ctx *LimitExprContext) {}

// EnterSortByExpr is called when production SortByExpr is entered.
func (s *BaseZitiQlListener) EnterSortByExpr(ctx *SortByExprContext) {}

// ExitSortByExpr is called when production SortByExpr is exited.
func (s *BaseZitiQlListener) ExitSortByExpr(ctx *SortByExprContext) {}

// EnterSortFieldExpr is called when production SortFieldExpr is entered.
func (s *BaseZitiQlListener) EnterSortFieldExpr(ctx *SortFieldExprContext) {}

// ExitSortFieldExpr is called when production SortFieldExpr is exited.
func (s *BaseZitiQlListener) ExitSortFieldExpr(ctx *SortFieldExprContext) {}

// EnterAndExpr is called when production AndExpr is entered.
func (s *BaseZitiQlListener) EnterAndExpr(ctx *AndExprContext) {}

// ExitAndExpr is called when production AndExpr is exited.
func (s *BaseZitiQlListener) ExitAndExpr(ctx *AndExprContext) {}

// EnterGroup is called when production Group is entered.
func (s *BaseZitiQlListener) EnterGroup(ctx *GroupContext) {}

// ExitGroup is called when production Group is exited.
func (s *BaseZitiQlListener) ExitGroup(ctx *GroupContext) {}

// EnterBoolConst is called when production BoolConst is entered.
func (s *BaseZitiQlListener) EnterBoolConst(ctx *BoolConstContext) {}

// ExitBoolConst is called when production BoolConst is exited.
func (s *BaseZitiQlListener) ExitBoolConst(ctx *BoolConstContext) {}

// EnterIsEmptyFunction is called when production IsEmptyFunction is entered.
func (s *BaseZitiQlListener) EnterIsEmptyFunction(ctx *IsEmptyFunctionContext) {}

// ExitIsEmptyFunction is called when production IsEmptyFunction is exited.
func (s *BaseZitiQlListener) ExitIsEmptyFunction(ctx *IsEmptyFunctionContext) {}

// EnterNotExpr is called when production NotExpr is entered.
func (s *BaseZitiQlListener) EnterNotExpr(ctx *NotExprContext) {}

// ExitNotExpr is called when production NotExpr is exited.
func (s *BaseZitiQlListener) ExitNotExpr(ctx *NotExprContext) {}

// EnterOperationOp is called when production OperationOp is entered.
func (s *BaseZitiQlListener) EnterOperationOp(ctx *OperationOpContext) {}

// ExitOperationOp is called when production OperationOp is exited.
func (s *BaseZitiQlListener) ExitOperationOp(ctx *OperationOpContext) {}

// EnterOrExpr is called when production OrExpr is entered.
func (s *BaseZitiQlListener) EnterOrExpr(ctx *OrExprContext) {}

// ExitOrExpr is called when production OrExpr is exited.
func (s *BaseZitiQlListener) ExitOrExpr(ctx *OrExprContext) {}

// EnterBoolSymbol is called when production BoolSymbol is entered.
func (s *BaseZitiQlListener) EnterBoolSymbol(ctx *BoolSymbolContext) {}

// ExitBoolSymbol is called when production BoolSymbol is exited.
func (s *BaseZitiQlListener) ExitBoolSymbol(ctx *BoolSymbolContext) {}

// EnterInStringArrayOp is called when production InStringArrayOp is entered.
func (s *BaseZitiQlListener) EnterInStringArrayOp(ctx *InStringArrayOpContext) {}

// ExitInStringArrayOp is called when production InStringArrayOp is exited.
func (s *BaseZitiQlListener) ExitInStringArrayOp(ctx *InStringArrayOpContext) {}

// EnterInNumberArrayOp is called when production InNumberArrayOp is entered.
func (s *BaseZitiQlListener) EnterInNumberArrayOp(ctx *InNumberArrayOpContext) {}

// ExitInNumberArrayOp is called when production InNumberArrayOp is exited.
func (s *BaseZitiQlListener) ExitInNumberArrayOp(ctx *InNumberArrayOpContext) {}

// EnterInDatetimeArrayOp is called when production InDatetimeArrayOp is entered.
func (s *BaseZitiQlListener) EnterInDatetimeArrayOp(ctx *InDatetimeArrayOpContext) {}

// ExitInDatetimeArrayOp is called when production InDatetimeArrayOp is exited.
func (s *BaseZitiQlListener) ExitInDatetimeArrayOp(ctx *InDatetimeArrayOpContext) {}

// EnterBetweenNumberOp is called when production BetweenNumberOp is entered.
func (s *BaseZitiQlListener) EnterBetweenNumberOp(ctx *BetweenNumberOpContext) {}

// ExitBetweenNumberOp is called when production BetweenNumberOp is exited.
func (s *BaseZitiQlListener) ExitBetweenNumberOp(ctx *BetweenNumberOpContext) {}

// EnterBetweenDateOp is called when production BetweenDateOp is entered.
func (s *BaseZitiQlListener) EnterBetweenDateOp(ctx *BetweenDateOpContext) {}

// ExitBetweenDateOp is called when production BetweenDateOp is exited.
func (s *BaseZitiQlListener) ExitBetweenDateOp(ctx *BetweenDateOpContext) {}

// EnterBinaryLessThanNumberOp is called when production BinaryLessThanNumberOp is entered.
func (s *BaseZitiQlListener) EnterBinaryLessThanNumberOp(ctx *BinaryLessThanNumberOpContext) {}

// ExitBinaryLessThanNumberOp is called when production BinaryLessThanNumberOp is exited.
func (s *BaseZitiQlListener) ExitBinaryLessThanNumberOp(ctx *BinaryLessThanNumberOpContext) {}

// EnterBinaryLessThanDatetimeOp is called when production BinaryLessThanDatetimeOp is entered.
func (s *BaseZitiQlListener) EnterBinaryLessThanDatetimeOp(ctx *BinaryLessThanDatetimeOpContext) {}

// ExitBinaryLessThanDatetimeOp is called when production BinaryLessThanDatetimeOp is exited.
func (s *BaseZitiQlListener) ExitBinaryLessThanDatetimeOp(ctx *BinaryLessThanDatetimeOpContext) {}

// EnterBinaryGreaterThanNumberOp is called when production BinaryGreaterThanNumberOp is entered.
func (s *BaseZitiQlListener) EnterBinaryGreaterThanNumberOp(ctx *BinaryGreaterThanNumberOpContext) {}

// ExitBinaryGreaterThanNumberOp is called when production BinaryGreaterThanNumberOp is exited.
func (s *BaseZitiQlListener) ExitBinaryGreaterThanNumberOp(ctx *BinaryGreaterThanNumberOpContext) {}

// EnterBinaryGreaterThanDatetimeOp is called when production BinaryGreaterThanDatetimeOp is entered.
func (s *BaseZitiQlListener) EnterBinaryGreaterThanDatetimeOp(ctx *BinaryGreaterThanDatetimeOpContext) {
}

// ExitBinaryGreaterThanDatetimeOp is called when production BinaryGreaterThanDatetimeOp is exited.
func (s *BaseZitiQlListener) ExitBinaryGreaterThanDatetimeOp(ctx *BinaryGreaterThanDatetimeOpContext) {
}

// EnterBinaryEqualToStringOp is called when production BinaryEqualToStringOp is entered.
func (s *BaseZitiQlListener) EnterBinaryEqualToStringOp(ctx *BinaryEqualToStringOpContext) {}

// ExitBinaryEqualToStringOp is called when production BinaryEqualToStringOp is exited.
func (s *BaseZitiQlListener) ExitBinaryEqualToStringOp(ctx *BinaryEqualToStringOpContext) {}

// EnterBinaryEqualToNumberOp is called when production BinaryEqualToNumberOp is entered.
func (s *BaseZitiQlListener) EnterBinaryEqualToNumberOp(ctx *BinaryEqualToNumberOpContext) {}

// ExitBinaryEqualToNumberOp is called when production BinaryEqualToNumberOp is exited.
func (s *BaseZitiQlListener) ExitBinaryEqualToNumberOp(ctx *BinaryEqualToNumberOpContext) {}

// EnterBinaryEqualToDatetimeOp is called when production BinaryEqualToDatetimeOp is entered.
func (s *BaseZitiQlListener) EnterBinaryEqualToDatetimeOp(ctx *BinaryEqualToDatetimeOpContext) {}

// ExitBinaryEqualToDatetimeOp is called when production BinaryEqualToDatetimeOp is exited.
func (s *BaseZitiQlListener) ExitBinaryEqualToDatetimeOp(ctx *BinaryEqualToDatetimeOpContext) {}

// EnterBinaryEqualToBoolOp is called when production BinaryEqualToBoolOp is entered.
func (s *BaseZitiQlListener) EnterBinaryEqualToBoolOp(ctx *BinaryEqualToBoolOpContext) {}

// ExitBinaryEqualToBoolOp is called when production BinaryEqualToBoolOp is exited.
func (s *BaseZitiQlListener) ExitBinaryEqualToBoolOp(ctx *BinaryEqualToBoolOpContext) {}

// EnterBinaryEqualToNullOp is called when production BinaryEqualToNullOp is entered.
func (s *BaseZitiQlListener) EnterBinaryEqualToNullOp(ctx *BinaryEqualToNullOpContext) {}

// ExitBinaryEqualToNullOp is called when production BinaryEqualToNullOp is exited.
func (s *BaseZitiQlListener) ExitBinaryEqualToNullOp(ctx *BinaryEqualToNullOpContext) {}

// EnterBinaryContainsOp is called when production BinaryContainsOp is entered.
func (s *BaseZitiQlListener) EnterBinaryContainsOp(ctx *BinaryContainsOpContext) {}

// ExitBinaryContainsOp is called when production BinaryContainsOp is exited.
func (s *BaseZitiQlListener) ExitBinaryContainsOp(ctx *BinaryContainsOpContext) {}

// EnterBinary_lhs is called when production binary_lhs is entered.
func (s *BaseZitiQlListener) EnterBinary_lhs(ctx *Binary_lhsContext) {}

// ExitBinary_lhs is called when production binary_lhs is exited.
func (s *BaseZitiQlListener) ExitBinary_lhs(ctx *Binary_lhsContext) {}

// EnterSetFunction is called when production SetFunction is entered.
func (s *BaseZitiQlListener) EnterSetFunction(ctx *SetFunctionContext) {}

// ExitSetFunction is called when production SetFunction is exited.
func (s *BaseZitiQlListener) ExitSetFunction(ctx *SetFunctionContext) {}

// EnterSet_expr is called when production set_expr is entered.
func (s *BaseZitiQlListener) EnterSet_expr(ctx *Set_exprContext) {}

// ExitSet_expr is called when production set_expr is exited.
func (s *BaseZitiQlListener) ExitSet_expr(ctx *Set_exprContext) {}

// EnterSubQuery is called when production SubQuery is entered.
func (s *BaseZitiQlListener) EnterSubQuery(ctx *SubQueryContext) {}

// ExitSubQuery is called when production SubQuery is exited.
func (s *BaseZitiQlListener) ExitSubQuery(ctx *SubQueryContext) {}

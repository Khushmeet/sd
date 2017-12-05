// Generated from sd.g4 by ANTLR 4.7.

package parser // sd

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BasesdListener is a complete listener for a parse tree produced by sdParser.
type BasesdListener struct{}

var _ sdListener = &BasesdListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasesdListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasesdListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasesdListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasesdListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterSymbol is called when production symbol is entered.
func (s *BasesdListener) EnterSymbol(ctx *SymbolContext) {}

// ExitSymbol is called when production symbol is exited.
func (s *BasesdListener) ExitSymbol(ctx *SymbolContext) {}

// EnterNumber is called when production number is entered.
func (s *BasesdListener) EnterNumber(ctx *NumberContext) {}

// ExitNumber is called when production number is exited.
func (s *BasesdListener) ExitNumber(ctx *NumberContext) {}

// EnterExpression is called when production expression is entered.
func (s *BasesdListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BasesdListener) ExitExpression(ctx *ExpressionContext) {}

// EnterCheckZero is called when production checkZero is entered.
func (s *BasesdListener) EnterCheckZero(ctx *CheckZeroContext) {}

// ExitCheckZero is called when production checkZero is exited.
func (s *BasesdListener) ExitCheckZero(ctx *CheckZeroContext) {}

// EnterAbstraction is called when production abstraction is entered.
func (s *BasesdListener) EnterAbstraction(ctx *AbstractionContext) {}

// ExitAbstraction is called when production abstraction is exited.
func (s *BasesdListener) ExitAbstraction(ctx *AbstractionContext) {}

// EnterIfElseExpr is called when production ifElseExpr is entered.
func (s *BasesdListener) EnterIfElseExpr(ctx *IfElseExprContext) {}

// ExitIfElseExpr is called when production ifElseExpr is exited.
func (s *BasesdListener) ExitIfElseExpr(ctx *IfElseExprContext) {}

// EnterApplicatio is called when production applicatio is entered.
func (s *BasesdListener) EnterApplicatio(ctx *ApplicatioContext) {}

// ExitApplicatio is called when production applicatio is exited.
func (s *BasesdListener) ExitApplicatio(ctx *ApplicatioContext) {}

// EnterArithmeicExps is called when production arithmeicExps is entered.
func (s *BasesdListener) EnterArithmeicExps(ctx *ArithmeicExpsContext) {}

// ExitArithmeicExps is called when production arithmeicExps is exited.
func (s *BasesdListener) ExitArithmeicExps(ctx *ArithmeicExpsContext) {}

// EnterArithmeticOps is called when production arithmeticOps is entered.
func (s *BasesdListener) EnterArithmeticOps(ctx *ArithmeticOpsContext) {}

// ExitArithmeticOps is called when production arithmeticOps is exited.
func (s *BasesdListener) ExitArithmeticOps(ctx *ArithmeticOpsContext) {}

// EnterIfThen is called when production ifThen is entered.
func (s *BasesdListener) EnterIfThen(ctx *IfThenContext) {}

// ExitIfThen is called when production ifThen is exited.
func (s *BasesdListener) ExitIfThen(ctx *IfThenContext) {}

// EnterFuncType is called when production funcType is entered.
func (s *BasesdListener) EnterFuncType(ctx *FuncTypeContext) {}

// ExitFuncType is called when production funcType is exited.
func (s *BasesdListener) ExitFuncType(ctx *FuncTypeContext) {}

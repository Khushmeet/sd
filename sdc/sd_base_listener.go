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

// EnterExpr is called when production expr is entered.
func (s *BasesdListener) EnterExpr(ctx *ExprContext) {}

// ExitExpr is called when production expr is exited.
func (s *BasesdListener) ExitExpr(ctx *ExprContext) {}

// EnterArithmeticOps is called when production arithmeticOps is entered.
func (s *BasesdListener) EnterArithmeticOps(ctx *ArithmeticOpsContext) {}

// ExitArithmeticOps is called when production arithmeticOps is exited.
func (s *BasesdListener) ExitArithmeticOps(ctx *ArithmeticOpsContext) {}

// EnterIfThen is called when production ifThen is entered.
func (s *BasesdListener) EnterIfThen(ctx *IfThenContext) {}

// ExitIfThen is called when production ifThen is exited.
func (s *BasesdListener) ExitIfThen(ctx *IfThenContext) {}

// EnterR_type is called when production r_type is entered.
func (s *BasesdListener) EnterR_type(ctx *R_typeContext) {}

// ExitR_type is called when production r_type is exited.
func (s *BasesdListener) ExitR_type(ctx *R_typeContext) {}

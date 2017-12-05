// Generated from sd.g4 by ANTLR 4.7.

package parser // sd

import "github.com/antlr/antlr4/runtime/Go/antlr"

// sdListener is a complete listener for a parse tree produced by sdParser.
type sdListener interface {
	antlr.ParseTreeListener

	// EnterExpr is called when entering the expr production.
	EnterExpr(c *ExprContext)

	// EnterArithmeticOps is called when entering the arithmeticOps production.
	EnterArithmeticOps(c *ArithmeticOpsContext)

	// EnterIfThen is called when entering the ifThen production.
	EnterIfThen(c *IfThenContext)

	// EnterR_type is called when entering the r_type production.
	EnterR_type(c *R_typeContext)

	// ExitExpr is called when exiting the expr production.
	ExitExpr(c *ExprContext)

	// ExitArithmeticOps is called when exiting the arithmeticOps production.
	ExitArithmeticOps(c *ArithmeticOpsContext)

	// ExitIfThen is called when exiting the ifThen production.
	ExitIfThen(c *IfThenContext)

	// ExitR_type is called when exiting the r_type production.
	ExitR_type(c *R_typeContext)
}

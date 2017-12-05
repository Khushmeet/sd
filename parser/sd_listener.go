// Generated from sd.g4 by ANTLR 4.7.

package parser // sd

import "github.com/antlr/antlr4/runtime/Go/antlr"

// sdListener is a complete listener for a parse tree produced by sdParser.
type sdListener interface {
	antlr.ParseTreeListener

	// EnterSymbol is called when entering the symbol production.
	EnterSymbol(c *SymbolContext)

	// EnterNumber is called when entering the number production.
	EnterNumber(c *NumberContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterCheckZero is called when entering the checkZero production.
	EnterCheckZero(c *CheckZeroContext)

	// EnterAbstraction is called when entering the abstraction production.
	EnterAbstraction(c *AbstractionContext)

	// EnterIfElseExpr is called when entering the ifElseExpr production.
	EnterIfElseExpr(c *IfElseExprContext)

	// EnterApplicatio is called when entering the applicatio production.
	EnterApplicatio(c *ApplicatioContext)

	// EnterArithmeicExps is called when entering the arithmeicExps production.
	EnterArithmeicExps(c *ArithmeicExpsContext)

	// EnterArithmeticOps is called when entering the arithmeticOps production.
	EnterArithmeticOps(c *ArithmeticOpsContext)

	// EnterIfThen is called when entering the ifThen production.
	EnterIfThen(c *IfThenContext)

	// EnterR_type is called when entering the r_type production.
	EnterR_type(c *R_typeContext)

	// ExitSymbol is called when exiting the symbol production.
	ExitSymbol(c *SymbolContext)

	// ExitNumber is called when exiting the number production.
	ExitNumber(c *NumberContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitCheckZero is called when exiting the checkZero production.
	ExitCheckZero(c *CheckZeroContext)

	// ExitAbstraction is called when exiting the abstraction production.
	ExitAbstraction(c *AbstractionContext)

	// ExitIfElseExpr is called when exiting the ifElseExpr production.
	ExitIfElseExpr(c *IfElseExprContext)

	// ExitApplicatio is called when exiting the applicatio production.
	ExitApplicatio(c *ApplicatioContext)

	// ExitArithmeicExps is called when exiting the arithmeicExps production.
	ExitArithmeicExps(c *ArithmeicExpsContext)

	// ExitArithmeticOps is called when exiting the arithmeticOps production.
	ExitArithmeticOps(c *ArithmeticOpsContext)

	// ExitIfThen is called when exiting the ifThen production.
	ExitIfThen(c *IfThenContext)

	// ExitR_type is called when exiting the r_type production.
	ExitR_type(c *R_typeContext)
}

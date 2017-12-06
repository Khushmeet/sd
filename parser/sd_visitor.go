// Generated from sd.g4 by ANTLR 4.7.

package parser // sd

import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by sdParser.
type sdVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by sdParser#symbol.
	VisitSymbol(ctx *SymbolContext) interface{}

	// Visit a parse tree produced by sdParser#number.
	VisitNumber(ctx *NumberContext) interface{}

	// Visit a parse tree produced by sdParser#expression.
	VisitExpression(ctx *ExpressionContext) interface{}

	// Visit a parse tree produced by sdParser#falseLiteral.
	VisitFalseLiteral(ctx *FalseLiteralContext) interface{}

	// Visit a parse tree produced by sdParser#checkZero.
	VisitCheckZero(ctx *CheckZeroContext) interface{}

	// Visit a parse tree produced by sdParser#abstraction.
	VisitAbstraction(ctx *AbstractionContext) interface{}

	// Visit a parse tree produced by sdParser#ifElseExpr.
	VisitIfElseExpr(ctx *IfElseExprContext) interface{}

	// Visit a parse tree produced by sdParser#applicatio.
	VisitApplicatio(ctx *ApplicatioContext) interface{}

	// Visit a parse tree produced by sdParser#trueLiteral.
	VisitTrueLiteral(ctx *TrueLiteralContext) interface{}

	// Visit a parse tree produced by sdParser#arithmeicExps.
	VisitArithmeicExps(ctx *ArithmeicExpsContext) interface{}

	// Visit a parse tree produced by sdParser#arithmeticOps.
	VisitArithmeticOps(ctx *ArithmeticOpsContext) interface{}

	// Visit a parse tree produced by sdParser#ifThen.
	VisitIfThen(ctx *IfThenContext) interface{}
}

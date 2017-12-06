// Generated from sd.g4 by ANTLR 4.7.

package parser // sd

import "github.com/antlr/antlr4/runtime/Go/antlr"

type BasesdVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BasesdVisitor) VisitSymbol(ctx *SymbolContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasesdVisitor) VisitNumber(ctx *NumberContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasesdVisitor) VisitExpression(ctx *ExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasesdVisitor) VisitFalseLiteral(ctx *FalseLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasesdVisitor) VisitCheckZero(ctx *CheckZeroContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasesdVisitor) VisitAbstraction(ctx *AbstractionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasesdVisitor) VisitIfElseExpr(ctx *IfElseExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasesdVisitor) VisitApplicatio(ctx *ApplicatioContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasesdVisitor) VisitTrueLiteral(ctx *TrueLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasesdVisitor) VisitArithmeicExps(ctx *ArithmeicExpsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasesdVisitor) VisitArithmeticOps(ctx *ArithmeticOpsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasesdVisitor) VisitIfThen(ctx *IfThenContext) interface{} {
	return v.VisitChildren(ctx)
}

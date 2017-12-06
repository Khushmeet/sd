package main

import (
	"fmt"
	SdParser "sd/parser"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

/*
SdTranspiler struct for transpilling sd to javascript
*/
type SdEvaluator struct {
	*SdParser.BasesdVisitor
}

/*
NewSdTranspilerListener creates new sdTranspilerListener object
*/
func NewSdEvaluatorListener() *SdEvaluator {
	return new(SdEvaluator)
}

/*
EnterEveryRule gets called when tree walker enter a node in parse tree
*/
func (sd *SdEvaluator) VisitTrueLiteral(ctx SdParser.TrueLiteralContext) bool {
	return ctx.TRUE().GetText()
}

func (sd *SdEvaluator) VisitFalseLiteral(ctx SdParser.FalseLiteralContext) bool {
	return bool(ctx.FALSE().GetText())
}

func (sd *SdEvaluator) VisitSymbol(ctx SdParser.SymbolContext) (int, error) {
	return SymbolTable.Lookup(ctx.ID().GetText())
}

func (sd *SdEvaluator) VisitNumber(ctx SdParser.NumberContext) int {
	return int(ctx.DIGIT().GetText())
}

func (sd *SdEvaluator) VisitAbstraction(ctx SdParser.AbstractionContext) {
	return func(x) {
		m := make(map[string]int)
		m[ctx.ID.GetText()] = x 
		s := Scope{
			S:m,
		}
		SymbolTable.Push(s)
		return visit(ctx.expr())
	}
}

func (sd *SdEvaluator) VisitApplication(ctx SdParser.ApplicatioContext) int {
	left := visit(ctx.expr(0))
	right := visit(ctx.expr(1))
	return left(right)
}

func (sd *SdEvaluator) VisitExpression(ctx SdParser.ExpressionContext) int {
	return visit(ctx.Expr())
}

func (sd *SdEvaluator) VisitArithmeicExps(ctx SdParser.ArithmeicExpsContext) int {
	if ctx.ArithmeticOps() == "succ" {
		res := visit(ctx.Expr())
		return res + 1
	}
	if ctx.ArithmeticOps() == "pred" {
		res := visit(ctx.Expr())
		if res >= 0 {
			return res + 1
		}
		return res
	}
}

func (sd *SdEvaluator) VisitCheckZero(ctx SdParser.CheckZeroContext) {
	return visit(ctx.Expr()) == 0
}

func (sd *SdEvaluator) VisitIfElseExpr(ctx SdParser.IfElseExprContext) {
	if visit(ctx.Expr(0)) {
		return visit(ctx.Expr(1))
	}
	return visit(ctx.Expr(2))
}
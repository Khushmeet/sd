// Generated from sd.g4 by ANTLR 4.7.

package parser // sd

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 20, 46, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3,
	2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3,
	2, 5, 2, 28, 10, 2, 3, 2, 3, 2, 7, 2, 32, 10, 2, 12, 2, 14, 2, 35, 11,
	2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 2, 3, 2,
	5, 2, 4, 6, 2, 4, 3, 2, 4, 5, 3, 2, 12, 13, 2, 51, 2, 27, 3, 2, 2, 2, 4,
	36, 3, 2, 2, 2, 6, 38, 3, 2, 2, 2, 8, 9, 8, 2, 1, 2, 9, 28, 7, 18, 2, 2,
	10, 28, 7, 19, 2, 2, 11, 28, 7, 14, 2, 2, 12, 28, 7, 15, 2, 2, 13, 14,
	7, 3, 2, 2, 14, 15, 7, 14, 2, 2, 15, 16, 9, 2, 2, 2, 16, 28, 5, 2, 2, 8,
	17, 18, 7, 6, 2, 2, 18, 19, 5, 2, 2, 2, 19, 20, 7, 7, 2, 2, 20, 28, 3,
	2, 2, 2, 21, 22, 5, 4, 3, 2, 22, 23, 5, 2, 2, 5, 23, 28, 3, 2, 2, 2, 24,
	25, 7, 11, 2, 2, 25, 28, 5, 2, 2, 4, 26, 28, 5, 6, 4, 2, 27, 8, 3, 2, 2,
	2, 27, 10, 3, 2, 2, 2, 27, 11, 3, 2, 2, 2, 27, 12, 3, 2, 2, 2, 27, 13,
	3, 2, 2, 2, 27, 17, 3, 2, 2, 2, 27, 21, 3, 2, 2, 2, 27, 24, 3, 2, 2, 2,
	27, 26, 3, 2, 2, 2, 28, 33, 3, 2, 2, 2, 29, 30, 12, 7, 2, 2, 30, 32, 5,
	2, 2, 8, 31, 29, 3, 2, 2, 2, 32, 35, 3, 2, 2, 2, 33, 31, 3, 2, 2, 2, 33,
	34, 3, 2, 2, 2, 34, 3, 3, 2, 2, 2, 35, 33, 3, 2, 2, 2, 36, 37, 9, 3, 2,
	2, 37, 5, 3, 2, 2, 2, 38, 39, 7, 8, 2, 2, 39, 40, 5, 2, 2, 2, 40, 41, 7,
	10, 2, 2, 41, 42, 5, 2, 2, 2, 42, 43, 7, 9, 2, 2, 43, 44, 5, 2, 2, 2, 44,
	7, 3, 2, 2, 2, 4, 27, 33,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'\u03BB'", "'=>'", "'\u2192'", "'('", "')'", "'if'", "'else'", "'then'",
	"'iszero'", "'succ'", "'pred'", "", "", "'Nat'", "'Bool'", "'true'", "'false'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "IF", "ELSE", "THEN", "ISZERO", "SUCC", "PRED",
	"ID", "DIGIT", "NAT", "BOOL", "TRUE", "FALSE", "WS",
}

var ruleNames = []string{
	"expr", "arithmeticOps", "ifThen",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type sdParser struct {
	*antlr.BaseParser
}

func NewsdParser(input antlr.TokenStream) *sdParser {
	this := new(sdParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "sd.g4"

	return this
}

// sdParser tokens.
const (
	sdParserEOF    = antlr.TokenEOF
	sdParserT__0   = 1
	sdParserT__1   = 2
	sdParserT__2   = 3
	sdParserT__3   = 4
	sdParserT__4   = 5
	sdParserIF     = 6
	sdParserELSE   = 7
	sdParserTHEN   = 8
	sdParserISZERO = 9
	sdParserSUCC   = 10
	sdParserPRED   = 11
	sdParserID     = 12
	sdParserDIGIT  = 13
	sdParserNAT    = 14
	sdParserBOOL   = 15
	sdParserTRUE   = 16
	sdParserFALSE  = 17
	sdParserWS     = 18
)

// sdParser rules.
const (
	sdParserRULE_expr          = 0
	sdParserRULE_arithmeticOps = 1
	sdParserRULE_ifThen        = 2
)

// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprContext() *ExprContext {
	var p = new(ExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = sdParserRULE_expr
	return p
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = sdParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) CopyFrom(ctx *ExprContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type SymbolContext struct {
	*ExprContext
}

func NewSymbolContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SymbolContext {
	var p = new(SymbolContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *SymbolContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SymbolContext) ID() antlr.TerminalNode {
	return s.GetToken(sdParserID, 0)
}

func (s *SymbolContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case sdVisitor:
		return t.VisitSymbol(s)

	default:
		return t.VisitChildren(s)
	}
}

type NumberContext struct {
	*ExprContext
}

func NewNumberContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NumberContext {
	var p = new(NumberContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *NumberContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumberContext) DIGIT() antlr.TerminalNode {
	return s.GetToken(sdParserDIGIT, 0)
}

func (s *NumberContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case sdVisitor:
		return t.VisitNumber(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExpressionContext struct {
	*ExprContext
}

func NewExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpressionContext {
	var p = new(ExpressionContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case sdVisitor:
		return t.VisitExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type FalseLiteralContext struct {
	*ExprContext
}

func NewFalseLiteralContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FalseLiteralContext {
	var p = new(FalseLiteralContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *FalseLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FalseLiteralContext) FALSE() antlr.TerminalNode {
	return s.GetToken(sdParserFALSE, 0)
}

func (s *FalseLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case sdVisitor:
		return t.VisitFalseLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

type CheckZeroContext struct {
	*ExprContext
}

func NewCheckZeroContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CheckZeroContext {
	var p = new(CheckZeroContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *CheckZeroContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CheckZeroContext) ISZERO() antlr.TerminalNode {
	return s.GetToken(sdParserISZERO, 0)
}

func (s *CheckZeroContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *CheckZeroContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case sdVisitor:
		return t.VisitCheckZero(s)

	default:
		return t.VisitChildren(s)
	}
}

type AbstractionContext struct {
	*ExprContext
}

func NewAbstractionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AbstractionContext {
	var p = new(AbstractionContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *AbstractionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AbstractionContext) ID() antlr.TerminalNode {
	return s.GetToken(sdParserID, 0)
}

func (s *AbstractionContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *AbstractionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case sdVisitor:
		return t.VisitAbstraction(s)

	default:
		return t.VisitChildren(s)
	}
}

type IfElseExprContext struct {
	*ExprContext
}

func NewIfElseExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IfElseExprContext {
	var p = new(IfElseExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *IfElseExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfElseExprContext) IfThen() IIfThenContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIfThenContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIfThenContext)
}

func (s *IfElseExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case sdVisitor:
		return t.VisitIfElseExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type ApplicatioContext struct {
	*ExprContext
}

func NewApplicatioContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ApplicatioContext {
	var p = new(ApplicatioContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *ApplicatioContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ApplicatioContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *ApplicatioContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ApplicatioContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case sdVisitor:
		return t.VisitApplicatio(s)

	default:
		return t.VisitChildren(s)
	}
}

type TrueLiteralContext struct {
	*ExprContext
}

func NewTrueLiteralContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TrueLiteralContext {
	var p = new(TrueLiteralContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *TrueLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TrueLiteralContext) TRUE() antlr.TerminalNode {
	return s.GetToken(sdParserTRUE, 0)
}

func (s *TrueLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case sdVisitor:
		return t.VisitTrueLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

type ArithmeicExpsContext struct {
	*ExprContext
}

func NewArithmeicExpsContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ArithmeicExpsContext {
	var p = new(ArithmeicExpsContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *ArithmeicExpsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArithmeicExpsContext) ArithmeticOps() IArithmeticOpsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArithmeticOpsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArithmeticOpsContext)
}

func (s *ArithmeicExpsContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ArithmeicExpsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case sdVisitor:
		return t.VisitArithmeicExps(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *sdParser) Expr() (localctx IExprContext) {
	return p.expr(0)
}

func (p *sdParser) expr(_p int) (localctx IExprContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 0
	p.EnterRecursionRule(localctx, 0, sdParserRULE_expr, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(25)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case sdParserTRUE:
		localctx = NewTrueLiteralContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(7)
			p.Match(sdParserTRUE)
		}

	case sdParserFALSE:
		localctx = NewFalseLiteralContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(8)
			p.Match(sdParserFALSE)
		}

	case sdParserID:
		localctx = NewSymbolContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(9)
			p.Match(sdParserID)
		}

	case sdParserDIGIT:
		localctx = NewNumberContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(10)
			p.Match(sdParserDIGIT)
		}

	case sdParserT__0:
		localctx = NewAbstractionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(11)
			p.Match(sdParserT__0)
		}
		{
			p.SetState(12)
			p.Match(sdParserID)
		}
		p.SetState(13)
		_la = p.GetTokenStream().LA(1)

		if !(_la == sdParserT__1 || _la == sdParserT__2) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
		{
			p.SetState(14)
			p.expr(6)
		}

	case sdParserT__3:
		localctx = NewExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(15)
			p.Match(sdParserT__3)
		}
		{
			p.SetState(16)
			p.expr(0)
		}
		{
			p.SetState(17)
			p.Match(sdParserT__4)
		}

	case sdParserSUCC, sdParserPRED:
		localctx = NewArithmeicExpsContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(19)
			p.ArithmeticOps()
		}
		{
			p.SetState(20)
			p.expr(3)
		}

	case sdParserISZERO:
		localctx = NewCheckZeroContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(22)
			p.Match(sdParserISZERO)
		}
		{
			p.SetState(23)
			p.expr(2)
		}

	case sdParserIF:
		localctx = NewIfElseExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(24)
			p.IfThen()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(31)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewApplicatioContext(p, NewExprContext(p, _parentctx, _parentState))
			p.PushNewRecursionContext(localctx, _startState, sdParserRULE_expr)
			p.SetState(27)

			if !(p.Precpred(p.GetParserRuleContext(), 5)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
			}
			{
				p.SetState(28)
				p.expr(6)
			}

		}
		p.SetState(33)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext())
	}

	return localctx
}

// IArithmeticOpsContext is an interface to support dynamic dispatch.
type IArithmeticOpsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArithmeticOpsContext differentiates from other interfaces.
	IsArithmeticOpsContext()
}

type ArithmeticOpsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArithmeticOpsContext() *ArithmeticOpsContext {
	var p = new(ArithmeticOpsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = sdParserRULE_arithmeticOps
	return p
}

func (*ArithmeticOpsContext) IsArithmeticOpsContext() {}

func NewArithmeticOpsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArithmeticOpsContext {
	var p = new(ArithmeticOpsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = sdParserRULE_arithmeticOps

	return p
}

func (s *ArithmeticOpsContext) GetParser() antlr.Parser { return s.parser }

func (s *ArithmeticOpsContext) SUCC() antlr.TerminalNode {
	return s.GetToken(sdParserSUCC, 0)
}

func (s *ArithmeticOpsContext) PRED() antlr.TerminalNode {
	return s.GetToken(sdParserPRED, 0)
}

func (s *ArithmeticOpsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArithmeticOpsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArithmeticOpsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case sdVisitor:
		return t.VisitArithmeticOps(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *sdParser) ArithmeticOps() (localctx IArithmeticOpsContext) {
	localctx = NewArithmeticOpsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, sdParserRULE_arithmeticOps)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(34)
	_la = p.GetTokenStream().LA(1)

	if !(_la == sdParserSUCC || _la == sdParserPRED) {
		p.GetErrorHandler().RecoverInline(p)
	} else {
		p.GetErrorHandler().ReportMatch(p)
		p.Consume()
	}

	return localctx
}

// IIfThenContext is an interface to support dynamic dispatch.
type IIfThenContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIfThenContext differentiates from other interfaces.
	IsIfThenContext()
}

type IfThenContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIfThenContext() *IfThenContext {
	var p = new(IfThenContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = sdParserRULE_ifThen
	return p
}

func (*IfThenContext) IsIfThenContext() {}

func NewIfThenContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfThenContext {
	var p = new(IfThenContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = sdParserRULE_ifThen

	return p
}

func (s *IfThenContext) GetParser() antlr.Parser { return s.parser }

func (s *IfThenContext) IF() antlr.TerminalNode {
	return s.GetToken(sdParserIF, 0)
}

func (s *IfThenContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *IfThenContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *IfThenContext) THEN() antlr.TerminalNode {
	return s.GetToken(sdParserTHEN, 0)
}

func (s *IfThenContext) ELSE() antlr.TerminalNode {
	return s.GetToken(sdParserELSE, 0)
}

func (s *IfThenContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfThenContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IfThenContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case sdVisitor:
		return t.VisitIfThen(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *sdParser) IfThen() (localctx IIfThenContext) {
	localctx = NewIfThenContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, sdParserRULE_ifThen)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(36)
		p.Match(sdParserIF)
	}
	{
		p.SetState(37)
		p.expr(0)
	}
	{
		p.SetState(38)
		p.Match(sdParserTHEN)
	}
	{
		p.SetState(39)
		p.expr(0)
	}
	{
		p.SetState(40)
		p.Match(sdParserELSE)
	}
	{
		p.SetState(41)
		p.expr(0)
	}

	return localctx
}

func (p *sdParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 0:
		var t *ExprContext = nil
		if localctx != nil {
			t = localctx.(*ExprContext)
		}
		return p.Expr_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *sdParser) Expr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 5)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 19, 51, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 3, 2, 3, 2, 3, 2, 3, 2, 3,
	2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3,
	2, 3, 2, 3, 2, 3, 2, 5, 2, 31, 10, 2, 3, 2, 3, 2, 7, 2, 35, 10, 2, 12,
	2, 14, 2, 38, 11, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3,
	4, 3, 5, 3, 5, 3, 5, 2, 3, 2, 6, 2, 4, 6, 8, 2, 5, 3, 2, 5, 6, 3, 2, 13,
	14, 3, 2, 17, 18, 2, 53, 2, 30, 3, 2, 2, 2, 4, 39, 3, 2, 2, 2, 6, 41, 3,
	2, 2, 2, 8, 48, 3, 2, 2, 2, 10, 11, 8, 2, 1, 2, 11, 31, 7, 15, 2, 2, 12,
	31, 7, 16, 2, 2, 13, 14, 7, 3, 2, 2, 14, 15, 7, 15, 2, 2, 15, 16, 7, 4,
	2, 2, 16, 17, 5, 8, 5, 2, 17, 18, 9, 2, 2, 2, 18, 19, 5, 2, 2, 8, 19, 31,
	3, 2, 2, 2, 20, 21, 7, 7, 2, 2, 21, 22, 5, 2, 2, 2, 22, 23, 7, 8, 2, 2,
	23, 31, 3, 2, 2, 2, 24, 25, 5, 4, 3, 2, 25, 26, 5, 2, 2, 5, 26, 31, 3,
	2, 2, 2, 27, 28, 7, 12, 2, 2, 28, 31, 5, 2, 2, 4, 29, 31, 5, 6, 4, 2, 30,
	10, 3, 2, 2, 2, 30, 12, 3, 2, 2, 2, 30, 13, 3, 2, 2, 2, 30, 20, 3, 2, 2,
	2, 30, 24, 3, 2, 2, 2, 30, 27, 3, 2, 2, 2, 30, 29, 3, 2, 2, 2, 31, 36,
	3, 2, 2, 2, 32, 33, 12, 7, 2, 2, 33, 35, 5, 2, 2, 8, 34, 32, 3, 2, 2, 2,
	35, 38, 3, 2, 2, 2, 36, 34, 3, 2, 2, 2, 36, 37, 3, 2, 2, 2, 37, 3, 3, 2,
	2, 2, 38, 36, 3, 2, 2, 2, 39, 40, 9, 3, 2, 2, 40, 5, 3, 2, 2, 2, 41, 42,
	7, 9, 2, 2, 42, 43, 5, 2, 2, 2, 43, 44, 7, 11, 2, 2, 44, 45, 5, 2, 2, 2,
	45, 46, 7, 10, 2, 2, 46, 47, 5, 2, 2, 2, 47, 7, 3, 2, 2, 2, 48, 49, 9,
	4, 2, 2, 49, 9, 3, 2, 2, 2, 4, 30, 36,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'\u03BB'", "':'", "'=>'", "'\u2192'", "'('", "')'", "'if'", "'else'",
	"'then'", "'iszero'", "'succ'", "'pred'", "", "", "'Nat'", "'Bool'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "IF", "ELSE", "THEN", "ISZERO", "SUCC", "PRED",
	"ID", "DIGIT", "NAT", "BOOL", "WS",
}

var ruleNames = []string{
	"expr", "arithmeticOps", "ifThen", "funcType",
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
	sdParserT__5   = 6
	sdParserIF     = 7
	sdParserELSE   = 8
	sdParserTHEN   = 9
	sdParserISZERO = 10
	sdParserSUCC   = 11
	sdParserPRED   = 12
	sdParserID     = 13
	sdParserDIGIT  = 14
	sdParserNAT    = 15
	sdParserBOOL   = 16
	sdParserWS     = 17
)

// sdParser rules.
const (
	sdParserRULE_expr          = 0
	sdParserRULE_arithmeticOps = 1
	sdParserRULE_ifThen        = 2
	sdParserRULE_funcType      = 3
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

func (s *SymbolContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(sdListener); ok {
		listenerT.EnterSymbol(s)
	}
}

func (s *SymbolContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(sdListener); ok {
		listenerT.ExitSymbol(s)
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

func (s *NumberContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(sdListener); ok {
		listenerT.EnterNumber(s)
	}
}

func (s *NumberContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(sdListener); ok {
		listenerT.ExitNumber(s)
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

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(sdListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(sdListener); ok {
		listenerT.ExitExpression(s)
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

func (s *CheckZeroContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(sdListener); ok {
		listenerT.EnterCheckZero(s)
	}
}

func (s *CheckZeroContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(sdListener); ok {
		listenerT.ExitCheckZero(s)
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

func (s *AbstractionContext) FuncType() IFuncTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFuncTypeContext)
}

func (s *AbstractionContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *AbstractionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(sdListener); ok {
		listenerT.EnterAbstraction(s)
	}
}

func (s *AbstractionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(sdListener); ok {
		listenerT.ExitAbstraction(s)
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

func (s *IfElseExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(sdListener); ok {
		listenerT.EnterIfElseExpr(s)
	}
}

func (s *IfElseExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(sdListener); ok {
		listenerT.ExitIfElseExpr(s)
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

func (s *ApplicatioContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(sdListener); ok {
		listenerT.EnterApplicatio(s)
	}
}

func (s *ApplicatioContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(sdListener); ok {
		listenerT.ExitApplicatio(s)
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

func (s *ArithmeicExpsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(sdListener); ok {
		listenerT.EnterArithmeicExps(s)
	}
}

func (s *ArithmeicExpsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(sdListener); ok {
		listenerT.ExitArithmeicExps(s)
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
	p.SetState(28)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
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
		{
			p.SetState(13)
			p.Match(sdParserT__1)
		}
		{
			p.SetState(14)
			p.FuncType()
		}
		p.SetState(15)
		_la = p.GetTokenStream().LA(1)

		if !(_la == sdParserT__2 || _la == sdParserT__3) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
		{
			p.SetState(16)
			p.expr(6)
		}

	case sdParserT__4:
		localctx = NewExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(18)
			p.Match(sdParserT__4)
		}
		{
			p.SetState(19)
			p.expr(0)
		}
		{
			p.SetState(20)
			p.Match(sdParserT__5)
		}

	case sdParserSUCC, sdParserPRED:
		localctx = NewArithmeicExpsContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(22)
			p.ArithmeticOps()
		}
		{
			p.SetState(23)
			p.expr(3)
		}

	case sdParserISZERO:
		localctx = NewCheckZeroContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(25)
			p.Match(sdParserISZERO)
		}
		{
			p.SetState(26)
			p.expr(2)
		}

	case sdParserIF:
		localctx = NewIfElseExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(27)
			p.IfThen()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(34)
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
			p.SetState(30)

			if !(p.Precpred(p.GetParserRuleContext(), 5)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
			}
			{
				p.SetState(31)
				p.expr(6)
			}

		}
		p.SetState(36)
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

func (s *ArithmeticOpsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(sdListener); ok {
		listenerT.EnterArithmeticOps(s)
	}
}

func (s *ArithmeticOpsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(sdListener); ok {
		listenerT.ExitArithmeticOps(s)
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
	p.SetState(37)
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

func (s *IfThenContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(sdListener); ok {
		listenerT.EnterIfThen(s)
	}
}

func (s *IfThenContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(sdListener); ok {
		listenerT.ExitIfThen(s)
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
		p.SetState(39)
		p.Match(sdParserIF)
	}
	{
		p.SetState(40)
		p.expr(0)
	}
	{
		p.SetState(41)
		p.Match(sdParserTHEN)
	}
	{
		p.SetState(42)
		p.expr(0)
	}
	{
		p.SetState(43)
		p.Match(sdParserELSE)
	}
	{
		p.SetState(44)
		p.expr(0)
	}

	return localctx
}

// IFuncTypeContext is an interface to support dynamic dispatch.
type IFuncTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFuncTypeContext differentiates from other interfaces.
	IsFuncTypeContext()
}

type FuncTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncTypeContext() *FuncTypeContext {
	var p = new(FuncTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = sdParserRULE_funcType
	return p
}

func (*FuncTypeContext) IsFuncTypeContext() {}

func NewFuncTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncTypeContext {
	var p = new(FuncTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = sdParserRULE_funcType

	return p
}

func (s *FuncTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncTypeContext) NAT() antlr.TerminalNode {
	return s.GetToken(sdParserNAT, 0)
}

func (s *FuncTypeContext) BOOL() antlr.TerminalNode {
	return s.GetToken(sdParserBOOL, 0)
}

func (s *FuncTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(sdListener); ok {
		listenerT.EnterFuncType(s)
	}
}

func (s *FuncTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(sdListener); ok {
		listenerT.ExitFuncType(s)
	}
}

func (p *sdParser) FuncType() (localctx IFuncTypeContext) {
	localctx = NewFuncTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, sdParserRULE_funcType)
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
	p.SetState(46)
	_la = p.GetTokenStream().LA(1)

	if !(_la == sdParserNAT || _la == sdParserBOOL) {
		p.GetErrorHandler().RecoverInline(p)
	} else {
		p.GetErrorHandler().ReportMatch(p)
		p.Consume()
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

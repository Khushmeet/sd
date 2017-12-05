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
	"expr", "arithmeticOps", "ifThen", "r_type",
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
	sdParserRULE_r_type        = 3
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

func (s *ExprContext) ID() antlr.TerminalNode {
	return s.GetToken(sdParserID, 0)
}

func (s *ExprContext) DIGIT() antlr.TerminalNode {
	return s.GetToken(sdParserDIGIT, 0)
}

func (s *ExprContext) R_type() IR_typeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IR_typeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IR_typeContext)
}

func (s *ExprContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *ExprContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ExprContext) ArithmeticOps() IArithmeticOpsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArithmeticOpsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArithmeticOpsContext)
}

func (s *ExprContext) ISZERO() antlr.TerminalNode {
	return s.GetToken(sdParserISZERO, 0)
}

func (s *ExprContext) IfThen() IIfThenContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIfThenContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIfThenContext)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(sdListener); ok {
		listenerT.EnterExpr(s)
	}
}

func (s *ExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(sdListener); ok {
		listenerT.ExitExpr(s)
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
		{
			p.SetState(9)
			p.Match(sdParserID)
		}

	case sdParserDIGIT:
		{
			p.SetState(10)
			p.Match(sdParserDIGIT)
		}

	case sdParserT__0:
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
			p.R_type()
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
		{
			p.SetState(22)
			p.ArithmeticOps()
		}
		{
			p.SetState(23)
			p.expr(3)
		}

	case sdParserISZERO:
		{
			p.SetState(25)
			p.Match(sdParserISZERO)
		}
		{
			p.SetState(26)
			p.expr(2)
		}

	case sdParserIF:
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
			localctx = NewExprContext(p, _parentctx, _parentState)
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

// IR_typeContext is an interface to support dynamic dispatch.
type IR_typeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsR_typeContext differentiates from other interfaces.
	IsR_typeContext()
}

type R_typeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyR_typeContext() *R_typeContext {
	var p = new(R_typeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = sdParserRULE_r_type
	return p
}

func (*R_typeContext) IsR_typeContext() {}

func NewR_typeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *R_typeContext {
	var p = new(R_typeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = sdParserRULE_r_type

	return p
}

func (s *R_typeContext) GetParser() antlr.Parser { return s.parser }

func (s *R_typeContext) NAT() antlr.TerminalNode {
	return s.GetToken(sdParserNAT, 0)
}

func (s *R_typeContext) BOOL() antlr.TerminalNode {
	return s.GetToken(sdParserBOOL, 0)
}

func (s *R_typeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *R_typeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *R_typeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(sdListener); ok {
		listenerT.EnterR_type(s)
	}
}

func (s *R_typeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(sdListener); ok {
		listenerT.ExitR_type(s)
	}
}

func (p *sdParser) R_type() (localctx IR_typeContext) {
	localctx = NewR_typeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, sdParserRULE_r_type)
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

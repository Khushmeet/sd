// Generated from sd.g4 by ANTLR 4.7.

package parser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 19, 106,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 6,
	3, 6, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 10,
	3, 10, 3, 10, 3, 10, 3, 10, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3,
	11, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13,
	3, 14, 6, 14, 82, 10, 14, 13, 14, 14, 14, 83, 3, 15, 6, 15, 87, 10, 15,
	13, 15, 14, 15, 88, 3, 16, 3, 16, 3, 16, 3, 16, 3, 17, 3, 17, 3, 17, 3,
	17, 3, 17, 3, 18, 6, 18, 101, 10, 18, 13, 18, 14, 18, 102, 3, 18, 3, 18,
	2, 2, 19, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19, 11,
	21, 12, 23, 13, 25, 14, 27, 15, 29, 16, 31, 17, 33, 18, 35, 19, 3, 2, 5,
	3, 2, 99, 124, 3, 2, 50, 59, 5, 2, 11, 12, 15, 15, 34, 34, 2, 108, 2, 3,
	3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11,
	3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2,
	19, 3, 2, 2, 2, 2, 21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 2, 25, 3, 2, 2, 2,
	2, 27, 3, 2, 2, 2, 2, 29, 3, 2, 2, 2, 2, 31, 3, 2, 2, 2, 2, 33, 3, 2, 2,
	2, 2, 35, 3, 2, 2, 2, 3, 37, 3, 2, 2, 2, 5, 39, 3, 2, 2, 2, 7, 41, 3, 2,
	2, 2, 9, 44, 3, 2, 2, 2, 11, 46, 3, 2, 2, 2, 13, 48, 3, 2, 2, 2, 15, 50,
	3, 2, 2, 2, 17, 53, 3, 2, 2, 2, 19, 58, 3, 2, 2, 2, 21, 63, 3, 2, 2, 2,
	23, 70, 3, 2, 2, 2, 25, 75, 3, 2, 2, 2, 27, 81, 3, 2, 2, 2, 29, 86, 3,
	2, 2, 2, 31, 90, 3, 2, 2, 2, 33, 94, 3, 2, 2, 2, 35, 100, 3, 2, 2, 2, 37,
	38, 7, 957, 2, 2, 38, 4, 3, 2, 2, 2, 39, 40, 7, 60, 2, 2, 40, 6, 3, 2,
	2, 2, 41, 42, 7, 63, 2, 2, 42, 43, 7, 64, 2, 2, 43, 8, 3, 2, 2, 2, 44,
	45, 7, 8596, 2, 2, 45, 10, 3, 2, 2, 2, 46, 47, 7, 42, 2, 2, 47, 12, 3,
	2, 2, 2, 48, 49, 7, 43, 2, 2, 49, 14, 3, 2, 2, 2, 50, 51, 7, 107, 2, 2,
	51, 52, 7, 104, 2, 2, 52, 16, 3, 2, 2, 2, 53, 54, 7, 103, 2, 2, 54, 55,
	7, 110, 2, 2, 55, 56, 7, 117, 2, 2, 56, 57, 7, 103, 2, 2, 57, 18, 3, 2,
	2, 2, 58, 59, 7, 118, 2, 2, 59, 60, 7, 106, 2, 2, 60, 61, 7, 103, 2, 2,
	61, 62, 7, 112, 2, 2, 62, 20, 3, 2, 2, 2, 63, 64, 7, 107, 2, 2, 64, 65,
	7, 117, 2, 2, 65, 66, 7, 124, 2, 2, 66, 67, 7, 103, 2, 2, 67, 68, 7, 116,
	2, 2, 68, 69, 7, 113, 2, 2, 69, 22, 3, 2, 2, 2, 70, 71, 7, 117, 2, 2, 71,
	72, 7, 119, 2, 2, 72, 73, 7, 101, 2, 2, 73, 74, 7, 101, 2, 2, 74, 24, 3,
	2, 2, 2, 75, 76, 7, 114, 2, 2, 76, 77, 7, 116, 2, 2, 77, 78, 7, 103, 2,
	2, 78, 79, 7, 102, 2, 2, 79, 26, 3, 2, 2, 2, 80, 82, 9, 2, 2, 2, 81, 80,
	3, 2, 2, 2, 82, 83, 3, 2, 2, 2, 83, 81, 3, 2, 2, 2, 83, 84, 3, 2, 2, 2,
	84, 28, 3, 2, 2, 2, 85, 87, 9, 3, 2, 2, 86, 85, 3, 2, 2, 2, 87, 88, 3,
	2, 2, 2, 88, 86, 3, 2, 2, 2, 88, 89, 3, 2, 2, 2, 89, 30, 3, 2, 2, 2, 90,
	91, 7, 80, 2, 2, 91, 92, 7, 99, 2, 2, 92, 93, 7, 118, 2, 2, 93, 32, 3,
	2, 2, 2, 94, 95, 7, 68, 2, 2, 95, 96, 7, 113, 2, 2, 96, 97, 7, 113, 2,
	2, 97, 98, 7, 110, 2, 2, 98, 34, 3, 2, 2, 2, 99, 101, 9, 4, 2, 2, 100,
	99, 3, 2, 2, 2, 101, 102, 3, 2, 2, 2, 102, 100, 3, 2, 2, 2, 102, 103, 3,
	2, 2, 2, 103, 104, 3, 2, 2, 2, 104, 105, 8, 18, 2, 2, 105, 36, 3, 2, 2,
	2, 6, 2, 83, 88, 102, 3, 8, 2, 2,
}

var lexerDeserializer = antlr.NewATNDeserializer(nil)
var lexerAtn = lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'\u03BB'", "':'", "'=>'", "'\u2192'", "'('", "')'", "'if'", "'else'",
	"'then'", "'iszero'", "'succ'", "'pred'", "", "", "'Nat'", "'Bool'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "", "", "IF", "ELSE", "THEN", "ISZERO", "SUCC", "PRED",
	"ID", "DIGIT", "NAT", "BOOL", "WS",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "IF", "ELSE", "THEN", "ISZERO",
	"SUCC", "PRED", "ID", "DIGIT", "NAT", "BOOL", "WS",
}

type sdLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var lexerDecisionToDFA = make([]*antlr.DFA, len(lexerAtn.DecisionToState))

func init() {
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

func NewsdLexer(input antlr.CharStream) *sdLexer {

	l := new(sdLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "sd.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// sdLexer tokens.
const (
	sdLexerT__0   = 1
	sdLexerT__1   = 2
	sdLexerT__2   = 3
	sdLexerT__3   = 4
	sdLexerT__4   = 5
	sdLexerT__5   = 6
	sdLexerIF     = 7
	sdLexerELSE   = 8
	sdLexerTHEN   = 9
	sdLexerISZERO = 10
	sdLexerSUCC   = 11
	sdLexerPRED   = 12
	sdLexerID     = 13
	sdLexerDIGIT  = 14
	sdLexerNAT    = 15
	sdLexerBOOL   = 16
	sdLexerWS     = 17
)

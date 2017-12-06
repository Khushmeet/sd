package main

import (
	"fmt"
	SdParser "sd/parser"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type SdTranspilerListener struct {
	*SdParser.BasesdListener
}

/*
NewSdTranspilerListener creates new sdTranspilerListener object
*/
func NewSdTranspilerListener() *SdTranspilerListener {
	return new(SdTranspilerListener)
}

// EnterEveryRule gets called when tree walker enter a node in parse tree
func (sd *SdTranspilerListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	fmt.Println(ctx.GetText())
}

package main

import (
	"fmt"
	"os"
	SdParser "sd/parser"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type sdTranspilerListener struct {
	*SdParser.BasesdListener
}

func newSdTranspilerListener() *sdTranspilerListener {
	return new(sdTranspilerListener)
}

func (sd *sdTranspilerListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	fmt.Println(ctx.GetText())
}

func main() {
	input, _ := antlr.NewFileStream(os.Args[1])
	lexer := SdParser.NewsdLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	parser := SdParser.NewsdParser(stream)
	parser.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	parser.BuildParseTrees = true
	tree := parser.Expr()
	antlr.ParseTreeWalkerDefault.Walk(newSdTranspilerListener(), tree)
}

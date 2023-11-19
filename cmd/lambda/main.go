package main

import (
	"github.com/n0c1337/lambda/internal/filesystem"
	"github.com/n0c1337/lambda/internal/lexer"
)

func main() {
	line2 := filesystem.ReadLine("./tests/lambda.S", 2)
	analysis := lexer.NewLexer()
	analysis.PerformLexicalAnalysis(line2)
}

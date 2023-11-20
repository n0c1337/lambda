package lexer

import (
	"errors"
	"fmt"
)

type Lexer struct{}

func NewLexer() Lexer {
	return Lexer{}
}

// 🧐
func (Lexer) PerformLexicalAnalysis(content string) (string, error) {
	StateMachineExpression()
	fmt.Println(content)
	switch content {
	default:
		return "Invalid Token Type", errors.New("invalid Token Type")
	}
}

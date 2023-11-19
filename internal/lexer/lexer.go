package lexer

// #cgo CFLAGS: -g -Wall
// #include "state_machine.h"
import "C"
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
	fmt.Println(content)
	switch content {
	default:
		return "Invalid Token Type", errors.New("invalid Token Type")
	}
}

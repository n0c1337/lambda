package compiler

import (
	"fmt"
)

type IfElseStatement struct {
	Condition string
}

type IfStatement struct {
	Condition string
}

type ElseIf struct {
	Condition string
}

func NewIfElseStatement(condition string) IfElseStatement {
	return IfElseStatement{Condition: condition}
}

func (ie IfElseStatement) Compile(comperand_1 string, comperand_2 string) string {
	var statement string

	trueLabel := NewLabel("conditionMatched")
	endLabel := NewLabel("end")

	statement += fmt.Sprintf("cmp %s, %s\n", comperand_1, comperand_2)
	switch ie.Condition {
	case "<":
		statement += fmt.Sprintf("jl %s\n", trueLabel.Name)
	case "<=":
		statement += fmt.Sprintf("jle %s\n", trueLabel.Name)
	case "==":
		statement += fmt.Sprintf("je %s\n", trueLabel.Name)
	case "!=":
		statement += fmt.Sprintf("jne %s\n", trueLabel.Name)
	case ">":
		statement += fmt.Sprintf("jg %s\n", trueLabel.Name)
	case ">=":
		statement += fmt.Sprintf("jge %s\n", trueLabel.Name)
	}

	statement += fmt.Sprintf("jmp %s\n", endLabel.Name)

	statement += fmt.Sprintf("%s\n", trueLabel.Compile())

	statement += fmt.Sprintf("%s\n", endLabel.Compile())
	statement += "ret"

	return statement
}

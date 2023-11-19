package compiler

import "fmt"

type Stack struct{}

func (Function) GetStack() Stack {
	return Stack{}
}

func (Stack) NewProlog() string {
	return "push rbp\nmov rbp, rsp"
}

func (Stack) NewEpilog() string {
	/*
		leave instruction will perform:
		mov rsp, rbp
		pop rbp
	*/
	return "leave"
}

func (Stack) Reserve(size int64) string {
	return fmt.Sprintf("sub rsp, %d", size)
}

func (Stack) Drop(size int64) string {
	return fmt.Sprintf("add rsp, %d", size)

}

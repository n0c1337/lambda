package lexer

/*
Go interface for "state_machine.c" the best non-deterministic state machine
*/

// #cgo CFLAGS: -g -Wall
// #include "state_machine.h"
import "C"

func NewStateMachine() C.StateMachine {
	var stateMachine C.StateMachine
	C.NewStateMachine(&stateMachine)
	return stateMachine
}

func AddState(stateMachine C.StateMachine, state int, accept int) {
	C.AddState(&stateMachine, C.int(state), C.int(accept))
}

func AddTransition(stateMachine C.StateMachine, start int, input int, end int) {
	C.AddTransition(&stateMachine, C.int(start), C.int(input), C.int(end))
}

func MatchTransition(stateMachine C.StateMachine, input []C.int) bool {
	result := C.MatchTransition(&stateMachine, &input[0])
	return bool(result)
}

package lexer

import "C"
import "fmt"

const (
	GO_TO_STATE_1 int = 1
	GO_TO_STATE_2 int = 2
	GO_TO_STATE_3 int = 3
	GO_TO_STATE_4 int = 4
)

func StateMachineExpression() {
	stateMachine := NewStateMachine()

	AddState(stateMachine, 0, 0)
	AddState(stateMachine, 1, 0)
	AddState(stateMachine, 2, 0)
	AddState(stateMachine, 3, 0)
	AddState(stateMachine, 4, 0) // End state

	AddTransition(stateMachine, 0, GO_TO_STATE_1, 1)
	AddTransition(stateMachine, 1, GO_TO_STATE_2, 2)
	AddTransition(stateMachine, 2, GO_TO_STATE_3, 3)
	AddTransition(stateMachine, 3, GO_TO_STATE_4, 4)

	input := []C.int{1, 2, 3, 4, -1}
	if MatchTransition(stateMachine, input) {
		fmt.Println("Match")
	} else {
		fmt.Println("No match")
	}
}

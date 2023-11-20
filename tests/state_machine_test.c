#include "../internal/lexer/state_machine.h"

// https://en.wikipedia.org/wiki/Nondeterministic_finite_automaton

#define GO_TO_STATE_1 1
#define GO_TO_STATE_2 2
#define GO_TO_STATE_3 3
#define GO_TO_STATE_4 4
#define END -1


int main()
{
    StateMachine stateMachine;

    NewStateMachine(&stateMachine);

    AddState(&stateMachine, 0, 0);
    AddState(&stateMachine, 1, 0);
    AddState(&stateMachine, 2, 0);
    AddState(&stateMachine, 3, 0);
    AddState(&stateMachine, 4, 1); // End state

    AddTransition(&stateMachine, 0, GO_TO_STATE_1, 1);
    AddTransition(&stateMachine, 1, GO_TO_STATE_2, 2);
    AddTransition(&stateMachine, 2, GO_TO_STATE_3, 3);
    AddTransition(&stateMachine, 3, GO_TO_STATE_4, 4);

    int input[] = {GO_TO_STATE_1, GO_TO_STATE_2, GO_TO_STATE_3, GO_TO_STATE_4, END};
    if (MatchTransition(&stateMachine, input))
    {
        printf("Match");
    }
    else
    {
        printf("No Match");
    }
    
    return 0;
}
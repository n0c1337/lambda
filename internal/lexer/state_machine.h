#ifndef _STATE_MACHINE_H
#define _STATE_MACHINE_H

#include <stdio.h>
#include <stdbool.h>

#define MAX_STATES 50
#define NO_TRANSITION -1
#define TRANSITION 1

typedef struct {
	int states[MAX_STATES];
	int accept;
	int transition[MAX_STATES][64][MAX_STATES];
} StateMachine;

void NewStateMachine(StateMachine *stateMachine);
void AddState(StateMachine *stateMachine, int state, bool accept);
void AddTransition(StateMachine *stateMachine, int start, int input, int end);
bool MatchTransition(StateMachine *stateMachine, int input[]);

#endif
#include "state_machine.h"
#include <stdio.h>

/*
Non-deterministic state machine
*/
void NewStateMachine(StateMachine *stateMachine)
{
	for (int i = 0; i < MAX_STATES; i++)
	{
		for (int j = 0; j < 64; j++)
		{
			for (int k = 0; k < MAX_STATES; k++)
			{
				stateMachine->transition[i][j][k] = NO_TRANSITION;
			}
		}
	}
}

void AddState(StateMachine *stateMachine, int state, int accept)
{
	stateMachine->states[state] = accept;
}

void AddTransition(StateMachine *stateMachine, int start, int input, int end)
{
	stateMachine->transition[start][input][end] = TRANSITION;
}

bool MatchTransition(StateMachine *stateMachine, int input[])
{
	printf("debug");
	int currentStates[MAX_STATES] = { 0 };
	currentStates[0] = 1;

	for (int i = 0; input[i] != -1; i++)
	{
		int nextStates[MAX_STATES] = { 0 };
		for (int j = 0; j < MAX_STATES; j++)
		{
			if (currentStates[j])
			{
				for (int k = 0; k < MAX_STATES; k++)
				{
					if (stateMachine->transition[j][input[i]][k] == TRANSITION) 
					{
						nextStates[k] = TRANSITION;
					}
				}
			}
		}

		for (int j = 0; j < MAX_STATES; j++)
		{
			currentStates[j] = nextStates[j];
		}
	}
	

	for (int i = 0; i < MAX_STATES; i++)
	{
		if (currentStates[i] && stateMachine->states[i]) 
		{
			return true;
		}
	}


	return false;
}
#ifndef STACK_H_INCLUDED
#define STACK_H_INCLUDED
#include <iostream>

#define info(S) S.info
#define top(S) S.top

typedef char infotype;
struct stack{
    infotype info[15];
    int top;
};

using namespace std;

void createStack_103022300064(stack &S);
bool isEmpty_103022300064(stack S);
bool isFull_103022300064(stack S) ;
void push_103022300064(stack &S, infotype x);
char pop_103022300064(stack &S);
void printInfo_103022300064(stack S);



#endif // STACK_H_INCLUDED

#ifndef stack_H
#define stack_H
#include <iostream>

#define info(P) P->info
#define next(P) P->next
#define top(S) S.top


typedef int infotype;
typedef struct elmStack *address;

struct elmStack {
    infotype info;
    address next;
};

struct stack{
    address top;
};

address createElmStack(int x);
void isEmpty(stack S);
void createStack(stack &S);
void allocation(stack &S, address x);
void deallocation(stack &S, address &x);
void printStack(stack S);

#endif


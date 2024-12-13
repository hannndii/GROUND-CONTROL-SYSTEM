#ifndef STACK_H_INCLUDED
#define STACK_H_INCLUDED
#include <iostream>
using namespace std;

struct infotype{
    string nama;
    int nilai;
};

struct stack{
    infotype info[3];
    int top;
};

void createStack(stack &S);
bool isEmpty(stack S);
bool isFull(stack S);
void pushStack(stack &S, infotype data);
infotype popStack(stack &S);
void printInfo(stack S);


#endif // STACK_H_INCLUDED

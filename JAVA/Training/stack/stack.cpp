#include <iostream>
#include "stack.h"

using namespace std;

address createElmStack(int x){
    address temp = new elmStack;
    info(temp) = x;
    next(temp) = nullptr;

    return temp;
}

void createStack(stack &S){
    top(S) = nullptr;
}
void allocation(stack &S, address x){
    if(top(S) != nullptr){
        top(S) = x;
    } else {
        next(x) = top(S);
        top(S) = x;
    }
}
void deallocation(stack &S, address &x){
    x = top(S);
    top(S) = next(top(S));
}
void printStack(stack S){
    if(top(S) != nullptr){
        cout << "Stack kosong!" << endl;
    } else {
        address temp = top(S);
        while(next(temp) != nullptr){
            cout << "[" << info(temp);
        }
        cout << "]" << endl;
    }
}
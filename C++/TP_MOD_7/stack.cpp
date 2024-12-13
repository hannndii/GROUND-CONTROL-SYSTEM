#include "stack.h"
#include <iostream>

using namespace std;

void createStack_103022300064(stack &S){
    top(S) = 0;
}
bool isEmpty_103022300064(stack S){
    if(top(S) == 0){
        return true;
    } else {
        return false;
    }
}
bool isFull_103022300064(stack S){
    if(top(S) == 15){
        return true;
    } else {
        return false;
    }
}
void push_103022300064(stack &S, infotype x){
    if(isFull_103022300064(S) == false){
        top(S) = top(S) + 1;
        info(S)[top(S)] = x;
    }
}
char pop_103022300064(stack &S){
    infotype x;
    x = info(S)[top(S)];
    top(S) = top(S) - 1;

    return top(S);

}
void printInfo_103022300064(stack S){
    for (int i = 0; i < top(S); i++){
        cout << info(S)[i] << " ";
    }
    cout << endl;
}

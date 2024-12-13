#include <iostream>
#include "stack.h"

using namespace std;

int main(){
    stack S;
    createStack(S);
    for(int i = 0; i < 5; i++){
        int input = 0;
        address temp = createElmStack(input);
        allocation(S, temp);
        input*2;
    }
    printStack(S);
    
}
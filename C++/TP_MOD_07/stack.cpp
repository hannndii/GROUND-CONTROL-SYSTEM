#include <iostream>
#include "stack.h"

using namespace std;

void createStack(stack &S){
    S.top = 0;
}
bool isEmpty(stack S){
    if(S.top == 0){
        return true;
    } else {
        return false;
    }
}
bool isFull(stack S){
    if(S.top == 3){
        return true;
    } else {
        return false;
    }
}
void pushStack(stack &S, infotype data) {
    if (!isFull(S)) {
        S.info[S.top] = data;
        S.top++;
    }
}
infotype popStack(stack &S){
    infotype temp;

    S.top--;
    temp = S.info[S.top];
    return temp;
}
void printInfo(stack S){
        for (int i = 0; i < S.top; i++){
            cout << "Nama: " << S.info[i].nama << endl;
            cout << "Nilai: " << S.info[i].nilai << endl;
        }
    cout << endl;
}

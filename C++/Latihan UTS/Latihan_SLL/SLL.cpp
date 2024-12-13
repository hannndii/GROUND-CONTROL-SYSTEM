#include "SLL.h"
#include <iostream>

using namespace std;

void createElm(List &L){
    first(L) = nullptr;
}
address createNewElm(infotype x){
    address temp = new elmlist;
    info(temp) = x;
    next(temp) = nullptr;

    return temp;
}
bool isEmpty(List L){
    if(first(L) == nullptr){
        return true;
    } else {
        return false;
    }
}
bool isFull(List L){
     if(first(L) != nullptr){
        return true;
    } else {
        return false;
    }
}
void insertFirst(List &L, address input){
    if(!isFull(L)){
        first(L) = input;
    } else {
        next(input) = first(L);
        first(L) = input;
    }
}
void deleteFirst(List &L, address &temp){
    temp = first(L);
    first(L) = next(first(L));
    next(temp) = nullptr;
}
void insertAfter(List &L, address &prec, address input){
    if(!isFull(L)){
        first(L) = input;
    } else {
        next(input) = next(prec);
        next(prec) = input;
    }
}
void deleteAfter(List &L, address prec, address &temp){
    temp = next(prec);
    next(prec) = next(temp);
    next(temp) = nullptr;
}
void insertLast(List &L, address input){
    if(!isFull(L)){
        first(L) = input;
    } else {
        address p = first(L);
        while(next(p) != nullptr){
            p = next(p);
        }
        next(p) = input;
    }
}
void deleteLast(List &L, address &temp){
    address p = first(L);
    while(next(p) != nullptr){
        p = next(p);
    }
    temp = p;
    p = nullptr;
    next(temp) = nullptr;
}
void printData(List L){
    if(!isFull(L)){
        cout << "List Kosong" << endl;
    } else {
        address p = first(L);
        while(p != nullptr){
            cout << info(p) << " ";
            p = next(p);
        }
        cout << endl;
    }
}

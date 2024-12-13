#ifndef SLL_H_INCLUDED
#define SLL_H_INCLUDED
#include <iostream>
#define first(L) L.first
#define next(P) P->next
#define info(P) P->info

typedef char infotype;
typedef struct elmlist *address;

struct List{
    address first;
};

struct elmlist{
    infotype info;
    address next;
};

void createElm(List &L);
address createNewElm(infotype x);
bool isEmpty(List L);
bool isFull(List L);
void insertFirst(List &L, address input);
void deleteFirst(List &L, address &temp);
void insertAfter(List &L, address &prec, address input);
void deleteAfter(List &L, address prec, address &temp);
void insertLast(List &L, address input);
void deleteLast(List &L, address &temp);
void printData(List L);

#endif // SLL_H_INCLUDED

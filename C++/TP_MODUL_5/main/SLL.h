#ifndef SLL_H_INCLUDED
#define SLL_H_INCLUDED
#include <iostream>
#define first(L) L.first
#define info(P) P->info
#define next(P) P->next

using namespace std;

typedef int infotype;
typedef struct elmlist *address;

struct List{
    address first;
};

struct elmlist{
    address next;
    infotype info;
};

void createList_103022300064(List &L);
address createNewElement_103022300064(infotype x);
void insertLast_103022300064(List &L, address p);
void showAllData_103022300064(List L);
address findMin(List L);
void insertMiddle_103022300064(List &L, int value);
int selectMenu_103022300064();
#endif // SLL_H_INCLUDED

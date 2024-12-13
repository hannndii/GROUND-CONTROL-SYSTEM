#ifndef DLL_H_INCLUDED
#define DLL_H_INCLUDED
#include <iostream>
#define first(L) L.first
#define last(L) L.last
#define next(P) P->next
#define prev(P) P->prev
#define info(P) P->info

using namespace std;

typedef string infotype;
typedef struct elmlist *address;

struct elmlist{
    infotype info;
    address next;
    address prev;
};

struct List {
    address first;
    address last;
};


bool isEmpty_103022300064(List L);
void createList_103022300064(List L);
address createNewList_103022300064(infotype x);
void insertFirst_103022300064(List &L, address p);
void inserAfter_103022300064(List &L, address &p, address &prec);
void inserLast_103022300064(List &L, address p);
void deleteFirst_103022300064(List &L, address p);
void deleteAfter_103022300064(List &L, address &prec, address p);
void deleteLast_103022300064(List &L, address p);
void concat_103022300064(List L1, List L2, List L3);
address findLagu_103022300064(List L, string judul);
void removeLagu_103022300064(string judul, List L);
void showAllData_103022300064(List L);
int selectMenu_103022300064();

#endif // DLL_H_INCLUDED

#ifndef DLL_H
#define DLL_H

#include <iostream>
using namespace std;

struct elmtList {
    string bandName;
    string songTitle;
    elmtList* prev;
    elmtList* next;
};

typedef elmtList* address;

struct List {
    address first;
    address last;
};

bool isEmpty(List L);

void createList(List &L);
void createNewElmt(string bandName, string songTitle, address &P);
void insertFirst(List &L, address P);
void insertLast(List &L, address P);
void insertAfter(List &L, address Prec, address P);
void deleteFirst(List &L, address &P);
void deleteLast(List &L, address &P);
void concat(List L1, List L2, List &L3);
address findLagu(string judul, List L);
void removeLagu(string judul, List &L);
int selectMenu_103022300064();
void showList(List L);
#endif

#ifndef SLL_H_INCLUDED
#define SLL_H_INCLUDED
#include "mahasiswa.h"
#define Nil NULL
#define info(P) (P)->info
#define next(P) (P)->next
#define first(L) L.first

typedef mahasiswa infotype;

typedef struct elmlist *address;

struct List{
    address first;
};

struct elmlist{
    infotype info;
    address next;
};


void createList(List &L);
address allocateElement(infotype baru);


#endif // SLL_H_INCLUDED

#include "SLL.h"

void createList(List &L){
    first(L) = NULL;
}
address allocateElement(infotype baru){
    address p = new elmlist;
    info(p) = baru;
    next(p) = NULL;
    return p;
}

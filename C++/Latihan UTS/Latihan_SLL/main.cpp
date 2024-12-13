#include <iostream>
#include "SLL.h"

using namespace std;

int main()
{
    List L;
    createElm(L);

    // FIRST
    char input;
    for(int i = 0; i < 5; i++){
        cin >> input;
        insertFirst(L, createNewElm(input));
    }

    printData(L);
    address temp1 = nullptr;
    deleteFirst(L, temp1);
    printData(L);
    cout << "= = = = = = = = = = " << endl;

    // AFTER
    char input1;
    address precT = nullptr;
    address l = first(L);

    for(int i = 0; i < 3; i++){
        precT = l;
        l = next(l);
    }

    cin >> input1;
    address p = createNewElm(input1);
    insertAfter(L, precT, p);
    printData(L);

    address temp2 = nullptr;
    deleteAfter(L, precT, temp2);
    printData(L);
    cout << "= = = = = = = = = = " << endl;

    // LAST
    char input2;
    cin >> input2;
    insertLast(L, createNewElm(input2));
    printData(L);

    address temp3 = nullptr;
    deleteLast(L, temp3);
    printData(L);

    return 0;
}

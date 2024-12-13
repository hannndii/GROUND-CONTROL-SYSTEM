#include <iostream>
#include "stack.h"

using namespace std;

int main()
{
    stack S;
    infotype data;
    infotype temp1;

    createStack(S);

    cout << "- - - - - MASUKKAN DATA - - - - -" << endl;
    int i = 0;
    for (i = 0; i < 3; i++){
        if(!isFull(S)){
            cout << "Input Nama: ";
            cin >> data.nama;
            cout << endl;
            cout << "Input Nilai: ";
            cin >> data.nilai;
            cout << endl;
            pushStack(S, data);
        } else {
            cout << "Data Penuh" << endl;
            break;
        }
    }
    printInfo(S);

    temp1 = popStack(S);
    cout << "DATA YANG DIHAPUS: " << endl;
    cout << temp1.nama << endl;
    cout << temp1.nilai << endl;
    cout << "DATA BARU: ";
    printInfo(S);

    return 0;
}

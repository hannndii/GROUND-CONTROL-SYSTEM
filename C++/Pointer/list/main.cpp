#include <iostream>
#include "list.h"

using namespace std;

int main() {
    // 1. Panggil create list
    List L;
    createList(L);

    // 2. Loop untuk memasukkan NIM per digit
    int nim_part;
    for (int i = 1; i <= 10; i++) {
        cout << "Masukkan NIM perdigit\n";
        cout << "Digit " << i << ": ";
        cin >> nim_part;

        // 3. Panggil fungsi allocate dan insertLast agar data dimasukkan sesuai urutan
        address P = allocate(nim_part);
        insertLast(L, P);
    }

    // 4. Tampilkan isi list
    cout << "Isi list: ";
    printInfo(L);

    return 0;
}

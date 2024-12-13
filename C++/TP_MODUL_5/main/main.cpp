#include <iostream>
#include "SLL.h"

using namespace std;

int main()
{
    List L;
    int pilihan = 0;

    pilihan = selectMenu_103022300064();
    while(pilihan != 0){
        switch(pilihan){
        case 1:
            createList_103022300064(L);
            int N;
            cout << "- - - - MASUKKAN DATA BARU - - - -" << endl;
            for(int i = 0; i < 10; i++){
                cout << "Angka ke-" << i+1 << ": ";
                cin >> N;
                address newN = createNewElement_103022300064(N);
                insertLast_103022300064(L, newN);
            }
            break;
        case 2:
            showAllData_103022300064(L);
            break;
        case 3:
            cout << "Lokasi alamat: " << findMin(L) << endl;
            break;
        case 4:
            int value;
            cout << "Masukkan data baru: ";
            cin >> value;
            insertMiddle_103022300064(L, value);
            showAllData_103022300064(L);
        }
        pilihan = selectMenu_103022300064();
    }
    cout << "ANDA TELAH KELUAR DARI PROGRAM" << endl;
    return 0;
}

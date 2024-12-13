#include <iostream>
#include "DLL.h"

using namespace std;

int main() {
    List L;
    createList(L);
    int menuChoice;

    do {
        menuChoice = selectMenu_103022300064();
        switch (menuChoice) {
            case 1: {
                // Input data baru (di akhir)
                string bandName, songTitle;
                cout << "Masukkan Nama Band: ";
                cin >> bandName;
                cout << "Masukkan Judul Lagu: ";
                cin >> songTitle;
                address P;
                createNewElmt(bandName, songTitle, P);
                insertLast(L, P);
                cout << "Data berhasil ditambahkan." << endl;
                break;
            }
            case 2: {
                // Tampilkan data
                cout << "Daftar Lagu: " << endl;
                showList(L);
                break;
            }
            case 3: {
                // Hapus data lagu
                string songTitle;
                cout << "Masukkan Judul Lagu yang akan dihapus: ";
                cin >> songTitle;
                removeLagu(songTitle, L);
                cout << "Data berhasil dihapus." << endl;
                break;
            }
            case 4: {
                // Input data di posisi tengah (Setelah elemen pertama)
                string bandName, songTitle;
                cout << "Masukkan Nama Band: ";
                cin >> bandName;
                cout << "Masukkan Judul Lagu: ";
                cin >> songTitle;
                address P;
                createNewElmt(bandName, songTitle, P);
                if (!isEmpty(L)) {
                    insertAfter(L, L.first, P);
                } else {
                    insertFirst(L, P);
                }
                cout << "Data berhasil ditambahkan di tengah." << endl;
                break;
            }
            case 5: {
                // Cari lagu
                string songTitle;
                cout << "Masukkan Judul Lagu yang dicari: ";
                cin >> songTitle;
                address found = findLagu(songTitle, L);
                if (found != NULL) {
                    cout << "Lagu ditemukan: " << found->bandName << " - " << found->songTitle << endl;
                } else {
                    cout << "Lagu tidak ditemukan." << endl;
                }
                break;
            }
            case 0: {
                // Keluar
                cout << "Keluar dari program." << endl;
                break;
            }
            default: {
                cout << "Pilihan tidak valid." << endl;
                break;
            }
        }
    } while (menuChoice != 0);

    return 0;
}

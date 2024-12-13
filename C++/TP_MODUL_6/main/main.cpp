#include <iostream>
#include "DLL.h"

using namespace std;

// ADA MASALAH PADA CASE PERTAMA KETIKA INPUT KE-2

int main() {
    List L;
    isEmpty_103022300064(L);
    int menu;
    string judul;
    address p;

    menu = selectMenu_103022300064();
    while (menu != 0) {
        switch (menu) {
            case 1: {
                for (int i = 0; i < 10; i++) {
                    cout << "Masukkan judul lagu baru: ";
                    cin >> judul;
                    p = createNewList_103022300064(judul);
                    insertFirst_103022300064(L, p);
                    cout << "Lagu berhasil ditambahkan." << endl;
                }
                system("pause"); // Pause sebelum kembali ke menu
                break;
            }

            case 2: {
                showAllData_103022300064(L);
                system("pause"); // Pause sebelum kembali ke menu
                break;
            }

            case 3: {
                cout << "Masukkan judul lagu yang ingin dihapus: ";
                cin >> judul;
                // removeLagu_103022300064(judul, L);
                cout << "Lagu berhasil dihapus (implementasi removeLagu belum ada)." << endl;
                system("pause"); // Pause sebelum kembali ke menu
                break;
            }

            case 4: {
                cout << "Masukkan judul lagu baru (tengah): ";
                cin >> judul;
                p = createNewList_103022300064(judul);
                address middle = first(L);
                int count = 0, i = 0;

                address temp = first(L);
                while (temp != NULL) {
                    count++;
                    temp = next(temp);
                }

                while (i < count / 2) {
                    middle = next(middle);
                    i++;
                }

                inserAfter_103022300064(L, p, middle);
                cout << "Lagu berhasil ditambahkan di posisi tengah." << endl;
                system("pause"); // Pause sebelum kembali ke menu
                break;
            }

            case 5: {
                cout << "Masukkan judul lagu yang ingin dicari: ";
                cin >> judul;
                address found = findLagu_103022300064(L, judul);
                if (found != NULL) {
                    cout << "Lagu dengan judul '" << judul << "' ditemukan di dalam list." << endl;
                } else {
                    cout << "Lagu dengan judul '" << judul << "' tidak ditemukan." << endl;
                }
                system("pause"); // Pause sebelum kembali ke menu
                break;
            }

            default:
                cout << "Pilihan tidak valid!" << endl;
                system("pause"); // Pause sebelum kembali ke menu
                break;

        }
        menu = selectMenu_103022300064(); // Menampilkan menu lagi setelah case selesai
    }

    cout << "Program keluar." << endl;
    return 0;
}

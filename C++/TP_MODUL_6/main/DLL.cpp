#include <iostream>
#include "DLL.h"

using namespace std;

bool isEmpty_103022300064(List L) {
    return (first(L) == NULL && last(L) == NULL);
}

address createNewList_103022300064(infotype x) {
    address temp = new elmlist;
    info(temp) = x;
    next(temp) = NULL;
    prev(temp) = NULL;
    return temp;
}

void insertFirst_103022300064(List &L, address p) {
    if (first(L) == NULL && last(L) == NULL) {
        first(L) = p;
        last(L) = p;
    } else {
        next(p) = first(L);
        prev(first(L)) = p;
        first(L) = p;
    }
}

void inserAfter_103022300064(List &L, address &p, address &prec) {
    if (first(L) == NULL && last(L) == NULL) {
        first(L) = p;
        last(L) = p;
    } else {
        next(p) = next(prec);
        if (next(prec) != NULL) {  // Jika bukan elemen terakhir
            prev(next(prec)) = p;
        } else {  // Jika elemen terakhir
            last(L) = p;
        }
        prev(p) = prec;
        next(prec) = p;
    }
}

void insertLast_103022300064(List &L, address p) {
    if (first(L) == NULL && last(L) == NULL) {
        first(L) = p;
        last(L) = p;
    } else {
        prev(p) = last(L);
        next(last(L)) = p;
        last(L) = p;
    }
}

void deleteFirst_103022300064(List &L, address p) {
    p = first(L);
    if (first(L) != last(L)) {
        first(L) = next(p);
        next(p) = NULL;
        prev(first(L)) = NULL;
    } else {
        first(L) = NULL;
        last(L) = NULL;
    }
}

void deleteAfter_103022300064(List &L, address &prec, address p) {
    p = next(prec);
    if (p != last(L)) {
        next(prec) = next(p);
        prev(next(p)) = prec;
    } else {
        next(prec) = NULL;
        last(L) = prec;
    }
    next(p) = NULL;
    prev(p) = NULL;
}

void deleteLast_103022300064(List &L, address p) {
    p = last(L);
    if (first(L) != last(L)) {
        last(L) = prev(p);
        next(last(L)) = NULL;
        prev(p) = NULL;
    } else {
        first(L) = NULL;
        last(L) = NULL;
    }
}

void concat_103022300064(List L1, List L2, List &L3) {
    if (first(L1) == NULL) {
        first(L3) = first(L2);
        last(L3) = last(L2);
    } else if (first(L2) == NULL) {
        first(L3) = first(L1);
        last(L3) = last(L1);
    } else {
        first(L3) = first(L1);
        last(L3) = last(L2);
        next(last(L1)) = first(L2);
        prev(first(L2)) = last(L1);
    }
}

address findLagu_103022300064(List L, string judul) {
    address temp = first(L);
    while (temp != NULL) {
        if (info(temp) == judul) {
            return temp;
        }
        temp = next(temp);
    }
    return NULL;
}

void removeLagu_103022300064(string judul, List &L) {
    address p = findLagu_103022300064(L, judul);
    if (p != NULL) {
        if (p == first(L)) {
            deleteFirst_103022300064(L, p);
        } else if (p == last(L)) {
            deleteLast_103022300064(L, p);
        } else {
            address prec = prev(p);
            deleteAfter_103022300064(L, p, prec);
        }
        cout << "Lagu dengan judul " << judul << " berhasil dihapus." << endl;
    } else {
        cout << "Lagu dengan judul " << judul << " tidak ditemukan." << endl;
    }
}

void showAllData_103022300064(List L){
    cout << "- - - - INFORMASI DATA TERBARU - - - -" << endl;
    cout << "Daftar lagu: " << endl;
    address temp = first(L);
    if (isEmpty_103022300064(L)) {
        cout << "List kosong." << endl;
    } else {
        while (temp != NULL) {
            cout << "- " << info(temp) << endl;
            temp = next(temp);
        }
    }
}

int selectMenu_103022300064() {
    cout << "===== MENU =====" << endl;
    cout << "0. Keluar" << endl;
    cout << "1. Input Data Baru" << endl;
    cout << "2. Tampilkan Data" << endl;
    cout << "3. Hapus Data Lagu" << endl;
    cout << "4. Input Data Lagu Baru (Posisi Tengah)" << endl;
    cout << "5. Cari Lagu" << endl;

    int input = 0;
    cout << "Pilih Menu: ";
    cin >> input;

    return input;
}

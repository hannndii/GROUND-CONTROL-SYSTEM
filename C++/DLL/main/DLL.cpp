#include "DLL.h"

bool isEmpty(List L) {
    return (L.first == NULL && L.last == NULL);
}

void createList(List &L) {
    L.first = NULL;
    L.last = NULL;
}

void createNewElmt(string bandName, string songTitle, address &P) {
    P = new elmtList;
    P->bandName = bandName;
    P->songTitle = songTitle;
    P->prev = NULL;
    P->next = NULL;
}

void insertFirst(List &L, address P) {
    if (isEmpty(L)) {
        L.first = P;
        L.last = P;
    } else {
        P->next = L.first;
        L.first->prev = P;
        L.first = P;
    }
}

void insertAfter(List &L, address Prec, address P) {
    if (Prec != NULL) {
        P->next = Prec->next;
        P->prev = Prec;

        if (Prec->next != NULL) {
            Prec->next->prev = P;
        } else {
            L.last = P;
        }

        Prec->next = P;
    }
}

void insertLast(List &L, address P) {
    if (isEmpty(L)) {
        L.first = P;
        L.last = P;
    } else {
        P->prev = L.last;
        L.last->next = P;
        L.last = P;
    }
}

void deleteFirst(List &L, address &P) {
    if (!isEmpty(L)) {
        P = L.first;
        if (L.first == L.last) {
            L.first = NULL;
            L.last = NULL;
        } else {
            L.first = L.first->next;
            L.first->prev = NULL;
        }
        delete P;
    }
}

void deleteLast(List &L, address &P) {
    if (!isEmpty(L)) {
        P = L.last;
        if (L.first == L.last) {
            L.first = NULL;
            L.last = NULL;
        } else {
            L.last = L.last->prev;
            L.last->next = NULL;
        }
        delete P;
    }
}

address findLagu(string judul, List L) {
    address P = L.first;
    while (P != NULL) {
        if (P->songTitle == judul) {
            return P;
        }
        P = P->next;
    }
    return NULL;
}

void removeLagu(string judul, List &L) {
    address P = findLagu(judul, L);
    if (P != NULL) {
        if (P == L.first) {
            deleteFirst(L, P);
        } else if (P == L.last) {
            deleteLast(L, P);
        } else {
            P->prev->next = P->next;
            P->next->prev = P->prev;
            delete P;
        }
    }
}

void concat(List L1, List L2, List &L3) {
    createList(L3);
    address P = L1.first;
    while (P != NULL) {
        address temp;
        createNewElmt(P->bandName, P->songTitle, temp);
        insertLast(L3, temp);
        P = P->next;
    }
    P = L2.first;
    while (P != NULL) {
        address temp;
        createNewElmt(P->bandName, P->songTitle, temp);
        insertLast(L3, temp);
        P = P->next;
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

// Fungsi untuk menampilkan data lagu dalam list
void showList(List L) {
    address P = L.first;
    if (P == NULL) {
        cout << "List kosong." << endl;
    } else {
        while (P != NULL) {
            cout << P->bandName << " - " << P->songTitle << endl;
            P = P->next;
        }
    }
}

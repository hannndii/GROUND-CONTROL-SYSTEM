#include <iostream>
#include "SLL.h"

using namespace std;

void createList_103022300064(List &L){
    first(L) = NULL;
}

address createNewElement_103022300064(infotype x){
    address p = new elmlist;
    info(p) = x;
    next(p) = NULL;
    return p;
}

void insertLast_103022300064(List &L, address x){
    if(first(L) == NULL){
        first(L) = x;
    } else{
        address last = first(L);
        while(next(last) != NULL){
            last = next(last);
        }
        next(last) = x;
    }
}

void showAllData_103022300064(List L){
    cout << "- - - - INFORMASI DATA TERBARU - - - -" << endl;
    address print = first(L);
    if(print == NULL){
        cout << "ANDA BELUM MEMASUKKAN DATA!";
    } else {
        cout << "Nomor NIM: ";
        while(print != NULL){
            cout << info(print) << " ";
            print = next(print);
        }
    }
    cout << endl;
}

address findMin(List L){
    address p = first(L);
    address temp = first(L);

    while(next(p) != NULL){
        if(info(temp) < info(next(p))){
            temp = p;
            p = next(p);
        } else {
            p = next(p);
        }
    }
    return temp;
}

void insertMiddle_103022300064(List &L, infotype value){
    address p = first(L);
    address newElement = createNewElement_103022300064(value);

    if (p == NULL) {
        // List kosong, tambahkan elemen pertama
        first(L) = newElement;
        return;
    }

    int count = 0;
    // Hitung jumlah elemen dalam list
    while (p != NULL) {
        count++;
        p = next(p);
    }

    // Hitung posisi tengah (dibulatkan ke bawah jika jumlah elemen genap)
    int middle = count / 2;

    p = first(L);
    address prec = NULL;
    for (int i = 0; i < middle; i++) {
        prec = p;
        p = next(p);
    }

    // Sisipkan elemen baru di posisi tengah
    next(newElement) = p;     // newElement menunjuk ke elemen di posisi tengah sebelumnya
    if (prec != NULL) {
        next(prec) = newElement;  // Prec menunjuk ke newElement
    } else {
        // Jika list hanya memiliki satu elemen
        first(L) = newElement;
    }
}


int selectMenu_103022300064(){
    cout << "===== MENU =====" << endl;
    cout << "0. Keluar" << endl;
    cout << "1. Input Data Baru" << endl;
    cout << "2. Tampilkan Data" << endl;
    cout << "3. Tampilkan Lokasi Data Terkecil" << endl;
    cout << "4. Input Data Baru (Posisi Tengah)"<< endl;

    int input = 0;
    cout << "Pilih Menu(1-4): ";
    cin >> input;

    return input;
}

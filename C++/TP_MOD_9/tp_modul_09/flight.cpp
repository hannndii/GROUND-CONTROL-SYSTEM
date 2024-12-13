#include <iostream>
#include "flight.h"

using namespace std;

void createListJadwal(ListJadwal &L){
    first(L) = nullptr;
}
adr_jadwalP createElemenJadwal(infotype x){
    adr_jadwalP temp = new elemenJadwal;
    info(temp) = x;
    next(temp) = nullptr;

    return temp;
}

void insertLast(ListJadwal &L, adr_jadwalP x){
    if(first(L) == nullptr){
        first(L) = x;
    } else {
        adr_jadwalP p = first(L);
        while(next(p) != nullptr){
            p = next(p);
        }
        next(p) = x;
    }
}

void showJadwal(ListJadwal L) {
    if (first(L) == nullptr) { // Periksa jika list kosong
        cout << "List Jadwal Kosong" << endl;
    } else {
        adr_jadwalP p = first(L);
        while (p != nullptr) {
            cout << "Kode Pesawat: " << info(p).Kode << endl;
            cout << "Pesawat: " << info(p).Jenis << endl;
            cout << "Tanggal Keberangkatan: " << info(p).Tanggal << endl;
            cout << "Waktu Keberangkatan: " << info(p).Waktu << endl;
            cout << "Dari: " << info(p).Asal << endl;
            cout << "Tujuan: " << info(p).Tujuan << endl;
            cout << "Kapasitas: " << info(p).Kapasitas << endl;
            cout << "--------------------------" << endl; // Pemisah antar jadwal
            p = next(p);
        }
    }
}
void deleteFirstJ(ListJadwal &L, adr_jadwalP &temp){
    temp = first(L);
    first(L) = next(first(L));
    next(temp) = nullptr;
}
adr_jadwalP searchJ(ListJadwal L, string dari, string ke, string tanggal){
    adr_jadwalP p = first(L);
    while(next(p) != nullptr){
        if(info(p).Asal == dari && info(p).Tujuan == ke && info(p).Tanggal == tanggal){
            return p;
        }
        p = next(p);
    }
    return nullptr;
}

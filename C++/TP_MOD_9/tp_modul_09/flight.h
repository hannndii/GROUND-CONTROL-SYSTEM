#ifndef FLIGHT_H_INCLUDED
#define FLIGHT_H_INCLUDED
#include <iostream>
#define first(L) L.first
#define info(P) P->info
#define next(P) P->next

using namespace std;

struct jadwalPenerbangan{
    string Kode;
    string Jenis;
    string Tanggal;
    string Waktu;
    string Asal;
    string Tujuan;
    int Kapasitas;
};

typedef struct elemenJadwal *adr_jadwalP;
typedef jadwalPenerbangan infotype;

struct elemenJadwal {
    infotype info;
    adr_jadwalP next;
};

struct ListJadwal {
    adr_jadwalP first;
};

void createListJadwal(ListJadwal &L);
adr_jadwalP createElemenJadwal(infotype x);
void insertLast(ListJadwal &L, adr_jadwalP x);
void showJadwal(ListJadwal L);
void deleteFirstJ(ListJadwal &L, adr_jadwalP &temp);
adr_jadwalP searchJ(ListJadwal L, string dari, string ke, string tanggal);




#endif // FLIGHT_H_INCLUDED

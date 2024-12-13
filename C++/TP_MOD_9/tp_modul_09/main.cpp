#include <iostream>
#include "flight.h"

using namespace std;

int main()
{
    ListJadwal L;
    int n;

    createListJadwal(L);

    cout << "= = = = = = = = = = = = = = = = = = = =" << endl;
    cout << "Masukkan jumlah jadwal penerbangan: ";
    cin >> n;
    cin.ignore();
    cout << "= = = = = = = = = = = = = = = = = = = =" << endl;
    for (int i = 0; i < n; i++){
        infotype data;
        cout << "Jadwal Penerbangan ke " << i+1 << endl;
        cout << "Dari: ";
        cin >> data.Asal;
        cout << "Tujuan: ";
        cin >> data.Tujuan;
        cout << "Pesawat: ";
        cin >> data.Jenis;
        cout << "Kode Pesawat: ";
        cin >> data.Kode;
        cout << "Tanggal Keberangkatan: ";
        cin >> data.Tanggal;
        cout << "Waktu Keberangkatan: ";
        cin >> data.Waktu;
        cout << "Kapasitas: ";
        cin >> data.Kapasitas;
        cin.ignore();

        adr_jadwalP temp = createElemenJadwal(data);
        insertLast(L, temp);

    }
    cout << "= = = = = = = = = = = = = = = = = = = =" << endl;
    cout << "\nDaftar Jadwal Penerbangan Husein Sastranegara\n" << endl;
    showJadwal(L);
    cout << "= = = = = = = = = = = = = = = = = = = =" << endl;

    return 0;
}

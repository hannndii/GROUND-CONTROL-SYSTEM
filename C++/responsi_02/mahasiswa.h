#ifndef MAHASISWA_H_INCLUDED
#define MAHASISWA_H_INCLUDED
#define nama(mahasiswa) mahasiswa->nama
#define semester(mahasiswa) mahasiswa->semester
#define ipk(mahasiswa) mahasiswa->ipk
#include <iostream>
using namespace std;

struct mahasiswa{
    string nama;
    int semester;
    float ipk;
};

//Pass by reference
void buatMahasiswa(mahasiswa *baru, string nama, int semester, float ipk);
bool cumlaude(float ipk);



#endif // MAHASISWA_H_INCLUDED

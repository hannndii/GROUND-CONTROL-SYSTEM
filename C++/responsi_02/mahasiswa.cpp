#include "mahasiswa.h"

void buatMahasiswa(mahasiswa *baru, string nama, int semester, float ipk){
    //Baru adalah pointer, bukan variabel
    baru->nama = nama;
    (*baru).semester = semester;
    ipk(baru) = ipk;
}

bool cumlaude(float ipk){
    return ipk >= 3.5;
}

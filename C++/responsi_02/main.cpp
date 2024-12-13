#include <iostream>
#include "mahasiswa.h"
#include "SLL.h"
using namespace std;

int main()
{
    List kelas;
    mahasiswa rpl;
    address p;

    string nama = "Budi";
    int semester = 3;
    float ipk = 3.9;
    buatMahasiswa(&rpl,nama,semester,ipk);
    cout << rpl.nama << endl;
    cout << rpl.semester << endl;
    cout << rpl.ipk << endl;
    string cld;
    if(cumlaude(rpl.ipk)){
        cld = "YA";
    }else{
        cld = "NA";
    }
    cout << cld << endl;
    createList(kelas);
    p = allocateElement(rpl);
    cout << info(p).nama << endl;
    cout << p<< endl;

    first(kelas) = p;
    cout << first(kelas) << endl;
    cout << info(first(kelas)).nama << endl;
}

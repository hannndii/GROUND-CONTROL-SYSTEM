#include <iostream>

using namespace std;

int main()
{
    int nilai = 12;
    string nilaiB = "handi";
    int *p_nilai = &nilai;
    string *p_nilaiB = &nilaiB;



    cout << "lokasi memory nama: " << p_nilaiB << " dan lokasi memory umur saya: " <<  p_nilai << endl;
    return 0;
}

// iostream(input ouput stream) adalah library dari c++
#include <iostream> 

// run manual di terminal dengan "./nama file.exe"

// jika menggunakan "using namespace" maka tidak perlu menuliskan "std" di main body lagi
using namespace std;

int Func(){
    /* variabel C++: int, double(float), const "nama variabel apa", string, char.*/
    int usia = 50;
    char nama[50] = "Muhammad";

    usia = 52;
  
    // Fungsi "endl" untuk mengakhiri line tersebut dan akan lanjut menambah line baru dibawah. 
    cout << "hello \nhello \n" << endl;
    
    cout << nama << endl;
    
    // multiple output
    cout << "Nama saya adalah " << nama << endl;
    cout << "Usia saya adalah " << usia;


    
    return 0;
}

// int main adalah body utama
int main(){
    Func();
}
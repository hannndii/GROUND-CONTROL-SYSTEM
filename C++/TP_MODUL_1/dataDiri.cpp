#include <iostream>

using namespace std;

int main(){
    string name, nim;
    cout << "Siapa nama pertama anda? ";
    getline(cin, name);
    cout << "Berapa NIM anda? ";
    cin >> nim;

    cout << "Nama saya: " << name << endl;
    cout << "NIM saya: " << nim << endl;
    return 0;    
}
#include <iostream>
#include "trial.h" // Memasukkan header file

using namespace std;

// Definisi constructor
Kerucut::Kerucut(double r, double t) : radius(r), tinggi(t) {}

// Definisi fungsi untuk menghitung volume
double Kerucut::volume() {
    return (1.0 / 3) * 3.14159 * radius * radius * tinggi;
}

int main() {
    // Membuat objek kerucut dengan radius 5 dan tinggi 10
    Kerucut kerucut(5.0, 10.0);

    // Menampilkan volume kerucut
    cout << "Volume kerucut: " << kerucut.volume() << std::endl;

    return 0;
}

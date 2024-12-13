#ifndef MAIN_H
#define MAIN_H

// Deklarasi class Kerucut
class Kerucut {
public:
    double radius;
    double tinggi;

    // Constructor
    Kerucut(double r, double t);

    // Fungsi untuk menghitung volume kerucut
    double volume();
};

#endif

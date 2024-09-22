// SOAL
// Adik mendapat tugas menghitung luas dan keliling dari sebanyak n persegi panjang. Buatlah progarm untuk menghitung luas dan keliling persegi panjang itu.

// Masukan terdiri dari bilangan bulat n yang menyatakan banyaknya persegi panjang. Selanjutnya diikuti sebanyak n pasangan panjang dan lebar.
// Keluaran berupa luas dan keliling persegi panjang.


package main

import "fmt"

func main() {
    var n int
    fmt.Scan(&n)

    for i := 0; i < n; i++ {
        var panjang, lebar int
        fmt.Scan(&panjang, &lebar)

        luas := panjang * lebar
        keliling := 2 * (panjang + lebar)

        fmt.Println(luas, keliling)
    }
}

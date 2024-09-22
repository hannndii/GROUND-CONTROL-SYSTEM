package main  

import "fmt"

func kakiHewan(N int) int {
    if N == 1 {
        return 2
    } else if N % 2 == 0 {
        return 3 + kakiHewan(N-1)
    } else if N % 2 == 1 {
        return 2 + kakiHewan(N-1)
    }
    return 0 // Tambahan untuk menghindari kesalahan kompilasi
}

func main() {
    var N int
    fmt.Scan(&N)
    jumlahKaki := kakiHewan(N)
    fmt.Println(jumlahKaki)
}

// Fungsi ini menerima satu argumen integer N yang mewakili nomor hewan terakhir yang akan dihitung jumlah kaki.
// Fungsi ini menggunakan rekursi untuk menghitung jumlah kaki dari hewan nomor 1 hingga N sesuai aturan yang diberikan:
// Jika N adalah 1, maka fungsi akan mengembalikan 2 kaki (karena hewan nomor 1 memiliki 2 kaki).
// Jika N genap, maka fungsi akan mengembalikan jumlah kaki dari hewan nomor N ditambah dengan jumlah kaki dari hewan nomor sebelumnya (N-1).
// Jika N ganjil, maka fungsi akan mengembalikan jumlah kaki dari hewan nomor N ditambah dengan jumlah kaki dari hewan nomor sebelumnya (N-1).
// Fungsi ini akan terus memanggil dirinya sendiri dengan nilai N yang semakin kecil sampai mencapai kondisi dasar (basis rekursi), yaitu ketika N adalah 1.
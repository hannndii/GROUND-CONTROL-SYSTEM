package main

import "fmt"

func displayOddNumbers(N int) int{
    if N < 1 {
        return 1
    } else if N%2 == 0 {
        displayOddNumbers(N - 1)
    } else {
        displayOddNumbers(N - 2)
        fmt.Println(N)
    }
	return N
}

func main() {
    var N int
    fmt.Print("Masukkan bilangan bulat positif N: ")
    fmt.Scan(&N)
    displayOddNumbers(N)
}

// PENJELASAN
// Dalam kode di atas, kita mendefinisikan fungsi displayOddNumbers yang mengambil argumen N. 
// Fungsi tersebut memeriksa apakah N kurang dari 1 atau genap. Jika N kurang dari 1, fungsi berhenti. Jika N genap, fungsi memanggil dirinya sendiri untuk N-1 untuk 
// memastikan bahwa N yang dilewatkan ke fungsi adalah bilangan ganjil. 
// Jika N adalah bilangan ganjil, fungsi terlebih dahulu memanggil dirinya sendiri untuk N-2 untuk melanjutkan ke bilangan ganjil sebelumnya, lalu mencetak N. 
// Kemudian, di dalam fungsi main, kita meminta pengguna untuk memasukkan nilai N dan memanggil fungsi displayOddNumbers dengan nilai yang dimasukkan.
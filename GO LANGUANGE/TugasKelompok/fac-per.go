package main

import (
    "fmt"
)

func factorial(n int) int {
    if n == 0 {
        return 1
    }
    return n * factorial(n-1)
}

func permutation(x, y int) int {
    if x >= y {
        return factorial(x) / factorial(x-y)
    }
    return factorial(y) / factorial(y-x)
}

func main() {
    var x, y int

    // Meminta masukan dari pengguna
    fmt.Println("Masukkan dua bilangan bulat:")
    fmt.Scan(&x, &y)

    // Menghitung nilai faktorial x dan y
    factX := factorial(x)
    factY := factorial(y)

    // Menghitung nilai permutasi x dan y
    perm := permutation(x, y)

    // Menampilkan hasilnya
    fmt.Printf("%d %d %d\n", factX, factY, perm)
}

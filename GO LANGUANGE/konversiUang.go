
package main

import "fmt"

func konversi(uangLumin int) (int, int, int, int, int) {
    // Konversi ke Stellaris
    uangStellaris := uangLumin / 20
    sisaLumin := uangLumin % 20

    // Konversi ke Nebula Crown
    uangNebula := uangStellaris / 2
    sisaStellaris := uangStellaris % 2

    // Konversi ke Galactum
    uangGalactum := uangNebula / 3
    sisaNebula := uangNebula % 3

    // Konversi ke Quantum Shard
    uangQuantum := uangGalactum / 10
    sisaGalactum := uangGalactum % 10

    return uangQuantum, sisaGalactum, sisaNebula, sisaStellaris, sisaLumin
}

func main() {
    var uangLumin int
    fmt.Print("Masukkan uang dalam koin Lumin: ")
    fmt.Scanln(&uangLumin)

    // Konversi uang
    hasilQuantum, hasilGalactum, hasilNebula, hasilStellaris, hasilLumin := konversi(uangLumin)
    fmt.Println(hasilQuantum, hasilGalactum, hasilNebula, hasilStellaris, hasilLumin)
}

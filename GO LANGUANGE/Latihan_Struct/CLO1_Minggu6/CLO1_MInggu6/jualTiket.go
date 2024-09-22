package main

import "fmt"

type Tiket struct {
    NamaFilm     string
    Hari         string
    TanggalMerah bool
}

func main() {
    var keterangan Tiket

	fmt.Scanf("%s %s %t", &keterangan.NamaFilm, &keterangan.Hari, &keterangan.TanggalMerah)

	result := HitungHarga(keterangan)
    fmt.Println(result)
    
}

func HitungHarga(tiket Tiket) int {
    harga := 50000
    if tiket.Hari == "Sabtu" || tiket.Hari == "Minggu" || tiket.TanggalMerah == true {
        harga += harga / 2
    } else {
		harga = harga
	}
    return harga
}

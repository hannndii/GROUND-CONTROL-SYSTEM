package main

import "fmt"

const NMAX int = 5

type tabInt struct {
	info [NMAX]int
	n    int
}

func main() {
	var nilaiAkhir tabInt
	bacaNilai(&nilaiAkhir)
	cetakNilai(nilaiAkhir)
}

func bacaNilai(NA *tabInt) {
	fmt.Println("Masukkan nilai-nilai (hingga ", NMAX, " angka):")
	for i := 0; i < NMAX; i++ {
		var nilai int
		fmt.Scan(&nilai)
		NA.info[i] = nilai
		NA.n++
	}
}

func cetakNilai(NA tabInt) {
	fmt.Print("Nilai yang terdapat pada array:")
	for i := 0; i < NA.n; i++ {
		fmt.Printf(" %d", NA.info[i])
	}
	fmt.Println()
}

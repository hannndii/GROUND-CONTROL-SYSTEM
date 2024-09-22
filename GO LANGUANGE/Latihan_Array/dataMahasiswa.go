package main

import "fmt"

const NMAX int = 100

type tabStudent[NMAX] DataMurid

type DataMurid struct{
	nama string
	umur int
	tinggi int
	berat int
}

func main(){
	var T tabStudent
	var n int
	var maxNama string

	max := T[0].berat

	fmt.Scan(&n)
	for i := 1; i <= n; i++{
		fmt.Scan(&T[i].nama)
		fmt.Scan(&T[i].umur)
		fmt.Scan(&T[i].tinggi)
		fmt.Scan(&T[i].berat)
		if max < T[i].berat{
			max = T[i].berat
			maxNama = T[i].nama
		}
	}

	fmt.Printf("Berat tertinggi adalah %s", maxNama)
}


package main

import "fmt"

const NMAX int = 10

type tableInt[NMAX] int

func bacaNilai(NA *int, n *int){
	fmt.Scan(&n)
	if n > NMAX {
		n = NMAX
	}
	NA = n

}

func cetakNilai(NA tableInt, n int){
	if NA == 0{
		fmt.Println("array kosong")
	} else {
		fmt.Println("Nilai yang terdapat pada array: ", NA)
	}
}

func main(){
	var nilaiAkhir tableInt
	var banyakNilai int

	bacaNilai(&nilaiAkhir, &banyakNilai)
	cetakNilai(nilaiAkhir, banyakNilai)
}


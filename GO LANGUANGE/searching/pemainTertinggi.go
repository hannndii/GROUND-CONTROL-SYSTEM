package main

import "fmt"

const NMAX = 50

type listPemain struct {
	nama    string
	posisi  string
	tinggi  int
}

type arrPemain [NMAX]listPemain

func main() {
	var dataPemain arrPemain
	var n int

	fmt.Println("Masukkan jumlah data:")
	fmt.Scanln(&n)
	read(&dataPemain, n)
	high := findPemain(dataPemain, n)
	print(dataPemain, n)

	fmt.Printf("Pemain tertinggi bernama: %s dengan tinggi %d\n", high.nama, high.tinggi)
}

func read(T *arrPemain, n int) {
	fmt.Println("Masukkan data pemain:")
	for i := 0; i < n; i++ {
		fmt.Scan(&T[i].nama, &T[i].posisi, &T[i].tinggi)
	}
}

func print(T arrPemain, n int) {
	fmt.Println("Data Pemain Barcelona: ")
	for i := 0; i < n; i++ {
		fmt.Println(T[i].nama, T[i].posisi, T[i].tinggi)
	}
}

func findPemain(T arrPemain, n int) listPemain {
	var highest listPemain
	max := T[0].tinggi

	for i := 1; i < n; i++ {
		if max < T[i].tinggi {
			max = T[i].tinggi
			highest = T[i]
		}
	}

	return highest
}

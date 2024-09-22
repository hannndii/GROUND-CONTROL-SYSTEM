package main

import "fmt"

const NMAX int = 20

type tabInt[NMAX] int

func main() {

	// Deklarasi variabel
	var data tabInt
	var nData int
	var x1 int
	
	fmt.Scan(&x1)

	fmt.Scan(&nData)
	bacaData(&data, &nData)

	cetakData(data, nData)

	fmt.Print("Hasil pencarian: ")
	if sequentialSearch(data, nData, x1) {
		fmt.Printf("Bilangan ditemukan. Terdapat %d bilangan %d.\n", frekuensiBilangan(data, nData, x1), x1)
	} else {
		fmt.Println("Bilangan tidak ditemukan.")
	}
}

func bacaData(A *tabInt, n *int) {
	for i := 0; i < *n; i++{
		fmt.Scan(&(*A)[i])
	}
}

func cetakData(A tabInt, n int) {
	fmt.Print("Data Bilangan: ")
	for i := 0; i < n; i++{
		if i == 10 {
			break
		}
		fmt.Print(A[i], " ")
	}
	fmt.Println()
}

func frekuensiBilangan(A tabInt, n, x int) int {
	ketemu := x
	total := 0
	i := 0
	for i < n-1{
		if ketemu == A[i] {
			total += 1
		} else {
			total += 0
		}
		i++
	}
	return total
}

func sequentialSearch(A tabInt, n int, x int) bool {
	var hasil bool
	i := 0
	for i < n-1{
		if x == A[i] {
			hasil = true
		} 
		i++
	}
	return hasil
}
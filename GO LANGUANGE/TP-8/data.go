// package main

// import (
// 	"fmt"
// )

// const NMAX = 10

// type tabInt [NMAX]int

// func main() {
// 	var data tabInt
// 	var nData int
// 	baca(&data, &nData)
// 	cetak(data, nData)
// 	fmt.Println("Jumlah:", jumlah(data, nData))
// 	fmt.Println("Rata-rata:", rata_rata(data, nData))
// }

// func baca(A *tabInt, n *int) {
// 	/*
// 		IS: Array A dan n terdefinisi sembarang
// 		Proses: Setiap bilangan yang dibaca ditampung dalam sebuah variabel. 
// 		    Jika bilangan negatif, maka ubah menjadi bilangan positif dan
// 		    masukan ke dalam array. Pembacaan berakhir jika terbaca bilangan 0.
// 		FS: Array A sebanyak n elemen berisi bilangan positif
// 	*/
// 	fmt.Println("Masukkan serangkaian bilangan (0 untuk berhenti):")
// 	for i := 0; i < NMAX; i++ {
// 		var bil int
// 		fmt.Scan(&bil)
// 		if bil == 0 {
// 			break
// 		}
// 		if bil < 0 {
// 			bil = -bil
// 		}
// 		A[i] = bil
// 		*n++
// 	}
// }

// func cetak(A tabInt, n int) {
// 	/*
// 		IS: Array A terdefinisi sebanyak n elemen
// 		FS: Array A tercetak di layar
// 	*/
// 	fmt.Println("Elemen array:")
// 	for i := 0; i < n; i++ {
// 		fmt.Print(A[i], " ")
// 	}
// 	fmt.Println()
// }

// func jumlah(A tabInt, n int) int {
// 	/* Mengembalikan jumlah dari nilai bilangan pada elemen array A */
// 	total := 0
// 	for i := 0; i < n; i++ {
// 		total += A[i]
// 	}
// 	return total
// }

// func rata_rata(A tabInt, n int) float64 {
// 	/* Mengembalikan rata-rata bilangan */
// 	if n == 0 {
// 		return 0
// 	}
// 	return float64(jumlah(A, n)) / float64(n)
// }


package main

import "fmt"

const NMAX = 10

type tabInt[NMAX]int

func baca(A *tabInt, n *int){
	for i := 0; i < NMAX; i++{
		var bil int
		fmt.Scan(&bil)
		if bil == 0 {
			break
		}
		if bil < 0{
			bil = bil * -1
		}
		A[i] = bil
		*n++
	}
}

func cetak(A tabInt, n int) {
	for i := 0; i < n; i++{
		fmt.Println(A[i])
	}
}

func jumlah(A tabInt, n int) int{
	total := 0
	for i := 0; i < n; i++{
		total += A[i]
	}
	return total
}

func rata_rata(A tabInt, n int) float64{
	if n == 0{
		return 0
	} else{
		return float64(jumlah(A, n)) / float64(n)
	}
}

func main(){
	var data tabInt
	var nData int

	baca(&data, &nData)
	cetak(data, nData)
	fmt.Println("Jumlah:", jumlah(data, nData))
	fmt.Println("Rata-rata:", rata_rata(data, nData))



}
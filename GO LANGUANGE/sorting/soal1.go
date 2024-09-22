package main

import "fmt"

const NMAX = 1024

type DataStudent struct{
	nama string
	id string
	stock int
	harga int
} 

type tabNilai[NMAX] DataStudent

func input(A *tabNilai, n int){
	for i := 0; i < n; i++{
		fmt.Scanf("%s %s %d %d \n", &(*A)[i].nama, &(*A)[i].id, &(*A)[i].stock, &(*A)[i].harga)
	}
}

func cetak(A tabNilai, n int){
	fmt.Printf("%20s %20s %20s %10s", "Nama Barang", "ID Barang" , "Jumlah Stock", "Harga")
	fmt.Println()
	for i := 0; i < n; i++{
		fmt.Printf("%20s %20s %20d %10d \n", A[i].nama, A[i].id, A[i].stock, A[i].harga)
	}
}

func urutHarga(A *tabNilai, n int){
	var pass, i, idx_min int
	var temp DataStudent

	pass = 1
	for pass < n{
		idx_min = pass-1
		i = pass
		for i < n{
			if A[idx_min].harga < A[i].harga{
				idx_min = i
			}
			i = i + 1
		}
		temp = (*A)[pass-1]
		(*A)[pass-1] = (*A)[idx_min]
		(*A)[idx_min] = temp
		pass = pass + 1
	}
}

func urutStock(A *tabNilai, n int){
	var pass, i, idx_min int
	var temp DataStudent

	pass = 1
	for pass < n{
		idx_min = pass-1
		i = pass
		for i < n{
			if A[idx_min].stock > A[i].stock{
				idx_min = i
			}
			i = i + 1
		}
		temp = (*A)[pass-1]
		(*A)[pass-1] = (*A)[idx_min]
		(*A)[idx_min] = temp
		pass = pass + 1
	}
}

func cariID(A tabNilai, n int, x string){
	i := 0
	for i < n{
		if A[i].id == x{
			fmt.Println("Data Ketemu!!")
			fmt.Println("==================================================================")
			fmt.Printf("%20s %20s %20s %10s", "Nama Barang", "ID Barang" , "Jumlah Stock", "Harga")
			fmt.Println()
			fmt.Printf("%20s %20s %20d %10d \n", A[i].nama, A[i].id, A[i].stock, A[i].harga)
			break

		} else {
			fmt.Println("Data Tidak Ada!!")
			break
		}
		i++
	}
}

// 1. Urut berdasarkan harga
// 1. Urut berdasarkan stock (descending)
// 1. cari ID, tampilkan harga/nama (descending)

func main(){
	var arrData tabNilai
	var nData int
	var findID string

	fmt.Scanln(&nData)
	input(&arrData, nData)
	fmt.Println("Data Perusahaan")
	cetak(arrData, nData)
	fmt.Println("Data Perusahaan terurut berdasarkan Harga")
	urutHarga(&arrData, nData)
	cetak(arrData, nData)
	fmt.Println("Data Perusahaan terurut berdasarkan stock")
	urutStock(&arrData, nData)
	cetak(arrData, nData)
	fmt.Print("Masukkan ID:")
	fmt.Scan(&findID)
	cariID(arrData, nData, findID)

}
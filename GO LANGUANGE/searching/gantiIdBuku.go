package main

import "fmt"

type daftarBuku struct{
	id int
	namaBuku string
	genre string
	tahun int
}

type arrBuku[10] daftarBuku

func input(dBuku *arrBuku, n int){
	k := 1
	for i := 0; i < n; i++{
		fmt.Printf("Buku Ke-%d", k)
		fmt.Println()
		fmt.Scan(&dBuku[i].id, &dBuku[i].namaBuku, &dBuku[i].genre, &dBuku[i].tahun)
		k++
	}
}


func gantiData(dBuku *arrBuku, n int, x int, update string){
	found := false
	for i := 0; i < n; i++ {
		if (*dBuku)[i].id == x {
			(*dBuku)[i].genre = update
			found = true
			break
		}
	}
	if !found {
		fmt.Println("Data tidak ada")
	} else {
		cetak(*dBuku, n)
	}
}

func cetak(dBuku arrBuku, n int){
	k := 1
	fmt.Println("Data Buku Perpustakaan:")
	for i := 0; i < n; i++{
		fmt.Println(dBuku[i].id, " ", dBuku[i].namaBuku, " ",dBuku[i].genre, " ",dBuku[i].tahun)
		k++
	}
}
var dataBuku arrBuku
	var nData, i int
	var DataX int
	var DataM string


func main(){
	
	fmt.Scan(&nData)

	input(&dataBuku, nData)
	fmt.Print("Masukkan ID buku:")
	fmt.Scan(&DataX)
	fmt.Print("Masukkan genre buku yang ingin diganti:")
	fmt.Scan(&DataM)
	cetak(dataBuku, nData)
	gantiData(&dataBuku, nData, DataX, DataM)
	
	

}
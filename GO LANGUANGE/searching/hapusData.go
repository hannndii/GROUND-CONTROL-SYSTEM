package main

import "fmt"

// Struktur Buku
type daftarBuku struct {
	id       int
	namaBuku string
	genre    string
	tahun    int
}

// Tipe array buku dengan ukuran tetap
type arrBuku [10]daftarBuku

// Fungsi untuk input data buku
func input(dBuku *arrBuku, n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("Buku Ke-%d\n", i+1)
		fmt.Print("Masukkan ID, Nama Buku, Genre, Tahun: ")
		fmt.Scan(&(*dBuku)[i].id, &(*dBuku)[i].namaBuku, &(*dBuku)[i].genre, &(*dBuku)[i].tahun)
	}
}

// Fungsi untuk mencetak data buku
func cetak(dBuku arrBuku, n int) {
	fmt.Println("Data Buku Perpustakaan:")
	for i := 0; i < n; i++ {
		fmt.Println(dBuku[i].id, " ", dBuku[i].namaBuku, " ", dBuku[i].genre, " ", dBuku[i].tahun)
	}
}

// Fungsi untuk menghapus buku berdasarkan ID
func hapusData(dBuku *arrBuku, n *int, x int) {
	found := false
	j := 0
	for i := 0; i < *n; i++ {
		if (*dBuku)[i].id != x {
			(*dBuku)[j] = (*dBuku)[i]
			j++
		} else {	
			found = true
		}
	}
	if found {
		*n--
		(*dBuku)[*n] = daftarBuku{} // Optional: Clear the last element
		fmt.Println("Data telah dihapus")
		cetak(*dBuku, *n)
	} else {
		fmt.Println("Data tidak ada")
	}
}

func main() {
	var dataBuku arrBuku
	var nData, dataX int

	fmt.Print("Masukkan jumlah data buku (maksimal 10): ")
	fmt.Scan(&nData)

	if nData > 10 {
		fmt.Println("Jumlah data buku tidak boleh lebih dari 10")
		return
	}

	input(&dataBuku, nData)

	fmt.Print("Masukkan ID buku yang ingin dihapus: ")
	fmt.Scan(&dataX)

	hapusData(&dataBuku, &nData, dataX)
}

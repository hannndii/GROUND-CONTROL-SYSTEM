package main

import "fmt"

const students = 100

type Students struct{
	nama [students] string 
	kelas [students] int
	nilai [students] int
}

func dataStudents(T Students, n int) string {
	var k int
	var hasil string

	k = 0
	for k < n{
		 if T.nama[k] == "handi" && T.nilai[k] == 85 {
			hasil = "Ada"
		 } else {
			hasil = "Tidak ada"
		 }
		 k = k+1
	}
	return hasil
}

func main(){
	var dataMahasiswa Students
	var n int
	
	fmt.Print("Masukkan jumlah data")
	fmt.Scan(&n)
	for i := 0; i < n; i++{
		fmt.Print("Masukkan nama mahasiswa: ")
		fmt.Scan(&dataMahasiswa.nama[i])
		fmt.Print("Masukkan kelas mahasiswa: ")
		fmt.Scan(&dataMahasiswa.kelas[i])
		fmt.Print("Masukkan nilai mahasiswa: ")
		fmt.Scan(&dataMahasiswa.nilai[i])
	}

	result := dataStudents(dataMahasiswa, n)

	fmt.Println("Hasil: ", result)
}
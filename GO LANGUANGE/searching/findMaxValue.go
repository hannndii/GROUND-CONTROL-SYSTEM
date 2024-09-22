package main

import "fmt"

const NMAX = 40

type dataMahasiswa struct{
	nama [NMAX]string
	nilai [NMAX]int
}

func valueMax(M dataMahasiswa, n int) int{
	var k, max int

	max = M.nilai[0]
	k = 1
	for k < n{
		if max < M.nilai[k]{
			max = M.nilai[k]
		}
		k = k + 1
	}
	return max
}

var mahasiswa [NMAX]dataMahasiswa

func main(){
	var n int
	var dataM dataMahasiswa

	fmt.Println("Tentukan jumlah mahasiswa: ")
	fmt.Scan(&n)
	for i := 0; i < n; i++{
		fmt.Print("Nama Mahasiswa:")
		fmt.Scan(&dataM.nama[i])
		fmt.Print("Nilai fisika:")
		fmt.Scan(&dataM.nilai[i])
	}
	result := valueMax(dataM, n)
	
	fmt.Print("nilai tertinggi: ", result)
}
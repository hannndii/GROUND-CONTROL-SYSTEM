package main

import "fmt"

const NMAX int = 100

type tabGolA[NMAX]Tim
type tabGolB[NMAX]Tim

type Tim struct{
	timA int
	timB int
}

func inputDataA(t *tabGolA, n int){
	fmt.Println("Masukkan jumlah Gol tim A")
	for i := 0; i < n; i++{
		fmt.Scan(&t[i].timA)
		if t[i].timA < 0{
			break
		}

	}

}
func inputDataB(t *tabGolB, n int){
	fmt.Println("Masukkan jumlah Gol tim B")
	for j := 0; j < n; j++{
		fmt.Scan(&t[j].timB)
		if t[j].timB < 0{
			break
		}

	}

}

func cetakDataTimA(t tabGolA, n int){
	for i := 0; i < n; i++{
		fmt.Println("Data tim A adalah ", t[i].timA)
	}
}

func cetakDataTimB(t tabGolB, n int){
	for i := 0; i < n; i++{
		fmt.Println("Data tim B adalah ", t[i].timA)
	}
}

// func rerata(t tabGolA, T tabGolB, n int){

// }



func main(){
	var jumlahGolA tabGolA
	var jumlahGolB tabGolB
	var jumlahData int

	fmt.Scan(&jumlahData)

	inputDataA(&jumlahGolA, jumlahData)
	inputDataB(&jumlahGolB, jumlahData)
	cetakDataTimA(jumlahGolA, jumlahData)
	cetakDataTimB(jumlahGolB, jumlahData)

	
}
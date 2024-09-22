package main

import (
	"fmt"
	"math"
)

const MAX_DATA = 1000

type Wisudawan struct {
	Nama     string
	NIM      string
	Eprt     int
	Semester int
	Ipk      float64
}

func bacaDataWisudawan() []Wisudawan {
	var dataWisudawan []Wisudawan
	var wisudawan Wisudawan

	for i := 0; i < MAX_DATA; i++ {
		fmt.Println("Masukkan nama: ")
		fmt.Scan(&wisudawan.Nama)
		fmt.Println("Masukkan nim: ")
		fmt.Scan(&wisudawan.NIM)
		fmt.Println("Masukkan eprt: ")
		fmt.Scan(&wisudawan.Eprt)
		fmt.Println("Masukkan semester: ")
		fmt.Scan(&wisudawan.Semester)
		fmt.Println("Masukkan ipk: ")
		fmt.Scan(&wisudawan.Ipk)

		if wisudawan.NIM == "none" {
			break
		}

		dataWisudawan = append(dataWisudawan, wisudawan)
	}

	return dataWisudawan
}

func cariEprtTertinggi(dataWisudawan []Wisudawan) int {
	maxEprt := 0
	for _, wisudawan := range dataWisudawan {
		if wisudawan.Eprt > maxEprt {
			maxEprt = wisudawan.Eprt
		}
	}
	return maxEprt
}

func cariIpkTerendah(dataWisudawan []Wisudawan) float64 {
	minIpk := math.MaxFloat64
	for _, wisudawan := range dataWisudawan {
		if wisudawan.Ipk < minIpk {
			minIpk = wisudawan.Ipk
		}
	}
	return minIpk
}

func hitungRataRataSemester(dataWisudawan []Wisudawan) float64 {
	totalSemester := 0
	for _, wisudawan := range dataWisudawan {
		totalSemester += wisudawan.Semester
	}
	return float64(totalSemester) / float64(len(dataWisudawan))
}

func main() {
	dataWisudawan := bacaDataWisudawan()

	if len(dataWisudawan) == 0 {
		fmt.Println("Tidak ada data wisudawan yang dimasukkan.")
		return
	}

	eprtTertinggi := cariEprtTertinggi(dataWisudawan)
	ipkTerendah := cariIpkTerendah(dataWisudawan)
	rataRataSemester := hitungRataRataSemester(dataWisudawan)

	fmt.Println("Eprt Tertinggi:", eprtTertinggi)
	fmt.Println("Ipk Terendah:", ipkTerendah)
	fmt.Printf("Rata-rata Semester: %.2f\n", rataRataSemester)
}

package main

import "fmt"

type mobil struct{
	merek string
	tahun int
	kecepatan int
}

func main(){
	var m1, m2, m3 mobil
	var rata_rata_kecepatan float64

	fmt.Scan(&m1.merek, &m1.tahun, &m1.kecepatan)
	fmt.Scan(&m2.merek, &m2.tahun, &m2.kecepatan)
	fmt.Scan(&m3.merek, &m3.tahun, &m3.kecepatan)

	rata_rata_kecepatan = (float64(m1.kecepatan) + float64(m2.kecepatan) + float64(m3.kecepatan)) / 3

	fmt.Printf("Rata rata kecepatan mobil %s (%d), ", m1.merek, m1.tahun)
	fmt.Printf("mobil %s (%d), dan mobil %s (%d) :", m2.merek, m2.tahun, m3.merek, m3.tahun)
	fmt.Printf("%.2f\n", rata_rata_kecepatan)
}
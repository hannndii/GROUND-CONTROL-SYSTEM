package main

import "fmt"

const NMAX = 100

type tabInt[NMAX] int

func findMax(T tabInt, n, x int) string{
	var k int
	var hasil string

	k = 0

	for k < n {
		if T[k] == x  {
			hasil = "Ketemu"
			break
		}
		k = k + 1
  }
  return hasil
}
 
func main() {
	var n, targetX int
	var T tabInt

	fmt.Println("Masukkan jumlah array n: ")
	fmt.Scan(&n)
	fmt.Println("Masukkan angka yang dicari: ")
	fmt.Scan(&targetX)

	for i := 0; i < n; i++{
		fmt.Scan(&T[i])
	}
	result := findMax(T, n, targetX)
	fmt.Println("hasil: ", result)
}
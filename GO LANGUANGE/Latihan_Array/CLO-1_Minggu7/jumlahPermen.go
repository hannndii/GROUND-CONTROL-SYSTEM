package main

import "fmt"

const nMax int = 1000
type tabInt [nMax]int

func isiArray(arr *tabInt, n int) {
	var i int
	for i = 0; i < n; i++ {
		fmt.Scan(&(*arr)[i])
	}
}

func pembagianPermen(arr tabInt, n int) tabInt {
	var permen tabInt
	var i, total int

	if n == 0 {
		return permen
	} else {
		for i = 0; i < n; i++ { 
			permen[i] = 1
			total += 1
		}
		for i = 1; i < n; i++ {
			if arr[i] > arr[i-1] {
				permen[i] = permen[i-1] + 1
				total += permen[i]
			}
		}
		i = n - 1
		for i > 0 {
			if arr[i] < arr[i-1] && permen[i] >= permen[i-1] {
				permen[i-1] = permen[i] + 1
				total += permen[i-1]
			}
			i--
		}
	}
	return permen
}

func permenMinimum(arr tabInt, n int) int {
	var i, total int
	total = 0
	for i = 0; i < n; i++ {
		total += arr[i]
	}
	return total
}

func cetak(arr tabInt, n int) {
	var i int
	var permen tabInt
	permen = pembagianPermen(arr, n)
	for i = 0; i < n; i++ {
		fmt.Printf("%d ", permen[i])
	}
	fmt.Println()
	fmt.Println(permenMinimum(permen, n))
}

func main() {
	var N int
	var rank tabInt
	fmt.Scan(&N)
	isiArray(&rank, N)
	cetak(rank, N)
}

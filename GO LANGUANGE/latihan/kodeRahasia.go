// digit pertama + digit terakhir

package main

import "fmt"

func hitungJumlah(n int) int{
	digitPertama := n / 1000
	digitKedua := n % 10

	hasil := digitPertama + digitKedua
	return hasil
}

func main(){
	var n int

	fmt.Scan(&n)

	hitungJumlah(n)
	fmt.Println(hitungJumlah(n))
}

// 2045
// 2045 div 1000 = 2
// 2045 mod 10 

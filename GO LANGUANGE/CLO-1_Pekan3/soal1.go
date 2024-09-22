package main

import "fmt"

func penjumlahan(N int){
	var i int
	hasil := 0
	for i = 1; i <= N*10; i++{
		hasil += i
		fmt.Println(hasil)
	}
}

func main(){
	var N int
	fmt.Scan(&N)
	
	penjumlahan(N)
}


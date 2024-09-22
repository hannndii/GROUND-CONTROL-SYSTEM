package main

import "fmt"

const N int = 1000

func lebihBesar(T [N]int, jumlah, target int) int {
	var find, k int

	find = 0
	k = 0
	if k == jumlah-1{
		return 1
	} else {
		if T[k] == target{
			// belum selesai
		} 
		return lebihBesar(T, jumlah, target+1 )
	}
	
}

func main(){
	var T [N]int
var jumlah, target int
fmt.Scanln(&jumlah)
for i := 0; i < jumlah; i++ {
	fmt.Scan(&T[i])
}
fmt.Scan(&target)
fmt.Println(lebihBesar(T, jumlah, target))
}

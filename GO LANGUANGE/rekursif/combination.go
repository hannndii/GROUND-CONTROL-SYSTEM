package main 

import "fmt"

func comb(n, r int) int{
	if r == 0 {
		return 1
	} else {
		return n*comb(n-1, r-1)
	}
}

func main(){
	var n, r int

	fmt.Scan(&n, &r)
	
	result := comb(n, r)
	fmt.Println(result)
}

// belum selesai
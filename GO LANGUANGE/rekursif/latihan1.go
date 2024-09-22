package main

import "fmt"

func sum_std(i, n int) int{
	if i == n{
		return i
	} else {
		return i + sum_std(i+1, n)
	}
}

func main(){
	var i, n int

	fmt.Scan(&n)
	result := sum_std(i, n)
	fmt.Println(result)
}


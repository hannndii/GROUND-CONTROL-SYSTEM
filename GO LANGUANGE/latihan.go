package main

import "fmt"

func fib(n, r int) int{
	if r == 0{
		return 1
	} else {
		return n * fib(n-1, r-1)
	}
}

func main(){
	var n, r int
    
	fmt.Scan(&n, &r)

	hasil := fib(n, r)

	fmt.Println(hasil)
}


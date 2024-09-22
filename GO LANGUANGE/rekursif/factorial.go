package main

import "fmt"

func fac(n int) int{
	if n == 1 {
		return 1
	} else {
		return n * fac(n-1)
	}

}

func main(){
	fac(5)
	fmt.Println(fac(5))
}


// 5 * 4 * 3 * 2 * 1
package main

import "fmt"

func cetak(x int) {
	if x == 10 {
		fmt.Println(x)
	} else {
	    fmt.Println(x)
		cetak(x+1)
		fmt.Println(x)
	}
}

func main(){
	cetak(6)
}

// package main

// import "fmt"

// func fact(n int) int {
//     if n == 0 {
//         return 1
//     }
//     return n * fact(n-1)
// }

// func main() {
//     fmt.Println(fact(7))

//     var fib func(n int) int

//     fib = func(n int) int {
//         if n < 2 {
//             return n
//         }

//         return fib(n-1) + fib(n-2)
//     }

//     fmt.Println(fib(7))
// }
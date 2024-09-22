package main 

import "fmt"

func perkalianLoop(n, m int) int{
	var i, hasil int

	hasil = 0
	for i = 1; i <= m; i++ {
		hasil += n  
	} 
	return hasil
}

func main(){
	perkalianLoop(2, 10)
	fmt.Println(perkalianLoop(2, 10))
}

// loop1 = hasil(0) + 2
// loop2 = hasil(2) + 2
// loop3 = hasil(4) + 2
// loop4 = hasil(6) + 2
// loop5 = hasil(8) + 2
// loop6 = hasil(10) + 2
// loop7 = hasil(12) + 2
// loop8 = hasil(14) + 2
// loop9 = hasil(16) + 2
// loop10 = hasil(18) + 2
// loop11 = hasil(20)

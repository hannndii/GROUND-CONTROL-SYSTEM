package main 

import "fmt"

const N int = 1000

func ganjilTerbesar(T [N]int, jumlah int) int {
	var max, k int

	k = 1
	max = -1
	for k < jumlah{
		if max < T[k] && T[k] % 2 != 0{
			max = T[k]
		}
		k = k + 1
	}

	return max
}

func main(){
	var T [N]int
    var jumlah int

	fmt.Scanln(&jumlah)

	for i := 0; i < jumlah; i++ {
	fmt.Scan(&T[i])
   }

   fmt.Println(ganjilTerbesar(T, jumlah))
}

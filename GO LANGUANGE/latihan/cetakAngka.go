package main  

import "fmt"

func cetakAngka(x, y, a, b int){
	for i := x; i <= y; i++{
		if i / (100) != a && i % 10 != b{
			fmt.Println(i)
		} 
	} 
}

func main(){
	var x, y, a, b int

	fmt.Scan(&x, &y, &a, &b)

	cetakAngka(x, y, a, b)
}


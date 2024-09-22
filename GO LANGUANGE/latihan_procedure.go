package main

import "fmt"

func totalDiskon(x float32, y bool){
	var hasil float32

	if x > 100000 && y == true{
		hasil = x - (x * 0.1)
	} else if x > 100000 && y == false{
		hasil = x - (x * 0.05)
	} else {
		fmt.Println("Anda tidak dapat diskon")
	}
	fmt.Println(hasil)
}

func main(){
	var x float32
	var y bool

	fmt.Scan(&x, &y)

	totalDiskon(x, y)
}
package main

import (
	"fmt"
	"math"
)

type titik struct{
	sumbuA float64
	sumbuB float64
	sumbuC float64
	sumbuD float64

}

func jarakG(p1, p2 titik) float64{
	hasil := (math.Pow(p1.sumbuA - p1.sumbuC, 2) + math.Pow(p2.sumbuB - p2.sumbuD, 2))
	return hasil
	  
}

func akar(x float64) float64{
	return math.Sqrt(x)
}

func main(){
	var p1, p2 titik

	fmt.Println("Masukkan nilai a, b, c, d: ")
	fmt.Scan(&p1.sumbuA, p1.sumbuC, p2.sumbuB, p2.sumbuD)
	
	result := akar(jarakG(p1, p2))

	fmt.Println("Hasil: ", result)
}

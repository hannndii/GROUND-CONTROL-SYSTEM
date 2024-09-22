package main

import (
	"fmt"
)

func kalkulator(operand1, operand2 float64, operator string) float64 {
	var hasil float64
	switch {
	case operator == "+":
		hasil = operand1 + operand2
	case operator == "-":
		hasil = operand1 - operand2
	case operator == "*":
		hasil = operand1 * operand2
	case operator == "/":
		if operand2 != 0 {
			hasil = operand1 / operand2
		} else {
			return 0
		}
	}
	return hasil
}

func main() {
	var operand1, operand2 float64
	var operator string

	fmt.Println("Masukkan 2 bilangan real dan operator (+, -, *, /): ")
	fmt.Scan(&operand1, &operand2, &operator)

	result := kalkulator(operand1, operand2, operator)
	fmt.Printf("Hasil: %.2f\n", result)
}



// package main

// import ("fmt")

// func kalkulator(x, y float64, operator rune) float64{
// 	var hasil int
// 	if operator == '+' {
// 		hasil = x + y
// 	} else if operator == '-'{
// 		hasil = x - y
// 	} else if operator == '*' {
// 		hasil = x * y
// 	} else if operator == '/'{
// 		hasil = x / y
// 	} else {
// 		fmt.Println("hahahah")
// 	}
// 	return hasil
	
// }

// func main(){
// 	var nilai1, nilai2 float64
// 	var operator rune
	
// 	fmt.Scan(&nilai1, &nilai2, &operator)
// 	result := kalkulator(nilai1, nilai2, operator)

// 	fmt.Printf("hasil: %f ", result)
// }
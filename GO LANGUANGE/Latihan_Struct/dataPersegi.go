package main

import "fmt"

type rectangle struct{
	length, width int
	color string
	property geometry
}

type geometry struct{
	luas, keliling int
}

func isiData(persegi *rectangle) {
	fmt.Println("Masukkan panjang persegi:")
	fmt.Scan(&persegi.length)
	fmt.Println("Masukkan lebar persegi:")
	fmt.Scan(&persegi.width)
	fmt.Println("Masukkan warna persegi:")
	fmt.Scan(&persegi.color)
}

func menu(){
	fmt.Println("--------------------")
	fmt.Printf("%13s\n", "M E N U")
	fmt.Println("--------------------")
	fmt.Println("1. Hitung Luas")
	fmt.Println("2. Hitung Keliling")
	fmt.Println("3. Exit")
	fmt.Println("--------------------")
}

func board(persegi *rectangle){
	persegi.property.luas = persegi.length * persegi.width
	fmt.Printf("Hasil: %d ,dengan warna persegi: %s", persegi.property.luas, persegi.color)
}


func roam(persegi *rectangle){
	persegi.property.keliling = 2*(persegi.length + persegi.width)
	fmt.Printf("Hasil: %d ,dengan warna persegi: %s", persegi.property.keliling, persegi.color)
}

func main(){
	var persegi rectangle
	var input int
	
	menu()
	fmt.Scan(&input)
	switch  {
	case input == 1:
		isiData(&persegi)
		board(&persegi)
	case input == 2:
		isiData(&persegi)
		roam(&persegi)
	case input == 3:
		fmt.Println("Terimakasih")
		break
	}
	
}


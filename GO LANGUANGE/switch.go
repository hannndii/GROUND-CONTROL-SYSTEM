package main

import "fmt"

func main(){
	var name string

	fmt.Scanln(&name)

	switch name{
	case "Endihan" :
		fmt.Println("Betul itu nama saya")

	case "Muhammad":
		fmt.Println("Itu nama depan saya")

	default:
		fmt.Println("Salah nama")
	}

	// switch length := len(name); length >= 7{
	// case true:
	// 	fmt.Println("Jumlah karakter disetujui")

	// case false:
	// 	fmt.Println("Jumlah karakter melebihi batas")
	// }

	length := len(name)

	switch{
	case length >= 7:
		fmt.Println("nama disetujui")

	case length < 7:
		fmt.Println("nama tidak disetujui")
	}

}
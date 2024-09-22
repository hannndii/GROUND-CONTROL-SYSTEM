package main

import "fmt"

func hurufVokal(kalimat string) int{
	jumlah := 0

	for _, huruf := range kalimat{
		switch huruf{
		case 'a', 'i', 'u', 'e', 'o':
			jumlah++
		}
	}
	return jumlah
}

func main(){
	var kalimat string

	fmt.Scan(&kalimat)

	result := hurufVokal(kalimat)
	fmt.Printf("%d", result)
}
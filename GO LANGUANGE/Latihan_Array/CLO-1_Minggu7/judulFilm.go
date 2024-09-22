package main

import "fmt"

type Film struct {
	Judul  string
	Rating float64
}

type tabFilm [10]Film

func main() {
	var listFilm tabFilm
	var input Film
	var count int

	for {
		fmt.Println("Masukkan judul film: ")
		fmt.Print("Film: ")
		fmt.Scan(&input.Judul)
		if input.Judul == "#" || count == len(listFilm) {
			break
		}

		fmt.Println("Masukkan rating film: ")
		fmt.Scan(&input.Rating)
		if input.Rating > 4.00 && input.Rating <= 5.00 {
			listFilm[count] = input
			count++
		}
	}

	for i := 0; i < count; i++ {
		fmt.Printf("%s ", listFilm[i].Judul)
	}
	fmt.Println()
}

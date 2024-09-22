package main

import ("fmt")

func main(){
	var hariKerja, genreAksi bool
	var jamNonton int

	fmt.Scan(&hariKerja, &jamNonton, &genreAksi)

	fmt.Println(mauNonton(hariKerja, jamNonton, genreAksi))
}


func mauNonton(x bool, y int, z bool) bool{
	var hasil bool
	if x == false && y >= 19 && z == true{
		hasil = true
	} else {
		hasil = false
	} 
	return hasil

}
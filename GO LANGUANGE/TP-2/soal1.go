

package main

import (
	"fmt"
	"strings"
)

func lowToUpper(kar byte) byte {
	// Mengecek apakah karakter berada dalam range huruf lowercase
	if kar >= 'a' && kar <= 'z' {
		// Mengonversi huruf lowercase menjadi huruf uppercase menggunakan fungsi strings.ToUpper()
		return strings.ToUpper(string(kar))[0]
	}
	// Jika karakter tidak berupa huruf lowercase, mengembalikan karakter aslinya
	return kar
}

func main() {
	var karakter byte
	fmt.Print("Masukkan karakter alfabet lowercase ['a'-'z']: ")
	fmt.Scanf("%c", &karakter)

	// Memanggil fungsi lowToUpper untuk mengonversi karakter
	hurufUppercase := lowToUpper(karakter)

	fmt.Println("Huruf uppercase:", string(hurufUppercase))
}


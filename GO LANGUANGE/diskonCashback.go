// SOAL	
// Sebuah swalayan di Bandung mengadakan promo berbentuk cashback dan diskon, yang ditentukan berdasarkan poin yang dikumpulkan oleh pembeli. 
// Cashback diperoleh, jika total 3 poin pertama sama dengan 3 poin terakhir dan memperoleh diskon apabila total poin pertama dan terakhir adalah faktor dari total 3 poin lainnya.

// Masukan terdiri dari lima bilangan bulat, yang menyatakan poin-poin yang terkumpul.
// Keluaran terdiri dari satu atau dua baris string yang menyatakan perolehan “cashback”, “diskon”., atau tidak sama sekali.

package main

import (
	"fmt"
)

func main() {
	var poin [5]int
	fmt.Println("Masukkan 5 poin yang terkumpul:")
	for i := 0; i < 5; i++ {
		fmt.Scan(&poin[i])
	}

	cashback := poin[0]+poin[1]+poin[2] == poin[2]+poin[3]+poin[4]
	diskon := (poin[1]+poin[2]+poin[3])%(poin[0]+poin[4]) == 0

	if cashback && diskon {
		fmt.Println("cashback\ndiskon")
	} else if cashback {
		fmt.Println("cashback")
	} else if diskon {
		fmt.Println("diskon")
	} else {
		fmt.Println("Tidak ada cashback atau diskon")
	}
}

// SOAL
// Buatlah program dalam bahasa Go untuk menghitung hasil perkalian bilangan bulat n dengan m menggunakan loop. 
// Masukan terdiri dari dua baris bilangan bulat n dan m.
// Keluaran berupa hasil perkalian.
// Petunjuk: Perkalian n x m = 5 x 2 dapat dinyatakan dengan: 2 + 2 + 2 + 2 + 2 = 10. Sedangkan perkalian n x m = 4 x 3 dapat dinyatakan dengan: 4 x 3 = 3 + 3 + 3 + 3 =  12

package main

import (
	"fmt"
)

func main() {
	var n, m int

	fmt.Scan(&n)
	fmt.Scan(&m)

	result := 0

	for i := 0; i < m; i++ {
		result += n
	}

	// Menampilkan hasil perkalian
	fmt.Println(result)
}

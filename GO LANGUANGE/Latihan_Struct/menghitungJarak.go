// Tentu, mari saya jelaskan setiap bagian dari kode tersebut:

// 1. **Package `main`**: Ini adalah baris pertama dalam kode Go yang menentukan bahwa file ini adalah bagian dari package `main`. Package `main` khusus digunakan untuk program yang dapat dieksekusi. 

// 2. **Import Statement**: Bagian ini mengimpor paket-paket yang diperlukan untuk program. Dalam kasus ini, kita menggunakan dua paket: `fmt` untuk masukan/keluaran dan `math` untuk operasi matematika.

// 3. **Definisi Struktur `titik`**: Dalam Go, kita dapat mendefinisikan tipe data baru dengan menggunakan `type`. 
// Di sini, kita mendefinisikan tipe data `titik` yang memiliki dua angka float64, `p1` dan `p2`, yang mewakili koordinat titik.

// 4. **Fungsi `jarak()`**: Ini adalah fungsi yang mengambil dua parameter bertipe `titik` dan mengembalikan jarak antara keduanya. 
// Fungsi ini menggunakan formula jarak Euclidean yang sederhana, yaitu akar kuadrat dari jumlah kuadrat dari selisih koordinat `x` dan `y` dari kedua titik.

// 5. **Fungsi `akar()`**: Ini adalah fungsi yang mengambil sebuah bilangan float64 dan mengembalikan akarnya menggunakan fungsi `math.Sqrt()`.

// 6. **Fungsi `main()`**: Ini adalah fungsi utama program. Di dalamnya, kita mendeklarasikan variabel-variabel yang diperlukan untuk menyimpan dua titik, yaitu `titik1` dan `titik2`. 
// Kemudian, kita meminta pengguna untuk memasukkan koordinat `x` dan `y` untuk kedua titik tersebut menggunakan `fmt.Scan()`. 
// Setelah mendapatkan koordinat untuk kedua titik, kita memanggil fungsi `jarak()` untuk menghitung jarak antara keduanya, dan hasilnya dicetak menggunakan `fmt.Println()`.

// Dengan cara ini, program mengambil dua titik sebagai masukan, menghitung jarak antara mereka menggunakan rumus jarak Euclidean, dan mencetak hasilnya.


package main

import (
	"fmt"
	"math"
)

type titik struct{
	p1 float64
	p2 float64
}

func jarak(p1, p2 titik) float64 {
	deltaX := p1.p1 - p2.p1
	deltaY := p1.p2 - p2.p2
	return akar(math.Pow(deltaX, 2) - math.Pow(deltaY, 2))
}

func akar(x float64) float64{
	return math.Sqrt(x)
}

func main(){
	var titik1, titik2 titik

	fmt.Scan(&titik1.p1, &titik2.p2)

	result := jarak(titik1, titik2)

	fmt.Println(result)
}
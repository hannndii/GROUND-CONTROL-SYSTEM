package main

import "fmt"

func main() {

	var bil, hasil int
	fmt.Scan(&bil)
	hasil = fibo(bil)
	fmt.Println(hasil)

}

func fibo(n int) int {
	var i, t1, t2, next, tot int
	t1 = 0
	t2 = 1
	tot = 0
	for i = 1; i <= n; i++ {
		if i == 1 {
			tot += t1
			continue
		}
		if i == 2 {
			tot += t2
			continue
		}
		next = t1 + t2
		t1 = t2
		t2 = next
		tot += next

	}
	return tot
}

// Program yang Anda berikan adalah dalam bahasa pemrograman Go (Golang) dan memiliki fungsi untuk menghitung total deret Fibonacci hingga n.

// Berikut adalah penjelasan untuk setiap bagian dari program:

// 1. `package main`: Ini adalah deklarasi paket yang menunjukkan bahwa file ini adalah bagian dari paket `main`. Setiap program Go dimulai dengan paket `main`.

// 2. `import "fmt"`: Ini adalah pernyataan impor yang mengimpor paket `fmt`, yang berisi fungsi-fungsi untuk masukan/keluaran format. Dalam program ini, fungsi `Scan()` dan `Println()` digunakan dari paket ini.

// 3. `func main() { ... }`: Ini adalah fungsi utama dari program. Eksekusi program dimulai dari sini.

// 4. `var bil, hasil int`: Ini mendeklarasikan dua variabel bertipe integer, `bil` dan `hasil`.

// 5. `fmt.Scan(&bil)`: Ini mengambil input dari pengguna dan menyimpannya di dalam variabel `bil`. Tanda `&` digunakan untuk mengambil alamat variabel, yang dibutuhkan oleh fungsi `Scan()`.

// 6. `hasil = fibo(bil)`: Ini memanggil fungsi `fibo()` dengan parameter `bil` dan menyimpan hasilnya di dalam variabel `hasil`.

// 7. `fmt.Println(hasil)`: Ini mencetak hasil perhitungan deret Fibonacci ke layar.

// 8. `func fibo(n int) int { ... }`: Ini adalah definisi fungsi `fibo()` yang menerima satu argumen bertipe integer (`n`) dan mengembalikan nilai bertipe integer.

// 9. `var i, t1, t2, next, tot int`: Ini mendeklarasikan beberapa variabel bertipe integer yang digunakan dalam fungsi `fibo()`.

// 10. `t1 = 0` dan `t2 = 1`: Ini menetapkan nilai awal untuk dua suku pertama dari deret Fibonacci.

// 11. `for i = 1; i <= n; i++ { ... }`: Ini adalah loop `for` yang akan berjalan dari 1 hingga `n` untuk menghitung deret Fibonacci hingga suku ke-n.

// 12. `if i == 1` dan `if i == 2`: Ini adalah kondisional yang menghandle kasus khusus untuk dua suku pertama dari deret Fibonacci.

// 13. `next = t1 + t2`, `t1 = t2`, `t2 = next`: Ini adalah perhitungan deret Fibonacci untuk suku-suku selanjutnya setelah dua suku pertama.

// 14. `tot += next`: Ini menambahkan nilai `next` ke dalam total.

// 15. `return tot`: Ini mengembalikan total dari deret Fibonacci.
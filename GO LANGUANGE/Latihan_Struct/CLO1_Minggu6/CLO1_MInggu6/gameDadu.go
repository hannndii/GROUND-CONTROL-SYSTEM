package main

import "fmt"

type SkorGame struct {
    TotalA   int
    TotalB   int
    TotalC   int
    max      int
    Pemenang byte
}

var skor SkorGame

func isGenap(n int) bool {
    return n%2 == 0
}

func maxSkor() {
    skor.max = skor.TotalA
    skor.Pemenang = 'A'

    if skor.TotalB > skor.max {
        skor.max = skor.TotalB
        skor.Pemenang = 'B'
    }
    if skor.TotalC > skor.max {
        skor.max = skor.TotalC
        skor.Pemenang = 'C'
    }
}

func gameDadu(n int, rolls [][]int) (byte, int) {
    for i := 0; i < n; i++ {
        d1, d2, d3 := rolls[i][0], rolls[i][1], rolls[i][2]

        if isGenap(d1) {
            skor.TotalA += d1
        }
        if isGenap(d2) {
            skor.TotalB += d2
        }
        if isGenap(d3) {
            skor.TotalC += d3
        }
    }
    maxSkor()
    return skor.Pemenang, skor.max
}

func main() {
    var n int
    fmt.Scanln(&n)

    rolls := make([][]int, n)
    for i := 0; i < n; i++ {
        rolls[i] = make([]int, 3)
        fmt.Scanln(&rolls[i][0], &rolls[i][1], &rolls[i][2])
    }

    pemenang, total := gameDadu(n, rolls)
    fmt.Printf("%c %d\n", pemenang, total)
}


// Tentu! Berikut adalah penjelasan singkat tentang kode yang telah disediakan:

// 1. Pertama-tama, kita mendefinisikan sebuah struktur bernama `SkorGame` yang akan digunakan untuk menyimpan total skor dari masing-masing asisten dosen (`TotalA`, `TotalB`, `TotalC`), skor maksimum (`max`), dan pemenang (`Pemenang`).

// 2. Kemudian, kita mendefinisikan sebuah fungsi `isGenap` yang menerima sebuah bilangan bulat dan mengembalikan `true` jika bilangan tersebut adalah angka genap, dan `false` jika tidak.

// 3. Fungsi `maxSkor` akan digunakan untuk menentukan pemenang dari game ini berdasarkan total skor yang diperoleh oleh setiap asisten dosen. Fungsi ini memperbarui nilai `max` dan `Pemenang` dalam struktur `SkorGame`.

// 4. Fungsi `gameDadu` adalah inti dari program ini. Fungsi ini menerima input berupa jumlah babak (`n`) dan hasil lemparan dadu untuk setiap babak dalam bentuk array dua dimensi (`rolls`). 
// Di dalam fungsi ini, kita iterasi melalui setiap babak, memeriksa apakah angka yang dihasilkan dari lemparan dadu adalah genap, dan jika ya, menambahkan angka tersebut ke total skor masing-masing asisten dosen.

// 5. Setelah semua babak selesai, fungsi `gameDadu` memanggil `maxSkor` untuk menentukan pemenang dari permainan ini.

// 6. Di dalam fungsi `main`, kita membaca input jumlah babak (`n`) dari pengguna dan menyimpan hasil lemparan dadu untuk setiap babak dalam array `rolls`.

// 7. Kemudian, kita memanggil fungsi `gameDadu` dengan parameter input yang telah kita dapatkan dari pengguna, dan menampilkan pemenang dan total skornya ke layar.

// Semoga penjelasan ini membantu Anda memahami implementasi dari kode tersebut! Jika Anda memiliki pertanyaan lebih lanjut, jangan ragu untuk bertanya.
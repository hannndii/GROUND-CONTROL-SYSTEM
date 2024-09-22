package main

import "fmt"

type mahasiswa struct {
	nama  string
	nim   string
	nilai float32
}

func main() {
	var mhs1, mhs2, mhs3 mahasiswa
	
	baca(&mhs1, &mhs2, &mhs3)
	cetak(mhs1, mhs2, mhs3)
	rata := rata_rata_nilai(mhs1, mhs2, mhs3)
	
	fmt.Printf("Rata-rata nilai: %2.6f\n", rata)
	mhs_max_nilai(mhs1, mhs2, mhs3)
}

func baca(m1, m2, m3 *mahasiswa) {
	fmt.Scanln(&m1.nama, &m1.nim, &m1.nilai)
	fmt.Scanln(&m2.nama, &m2.nim, &m2.nilai)
	fmt.Scanln(&m3.nama, &m3.nim, &m3.nilai)
}

func cetak(m1, m2, m3 mahasiswa) {
	fmt.Println("Data mahasiswa:")
	fmt.Println(m1.nama, m1.nim, m1.nilai)
	fmt.Println(m2.nama, m2.nim, m2.nilai)
	fmt.Println(m3.nama, m3.nim, m3.nilai)
}

func rata_rata_nilai(m1, m2, m3 mahasiswa) float32 {
	total := float32(m1.nilai + m2.nilai + m3.nilai)
	return total / 3
}

func mhs_max_nilai(m1, m2, m3 mahasiswa) {
	max := m1
	if m2.nilai > max.nilai {
		max = m2
	}
	if m3.nilai > max.nilai {
		max = m3
	}
	fmt.Printf("Mahasiswa dengan nilai tertinggi: %s %g\n", max.nama, max.nilai)
}

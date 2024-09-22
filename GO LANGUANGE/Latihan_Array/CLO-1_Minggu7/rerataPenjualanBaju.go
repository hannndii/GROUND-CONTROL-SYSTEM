package main

import "fmt"

const NMax int = 1000
type Baju [NMax]int

func inputB(T, P *Baju, jHari *int) {
	var terjual, pendapatan int
	fmt.Scan(&terjual)
	fmt.Scan(&pendapatan)
	*jHari = 0
	for terjual != 0 && pendapatan != 0 {
		T[*jHari] = terjual
		P[*jHari] = pendapatan
		*jHari = *jHari + 1
		fmt.Scan(&terjual)
		fmt.Scan(&pendapatan)
	}
}

func rata2(T, P Baju, RT, RP *float64, jHari int) {
	var i, totalT, totalP int
	totalT = 0
	totalP = 0
	for i = 0; i < jHari; i++ {
		totalT = totalT + T[i]
		totalP = totalP + P[i]
	}
	*RT = float64(totalT) / float64(jHari)
	*RP = float64(totalP) / float64(jHari)
}

func tampilkan(RT, RP float64, jHari int) {
	fmt.Printf("Rata-rata barang yang terjual selama %d hari adalah: %.2f\n", jHari, RT)
	fmt.Printf("Rata-rata pendapatan yang diperoleh selama %d hari adalah: %.2f\n", jHari, RP)
}

func main() {
	var t, p Baju
	var rata2_T, rata2_P float64
	var jHari int
	inputB(&t, &p, &jHari)
	rata2(t, p, &rata2_T, &rata2_P, jHari)
	tampilkan(rata2_T, rata2_P, jHari)
}

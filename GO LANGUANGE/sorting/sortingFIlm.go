package main

import (
	"fmt"
)

const NMAX = 5

type tDrakor struct {
	Judul   string
	Rating  float64
	JumEp   int
	Durasi  int
}

type tabDrakor [NMAX]tDrakor

func main() {
	var drakor tabDrakor
	var nDrakor int

	bacaDataDrakor(&drakor, &nDrakor)
	fmt.Println("\nData sebelum diurutkan:")
	cetakDataDrakor(drakor, nDrakor)

	urutMenurun(&drakor, nDrakor)
	fmt.Println("\nData setelah diurutkan menurun berdasarkan rating:")
	cetakDataDrakor(drakor, nDrakor)

	urutMenaik(&drakor, nDrakor)
	fmt.Println("\nData setelah diurutkan menaik berdasarkan durasi:")
	cetakDataDrakor(drakor, nDrakor)
}

func bacaDataDrakor(A *tabDrakor, n *int) {
	fmt.Scan(n)
	for i := 0; i < *n; i++ {
		fmt.Scan(&A[i].Judul, &A[i].Rating, &A[i].JumEp, &A[i].Durasi)
	}
}

func cetakDataDrakor(A tabDrakor, n int) {
	fmt.Printf("%20s %6s %6s %6s\n", "Judul", "Rating", "Jum Ep", "Durasi")
	for i := 0; i < n; i++ {
		fmt.Printf("%20s %6.1f %6d %6d\n", A[i].Judul, A[i].Rating, A[i].JumEp, A[i].Durasi)
	}
}

func urutMenurun(A *tabDrakor, n int) {
	for i := 0; i < n-1; i++ {
		maxIdx := i
		for j := i + 1; j < n; j++ {
			if A[j].Rating > A[maxIdx].Rating {
				maxIdx = j
			}
		}
		A[i], A[maxIdx] = A[maxIdx], A[i]
	}
}

func urutMenaik(A *tabDrakor, n int) {
	for i := 1; i < n; i++ {
		key := A[i]
		j := i - 1
		for j >= 0 && A[j].Durasi > key.Durasi {
			A[j+1] = A[j]
			j--
		}
		A[j+1] = key
	}
}
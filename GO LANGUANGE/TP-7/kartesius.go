package main

import "fmt"

type titik struct{
	koordinatX, koordinatY float64
	warna string
}

func main(){
	var x1, x2, y1, y2 float64
	var w1, w2 string
	var t1, t2 titik

	fmt.Scan(&x1, &y1, &w1)
	fmt.Scan(&x2, &y2, &w2)

	t1 = pembuatan_titik_baru(x1, y1, w1)
	t2 = pembuatan_titik_baru(x2, y2, w2)

	fmt.Printf("Data titik 1: Koordinat (%v, %v), warna %s\n", t1.koordinatX, t1.koordinatY, t1.warna)
	fmt.Printf("Data titik 2: Koordinat (%v, %v), warna %s\n", t2.koordinatX, t2.koordinatY, t2.warna)
}

func pembuatan_titik_baru(x, y float64, w string) titik{
	return titik{x, y, w}
}

// package main

// import "fmt"

// type titik struct{
// 	koordinatX, koordinatY float64
// 	warna string
// }

// func main(){
// 	var t1, t2 titik

// 	fmt.Scan(&t1.koordinatX, &t1.koordinatY, &t1.warna)
// 	fmt.Scan(&t2.koordinatX, &t2.koordinatY, &t2.warna)
	

// 	t1 = pembuatan_titik_baru(t1.koordinatX, t1.koordinatX, t1.warna)
// 	t2 = pembuatan_titik_baru(t2.koordinatX, t2.koordinatY, t2.warna)

// 	fmt.Printf("Data titik 1: Koordinat (%v, %v), warna %s\n", t1.koordinatX, t1.koordinatY, t1.warna)
// 	fmt.Printf("Data titik 2: Koordinat (%v, %v), warna %s\n", t2.koordinatX, t2.koordinatY, t2.warna)
// }

// func pembuatan_titik_baru(x, y float64, w string) titik{
// 	return titik{x, y, w}
// }
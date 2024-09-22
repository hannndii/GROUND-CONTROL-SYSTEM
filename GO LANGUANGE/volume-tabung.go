package main

import ("fmt")

func main(){
	var jariJari, tinggi, luas float32
	fmt.Scan(&jariJari, &tinggi)

	luas = 2 * luasLingkaran(jariJari) + luasSelimut(jariJari, tinggi)

	fmt.Println(luas)

}
func luasLingkaran(r float32) float32 {
	return 3.14 * (r * r)
}
func luasSelimut(r, t float32) float32 {
	return 3.14 * (r * t)
}
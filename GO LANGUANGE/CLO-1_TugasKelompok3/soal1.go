package main

import (
	"fmt"
)

func main(){
	var uang, k10, k2, k3 int

	fmt.Scan(&uang)
	hitungUang(uang, &k10, &k2, &k3)
	fmt.Println(k10, "lembar 10000")
	fmt.Println(k2, "lembar 2000")
	fmt.Println(k3, "lembar 1000")
	

}

func hitungUang(uang int, k10, k2, k3 *int){
	*k10 = uang / 10000
	*k2 = uang % 10000 / 2000
 	*k3 = uang % 2000 / 1000
}

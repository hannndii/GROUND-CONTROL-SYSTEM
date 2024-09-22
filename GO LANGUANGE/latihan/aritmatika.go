package main

import "fmt"

func aritmatika(x, y, z int){
	if y - x == z - y {
		fmt.Println("ya")
	} else {
		fmt.Println("tidak")
	}
}

func main(){
	var x, y, z int

	fmt.Scan(&x, &y, &z)

	aritmatika(x, y, z)

}
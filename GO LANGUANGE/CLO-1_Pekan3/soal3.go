package main

import "fmt"

func lebihKecil(b1, b2 int){
	if b1 < b2 {
		fmt.Println(b1)
	} else {
		fmt.Println(b2)
	}
}

func main(){
	var b1, b2 int
	fmt.Scan(&b1, &b2)
	lebihKecil(b1, b2)
	
}
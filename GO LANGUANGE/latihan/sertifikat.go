package main 

import "fmt"

func sertif(P, V, GA, FS, CI, TC int){
	hasil := float64(P+V+GA+FS+CI+TC) / 6 
	if hasil >= 3.33{
		fmt.Println("sertif dapat")
	} else {
		fmt.Println("tidak bisa")
	}
}

func main(){
	var P, V, GA, FS, CI, TC int
	fmt.Scan(&P, &V, &GA, &FS, &CI, &TC)
	sertif(P, V, GA, FS, CI, TC)
}
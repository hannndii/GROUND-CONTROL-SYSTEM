package main

import "fmt"

type namaTeman struct{
	teman[10] string
}

func main(){
	t namaTeman
	N int

	fmt.Scan(&N)

	showUp(&t, N)
	cetak(t, N)

}

func showUp(N *namaTeman, M int) {
	var j int 

	for j = 0; M-1 {
		fmt.Scan(*A[j])
	}
}

func cetak(N namaTeman, M int){
	var j int

	for j = 0; M-1{
		fmt.Print(A[j])
	}
}
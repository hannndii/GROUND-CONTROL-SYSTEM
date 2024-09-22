package main

import "fmt"

const NMAX = 10

type tabArr[NMAX] int

func main(){
	var arrT tabArr
	var nData int

	fmt.Scanln(&nData)
	input(&arrT, nData)
	selectionSort(&arrT, nData)
	cetak(arrT, nData)
}

func input(T *tabArr, n int){
	for i := 0; i < n; i++{
		fmt.Scan(&(*T)[i])
	}
}

func cetak(T tabArr, n int){
	for i := 0; i < n; i++{
		fmt.Print(T[i], " ")
	}
}

func selectionSort(T *tabArr, n int){
	var pass, i, idx, temp int

	pass = 1
	for pass < n{
		idx = pass-1
		i = pass
		for i < n{
			if (*T)[idx] > (*T)[i]{
				idx = i
			}
			i = i + 1
		}
		temp = (*T)[pass-1]
		(*T)[pass-1] = (*T)[idx]
		(*T)[idx] = temp
		pass = pass + 1
	}
}
package main

import "fmt"

const NMAX = 1024

type data struct{
	nama string
	nilai int
}

type tabData[NMAX] data

func inputData(T *tabData, n int){
	for i := 0; i < n; i++{
		fmt.Scanln(&(*T)[i].nama, &(*T)[i].nilai)
	}
}

func cetakData(T tabData, n int, x string){ 
	for i := 0; i < n; i++{
		if T[i].nilai > 80 && x == "A"{
			fmt.Println(T[i].nama, T[i].nilai, "A")
		} else if T[i].nilai <= 80 && T[i].nilai > 70 && x == "AB" {
			fmt.Println(T[i].nama, T[i].nilai, "AB")
		} else if T[i].nilai <= 70 && T[i].nilai > 65 && x == "B"{
			fmt.Println(T[i].nama, T[i].nilai, "B")
		} else if T[i].nilai <= 65 && T[i].nilai > 60 && x == "BC"{
			fmt.Println(T[i].nama, T[i].nilai, "BC")
		} else if T[i].nilai <= 60 && T[i].nilai > 50 && x == "C"{
			fmt.Println(T[i].nama, T[i].nilai, "C")
		} else if T[i].nilai <= 50 && T[i].nilai > 40 && x == "D"{
			fmt.Println(T[i].nama, T[i].nilai, "D")
		} else if T[i].nilai <= 40 && x == "E"{
			fmt.Println(T[i].nama, T[i].nilai, "E")
		}
		
	}
}

func urutData(T *tabData, n int){
	var pass, i, idx_min int
	var temp data

	pass = 1
	for pass < n{
		idx_min = pass-1
		i = pass
		for i < n{
			if T[idx_min].nilai > T[i].nilai{
				idx_min = i
			}
			i++
		}
		temp = T[pass-1]
		T[pass-1] = T[idx_min]
		T[idx_min] = temp
		pass++
	}
}

func main(){
	var arrData tabData
	var nData int
	var target string

	fmt.Scanln(&nData)
	fmt.Scanln(&target)
	inputData(&arrData, nData)
	urutData(&arrData, nData)
	cetakData(arrData, nData, target)
}
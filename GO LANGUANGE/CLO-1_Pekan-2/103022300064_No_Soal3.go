package main

import ("fmt")

func main(){
	var a ,b ,c float64
    fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Scan(&c)
	
	result := hitungD(b ,a ,c)
	fmt.Println(cekTipot(b, a, c), result)
}

func hitungD(b, a, c float64) float64{
	return b*b - 4*a*c
}

func cekTipot(b, a, c  float64) string{
	D := hitungD(b, a, c)
    if D >= 0{
	    return "memiliki titik potong dengan sumbu-x"
	} else {
		return "tidak memiliki titik potong dengan sumbu-x"
	}
	
}
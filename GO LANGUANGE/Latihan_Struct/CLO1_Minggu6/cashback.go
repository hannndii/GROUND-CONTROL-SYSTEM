package main

import "fmt"

type dataPelanggan struct {
	member bool
	totalBelanja int

}

func main() {
   var pembeli dataPelanggan

   fmt.Scanf("%t %d", &pembeli.member, &pembeli.totalBelanja)

   result := hitungBiaya(pembeli.member, pembeli.totalBelanja)
   fmt.Println(result)
}
func hitungBiaya(member bool, totalBelanja int) int{
	var total int
    if totalBelanja >= 400000 && member == true {
        total = (totalBelanja - int(float64(totalBelanja)*0.2)) - 50000
    } else if totalBelanja >= 300000 {
        total = totalBelanja - int(float64(totalBelanja)*0.2)
    } else if totalBelanja >= 200000 && member == true{
        total = (totalBelanja - int(float64(totalBelanja)*0.1)) - 25000
    } else if totalBelanja >= 100000 {
        total = totalBelanja - int(float64(totalBelanja)*0.1)
    } else {
        total = totalBelanja
    }
    return total
}
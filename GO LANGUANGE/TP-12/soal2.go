package main
import "fmt"
func main(){
	var search, inputVar int
	var arr, arrSorted []int

	fmt.Scanln(&search)
	for {
		fmt.Scan(&inputVar)
		arr = append(arr, inputVar)

		if len(arr) > 2{
			if arr[len(arr)-1] < arr[len(arr)-2] || len(arr)>10 {
				break
			}
		}
	}

	arrSorted = arr[0:len(arr)-1]
	// printData
	fmt.Print("Data Bilangan: ")
	for i:=0; i<len(arrSorted); i++{
		fmt.Printf("%v ", arrSorted[i])
	}

	// positionSearch
	position:=-1
	for i:=0; i<len(arrSorted); i++{
		if arrSorted[i] == search {
			position = i
			break
		}
	}

	

	fmt.Println()
	if position==-1 {
		fmt.Printf("Hasil pencarian: Bilangan tidak ditemukan.")
		
	} else {
		fmt.Printf("Hasil pencarian: Bilangan %v ditemukan pada posisi ke-%v", search, position+1)
	}
	
	


}
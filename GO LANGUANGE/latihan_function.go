// package main

// import (
// 	"fmt"
// 	"strings"
// )

// func main(){
// 	var karakter byte

// 	fmt.Scanf("%c", &karakter)
// 	UpperCase := hurufUppercase(karakter)

// 	fmt.Println(string(UpperCase))
// }

// func hurufUppercase(kar  byte) byte{
// 	if kar >= 'a' && kar <= 'z'{
// 		return strings.ToUpper(string(kar))[0]
// 	}
// 	return kar

// }

package main

import (
	"fmt"
	"strings"
)

func hitungKar(sentence string){
	word := strings.Fields(sentence)
	jumlahKata := len(word)
	fmt.Println(jumlahKata)
}

func main(){
	var sentence string

	fmt.Scanf("%s", &sentence)
	hitungKar(sentence)

}
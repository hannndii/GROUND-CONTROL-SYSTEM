package main

import (
	"fmt"
)

func cariGo(str string) {
	count := 0
	length := len(str)

	for i := 0; i < length-1; i++ {
		if str[i] == 'g' && str[i+1] == 'o' {
			count++
		}
	}
	fmt.Println(count)
}

func main() {
	var input string
	fmt.Scanln(&input)

	cariGo(input)
}

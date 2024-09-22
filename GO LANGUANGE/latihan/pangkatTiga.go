package main

import "fmt"

func pangkatTiga(n int) int{
	hasil := n * n * n
	return hasil
}

func main(){
	pangkatTiga(-3)
	fmt.Println(pangkatTiga(-3))
}
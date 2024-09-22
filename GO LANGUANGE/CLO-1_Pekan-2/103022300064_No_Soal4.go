package main

import ("fmt")

func main (){
	var s1, s2 string
	fmt.Scan(&s1, &s2)
	fmt.Println(penggabunganString(s1, s2))
}

func penggabunganString(s1, s2 string) string{
	hasil := s1 + s2
	return hasil
}
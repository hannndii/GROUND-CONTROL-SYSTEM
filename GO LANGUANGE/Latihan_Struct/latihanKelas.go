package main

import "fmt"

type Name struct {
	firstName  string
	middleName string
	lastName   string
}

type Address struct {
	street string
	city   string
	zip    string
}

type Student struct {
	fullName Name
	address  Address
}

func main() {
	var siswa Student

	siswa.fullName.firstName = "Muhammad"
	siswa.fullName.middleName = "Endihan"
	siswa.address.city = "Bandung"
	siswa.address.street = "Sukabirus"
	fmt.Println(siswa.fullName.firstName)
	fmt.Println(siswa.fullName.middleName)
	fmt.Println(siswa.address.city)
	fmt.Println(siswa.address.street)
}
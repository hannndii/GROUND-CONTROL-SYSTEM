package main

import "fmt"

type Name struct {
	firstName  string
	middleName string
	lastName   string
}

type Address struct {
	nameStreet string
	numberHome int
	city   string
	Country    string
}

type Student struct {
	fullName Name
	address  Address
}

func main() {
	var siswa Student

	siswa.fullName.firstName = "Muhammad"
	siswa.fullName.middleName = "Endihan"
	siswa.fullName.lastName = "Nasution"
	siswa.address.numberHome = 26
	siswa.address.nameStreet = "Tupolev"
	siswa.address.city = "Jakarta Timur"
	siswa.address.Country = "Indonesia"

	fmt.Printf("Nama lengkap: %s %s %s \n", siswa.fullName.firstName, siswa.fullName.middleName, siswa.fullName.lastName)
	fmt.Printf("Alamat: Jalan %s, No. %d, %s, %s", siswa.address.nameStreet, siswa.address.numberHome, siswa.address.city, siswa.address.Country)
}
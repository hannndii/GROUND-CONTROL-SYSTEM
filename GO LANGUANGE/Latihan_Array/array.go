package main
import "fmt"

type Mahasiswa struct{
    Nama string
    Umur int
    Nilai [3] int
}

func main() {
  var numStudents int
  fmt.Println("Masukkan jumlah student: ")
  fmt.Scanln(&numStudents)

  students := make([]Mahasiswa, numStudents)
  
  for i := 0; i < len(students); i++{
	fmt.Printf("\n Input data untuk student %d\n: ", i+1)
	fmt.Println("Masukkan nama:")
	fmt.Scanln(&students[i].Nama)
	fmt.Println("Masukkan umur:")
	fmt.Scanln(&students[i].Umur)
	fmt.Println("Masukkan nilai(sebanyak 3 buah):")
	for j := 0; j < 3; i++ {
		fmt.Printf("Mata Pelajaran %d: ", j+1)
		fmt.Scanln(&students[i].Nilai[j])
	}
  }
  for i := 0; i < len(students); i++ {
	fmt.Println("\nDaftar isi mahasiswa %d: \n", i+1)
    fmt.Println("Nama:", students[i].Nama)
    fmt.Println("Umur:", students[i].Umur)
    for j := 0; j < len(students); i++{
        fmt.Printf("Nilai kamu %d: %d\n", i+1, students.Nilai[i])
  }
  }
  
}
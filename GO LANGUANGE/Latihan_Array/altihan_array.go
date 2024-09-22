package main
import "fmt"

type Mahasiswa struct{
    Nama string
    Umur int
    Nilai [3] int
}

func inputData(student[]Mahasiswa) {
    for i := 0; i < len(student); i++{
		fmt.Printf("\n Input data untuk student %d\n: ", i+1)
		fmt.Println("Masukkan nama:")
		fmt.Scanln(&student[i].Nama)
		fmt.Println("Masukkan umur:")
		fmt.Scanln(&student[i].Umur)
		fmt.Println("Masukkan nilai(sebanyak 3 buah):")
		for j := 0; j < len(student[i].Nilai); j++ {
			fmt.Printf("Nilai mata kuliah ke- %d: ", j+1)
			fmt.Scanln(&student[i].Nilai[j])
		}
	}
}

func outputData(student[]Mahasiswa) {
	for i := 0; i < len(student); i++ {
		fmt.Printf("\nDaftar isi mahasiswa %d: ", i+1)
		fmt.Println("Nama:", student[i].Nama)
		fmt.Println("Umur:", student[i].Umur)
		for j := 0; j < len(student[i].Nilai); j++{
			fmt.Printf("Nilai kamu %d: %d\n", j+1, student[i].Nilai[j])
	  }
	}
}

func main() {
  var jumlahStudent int

  fmt.Println("Masukkan jumlah student: ")
  fmt.Scanln(&jumlahStudent)


  student := make([]Mahasiswa, jumlahStudent)

  inputData(student)
  outputData(student)
  
}
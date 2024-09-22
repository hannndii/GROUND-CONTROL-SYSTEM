package main
import "fmt"

// Definisikan tipe bentukan untuk review series
type dataTubuh struct {
   tinggi float64
	berat float64
	bmi float64
}

func main() {
   var people dataTubuh

   fmt.Scanf("%f %f", &people.tinggi, &people.berat)

   hitungBMI := people.berat / (people.tinggi * people.tinggi)
   var kategori string
   if hitungBMI < 18.5{
	kategori = "Underweight"
   } else if hitungBMI >= 18.5 && hitungBMI <= 24.9 {
	kategori = "Normal Weight"
   } else if hitungBMI >= 25 && hitungBMI <= 29.9{
	kategori = "Overweight"
   } else {
	kategori = "Obesity"
   }

   fmt.Printf("BMI: %2.2f \n", hitungBMI)
   fmt.Println("Kategori: ", kategori)
}

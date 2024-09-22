package main

import "fmt"

func reamur(C float64) float64 {
    return 0.8 * C
}

func fahrenheit(C float64) float64 {
    return (C * 9 / 5) + 32
}

func kelvin(C float64) float64 {
    return C + 273.15
}

func main() {
    var startC, endC, step float64

    fmt.Println("Masukkan temperatur awal (Celsius):")
    fmt.Scanln(&startC)
    fmt.Println("Masukkan temperatur akhir (Celsius):")
    fmt.Scanln(&endC)
    fmt.Println("Masukkan langkah (step):")
    fmt.Scanln(&step)

    fmt.Printf("%10s %10s %10s %10s \n", "C", "R", "F", "K")
    for C := startC; C <= endC; C += step {
        R := reamur(C)
        F := fahrenheit(C)
        K := kelvin(C)
        fmt.Printf("%10.2f %10.2f %10.2f %10.2f \n", C, R, F, K)
    }
}

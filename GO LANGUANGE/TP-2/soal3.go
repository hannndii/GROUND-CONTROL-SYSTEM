package main

import (
    "fmt"
    "math"
)

func panjang(R, T float64) float64 {
    return R * math.Cos(T * math.Pi / 180)
}

func panjangy(R, T float64) float64 {
    return R * math.Sin(T * math.Pi / 180)
}

func main() {
    var R, T float64
    fmt.Println("Masukkan panjang (R):")
    fmt.Scanln(&R)
    fmt.Println("Masukkan sudut (T) dalam derajat:")
    fmt.Scanln(&T)

    x := panjang(R, T)
    y := panjangy(R, T)

    fmt.Printf("Panjang R pada sumbu x: %.2f\n", x)
    fmt.Printf("Panjang R pada sumbu y: %.2f\n", y)
}

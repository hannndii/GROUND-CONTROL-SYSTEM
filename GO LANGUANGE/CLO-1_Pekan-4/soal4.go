package main

import "fmt"

func main() {
    var N int
    fmt.Scan(&N)
    printBinary(N)
    fmt.Println()
}

func printBinary(N int) {
    if N > 1 {
        printBinary(N / 2)
    }
    fmt.Print(N % 2)
}

package main

import "fmt"

func main() {
    var n int
    fmt.Scan(&n)
    fmt.Println(jumlahtiapdigit(n))
}

func jumlahtiapdigit(N int) int {
    if N < 10 {
        return N
    }
    return N%10 + jumlahtiapdigit(N/10)
}

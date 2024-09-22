package main

import "fmt"

func main() {
    var a, b, n int
    fmt.Scan(&a, &b, &n)
    fmt.Println(UnAritmatika(a, b, n))
}

func UnAritmatika(a, b, n int) int {
    if n == 1 {
        return a
    } else {
    return UnAritmatika(a, b, n-1) + b
  }
}

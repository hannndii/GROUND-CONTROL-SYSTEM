package main

import "fmt"

func main() {
    var n int
    fmt.Scan(&n)
    fmt.Println(kali(n))
}

func kali(n int) int {
    if n == 1 {
        return 1
    } else {
    return n * kali(n-1)
  }
}

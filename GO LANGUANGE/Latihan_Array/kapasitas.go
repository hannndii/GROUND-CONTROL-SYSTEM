package main

import (
    "fmt"
)
const maxSize = 256

func isiArray(arr [maxSize]int, n int) {
    for i := 0; i < n; i++ {
        (arr)[i] = i + 1
    }
}

func balikArray(arr [maxSize]int, n int) {
    for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
        arr[i], arr[j] = arr[j], arr[i]
    }
}

func palindrome(arr[maxSize]int, n int) bool {
    for i := 0; i < n/2; i++ {
        if arr[i] != arr[n-1-i] {
            return false
        }
    }
    return true
}

func main() {
    var array [maxSize]int
    var n int
    fmt.Print("Masukkan  array: ")
    fmt.Scanln(&n)

    if n > maxSize {
        fmt.Println("melebihi kapasitas maksimal.")
        return
    }

    isiArray(&array, n)

    fmt.Println("sebelum dibalik:")
    fmt.Println(array[:n])


    balikArray(&array, n)

    fmt.Println("setelah dibalik:")
    fmt.Println(array[:n])


    if palindrome(&array, n) {
        fmt.Println("palindrom.")
    } else {
        fmt.Println("bukan palindrom.")
    }
}
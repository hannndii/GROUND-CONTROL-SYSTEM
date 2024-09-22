// package main

// import "fmt"

// func cetakPolaBintang(n int) {
//     if n%2 == 0 || n < 1 {
//         fmt.Println("Masukkan harus bilangan asli positif ganjil.")
//         return
//     }

//     for i := 0; i < n; i++ {
//         for j := 0; j < n; j++ {
//             if j == n-1-i {
//                 fmt.Print(" ")
//             } else {
//                 fmt.Print("*")
//             }
//         }
//         fmt.Println()
//     }
// }

// func main() {
//     var input int
//     fmt.Scan(&input)
//     cetakPolaBintang(input)
// }

package main

import "fmt"

func cetakPolaBintang(n int) {
    if n%2 == 0 || n < 1 {
        fmt.Println("Masukkan harus bilangan asli positif ganjil.")
        return
    }

    for i := 0; i < n; i++ {
        for j := 0; j < n; j++ {
            if j == i {
                fmt.Print(" ")
            } else {
                fmt.Print("*")
            }
        }
        fmt.Println()
    }
}

func main() {
    var input int
    fmt.Scan(&input)
    cetakPolaBintang(input)
}

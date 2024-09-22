// package main

// import "fmt"

// func main() {
//     var N int
//     fmt.Scan(&N)
//     result := calculateSequence(N)
//     fmt.Println(result)
// }

// func calculateSequence(N int) int {
//     if N == 1 || N == 2 || N == 3 {
//         return 1
//     } else if N == 4 {
//         return 3
//     } else if N == 5 {
//         return 5
//     } else {
//         return calculateSequence(N-1) + calculateSequence(N-2) + calculateSequence(N-3)
//     }
// }


package main

import "fmt"

func calculateSequence(N int) int{
    if N == 1 || N == 2 || N == 3{
        return 1
    } else if N == 4{
        return 3
    } else if N == 5{
        return 5
    } else {
        return calculateSequence(N-1) + calculateSequence(N-2) + calculateSequence(N-3)
    }
}

func main(){
    var N int

    fmt.Scan(&N)

    result := calculateSequence(N)
    fmt.Println(result)
}
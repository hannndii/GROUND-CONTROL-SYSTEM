

package main

import (
    "fmt"
)

const NMAX = 20 // Maximum number of characters

type tabChar [NMAX]byte // Array of characters (maximum 20 elements)

func main() {
    // Declare an array of characters (k)
    var kata tabChar

    // Get the number of characters (n) from the user
    var nChar int
    fmt.Print("Enter the number of characters: ")
    fmt.Scanln(&nChar)

    // Check if the number of characters is within the maximum limit
    if nChar > NMAX {
        fmt.Println("Error: Number of characters cannot exceed", NMAX)
        return
    }

    // Read the characters from the user
    baca(&kata, &nChar) // Call the baca function to read characters

    // Print the characters in reverse order
    cetak(kata, nChar) // Call the cetak function to print characters in reverse order
}

func baca(k *tabChar, n *int) {
    /*
        IS: Array karakter (k) dan banyak elemen (n) terdefinisi sembarang
        FS: Array karakter (k) dan banyak elemen (n) terdefinisi
     */
    for i := 0; i < *n; i++ {
        fmt.Scanf("%c", &k[*i]) // Read a character into the array
    }
}

func cetak(k tabChar, n int) {
    /*
        IS: Array k dan n terdefinisi
        FS: Array k tercetak di layar dengan urutan terbalik
     */
    for i := n - 1; i >= 0; i-- {
        fmt.Printf("%c", k[i]) // Print the character in reverse order
    }
}


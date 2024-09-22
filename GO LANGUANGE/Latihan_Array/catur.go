// package main
// import "fmt"


// func main(){
//     var chess [64]byte

//     for i:=0; i<64; i++{
//         fmt.Scanf("%c", &chess[i])
//     }

//     count:=1
//     whitePoint:=0
//     blackPoint:=0
//     for chessPiece := range chess {
//         if count%8 == 0 {
//             fmt.Printf("%c", chess[chessPiece])
//             fmt.Println()
//         } else {
//             fmt.Printf("%c", chess[chessPiece])
//         }
//         count++

//         // blackPoint counter
//         if chess[chessPiece] == 'P' {
//             blackPoint++
//         } else if chess[chessPiece] == 'R' {
//             blackPoint+=5
//         } else if chess[chessPiece] == 'N' {
//             blackPoint+=3
//         } else if chess[chessPiece] == 'B' {
//             blackPoint+=3
//         } else if chess[chessPiece] == 'Q' {
//             blackPoint+=9
//         } else if chess[chessPiece] == 'K' {
//             blackPoint+=10
//         }

//         // whitePoint counter
//         if chess[chessPiece] == 'p' {
//             whitePoint++
//         } else if chess[chessPiece] == 'r' {
//             whitePoint+=5
//         } else if chess[chessPiece] == 'n' {
//             whitePoint+=3
//         } else if chess[chessPiece] == 'b' {
//             whitePoint+=3
//         } else if chess[chessPiece] == 'q' {
//             whitePoint+=9
//         } else if chess[chessPiece] == 'k' {
//             whitePoint+=10
//         }
//     }
//     fmt.Println(whitePoint, blackPoint)
// }

package main

import "fmt"

const NMAX int = 64

type tabInt[NMAX] int

func main(){
	var T tabInt
	whiteCess := 0
	blackChess := 0
	for i := 0; i < len(T); i++{
		fmt.Scanf("%c", &T[i])
		if T[i] == 'R'{
			blackChess += 5
		} else if T[i] == 'N'{
			blackChess += 3
		} else if T[i] == 'B'{
			blackChess += 3
		} else if T[i] == 'Q'{
			blackChess += 9
		} else if T[i] == 'K'{
			blackChess += 10
		} else {
			blackChess += 0
		}

		if T[i] == 'r'{
			whiteCess += 3
		} else if T[i] == 'n'{
			whiteCess += 9
		} else if T[i] == 'b'{
			whiteCess += 10
		} else if T[i] == 'q'{
			whiteCess += 10
		} else if T[i] == 'k'{
			whiteCess += 3
		} else {
			whiteCess += 0
		}
	}
	fmt.Println(blackChess, whiteCess)
	
}
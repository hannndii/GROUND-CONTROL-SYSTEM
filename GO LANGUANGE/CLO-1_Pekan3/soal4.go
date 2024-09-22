package main

import "fmt"

func game(m1, m2, m3 *int) {
	var p1, p2, p3 int
	for {
		fmt.Scan(&p1, &p2, &p3)
		if p1 == 0 && p2 == 0 && p3 == 0 {
			break
		}

		if (p1%2 == 0 && p2%2 != 0 && p3%2 != 0) || (p1%2 != 0 && p2%2 == 0 && p3%2 == 0) {
			*m1++
		} else if (p2%2 == 0 && p1%2 != 0 && p3%2 != 0) || (p2%2 != 0 && p1%2 == 0 && p3%2 == 0) {
			*m2++
		} else if (p3%2 == 0 && p1%2 != 0 && p2%2 != 0) || (p3%2 != 0 && p1%2 == 0 && p2%2 == 0) {
			*m3++
		}
	}
}

func main() {
	var m1, m2, m3 int
	game(&m1, &m2, &m3)
	fmt.Println(m1, m2, m3)
}

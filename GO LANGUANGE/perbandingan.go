package main

import "fmt"

func main() {
	var n, kekuatan, kecepatan int
	fmt.Scanln(&n)
	fmt.Scanln(&kekuatan, &kecepatan)

	kemenangan := 0
	for i := 0; i < n; i++ {
		var kekuatanLawan, kecepatanLawan int
		fmt.Scanln(&kekuatanLawan, &kecepatanLawan)

		if kekuatan >= 3 && kecepatan >= 3 {
			if kekuatanLawan > kekuatan && kecepatanLawan > kecepatan {
				continue
			}
			if kekuatanLawan >= kekuatan {
				if kecepatanLawan < kecepatan {
					kecepatan -= 6
				}
			} else if kecepatanLawan >= kecepatan {
				kekuatan -= 6
			} else {
				if kekuatan > kecepatan {
					kecepatan -= 6
				} else {
					kekuatan -= 6
				}
			}
			kemenangan++
		}
	}

	fmt.Println(kemenangan, kekuatan, kecepatan)
}

// package main

// import (
// 	"fmt"
// 	"math"
// )

// func hitungLuasKelilingLingkaran(r float64, l, k *float64){
// 	*l = math.Pi * r * r
// 	*k = 2 * math.Pi * r
// }

// func hitungLuasKelilingPersegi(s float64, l, k *float64){
// 	*l = s * s
// 	*k = 4 * s
// }
// func hitungTotal(lL, lP, kL, kP float64, toLuas, totKel *float64){
// 	LL = hitungLuasKelilingLingkaran(r, l, k)
// 	lP = hitungLuasKelilingPersegi(s, l, k)
// 	kL = hitungLuasKelilingLingkaran(r, l, k)
// 	kP = hitungLuasKelilingPersegi(s, k, l)

// 	*toLuas = lL + lP
// 	*totKel = kL + kP

// }

// func main(){
// 	var r, s float64
// 	fmt.Scan(&r, &s)

// 	fmt.Printf("%7s %7s %7s %7s %7s %7s %7s %7s\n", "r", "s", "lL", "lP", "kL", "kP", "toLuas", "totKel" )
// 	for i := 1; i < 10; i++ {
// 		if r != 0 && s != 0{
// 			LL = hitungLuasKelilingLingkaran(r, l, k)
// 			lP = hitungLuasKelilingPersegi(s, l, k)
// 			kL = hitungLuasKelilingLingkaran(r, l, k)
// 			kP = hitungLuasKelilingPersegi(s, k, l)
// 			&toLuas = lL + lP
// 			&totKel = kL + kP
// 		} else {
// 			break
// 		}
// 	} 
// 	fmt.Printf("%2.f %2.f %2.f %2.f %2.f %2.f %2.f %2.f\n", r, s, lL, lP, kL, kP, toLuas, totKel )
	
// }


package main

import (
	"fmt"
)
const pi = 3.14

func hitungLuasKelilingLingkaran(r float64, I, k *float64) {
	*I = pi * r * r

	*k = 2 * pi * r
}

func hitungLuasKelilingPersegi(s float64, I, k *float64) {
	*I = s * s	

	*k = 4 * s
}

func hitungTotal(IL, IP, KL, KP float64, totalLuas, totalKeliling *float64) {
	*totalLuas = IL + IP

	*totalKeliling = KL + KP
}

func main() {
	var (
		r, s, luasLingkaran, luasPersegi, kelilingLingkaran, kelilingPersegi, totalLuas, totalKeliling float64
	)
     
	fmt.Scanf("%f %f", &r, &s)
	fmt.Printf("%7s %7s %7s %7s %7s %7s %7s %7s\n", "R", "S", "LL", "LP", "KL", "KP", "TL", "TP")
    
	for {
        fmt.Scanf("%f %f", &r, &s)
		if r == 0 && s == 0 {
			break
		}

		hitungLuasKelilingLingkaran(r, &luasLingkaran, &kelilingLingkaran)
		hitungLuasKelilingPersegi(s, &luasPersegi, &kelilingPersegi)
		hitungTotal(luasLingkaran, luasPersegi, kelilingLingkaran, kelilingPersegi, &totalLuas, &totalKeliling)

		fmt.Printf("%7.2f %7.2f %7.2f %7.2f %7.2f %7.2f %7.2f %7.2f\n", r, s, luasLingkaran, luasPersegi, kelilingLingkaran, kelilingPersegi, totalLuas, totalKeliling)
	}
}

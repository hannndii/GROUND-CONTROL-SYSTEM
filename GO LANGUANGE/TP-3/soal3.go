package main

import "fmt"

func main() {
    var N, menang, draw, kalah, gol, kegolan, selisihGol, poin int

    fmt.Scanln(&N)

    for i := 0; i < N; i++ {
        var golTim, kegolanTim int
        fmt.Scanln(&golTim)
        fmt.Scanln(&kegolanTim)

        hitungMenang(golTim, kegolanTim, &menang)
        hitungDraw(golTim, kegolanTim, &draw)
        hitungKalah(golTim, kegolanTim, &kalah)
        hitungJumGolKegolanSelisih(golTim, kegolanTim, &gol, &kegolan, &selisihGol)
    }

    hitungJumPoint(menang, draw, &poin)

    fmt.Println(N, menang, draw, kalah, gol, kegolan, selisihGol, poin)
}

func hitungMenang(gol, kegolan int, jm *int) {
    if gol > kegolan {
        *jm++
    }
}

func hitungDraw(gol, kegolan int, jd *int) {
    if gol == kegolan {
        *jd++
    }
}

func hitungKalah(gol, kegolan int, jk *int) {
    if gol < kegolan {
        *jk++
    }
}

func hitungJumGolKegolanSelisih(gol, kegolan int, jg, jk, jsg *int) {
    *jg += gol
    *jk += kegolan
    *jsg += gol - kegolan
}

func hitungJumPoint(menang, draw int, jp *int) {
    *jp = (menang * 3) + draw
}

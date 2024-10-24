package main

import (
	"fmt"
	"math"
)

func main() {

	var ax, ay, bx, by, cx, cy float64
	fmt.Print("Masukkan koordinat titik A (x y):")
	fmt.Scan(&ax, &ay)

	fmt.Print("Masukkan koordinat titik B (x y):")
	fmt.Scan(&bx, &by)

	fmt.Print("Masukkan koordinat titik C (x y):")
	fmt.Scan(&cx, &cy)

	AB := math.Sqrt(math.Pow(bx-ax, 2) + math.Pow(by-ay, 2))
	BC := math.Sqrt(math.Pow(cx-bx, 2) + math.Pow(cy-by, 2))
	CA := math.Sqrt(math.Pow(cx-ax, 2) + math.Pow(cy-ay, 2))

	hasil := math.Max(AB, math.Max(BC, CA))

	fmt.Print("Panjang sisi terpanjang dari segitiga: ", hasil)
}

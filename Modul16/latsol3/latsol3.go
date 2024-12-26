package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var totalDrops int
	fmt.Print("Masukkan jumlah tetesan air hujan: ")
	fmt.Scan(&totalDrops)

	countA, countB, countC, countD := 0, 0, 0, 0

	for i := 0; i < totalDrops; i++ {
		x := rand.Float64()
		y := rand.Float64()

		if x <= 0.5 && y <= 0.5 {
			countA++
		} else if x > 0.5 && y <= 0.5 {
			countB++
		} else if x <= 0.5 && y > 0.5 {
			countC++
		} else if x > 0.5 && y > 0.5 {
			countD++
		}
	}

	fmt.Printf("Curah hujan daerah A: %.4f milimeter\n", float64(countA)*0.0001)
	fmt.Printf("Curah hujan daerah B: %.4f milimeter\n", float64(countB)*0.0001)
	fmt.Printf("Curah hujan daerah C: %.4f milimeter\n", float64(countC)*0.0001)
	fmt.Printf("Curah hujan daerah D: %.4f milimeter\n", float64(countD)*0.0001)
}

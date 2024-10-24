package main

import "fmt"

func main() {
	var bb, tb, bmi float64
	fmt.Print("masukkan bb (kg): ")
	fmt.Scan(&bb)
	fmt.Print("masukkan tb (m): ")
	fmt.Scan(&tb)
	bmi = bb / (tb * tb)
	fmt.Printf("BMI anda: %.2f", bmi)

}

package main

import "fmt"

func main() {
	var bb, tb, bmi float64
	fmt.Print("masukkan bmi: ")
	fmt.Scan(&bmi)
	fmt.Print("masukkan tb (m): ")
	fmt.Scan(&tb)
	bb = bmi * (tb * tb)
	fmt.Printf("berat badan anda: %.f", bb)

}

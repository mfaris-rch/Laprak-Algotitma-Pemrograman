package main

import "fmt"

func main() {
	var x int
	var hasil string
	fmt.Scan(&x)
	hasil = "bukan"
	if x < 0 && x%2 == 0 {
		hasil = "genap negatif"
	}
	fmt.Print(hasil)
}

package main

import "fmt"

func main() {
	var j, hasil, bil, pangkat int
	fmt.Scan(&bil, &pangkat)
	hasil = 1
	for j = 0; j < pangkat; j++ {
		hasil = hasil * bil
	}
	fmt.Print(hasil)
}

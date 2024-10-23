package main

import "fmt"

func main() {
	var n, j, hasil int
	fmt.Scan(&n)
	hasil = 1
	for j = 1; j <= n; j++ {
		hasil = hasil * j
	}
	fmt.Print(hasil)
}

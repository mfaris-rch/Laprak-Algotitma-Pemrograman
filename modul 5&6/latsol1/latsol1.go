package main

import "fmt"

func main() {
	var hasil, n, j int
	fmt.Scan(&n)
	for j = 1; j <= n; j++ {
		hasil = hasil + j
	}
	fmt.Println(hasil)
}

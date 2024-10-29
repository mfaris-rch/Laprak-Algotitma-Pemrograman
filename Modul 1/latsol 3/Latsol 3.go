package main

import "fmt"

func main() {
	var r int
	var phi, L float64
	fmt.Scan(&r)
	phi = 3.14
	L = phi * (float64(r)) * (float64(r))
	fmt.Print(L)
}

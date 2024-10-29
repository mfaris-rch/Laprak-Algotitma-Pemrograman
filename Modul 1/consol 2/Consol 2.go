package main

import "fmt"

func main() {
	// fx = 2 / (x+5) + 5
	// masukan input x
	var x, fx float64
	fmt.Scan(&x)
	fx = 2/(x+5) + 5
	fmt.Print(fx)
}

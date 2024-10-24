package main

import "fmt"

func main() {
	var fx, x float64
	fmt.Print("masukkan nilai f(x): ")
	fmt.Scan(&x)
	fx = 2/(x+5) + 5
	fmt.Println(fx)
}

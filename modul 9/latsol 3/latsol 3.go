package main

import "fmt"

func main() {
	var x, y int
	var hasil1, hasil2 bool
	fmt.Scan(&x, &y)
	hasil1 = y%x == 0
	if x%y == 0 {
		hasil2 = true
	}
	fmt.Print(hasil1, hasil2)

}

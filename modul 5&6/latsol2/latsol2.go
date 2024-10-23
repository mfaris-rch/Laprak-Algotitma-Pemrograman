package main

import "fmt"

func main() {
	var n, j int
	var phi, volume float64
	fmt.Scan(&n)
	phi = 3.14159265358979323846
	for j = 0; j < n; j++ {
		var r, t float64
		fmt.Scan(&r, &t)
		volume = 1.0 / 3.0 * phi * r * r * t
		fmt.Printf("%.14f\n", volume)
	}
}

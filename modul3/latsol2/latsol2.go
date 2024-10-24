package main

import "fmt"

func main() {
	var r, v, l float64
	fmt.Print("jejari:")
	fmt.Scan(&r)
	v = (4.0 / 3.0) * 3.1415926535 * (r * r * r)
	l = 4 * 3.1415926535 * (r * r)
	fmt.Print("Bola dengan jari-jari: ", r)
	fmt.Print("memiliki volume : ", v)
	fmt.Print("luas kulit: ", l)
}

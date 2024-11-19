package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	jumlahmotor := n / 2
	if n%2 != 0 {
		jumlahmotor += 1
	}
	fmt.Println(jumlahmotor)

}

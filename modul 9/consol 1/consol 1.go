package main

import "fmt"

func main() {
	var bil int
	fmt.Scan(&bil)
	if bil < 0 {
		bil = -bil
	}
	fmt.Print(bil)
}

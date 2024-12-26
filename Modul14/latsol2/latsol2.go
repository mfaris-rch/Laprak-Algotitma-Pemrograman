package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	if n <= 1 {
		fmt.Println("Bukan Prima")
	} else {
		prima := true
		for i := 2; i*i <= n; i++ {
			if n%i == 0 {
				prima = false
				break
			}
		}
		if prima {
			fmt.Println("Prima")
		} else {
			fmt.Println("Bukan Prima")
		}
	}
}

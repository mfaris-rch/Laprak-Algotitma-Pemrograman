package main

import "fmt"

func main() {
	var ipk, tak float64
	var predikat string
	fmt.Scan(&ipk)
	if ipk > 3.00 {
		fmt.Scan(&tak)
		if tak < 2.00 {
			predikat = "poor"
		} else if 2.00 <= tak && tak <= 2.75 {
			predikat = "Fair"
		} else if 2.76 <= tak && tak <= 3.00 {
			predikat = "Sastifactory"
		} else if 3.01 <= tak && tak <= 3.50 {
			predikat = "Very Good"
		} else {
			predikat = "Excellent"
		}
		fmt.Println("Cumlaude dengan IPK", ipk, "dan status predikat TAK adalah", predikat)
	} else {
		fmt.Println("Tidak Cumlaude")
	}
}

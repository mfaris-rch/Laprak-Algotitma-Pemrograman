package main

import "fmt"

func main() {
	var bil, d1, d2, d3, d4 int
	var teks string
	fmt.Print("Bilangan: ")
	fmt.Scan(&bil)
	d4 = bil % 10
	d3 = (bil % 100) / 10
	d2 = (bil % 1000) / 100
	d1 = bil / 1000

	if d1 < d2 && d2 < d3 && d3 < d4 {
		teks = "urut membesar"
	} else if d1 > d2 && d2 > d3 && d3 > d4 {
		teks = "urut mengecil"
	} else {
		teks = "tidak urut"
	}
	fmt.Println("digit pada bilangan", bil, teks)
}

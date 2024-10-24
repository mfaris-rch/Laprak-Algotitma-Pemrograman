package main

import "fmt"

func main() {
	var totalbelanja, diskon int
	fmt.Print("masukkan total harga belanja: ")
	fmt.Scanln(&totalbelanja)
	fmt.Print("masukkan diskon: ")
	fmt.Scanln(&diskon)

	totaldiskon := totalbelanja - (totalbelanja * diskon / 100)
	fmt.Print("total harga belanja setelah diskon: ", totaldiskon)
}

package main

import "fmt"

func main() {
	var fahrenheit float64
	var celcius float64

	fmt.Print("Masukkan suhu dalam Fahrenheit: ")
	fmt.Scanln(&fahrenheit)

	// Rumus konversi
	celcius = (fahrenheit - 32) * 5 / 9

	fmt.Printf("%.2f derajat Fahrenheit sama dengan %.2f derajat Celcius\n", fahrenheit, celcius)
}

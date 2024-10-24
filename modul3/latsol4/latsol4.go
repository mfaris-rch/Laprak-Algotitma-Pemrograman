package main

import "fmt"

func main() {
	var celsius float64

	fmt.Print("Masukkan suhu dalam Celsius: ")
	fmt.Scanln(&celsius)

	fahrenheit := (celsius * 9 / 5) + 32
	fmt.Println("Suhu dalam Fahrenheit:", fahrenheit)

	reamur := celsius * 4 / 5
	fmt.Println("Suhu dalam Reamur:", reamur)

	kelvin := celsius + 273.15
	fmt.Println("Suhu dalam Kelvin:", kelvin)

}

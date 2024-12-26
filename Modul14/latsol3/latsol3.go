package main

import "fmt"

func main() {
	var pita string
	var input string
	totalBunga := 0

	fmt.Println("Masukkan nama bunga satu per satu (ketik 'SELESAI' untuk berhenti):")

	for {
		fmt.Printf("Bunga %d: ", totalBunga+1)
		fmt.Scanln(&input)

		if input == "SELESAI" {
			break
		}

		if pita == "" {
			pita = input
		} else {
			pita += " – " + input
		}

		totalBunga++
	}

	fmt.Println("\nPita :", pita)
	fmt.Println("Bunga:", totalBunga)
}

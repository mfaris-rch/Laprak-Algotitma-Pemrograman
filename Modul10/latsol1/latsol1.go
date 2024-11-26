package main

import "fmt"

func main() {
	var gram, beratGram, beratKg, dBiayaGram, dBiayaKg, tBiaya int
	fmt.Print("Berat parsel (gram): ")
	fmt.Scan(&gram)
	beratKg = gram / 1000
	beratGram = gram % 1000
	dBiayaKg = beratKg * 10000
	if beratGram < 500 {
		dBiayaGram = beratGram * 15
	} else if beratGram >= 500 {
		dBiayaGram = beratGram * 5
	}
	if gram >= 10000 {
		tBiaya = dBiayaKg
	} else if gram <= 10000 {
		tBiaya = dBiayaKg + dBiayaGram
	}
	fmt.Println("Detail berat:", beratKg, "kg +", beratGram, "gr")
	fmt.Println("Detail biaya: Rp.", dBiayaKg, "+ Rp.", dBiayaGram)
	fmt.Println("Total biaya: Rp.", tBiaya)
}

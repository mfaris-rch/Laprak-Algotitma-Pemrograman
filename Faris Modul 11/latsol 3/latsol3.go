package main

import "fmt"

func main() {
	var n, result int
	fmt.Scan(&n)
	switch {
	case n%10 == 0:
		result = n / 10
		fmt.Println("kategori:Bilangan Kelipatan 10")
		fmt.Printf("Hasil penjumlahan dengan bilangan berikutnya %d/10 = %d", n, result)
	case n%5 == 0:
		result = n * n
		fmt.Println("kategori:Bilangan Kelipatan 5")
		fmt.Printf("Hasil penjumlahan dengan bilangan berikutnya %d ^2 = %d", n, result)
	case n%2 == 0:
		result = n * (n + 1)
		fmt.Println("kategori:Bilangan genap")
		fmt.Printf("Hasil penjumlahan dengan bilangan berikutnya %d * %d = %d", n, n+1, result)
	case n%2 != 0:
		result = n + (n + 1)
		fmt.Println("kategori:Bilangan ganjil")
		fmt.Printf("Hasil penjumlahan dengan bilangan berikutnya %d + %d = %d", n, n+1, result)

	}
}
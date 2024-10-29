package main

import "fmt"

func main() {
	var nama, kelas, nim string
	fmt.Print("masukan nama:")
	fmt.Scan(&nama)
	fmt.Print("masukan kelas:")
	fmt.Scan(&kelas)
	fmt.Print("masukan nim:")
	fmt.Scan(&nim)

	fmt.Println("\nPerkenalkan saya adalah", nama, ", salah satu mahasiswa Prodi S1-IF dari kelas", kelas, "dengan NIM", nim, ".")
}

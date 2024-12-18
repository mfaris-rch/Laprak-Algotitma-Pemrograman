package main

import "fmt"

func main() {
	var username, password string
	fmt.Scan(&username, &password)
	percobaan := 0
	for username != "Admin" || password != "Admin" {
		fmt.Scan(&username, &password)
		percobaan++
	}
	fmt.Println(percobaan, "percobaan gagal login")
}

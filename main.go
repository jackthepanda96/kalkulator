package main

import "fmt"

func main() {
	var input int
	fmt.Println("Pilihan menu")
	fmt.Println("1. Penjumlahan")
	fmt.Println("2. Pengurangan")
	fmt.Println("3. Perkalian")
	fmt.Println("4. Pembagian")
	fmt.Println("5. Riwayat")
	fmt.Println("9. EXIT")
	fmt.Print("Masukkan menu: ")
	fmt.Scanln(&input)
}

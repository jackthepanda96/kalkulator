package main

import "fmt"

func main() {
	var input, angka1, angka2, hasil int
	for input < 9 {
		fmt.Println("Pilihan menu")
		fmt.Println("1. Penjumlahan")
		fmt.Println("2. Pengurangan")
		fmt.Println("3. Perkalian")
		fmt.Println("4. Pembagian")
		fmt.Println("5. Riwayat")
		fmt.Println("9. EXIT")
		fmt.Print("Masukkan menu: ")
		fmt.Scanln(&input)
		if input < 5 && input > 0 {
			fmt.Print("Masukkan angka pertama")
			fmt.Scanln(&angka1)
			fmt.Print("Masukkan angka kedua")
			fmt.Scanln(&angka2)
			switch input {
			case 1:
				hasil = angka1 + angka2
			case 2:
				hasil = angka1 - angka2
			case 3:
				hasil = angka1 * angka2
			case 4:
				if angka2 > 0 {
					hasil = angka1 / angka2
				} else {
					hasil = 0
				}
			}
			fmt.Println("Hasil :", hasil)
		} else if input == 5 {

		}
	}
}

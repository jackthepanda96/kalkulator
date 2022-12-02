package main

import "fmt"

func main() {
	var input, angka1, angka2, hasil int
	var riwayat []string
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
				riwayat = append(riwayat, fmt.Sprintf("%d + %d = %+d\n", angka1, angka2, hasil))
			case 2:
				hasil = angka1 - angka2
				riwayat = append(riwayat, fmt.Sprintf("%d - %d = %04d\n", angka1, angka2, hasil))
			case 3:
				hasil = angka1 * angka2
				riwayat = append(riwayat, fmt.Sprintf("%d * %d = %o\n", angka1, angka2, hasil))
			case 4:
				if angka2 != 0 {
					hasil = angka1 / angka2
				} else {
					hasil = 0
				}

				riwayat = append(riwayat, fmt.Sprintf("%d / %d = %b\n", angka1, angka2, hasil))
			}
			fmt.Println("Hasil :", hasil)
		} else if input == 5 {
			for i := 0; i < len(riwayat); i++ {
				fmt.Print(riwayat[i])
			}
		}
	}
}

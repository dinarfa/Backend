package main

import (
	"fmt"
)

// Semacam increment
// const(
// 	first = iota
// 	second;
// )

func main() {
	// Deklarasi variabel
	var nama string
	nama = "Dinar"
	fmt.Println(nama)
	// Bisa diubah
	nama = "Fauziah"

	// Tidak perlu tulis tipe data karena sudah di inisialisasikan
	var umur = 19
	fmt.Println(umur)

	// Tidak perlu tulis var dan tipe data.
	agama := "Islam"
	fmt.Println(agama)

	// Deklarasi Multiple variable

	matrix := [2][3]int{
		{5, 6, 7},
		{8, 9, 10},
	}
	for i, row := range matrix {
		for j, value := range row {
			fmt.Printf("matrix[%d][%d] = %d\n", i, j, value)
		}
	}

}

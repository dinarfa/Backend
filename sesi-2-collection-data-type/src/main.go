package main

import "fmt"

func main() {
	// Deklarasi array
	//var variable_name [SIZE]variable_type

	// Contoh
	var countries [2]string

	// Inisialisasi
	countries = [2]string{"Indonesia", "France"}

	fmt.Println(countries)

	// Mengubah Elemen melalui indeks
	numbers := [5]int{10, 20, 30, 40, 50}
	numbers[2] = 99
	numbers[0] = 100

	fmt.Println("Array setelah perubahan :", numbers)

	// Loop Elemen Array
	for i, v := range numbers {
		fmt.Printf("Indeks %d: %d\n", i, v)
	}
	// Kalau tdk mau print indeksnya, pake underscore variabel.

	// Multidimensional Array
	// Seperti baris dan kolom
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

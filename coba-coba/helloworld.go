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
	var tb, jumlahSodara int = 149, 5
	fmt.Println(tb)
	fmt.Println(jumlahSodara)

	// var (
	// 	country = "Indonesia"
	// 	wallet  = 20
	// )
	// fmt.Println(country)
	// fmt.Println(wallet)

	var countries = [3]string{"Indonesia", "Inggris"}
	fmt.Println(countries)
	// members := [3] string{"Dinar", "Dini"}
	// fmt.Println(members(len(members))-1)
	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			if i%3 == 0 && i%5 == 0 {
				fmt.Println("FizzBuzz")
			} else {
				fmt.Println("Fizz")
			}
		} else if i%5 == 0 {
			if i%3 == 0 && i%5 == 0 {
				fmt.Println("FizzBuzz")
			} else {
				fmt.Println("Buzz")
			}
		} else {
			fmt.Println(i)
		}
	}

}

package main

import "fmt"

type Person struct {
	name     string
	position string
}

func main() {
	// orang1 := Person{"Dinar Fauziah", "Backend"}
	// orang2 := orang1

	// orang2.name = "Dinaro"

	// fmt.Println(orang1) // orang1 tidak berubah
	// fmt.Println(orang2)

	// Operator &, gunanya menghemat memori.
	// kita bisa oper alamat variabel orang1 ke orang2 tanpa membuat yang aslinya
	// orang2 := &orang1
	// orang2.name = "Dinaro"
	// fmt.Println(orang1)
	// fmt.Println(orang2)

	// Operator * digunakan untuk mengakses atau memodifikasi nilai yang ditunjuk oleh pointer.
	orang1 := Person{"Dinar", "Backend"}
	orang2 := &orang1

	orang2.name = "Dini"

	*orang2 = Person{"Fahmi", "Frontend"}

	fmt.Println(orang1)
	fmt.Println(orang2)



}

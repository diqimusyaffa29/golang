package main

import "fmt"

func main() {
	var name string

	name = "Muhammad Shiddiqi Musyaffa"
	fmt.Println(name)

	name = "Diqi" //akan menimpa nilai sebelumnya dari variabel yang di panggil
	fmt.Println(name)

	//=====================================================================================
	var name2 = "Anjay Mabar" //golang akan secara otomatis akan memasukkan tipe datanya, yaitu string
	fmt.Println(name2)

	//======================================================================================
	name3 := "Kocak Gaming" //menggunakan := hanya pada awal inisialisasi sebuah variabel baru
	fmt.Println(name3)

	//======================================================================================
	var (
		firstName  = "Muhammad"
		middleName = "Shiddiqi"
		LastName   = "Musyaffa'"
	)

	fmt.Println(firstName, middleName, LastName)
}

package main

import "fmt"

func main() {
	names := [...]string{"Budi", "Rio", "Kocak", "Gaming", "Mantap", "Jiwa"}
	slice := names[4:6]
	fmt.Println(slice)
	fmt.Println(slice[0])
	fmt.Println(slice[1])

	slice2 := names[:3]
	fmt.Println(slice2)

	slice3 := names[3:]
	fmt.Println(slice3)

	slice4 := names[:]
	fmt.Println(slice4)
	fmt.Println(len(slice4))
	fmt.Println(cap(slice3))

	// Append Slice

	days := [...]string{
		"senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu", "Minggu",
	}
	daySlice1 := days[5:]
	daySlice1[0] = "Sabtu Baru" //ini akan menimpa data dari slice days pada pointer 5
	daySlice1[1] = "Minggu Mantap Uhuy"
	fmt.Println(cap(daySlice1)) //capacity nya hanya memiliki 2, maka tidak bisa lagi memasukkan data lagi ke dalam daySlice1
	fmt.Println(daySlice1)
	fmt.Println(days)

	daySlice2 := append(daySlice1, "Libur Baru") //karena pada daySlice1 cap nya sudah habis, maka secara otomatis daySlice2 akan membuat array baru
	// daysBaru := [...]string{
	// 	"senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu", "Minggu", "Libur Baru"
	// }
	fmt.Println(daySlice2)
	fmt.Println(days)

	// Make Slice

	newSlice := make([]string, 2, 5)
	newSlice[0] = "Diqi"
	newSlice[1] = "Diqi2"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	newSlice2 := append(newSlice, "Diqi3")
	fmt.Println(newSlice2)
	fmt.Println(len(newSlice2))
	fmt.Println(cap(newSlice2))

	fromSlice := days[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))

	copy(toSlice, fromSlice)

	fmt.Println(fromSlice)
	fmt.Println(toSlice)

	iniArray := [...]string{"Kocak", "Gaming", "Uhuy"}
	iniSlice := []string{"Kocak", "Gaming", "Uhuy"}

	fmt.Println(iniArray)
	fmt.Println(iniSlice)
}

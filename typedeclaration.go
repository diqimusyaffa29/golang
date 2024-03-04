package main

import "fmt"

func main()  {
	type NoKTP string

	var ktpDiqi NoKTP = "11111111111111"
	fmt.Println(ktpDiqi)
	fmt.Println(NoKTP("22222222222222"))
	fmt.Println(NoKTP("22222222222222"))
}
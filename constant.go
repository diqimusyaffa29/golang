package main

func main() {
	const firstName string = "Diqi"
	const lastName = "Musyaffa'"

	//ERRROR karena tipe variablenya adalah constant
	//firstName = "Kocak"
	//lastName = "Kocakuhuyy"
	const (
		address string = "Jl. Kocak"
		country        = "Indonesia"
	)

	//ERROT karena tipe variablenya adalah constant
	address = "Mantap jiwa"
	country = "Mantap jiwa"
}

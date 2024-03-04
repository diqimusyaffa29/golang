package main

import "fmt"

func main()  {
	var names [3] string
	names[0] = "Muhammad"
	names[1] = "Shiddiqi"
	names[2] = "Musyaffa'"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])


	var values = [...]int{
		90,80,95,100,8282,372,
	}

	fmt.Println(values)

	// Function Array
	fmt.Println(len(values))
}
package main

import "fmt"

func main()  {
	type person struct {
		firstName string
		lastName string
		age uint8
	}
	p := person{firstName: "Tony", lastName: "Aziz", age: 28}
	fmt.Println(p)
	// learn value type
}

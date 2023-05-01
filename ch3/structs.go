package main

import "fmt"

func main() {
	type person struct {
		name string
		age  int
		per  string
	}

	var fred person
	fmt.Println(fred)

	bob := person{}
	fmt.Println(bob)

	julia := person{
		"Julia",
		40,
		"cat",
	}
	fmt.Println(julia)

	beth := person{
		age:  40,
		name: "Beth",
	}
	fmt.Println(beth)

	bob.name = "Bob"
	fmt.Println(bob.name)
}

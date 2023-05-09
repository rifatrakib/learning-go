package main

import "fmt"

type MyFuncOpts struct {
	FirstName string
	LastName  string
	Age       int
}

func MyFunc(opts MyFuncOpts) error {
	fmt.Println(opts.FirstName)
	fmt.Println(opts.LastName)
	fmt.Println(opts.Age)
	return nil
}

func main() {
	MyFunc(MyFuncOpts{
		LastName: "Patel",
		Age:      40,
	})
	MyFunc(MyFuncOpts{
		FirstName: "Joe",
		LastName:  "Smith",
	})
}

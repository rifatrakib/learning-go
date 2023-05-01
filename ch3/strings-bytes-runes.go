package main

import "fmt"

func main() {
	var s string = "Hello, 世界"
	var bs []byte = []byte(s)
	var rs []rune = []rune(s)
	fmt.Println(bs)
	fmt.Println(rs)
}

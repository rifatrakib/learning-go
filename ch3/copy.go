package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4}
	y := make([]int, 4)
	num := copy(y, x)
	fmt.Println(num, y)

	z := make([]int, 2)
	num2 := copy(z, x)
	fmt.Println(num2, z)

	w := make([]int, 2)
	num3 := copy(w, x[2:])
	fmt.Println(num3, w)

	s := []int{1, 2, 3, 4}
	num4 := copy(s[:3], s[1:])
	fmt.Println(num4, s)

	a := []int{1, 2, 3, 4}
	b := [4]int{5, 6, 7, 8}
	c := make([]int, 2)
	copy(c, b[:])
	fmt.Println(c)
	copy(b[:], a)
	fmt.Println(b)
}

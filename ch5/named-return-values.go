package main

import (
	"errors"
	"fmt"
)

func divAndRemainderNamed(numerator int, denominator int) (result int, remainder int, err error) {
	if denominator == 0 {
		return 0, 0, errors.New("cannot divide by zero")
	}
	result, remainder = numerator/denominator, numerator%denominator
	return result, remainder, err
}

func main() {
	x, y, z := divAndRemainderNamed(5, 2)
	fmt.Println(x, y, z)
}

package main

import "fmt"

func main() {
	// complete for loop
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// condition-only for loop
	j := 1
	for j < 10 {
		fmt.Println(j)
		j = j * 2
	}

	// infinite for loop
	for {
		if j > 20 {
			break
		}
		fmt.Println(j)
		j += 2
	}

	// using continue for better readability
	for i := 0; i < 30; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
			continue
		}

		if i%3 == 0 {
			fmt.Println("Fizz")
			continue
		}

		if i%5 == 0 {
			fmt.Println("Buzz")
			continue
		}

		fmt.Println(i)
	}

	// for-range loop
	evenVals := []int{2, 4, 6, 8, 10, 12}
	for i, v := range evenVals {
		fmt.Println(i, v)
	}

	for _, v := range evenVals {
		fmt.Println(v)
	}

	uniqueNames := map[string]bool{"Fred": true, "Raul": true, "Wilma": true}
	for k, _ := range uniqueNames {
		fmt.Println(k)
	}
}

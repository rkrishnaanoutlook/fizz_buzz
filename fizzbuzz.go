package main

import (
	"fmt"
	"strconv"
)

func fizzbuzz(n int) []string {
	output := make([]string, 0, n)
	for i := 1; i <= n; i++ {
		switch {
		case i%15 == 0:
			output = append(output, "fizzbuzz")
		case i%5 == 0:
			output = append(output, "buzz")
		case i%3 == 0:
			output = append(output, "fizz")
		default:
			output = append(output, strconv.Itoa(i))
		}
	}
	return output
}

func main() {
	fmt.Println(fizzbuzz(20))
}

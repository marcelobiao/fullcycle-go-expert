package main

import (
	"fmt"
)

func main() {
	total := func() int {
		return sum(5, 10, 50)
	}

	result := total()

	fmt.Println(result)
}

func sum(numeros ...int) int {
	total := 0
	for _, num := range numeros {
		total += num
	}
	return total
}

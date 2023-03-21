package main

import (
	"errors"
	"fmt"
)

func main() {
	valor, err := sum(5, 10)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(valor)
}

func sum(a int, b int) (int, error) {
	sum := a + b
	if sum >= 50 {
		return 0, errors.New("Soma maior que 50")
	}
	return sum, nil
}

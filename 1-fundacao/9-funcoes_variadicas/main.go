package main

import (
	"errors"
	"fmt"
)

func main() {
	valor, err := sum(5, 10, 50)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(valor)
}

func sum(numeros ...int) (int, error) {
	total := 0
	for _, num := range numeros {
		total += num
	}

	if total >= 50 {
		return 0, errors.New("Soma maior que 50")
	}
	return total, nil
}

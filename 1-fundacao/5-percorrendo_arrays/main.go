package main

import "fmt"

type ID int

func main() {
	var meuArray [3]int
	meuArray[0] = 1
	meuArray[1] = 2
	meuArray[2] = 3

	fmt.Println("Tamanho do array:", len(meuArray))

	for i, v := range meuArray {
		fmt.Println("Index:", i, " Valor:", v)
	}
}

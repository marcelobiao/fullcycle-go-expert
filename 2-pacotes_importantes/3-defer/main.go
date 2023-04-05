package main

import "fmt"

// https://go.dev/tour/flowcontrol/12
func main() {
	// defer adia a execução de um comando
	// As instruções defer são empilhadas e resolvidas ao fim do escopo
	defer fmt.Println("5")
	defer fmt.Println("4")
	fmt.Println("1")
	fmt.Println("2")
	fmt.Println("3")
}

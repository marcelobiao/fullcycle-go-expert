package main

import "fmt"

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
}

func main() {
	marcelo := Cliente{
		Nome:  "Marcelo",
		Idade: 30,
		Ativo: true,
	}
	fmt.Println(marcelo)
	marcelo.Ativo = false
	fmt.Println(marcelo)
}

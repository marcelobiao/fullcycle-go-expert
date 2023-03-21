package main

import "fmt"

type Endereco struct {
	Lagradouro string
	Numero     int
	Cidade     string
	Estado     string
}

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
	Endereco
}

func main() {
	marcelo := Cliente{
		Nome:  "Marcelo",
		Idade: 30,
		Ativo: true,
	}
	fmt.Println(marcelo)
	marcelo.Ativo = false
	marcelo.Numero = 20
	fmt.Println(marcelo)
}

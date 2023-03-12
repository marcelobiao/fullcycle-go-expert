package main

import "fmt"

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
}

func (c Cliente) Desativar() {
	c.Ativo = false
	fmt.Println("Cliente desativado")
}

func main() {
	marcelo := Cliente{
		Nome:  "Marcelo",
		Idade: 30,
		Ativo: true,
	}
	fmt.Println(marcelo)
	marcelo.Desativar()
}

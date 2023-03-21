package main

import "fmt"

type Pessoa interface {
	Desativar()
}
type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
}

func (c Cliente) Desativar() {
	c.Ativo = false
	fmt.Println("Cliente desativado")
}

func Desativacao(pessoa Pessoa) {
	pessoa.Desativar()
}

func main() {
	marcelo := Cliente{
		Nome:  "Marcelo",
		Idade: 30,
		Ativo: true,
	}
	fmt.Println(marcelo)
	Desativacao(marcelo)
}

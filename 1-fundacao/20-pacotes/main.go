package main

import (
	"fmt"

	"pacotes/matematica"
)

func main() {
	s := matematica.Soma(10, 20)
	fmt.Println("Resultado: %v", s)
}

/*
Cria um novo pacote
$ go mod init name

Importa um novo pacote
$ go get nome

Sincroniza seus pacotes do go.mod
$ go mod tidy
*/

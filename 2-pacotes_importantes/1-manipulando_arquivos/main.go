package main

import (
	"bufio"
	"fmt"
	"os"
)

// https://pkg.go.dev/os
func main() {
	// Escrita
	f, err := os.Create("arquivo.txt")
	if err != nil {
		panic(err)
	}

	tamanho, err := f.WriteString("Hello, World!\nHello, World!")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Arquivo criado com sucesso, tamanho %d bytes", tamanho)

	f.Close()

	// Leitura

	arquivoCompleto, err := os.ReadFile("arquivo.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(arquivoCompleto))

	// Leitura de pouco em pouco abrindo o arquivo

	arquivo, err := os.Open("arquivo.txt")
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(arquivo)

	for {
		line, a, err := reader.ReadLine()
		if err != nil {
			break
		}
		fmt.Println(string(line))
		fmt.Println(a)
	}
}

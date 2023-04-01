package main

/*
https://pkg.go.dev/go/build
Compilando o projeto
$ go build main.go
*/
func main() {
	numeros := []string{"um", "dois", "três"}
	for k, v := range numeros {
		println(k, v)
	}
}

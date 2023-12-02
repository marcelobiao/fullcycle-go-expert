package main

import "fmt"

// Thread 1
func main() {
	hello := make(chan string)
	go recebe("hello", hello)
	ler(hello)
}

// send-only channel chan<-
func recebe(nome string, hello chan<- string) {
	hello <- nome
}

// receive-only channel chan<-
func ler(data <-chan string) {
	fmt.Println(<-data)
}

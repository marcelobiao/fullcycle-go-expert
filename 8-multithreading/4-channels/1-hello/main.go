package main

import "fmt"

// Thread 1
func main() {
	ch := make(chan string) // Vazio

	// Thread 2
	go func() {
		ch <- "olá Mundo!" // Cheio
	}()

	// Thread 1
	msg := <-ch // Vazio
	fmt.Println(msg)
}

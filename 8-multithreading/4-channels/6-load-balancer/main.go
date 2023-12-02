package main

import (
	"fmt"
	"time"
)

// Thread 1
func main() {
	data := make(chan int)
	qtdWorkers := 100000

	// Inicializa os workers
	for i := 0; i < qtdWorkers; i++ {
		go worker(i, data)
	}

	for i := 0; i < 1000000; i++ {
		data <- i
	}
}

func worker(workerId int, data chan int) {
	for x := range data {
		fmt.Printf("worker %d received %d\n", workerId, x)
		time.Sleep(time.Second)
	}
}

package main

import (
	"fmt"
	"sync"
	"time"
)

func task(name string, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		fmt.Printf("%d: Task %s is running \n", i, name)
		time.Sleep(100 * time.Millisecond)
	}
}

// Thread 1
func main() {
	wg := sync.WaitGroup{}
	// Thread 2
	wg.Add(1)
	go task("A", &wg)
	// Thread 3
	wg.Add(1)
	go task("B", &wg)

	wg.Wait()
	fmt.Println("Fim")
}

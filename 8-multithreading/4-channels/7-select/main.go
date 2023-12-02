package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

type Message struct {
	id   int64
	body string
}

// Thread 1
func main() {
	c1 := make(chan Message)
	c2 := make(chan Message)
	var i int64 = 0
	go func() { // RabbitMQ
		for {
			atomic.AddInt64(&i, 1)
			// time.Sleep(time.Second)
			msg := Message{i, "Hello from RabbitMQ"}
			c1 <- msg
		}
	}()
	go func() { // Kafka
		for {
			atomic.AddInt64(&i, 1)
			// time.Sleep(time.Second)
			msg := Message{i, "Hello from Kafka"}
			c2 <- msg
		}
	}()
	for {
		select {
		case msg := <-c1:
			fmt.Printf("Received from RabbitMQ: %d - %s\n", msg.id, msg.body)
		case msg := <-c2:
			fmt.Printf("Received from Kafka: %d - %s\n", msg.id, msg.body)
		case <-time.After(time.Second * 5):
			println("timeout")
		}
	}
}

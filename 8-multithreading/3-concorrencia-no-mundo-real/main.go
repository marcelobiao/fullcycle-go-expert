package main

import (
	"fmt"
	"net/http"
	"sync"
	"sync/atomic"
)

// number Race condition
var number uint64 = 0

// Solução 1
func main() {
	m := sync.Mutex{}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		m.Lock()
		number++
		m.Unlock()
		w.Write([]byte(fmt.Sprintf("Hello World! %d", number)))
	})
	http.ListenAndServe(":3000", nil)
}

// Solução 2
func main2() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		atomic.AddUint64(&number, 1)
		w.Write([]byte(fmt.Sprintf("Hello World! %d", number)))
	})
	http.ListenAndServe(":3000", nil)
}

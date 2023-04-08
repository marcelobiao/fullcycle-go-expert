package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log.Println("request Iniciada")

	defer log.Println("request finalizada")

	select {
	case <-time.After(5 * time.Second):
		// Imprime no comand line stdout
		log.Println("request processada com sucesso")
		// Imprime no browser
		w.Write([]byte("request processada com sucesso"))
	case <-ctx.Done():
		// Imprime no comand line stdout
		log.Println("request cancelada pelo cliente")
		// Imprime no browser
		http.Error(w, "request cancelada pelo cliente", http.StatusRequestTimeout)
	}
}

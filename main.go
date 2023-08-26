package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Привет, это простой HTTP-сервер на Go!")
	})

	// Запуск HTTP-сервера на порту 8080
	http.ListenAndServe(":8080", nil)
}

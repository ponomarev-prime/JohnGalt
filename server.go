package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		clientIP := r.RemoteAddr
		fmt.Printf("HTTP reqest: %s %s %s\n", clientIP, r.Method, r.URL.Path)
		fmt.Fprintf(w, "Привет, это простой HTTP-сервер на Go!")
	})

	http.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
		clientIP := r.RemoteAddr
		// Чтение данных из тела запроса
		body := r.URL.Query().Get("data")

		// Отправка данных клиенту в ответ на запрос
		fmt.Fprintf(w, "Вы отправили: %s", body)
		fmt.Printf("HTTP reqest: %s %s %s\n", clientIP, r.Method, r.URL.Path)
	})

	// Запуск HTTP-сервера на порту 8080
	http.ListenAndServe(":8080", nil)
}

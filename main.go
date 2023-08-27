package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/echo", echoHandler)
	http.HandleFunc("/html", pageHandler)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	// server go
	err := http.ListenAndServe(":8080", nil) // устанавливаем порт веб-сервера
	//err := http.ListenAndServeTLS(":8080", "cert.pem", "key.pem", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	// client IP
	clientIP := r.RemoteAddr

	// Terminal output
	fmt.Printf("HTTP reqest to /: %s %s %s\n", clientIP, r.Method, r.URL.Path)

	// http response
	fmt.Fprintf(w, "Yo, this is GO simple server!")
}

func echoHandler(w http.ResponseWriter, r *http.Request) {
	// client IP
	clientIP := r.RemoteAddr

	// read data from reqest
	body := r.URL.Query().Get("data")

	// Terminal output
	fmt.Printf("HTTP reqest to /echo: %s %s %s\n", clientIP, r.Method, r.URL.Path)

	// Response
	fmt.Fprintf(w, "You send me : %s", body)
}

func pageHandler(w http.ResponseWriter, r *http.Request) {
	// client IP
	clientIP := r.RemoteAddr

	fmt.Printf("HTTP reqest to /: %s %s %s\n", clientIP, r.Method, r.URL.Path)

	http.ServeFile(w, r, "index.html")
}

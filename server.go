package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/echo", echoHandler)

	// server go
	http.ListenAndServe(":8080", nil)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	// client IP
	clientIP := r.RemoteAddr

	// Terminal output
	fmt.Printf("HTTP reqest: %s %s %s\n", clientIP, r.Method, r.URL.Path)

	// http response
	fmt.Fprintf(w, "Yo, this is GO simple server!")
}

func echoHandler(w http.ResponseWriter, r *http.Request) {
	// client IP
	clientIP := r.RemoteAddr

	// read data from reqest
	body := r.URL.Query().Get("data")

	// Terminal output
	fmt.Printf("HTTP reqest: %s %s %s\n", clientIP, r.Method, r.URL.Path)

	// Response
	fmt.Fprintf(w, "You send me : %s", body)
}

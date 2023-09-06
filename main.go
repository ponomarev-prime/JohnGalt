package main

import (
	"JohnGalt/handlerman"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handlerman.RootHandler)
	http.HandleFunc("/echo", handlerman.EchoHandler)
	http.HandleFunc("/html", handlerman.PageHandler)
	http.HandleFunc("/time", handlerman.TimesHandler)
	http.HandleFunc("/proc", handlerman.ProcinfoHandler)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	// server go
	err := http.ListenAndServe(":8080", nil) // устанавливаем порт веб-сервера
	//err := http.ListenAndServeTLS(":8080", "cert.pem", "key.pem", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

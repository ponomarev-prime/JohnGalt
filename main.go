package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

const bdate string = "04.08.1992"

func main() {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/echo", echoHandler)
	http.HandleFunc("/html", pageHandler)
	http.HandleFunc("/time", timesHandler)
	http.HandleFunc("/proc", procinfoHandler)

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
	// Terminal output
	fmt.Printf("HTTP reqest to /echo: %s %s %s\n", clientIP, r.Method, r.URL.Path)

	// read data from reqest
	body := r.URL.Query().Get("data")

	// Response
	fmt.Fprintf(w, "You send me : %s", body)
}

func pageHandler(w http.ResponseWriter, r *http.Request) {
	// client IP
	clientIP := r.RemoteAddr

	fmt.Printf("HTTP reqest to /: %s %s %s\n", clientIP, r.Method, r.URL.Path)

	http.ServeFile(w, r, "index.html")
}

func timesHandler(w http.ResponseWriter, r *http.Request) {
	clientIP := r.RemoteAddr
	fmt.Printf("HTTP reqest to /: %s %s %s\n", clientIP, r.Method, r.URL.Path)

	now := time.Now()

	birthDate, err := time.Parse("02.01.2006", bdate)
	if err != nil {
		http.Error(w, "Error parse date", http.StatusInternalServerError)
		return
	}

	dur := now.Sub(birthDate)
	sec := int(dur.Seconds())

	fmt.Fprintf(w, "Duration from %s to Now :: %d", birthDate, sec)
}

func procinfoHandler(w http.ResponseWriter, r *http.Request) {
	clientIP := r.RemoteAddr
	fmt.Printf("HTTP reqest to /: %s %s %s\n", clientIP, r.Method, r.URL.Path)

	cpuInfo, err := os.ReadFile("/proc/cpuinfo")
	if err != nil {
		fmt.Println("Ошибка при чтении файла cpuinfo:", err)
		return
	}

	modelName := getModelName(string(cpuInfo))

	fmt.Fprintf(w, "PROC :: %s", modelName)
}

func getModelName(cpuInfo string) string {
	lines := strings.Split(cpuInfo, "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "model name") {
			parts := strings.SplitN(line, ":", 2)
			if len(parts) > 1 {
				return strings.TrimSpace(parts[1])
			}
		}
	}
	return "Информация о модели процессора не найдена"
}

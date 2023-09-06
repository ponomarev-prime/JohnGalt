package handlerman

import (
	"JohnGalt/procmodel"
	"fmt"
	"net/http"
	"os"
	"time"
)

const bdate string = "04.08.1992"

func RootHandler(w http.ResponseWriter, r *http.Request) {
	// client IP
	clientIP := r.RemoteAddr

	// Terminal output
	fmt.Printf("HTTP reqest to /: %s %s %s\n", clientIP, r.Method, r.URL.Path)

	// http response
	fmt.Fprintf(w, "Yo, this is GO simple server!")
}

func EchoHandler(w http.ResponseWriter, r *http.Request) {
	// client IP
	clientIP := r.RemoteAddr
	// Terminal output
	fmt.Printf("HTTP reqest to /echo: %s %s %s\n", clientIP, r.Method, r.URL.Path)

	// read data from reqest
	body := r.URL.Query().Get("data")

	// Response
	fmt.Fprintf(w, "You send me : %s", body)
}

func PageHandler(w http.ResponseWriter, r *http.Request) {
	// client IP
	clientIP := r.RemoteAddr

	fmt.Printf("HTTP reqest to /: %s %s %s\n", clientIP, r.Method, r.URL.Path)

	http.ServeFile(w, r, "index.html")
}

func TimesHandler(w http.ResponseWriter, r *http.Request) {
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

func ProcinfoHandler(w http.ResponseWriter, r *http.Request) {
	clientIP := r.RemoteAddr
	fmt.Printf("HTTP reqest to /: %s %s %s\n", clientIP, r.Method, r.URL.Path)

	cpuInfo, err := os.ReadFile("/proc/cpuinfo")
	if err != nil {
		fmt.Println("Ошибка при чтении файла cpuinfo:", err)
		return
	}

	modelName := procmodel.GetModelName(string(cpuInfo))

	fmt.Fprintf(w, "PROC :: %s", modelName)
}

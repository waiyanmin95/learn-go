package main

import (
	"log"
	"net/http"
)

func write(writer http.ResponseWriter, message string) {
	_, err := writer.Write([]byte(message))
	if err != nil {
		log.Fatal(err)
	}
}

func englishHandler(w http.ResponseWriter, req *http.Request) {
	write(w, "Hello, Web!")
}

func frenchHandler(w http.ResponseWriter, req *http.Request) {
	write(w, "Salut, Web!")
}

func hindiHandler(w http.ResponseWriter, req *http.Request) {
	write(w, "Namaste, Web!")
}

func main() {
	http.HandleFunc("/hello", englishHandler)
	http.HandleFunc("/salut", frenchHandler)
	http.HandleFunc("/namaste", hindiHandler)

	log.Fatal(http.ListenAndServe(":3333", nil))
}

package main

import (
	"io"
	"log"
	"net/http"
)

//h1 := func(writer http.ResponseWriter,_ *http.Request) {
//	message := []byte("Hello World")
//	_, err := writer.Write(message)
//	if err != nil {
//		log.Fatal(err)
//	}
//}
//
//h2 := func(writer http.ResponseWriter,request *http.Request) {
//	message := []byte("YOSH")
//	_, err := writer.Write(message)
//	if err != nil {
//		log.Fatal(err)
//	}
//}

func main() {
	h1 := func(w http.ResponseWriter, _ *http.Request) {
		io.WriteString(w, "Hello from a HandleFunc #1!\n")
	}

	h2 := func(w http.ResponseWriter, _ *http.Request) {
		io.WriteString(w, "Hello from a HandleFunc TWOOO!\n")
	}

	http.HandleFunc("/", h1)
	http.HandleFunc("/endpoint", h2)
	err := http.ListenAndServe(":8080", nil)
	log.Fatal(err)
}

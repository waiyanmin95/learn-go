package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	log.Fatal(http.ListenAndServeTLS("localhost:8443", "cert.pem", "key.pem", nil))
}

func index(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("Hell World, This is TLS Server!"))
}

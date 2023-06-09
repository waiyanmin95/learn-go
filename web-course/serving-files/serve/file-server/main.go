package main

import (
	"io"
	"net/http"
)

func eks(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	io.WriteString(w, `<img src="/eks.png">`)
}

func main() {
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.HandleFunc("/eks", eks)

	http.ListenAndServe("localhost:8080", nil)
}

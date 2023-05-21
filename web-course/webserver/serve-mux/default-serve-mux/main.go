package main

import (
	"io"
	"net/http"
)

func d(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Woof Woof Woof")
}

func c(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Meow Meow Meow")
}

func main() {
	http.HandleFunc("/dog/", d)
	http.HandleFunc("/cat", c)

	http.ListenAndServe("localhost:8080", nil)
}

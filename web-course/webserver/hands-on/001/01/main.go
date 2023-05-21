package main

import (
	"io"
	"net/http"
)

func index(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "This is Index page.")
}

func dog(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Woof Woof Woof.")
}

func me(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Biology is the study of living things.")
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/me/", me)

	http.ListenAndServe("localhost:8080", nil)
}

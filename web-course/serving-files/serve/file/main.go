package main

import (
	"io"
	"net/http"
)

func dog(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	io.WriteString(w, `<img src="/eks.png">`)
}

func dogPic(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "eks.png")
}

func main() {
	http.HandleFunc("/", dog)
	http.HandleFunc("/eks.png", dogPic)

	http.ListenAndServe("localhost:8080", nil)
}

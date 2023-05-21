package main

import (
	"io"
	"net/http"
	"os"
)

func dog(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	io.WriteString(w, `
	<img src="/golden.jpg">
	`)
}

func dogPic(w http.ResponseWriter, req *http.Request) {
	f, err := os.Open("eks.png")
	if err != nil {
		http.Error(w, "Image not found", http.StatusNotFound)
		return
	}

	defer f.Close()

	fi, err := f.Stat()
	if err != nil {
		http.Error(w, "Image not found", http.StatusNotFound)
		return
	}
	http.ServeContent(w, req, f.Name(), fi.ModTime(), f) // better option over "io.Copy"
}

func main() {
	http.HandleFunc("/", dog)
	http.HandleFunc("/golden.jpg", dogPic)

	http.ListenAndServe("localhost:8080", nil)
}

package main

import (
	"io"
	"net/http"
)

type hotdog int

func (m hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/dog":
		io.WriteString(w, "Woof Woof Woof")
	case "/cat":
		io.WriteString(w, "Meow Meow")
	}
}

func main() {
	var d hotdog
	http.ListenAndServe("localhost:8080", d)
}

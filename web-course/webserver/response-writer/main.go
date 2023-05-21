package main

import (
	"fmt"
	"net/http"
)

type hotdog int

func (m hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("WYM-Key", "This is from Learn Go")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintln(w, "<h1> Any code you want to run </h1>")
}

func main() {
	var d hotdog
	http.ListenAndServe("localhost:8080", d)
}

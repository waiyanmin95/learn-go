package main

import (
	"io"
	"log"
	"net/http"
)

func helper(w http.ResponseWriter, req *http.Request) {
	v := req.FormValue("q")
	io.WriteString(w, "Do my search : "+v)
}

func main() {
	http.HandleFunc("/", helper)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

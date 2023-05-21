package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/bar", bar)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func foo(w http.ResponseWriter, req *http.Request) {
	fmt.Println("You request method on FOO is: ", req.Method)
}

func bar(w http.ResponseWriter, req *http.Request) {
	fmt.Println("You request method on BAR is: ", req.Method)
	http.Redirect(w, req, "/", http.StatusMovedPermanently)
}

package main

import "net/http"

func main() {
	//http.ListenAndServe("localhost:8080", http.FileServer(http.Dir(".")))
	http.Handle("/", http.FileServer(http.Dir(".")))
}

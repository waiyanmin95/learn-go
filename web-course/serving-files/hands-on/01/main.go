package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
)

func foo(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "foo ran")
}

func dog(w http.ResponseWriter, req *http.Request) {
	tpl, err := template.ParseFiles("dog.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	tpl.ExecuteTemplate(w, "dog.gohtml", nil)
}

func doggo(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "dog.png")
}
func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/dog", dog)
	http.HandleFunc("/doggo", doggo)
	http.ListenAndServe("localhost:8080", nil)
}

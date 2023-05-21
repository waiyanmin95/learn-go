package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
)

const data = "Biology"

func index(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "<h1>This is Index page.</h1>")
}

func dog(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "<h1>Woof Woof Woof.</h1>")
}

func me(w http.ResponseWriter, req *http.Request) {
	tpl, err := template.ParseFiles("index.gohtml")
	if err != nil {
		log.Fatalln("Error Parsing Template", err)
	}

	err = tpl.ExecuteTemplate(w, "index.gohtml", data)
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/me/", me)

	http.ListenAndServe("localhost:8080", nil)
}

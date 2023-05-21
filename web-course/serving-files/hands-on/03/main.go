package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("templates/index.gohtml"))
}

func main() {
	fs := http.FileServer(http.Dir("public"))

	http.HandleFunc("/", hotdog)
	http.Handle("/pics/", fs)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func hotdog(w http.ResponseWriter, req *http.Request) {
	err := tpl.Execute(w, nil)
	if err != nil {
		log.Fatalln("Template didn't execute", err)
	}
}

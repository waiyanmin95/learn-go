package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func foo(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Your request method at FOO is:", req.Method)
}

func bar(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Your request method at BAR is:", req.Method)
	// process the redirect
	http.Redirect(w, req, "/", http.StatusTemporaryRedirect)
}

func barred(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Your request method at BARRED is:", req.Method)
	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/bar", bar)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.HandleFunc("/barred", barred)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

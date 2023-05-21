package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", index)
	router.Handle("/favicon.ico", http.NotFoundHandler())
	log.Fatal(http.ListenAndServe("localhost:8080", router))

}

func index(w http.ResponseWriter, req *http.Request) {
	c := getCookie(w, req)
	c = appendValue(w, c)
	xs := strings.Split(c.Value, "|")
	fmt.Println(xs)
	err := tpl.ExecuteTemplate(w, "index.gohtml", xs)
	if err != nil {
		return
	}
}

func getCookie(w http.ResponseWriter, req *http.Request) *http.Cookie {
	c, err := req.Cookie("session")
	if err != nil {
		sID := uuid.New()
		c = &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
		http.SetCookie(w, c)
	}
	return c
}

func appendValue(w http.ResponseWriter, c *http.Cookie) *http.Cookie {
	// Define values
	p1 := "apple.jpg"
	p2 := "ball.jpg"
	p3 := "cat.jpg"

	// Append the values
	s := c.Value

	// Debug
	fmt.Println(s)

	if !strings.Contains(s, p1) {
		s += "|" + p1
	}

	// Debug
	fmt.Println(s)

	if !strings.Contains(s, p2) {
		s += "|" + p2
	}

	// Debug
	fmt.Println(s)

	if !strings.Contains(s, p3) {
		s += "|" + p3
	}

	// Debug
	fmt.Println(s)
	c.Value = s
	http.SetCookie(w, c)
	return c
}

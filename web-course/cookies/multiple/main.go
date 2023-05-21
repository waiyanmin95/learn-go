package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", set)
	http.HandleFunc("/read", read)
	http.HandleFunc("/bundle", bundle)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func set(w http.ResponseWriter, req *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:  "application",
		Value: "hello-world",
	})
	fmt.Fprintln(w, "Your Cookies Written - Check the browser")
	fmt.Fprintln(w, "It's under the developer tools")
}

func read(w http.ResponseWriter, req *http.Request) {

	c1, err := req.Cookie("application")
	if err != nil {
		log.Println(err)
	} else {
		fmt.Fprintln(w, "Your Cookie #1", c1)
	}

	c2, err := req.Cookie("general")
	if err != nil {
		log.Println(err)
	} else {
		fmt.Fprintln(w, "Your Cookie #2", c2)
	}

	c3, err := req.Cookie("mobile")
	if err != nil {
		log.Println(err)
	} else {
		fmt.Fprintln(w, "Your Cookie #3", c3)
	}
}

func bundle(w http.ResponseWriter, req *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:  "general",
		Value: "This is something about general",
	})

	http.SetCookie(w, &http.Cookie{
		Name:   "mobile",
		Value:  "This is something about mobile",
		Secure: false,
	})

	fmt.Fprintln(w, "Your Cookies Written - Check the browser")
	fmt.Fprintln(w, "It's under the developer tools")
}

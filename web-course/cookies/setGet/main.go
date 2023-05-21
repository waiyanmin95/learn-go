package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", set)
	http.HandleFunc("/read", read)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func set(w http.ResponseWriter, req *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:  "ivan-cookies",
		Value: "WarrHarrHarr",
	})
	fmt.Fprintln(w, "Cookies Written - Check your browser")
	fmt.Fprintln(w, "It's under the dev tools")
}

func read(w http.ResponseWriter, req *http.Request) {

	c, err := req.Cookie("ivan-cookies")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprintln(w, "Your cookie", c)
}

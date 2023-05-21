package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
)

func main() {
	http.HandleFunc("/", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func index(w http.ResponseWriter, req *http.Request) {
	cookie, err := req.Cookie("session")
	if err != nil {
		id := uuid.New()
		cookie = &http.Cookie{
			Name:  "session",
			Value: id.String(),
			Path:  "/",
			// Secure:   true,
			HttpOnly: true,
			MaxAge:   int(time.Second * 10),
		}
		http.SetCookie(w, cookie)
	}
	fmt.Fprintln(w, "Hello World")
	fmt.Println(cookie)
}

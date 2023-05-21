package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

const (
	hmacKey = "ourkey"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/authenticate", auth)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func index(w http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("session")
	if err != nil {
		c = &http.Cookie{
			Name:  "session",
			Value: "",
		}
	}

	if req.Method == http.MethodPost {
		e := req.FormValue("email")
		code := getCode(e)
		c.Value = fmt.Sprintf("%s|%s", e, code)
	}

	http.SetCookie(w, c)

	w.Write([]byte(`<!DOCTYPE html>
	<html>
	  <body>
	    <form method="POST">
	      <input type="email" name="email">
	      <input type="submit">
	    </form>
	    <a href="/authenticate">Validate This ` + c.Value + `</a>
	  </body>
	</html>`))
}

func auth(w http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("session")
	if err != nil || c.Value == "" {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	xs := strings.Split(c.Value, "|")
	email := xs[0]
	codeRcvd := xs[1]
	codeCheck := getCode(email)

	if !hmac.Equal([]byte(codeRcvd), []byte(codeCheck)) {
		fmt.Println("HMAC codes didn't match")
		fmt.Println(codeRcvd)
		fmt.Println(codeCheck)
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	w.Write([]byte(`<!DOCTYPE html>
	<html>
	  <body>
		<h1>` + codeRcvd + ` - RECEIVED </h1>
		<h1>` + codeCheck + ` - RECALCULATED </h1>
	  </body>
	</html>`))
}

func getCode(data string) string {
	h := hmac.New(sha256.New, []byte(hmacKey))
	io.WriteString(h, data)
	return fmt.Sprintf("%x", h.Sum(nil))
}

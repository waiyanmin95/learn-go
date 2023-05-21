package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"io"
)

func main() {
	c := getCode("wai@amity.co")

	g := getCode("wai@amity.co")

	if c == g {
		fmt.Println("It's match")
	} else {
		fmt.Println("It's not match")
	}
}

func getCode(s string) string {
	h := hmac.New(sha256.New, []byte("holy"))
	io.WriteString(h, s)
	return fmt.Sprintf("%x", h.Sum(nil))
}
